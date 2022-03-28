// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package enddevices provides functions to configure End Device claiming clients.
package enddevices

import (
	"context"
	"path/filepath"
	"strings"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/deviceclaimingserver/enddevices/ttjs"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/qrcode"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

// EndDeviceClaimer provides methods for Claiming End Devices on (external) Join Server.
type EndDeviceClaimer interface {
	// SupportsJoinEUI returns whether the Join Server supports this JoinEUI.
	SupportsJoinEUI(joinEUI types.EUI64) bool
	// Claim claims an End Device.
	Claim(ctx context.Context, req *ttnpb.ClaimEndDeviceRequest) (*ttnpb.EndDeviceIdentifiers, error)
	// GetClaimStatus returns the claim status an End Device.
	GetClaimStatus(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (*ttnpb.GetClaimStatusResponse, error)
	// Unclaim releases the claim on an End Device.
	Unclaim(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (err error)
}

// Component abstracts the underlying *component.Component.
type Component interface {
	httpclient.Provider
	GetBaseConfig(ctx context.Context) config.ServiceBase
	GetPeerConn(ctx context.Context, role ttnpb.ClusterRole, ids cluster.EntityIdentifiers) (*grpc.ClientConn, error)
	AllowInsecureForCredentials() bool
}

const ttJSType = "ttjs"

var errInvalidUpstream = errors.DefineInvalidArgument("invalid_upstream", "upstream `{type}` is invalid")

// Upstream abstracts EndDeviceClaimingServer.
type Upstream struct {
	Component
	deviceRegistry ttnpb.EndDeviceRegistryClient
	servers        map[string]EndDeviceClaimer
}

// NewUpstream returns a new Upstream.
func NewUpstream(ctx context.Context, conf Config, c Component, opts ...Option) (*Upstream, error) {
	upstream := &Upstream{
		Component: c,
		servers:   make(map[string]EndDeviceClaimer),
	}
	fetcher, err := conf.Fetcher(ctx, c.GetBaseConfig(ctx).Blob, c)
	if err != nil {
		return nil, err
	}
	if fetcher == nil {
		return upstream, nil
	}
	baseConfigBytes, err := fetcher.File(JSClientConfigurationName)
	if err != nil {
		return nil, err
	}
	var baseConfig baseConfig
	if err := yaml.UnmarshalStrict(baseConfigBytes, &baseConfig); err != nil {
		return nil, err
	}

	// Setup upstreams.
	for _, js := range baseConfig.JoinServers {
		var (
			s          EndDeviceClaimer
			clientName string
		)
		switch js.Type {
		case ttJSType:
			// Fetch and parse configuration.
			fileParts := strings.Split(filepath.ToSlash(js.File), "/")
			fetcher := fetch.WithBasePath(fetcher, fileParts[:len(fileParts)-1]...)
			fileName := fileParts[len(fileParts)-1]
			configBytes, err := fetcher.File(fileName)
			if err != nil {
				return nil, err
			}

			var ttjsConfig ttjs.Config
			if err := yaml.UnmarshalStrict(configBytes, &ttjsConfig); err != nil {
				return nil, err
			}
			ttjsConfig.NetID = conf.NetID
			ttjsConfig.JoinEUIPrefixes = js.JoinEUIs
			s, err = ttjsConfig.NewClient(ctx, c)
			if err != nil {
				return nil, err
			}
			// The file for each client will be unique.
			clientName = strings.Trim(fileName, filepath.Ext(fileName))
		default:
			return nil, errInvalidUpstream.WithAttributes("type", js.Type)
		}
		upstream.servers[clientName] = s
	}

	for _, opt := range opts {
		opt(upstream)
	}

	return upstream, nil
}

// Option configures Upstream.
type Option func(*Upstream)

// WithDeviceRegistry overrides the device registry of the Upstream.
func WithDeviceRegistry(reg ttnpb.EndDeviceRegistryClient) Option {
	return func(upstream *Upstream) {
		upstream.deviceRegistry = reg
	}
}

var (
	errParseQRCode          = errors.Define("parse_qr_code", "parse QR code failed")
	errQRCodeData           = errors.DefineInvalidArgument("qr_code_data", "invalid QR code data")
	errNoJoinEUI            = errors.DefineInvalidArgument("no_join_eui", "failed to extract JoinEUI from request")
	errNoEUI                = errors.DefineInvalidArgument("no_eui", "DevEUI/JoinEUI not found in request")
	errClaimingNotSupported = errors.DefineAborted("claiming_not_supported", "claiming not supported for JoinEUI `{eui}`")
)

func (upstream *Upstream) joinEUIClaimer(ctx context.Context, joinEUI types.EUI64) EndDeviceClaimer {
	for name, srv := range upstream.servers {
		if !srv.SupportsJoinEUI(joinEUI) {
			continue
		}
		log.FromContext(ctx).WithFields(log.Fields(
			"name", name,
			"join_eui", joinEUI,
		)).Debug("JoinEUI supported by upstream")
		return srv
	}
	return nil
}

// Claim implements EndDeviceClaimingServer.
func (upstream *Upstream) Claim(ctx context.Context, req *ttnpb.ClaimEndDeviceRequest) (ids *ttnpb.EndDeviceIdentifiers, err error) {
	// Check that the collaborator has necessary rights before attempting to claim it on an upstream.
	// Since this is part of the create device flow, we check that the collaborator has the rights to create devices in the application.
	targetAppID := req.GetTargetApplicationIds()
	if err := rights.RequireApplication(ctx, *targetAppID,
		ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE,
	); err != nil {
		return nil, err
	}

	var joinEUI types.EUI64
	if authenticatedIDs := req.GetAuthenticatedIdentifiers(); authenticatedIDs != nil {
		joinEUI = req.GetAuthenticatedIdentifiers().JoinEui
	} else if qrCode := req.GetQrCode(); qrCode != nil {
		data, err := qrcode.Parse(qrCode)
		if err != nil {
			return nil, errParseQRCode.WithCause(err)
		}
		authIDs, ok := data.(qrcode.AuthenticatedEndDeviceIdentifiers)
		if !ok {
			return nil, errQRCodeData.New()
		}
		joinEUI, _, _ = authIDs.AuthenticatedEndDeviceIdentifiers()
	} else {
		return nil, errNoJoinEUI.New()
	}

	claimer := upstream.joinEUIClaimer(ctx, joinEUI)
	if claimer == nil {
		return nil, errClaimingNotSupported.WithAttributes("eui", joinEUI)
	}
	return claimer.Claim(ctx, req)
}

// Unclaim implements EndDeviceClaimingServer.
func (upstream *Upstream) Unclaim(ctx context.Context, in *ttnpb.EndDeviceIdentifiers) (*pbtypes.Empty, error) {
	if in.DevEui == nil || in.JoinEui == nil {
		return nil, errNoEUI.New()
	}
	err := upstream.requireRights(ctx, in, ttnpb.Rights{
		Rights: []ttnpb.Right{
			ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE,
		},
	})
	if err != nil {
		return nil, err
	}
	claimer := upstream.joinEUIClaimer(ctx, *in.JoinEui)
	if claimer == nil {
		return nil, errClaimingNotSupported.WithAttributes("eui", in.JoinEui)
	}
	err = claimer.Unclaim(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pbtypes.Empty{}, nil

}

// GetInfoByJoinEUI implements EndDeviceClaimingServer.
func (upstream *Upstream) GetInfoByJoinEUI(ctx context.Context, in *ttnpb.GetInfoByJoinEUIRequest) (*ttnpb.GetInfoByJoinEUIResponse, error) {
	claimer := upstream.joinEUIClaimer(ctx, *in.JoinEui)
	return &ttnpb.GetInfoByJoinEUIResponse{
		JoinEui:          in.JoinEui,
		SupportsClaiming: (claimer != nil),
	}, nil
}

// GetClaimStatus implements EndDeviceClaimingServer.
func (upstream *Upstream) GetClaimStatus(ctx context.Context, in *ttnpb.EndDeviceIdentifiers) (*ttnpb.GetClaimStatusResponse, error) {
	if in.DevEui == nil || in.JoinEui == nil {
		return nil, errNoEUI.New()
	}
	err := upstream.requireRights(ctx, in, ttnpb.Rights{
		Rights: []ttnpb.Right{
			ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ,
		},
	})
	if err != nil {
		return nil, err
	}
	claimer := upstream.joinEUIClaimer(ctx, *in.JoinEui)
	if claimer == nil {
		return nil, errClaimingNotSupported.WithAttributes("eui", in.JoinEui)
	}
	return claimer.GetClaimStatus(ctx, in)
}

func (upstream *Upstream) requireRights(ctx context.Context, in *ttnpb.EndDeviceIdentifiers, appRights ttnpb.Rights) error {
	// Collaborator must have the required rights on the application.
	if err := rights.RequireApplication(ctx, *in.ApplicationIds,
		appRights.Rights...,
	); err != nil {
		return err
	}
	// Check that the device actually exists in the application.
	// If the EUIs are set in the request, the IS also checks that they match the stored device.
	callOpt, err := rpcmetadata.WithForwardedAuth(ctx, upstream.Component.AllowInsecureForCredentials())
	if err != nil {
		return err
	}
	er, err := upstream.getDeviceRegistry(ctx)
	if err != nil {
		return err
	}
	_, err = er.Get(ctx, &ttnpb.GetEndDeviceRequest{
		EndDeviceIds: in,
	}, callOpt)
	return err
}

func (upstream *Upstream) getDeviceRegistry(ctx context.Context) (ttnpb.EndDeviceRegistryClient, error) {
	if upstream.deviceRegistry != nil {
		return upstream.deviceRegistry, nil
	}
	conn, err := upstream.Component.GetPeerConn(ctx, ttnpb.ClusterRole_ENTITY_REGISTRY, nil)
	if err != nil {
		return nil, err
	}
	return ttnpb.NewEndDeviceRegistryClient(conn), nil
}