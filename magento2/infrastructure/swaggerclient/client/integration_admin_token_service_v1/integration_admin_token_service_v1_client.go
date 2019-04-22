// Code generated by go-swagger; DO NOT EDIT.

package integration_admin_token_service_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new integration admin token service v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for integration admin token service v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IntegrationAdminTokenServiceV1CreateAdminAccessTokenPost Create access token for admin given the admin credentials.
*/
func (a *Client) IntegrationAdminTokenServiceV1CreateAdminAccessTokenPost(params *IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostParams) (*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIntegrationAdminTokenServiceV1CreateAdminAccessTokenPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "integrationAdminTokenServiceV1CreateAdminAccessTokenPost",
		Method:             "POST",
		PathPattern:        "/V1/integration/admin/token",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
