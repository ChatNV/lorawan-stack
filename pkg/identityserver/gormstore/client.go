// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package store

import (
	"github.com/lib/pq"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Client model.
type Client struct {
	Model
	SoftDelete

	// BEGIN common fields
	ClientID    string `gorm:"unique_index:client_id_index;type:VARCHAR(36);not null"`
	Name        string `gorm:"type:VARCHAR"`
	Description string `gorm:"type:TEXT"`

	Attributes  []Attribute  `gorm:"polymorphic:Entity;polymorphic_value:client"`
	Memberships []Membership `gorm:"polymorphic:Entity;polymorphic_value:client"`

	AdministrativeContactID *string  `gorm:"type:UUID;index"`
	AdministrativeContact   *Account `gorm:"save_associations:false"`

	TechnicalContactID *string  `gorm:"type:UUID;index"`
	TechnicalContact   *Account `gorm:"save_associations:false"`
	// END common fields

	ClientSecret       string         `gorm:"type:VARCHAR"`
	RedirectURIs       pq.StringArray `gorm:"type:VARCHAR ARRAY;column:redirect_uris"`
	LogoutRedirectURIs pq.StringArray `gorm:"type:VARCHAR ARRAY;column:logout_redirect_uris"`

	State            int    `gorm:"not null"`
	StateDescription string `gorm:"type:VARCHAR"`

	SkipAuthorization bool `gorm:"not null"`
	Endorsed          bool `gorm:"not null"`

	Grants Grants `gorm:"type:INT ARRAY"`
	Rights Rights `gorm:"type:INT ARRAY"`
}

func init() {
	registerModel(&Client{})
}

// functions to set fields from the client model into the client proto.
var clientPBSetters = map[string]func(*ttnpb.Client, *Client){
	nameField: func(pb *ttnpb.Client, cli *Client) {
		pb.Name = cli.Name
	},
	descriptionField: func(pb *ttnpb.Client, cli *Client) {
		pb.Description = cli.Description
	},
	attributesField: func(pb *ttnpb.Client, cli *Client) {
		pb.Attributes = attributes(cli.Attributes).toMap()
	},
	administrativeContactField: func(pb *ttnpb.Client, cli *Client) {
		if cli.AdministrativeContact != nil {
			pb.AdministrativeContact = cli.AdministrativeContact.OrganizationOrUserIdentifiers()
		}
	},
	technicalContactField: func(pb *ttnpb.Client, cli *Client) {
		if cli.TechnicalContact != nil {
			pb.TechnicalContact = cli.TechnicalContact.OrganizationOrUserIdentifiers()
		}
	},
	secretField: func(pb *ttnpb.Client, cli *Client) {
		pb.Secret = cli.ClientSecret
	},
	redirectURIsField: func(pb *ttnpb.Client, cli *Client) {
		pb.RedirectUris = cli.RedirectURIs
	},
	logoutRedirectURIsField: func(pb *ttnpb.Client, cli *Client) {
		pb.LogoutRedirectUris = cli.LogoutRedirectURIs
	},
	stateField: func(pb *ttnpb.Client, cli *Client) {
		pb.State = ttnpb.State(cli.State)
	},
	stateDescriptionField: func(pb *ttnpb.Client, cli *Client) {
		pb.StateDescription = cli.StateDescription
	},
	skipAuthorizationField: func(pb *ttnpb.Client, cli *Client) {
		pb.SkipAuthorization = cli.SkipAuthorization
	},
	endorsedField: func(pb *ttnpb.Client, cli *Client) {
		pb.Endorsed = cli.Endorsed
	},
	grantsField: func(pb *ttnpb.Client, cli *Client) {
		pb.Grants = cli.Grants
	},
	rightsField: func(pb *ttnpb.Client, cli *Client) {
		pb.Rights = cli.Rights
	},
}

// functions to set fields from the client proto into the client model.
var clientModelSetters = map[string]func(*Client, *ttnpb.Client){
	nameField: func(cli *Client, pb *ttnpb.Client) {
		cli.Name = pb.Name
	},
	descriptionField: func(cli *Client, pb *ttnpb.Client) {
		cli.Description = pb.Description
	},
	attributesField: func(cli *Client, pb *ttnpb.Client) {
		cli.Attributes = attributes(cli.Attributes).updateFromMap(pb.Attributes)
	},
	administrativeContactField: func(cli *Client, pb *ttnpb.Client) {
		if pb.AdministrativeContact == nil {
			cli.AdministrativeContact = nil
			return
		}
		cli.AdministrativeContact = &Account{
			AccountType: pb.AdministrativeContact.EntityType(),
			UID:         pb.AdministrativeContact.IDString(),
		}
	},
	technicalContactField: func(cli *Client, pb *ttnpb.Client) {
		if pb.TechnicalContact == nil {
			cli.TechnicalContact = nil
			return
		}
		cli.TechnicalContact = &Account{
			AccountType: pb.TechnicalContact.EntityType(),
			UID:         pb.TechnicalContact.IDString(),
		}
	},
	secretField: func(cli *Client, pb *ttnpb.Client) {
		cli.ClientSecret = pb.Secret
	},
	redirectURIsField: func(cli *Client, pb *ttnpb.Client) {
		cli.RedirectURIs = pq.StringArray(pb.RedirectUris)
	},
	logoutRedirectURIsField: func(cli *Client, pb *ttnpb.Client) {
		cli.LogoutRedirectURIs = pq.StringArray(pb.LogoutRedirectUris)
	},
	stateField: func(cli *Client, pb *ttnpb.Client) {
		cli.State = int(pb.State)
	},
	stateDescriptionField: func(cli *Client, pb *ttnpb.Client) {
		cli.StateDescription = pb.StateDescription
	},
	skipAuthorizationField: func(cli *Client, pb *ttnpb.Client) {
		cli.SkipAuthorization = pb.SkipAuthorization
	},
	endorsedField: func(cli *Client, pb *ttnpb.Client) {
		cli.Endorsed = pb.Endorsed
	},
	grantsField: func(cli *Client, pb *ttnpb.Client) {
		cli.Grants = pb.Grants
	},
	rightsField: func(cli *Client, pb *ttnpb.Client) {
		cli.Rights = pb.Rights
	},
}

// fieldMask to use if a nil or empty fieldmask is passed.
var defaultClientFieldMask store.FieldMask

func init() {
	paths := make([]string, 0, len(clientPBSetters))
	for _, path := range ttnpb.ClientFieldPathsNested {
		if _, ok := clientPBSetters[path]; ok {
			paths = append(paths, path)
		}
	}
	defaultClientFieldMask = paths
}

// fieldmask path to column name in clients table.
var clientColumnNames = map[string][]string{
	attributesField:            {},
	contactInfoField:           {},
	nameField:                  {nameField},
	descriptionField:           {descriptionField},
	secretField:                {"client_secret"},
	redirectURIsField:          {redirectURIsField},
	logoutRedirectURIsField:    {logoutRedirectURIsField},
	stateField:                 {stateField},
	stateDescriptionField:      {stateDescriptionField},
	skipAuthorizationField:     {skipAuthorizationField},
	endorsedField:              {endorsedField},
	grantsField:                {grantsField},
	rightsField:                {rightsField},
	administrativeContactField: {administrativeContactField + "_id"},
	technicalContactField:      {technicalContactField + "_id"},
}

func (cli Client) toPB(pb *ttnpb.Client, fieldMask store.FieldMask) {
	pb.Ids = &ttnpb.ClientIdentifiers{ClientId: cli.ClientID}
	pb.CreatedAt = ttnpb.ProtoTimePtr(cleanTime(cli.CreatedAt))
	pb.UpdatedAt = ttnpb.ProtoTimePtr(cleanTime(cli.UpdatedAt))
	pb.DeletedAt = ttnpb.ProtoTime(cleanTimePtr(cli.DeletedAt))
	if len(fieldMask) == 0 {
		fieldMask = defaultClientFieldMask
	}
	for _, path := range fieldMask {
		if setter, ok := clientPBSetters[path]; ok {
			setter(pb, &cli)
		}
	}
}

func (cli *Client) fromPB(pb *ttnpb.Client, fieldMask store.FieldMask) (columns []string) {
	if len(fieldMask) == 0 {
		fieldMask = defaultClientFieldMask
	}
	for _, path := range fieldMask {
		if setter, ok := clientModelSetters[path]; ok {
			setter(cli, pb)
			if columnNames, ok := clientColumnNames[path]; ok {
				columns = append(columns, columnNames...)
			}
			continue
		}
	}
	return columns
}
