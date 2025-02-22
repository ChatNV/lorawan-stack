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

package templates

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "client_requested",
		email.FSTemplate{
			SubjectTemplate:      "A new OAuth client has been requested",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "client_requested.html.tmpl",
			TextTemplateFile:     "client_requested.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("client_requested", &email.NotificationBuilder{
		EmailTemplateName: "client_requested",
		DataBuilder:       newClientRequestedData,
	})
}

func newClientRequestedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var nData ttnpb.CreateClientRequest
	if err := ttnpb.UnmarshalAny(data.Notification().GetData(), &nData); err != nil {
		return nil, err
	}
	return &ClientRequestedData{
		NotificationTemplateData: data,
		CreateClientRequest:      &nData,
	}, nil
}

// ClientRequestedData is the data for the client_requested email.
type ClientRequestedData struct {
	email.NotificationTemplateData
	*ttnpb.CreateClientRequest
}
