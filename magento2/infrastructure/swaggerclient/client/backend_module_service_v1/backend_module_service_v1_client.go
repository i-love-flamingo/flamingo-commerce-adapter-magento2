// Code generated by go-swagger; DO NOT EDIT.

package backend_module_service_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new backend module service v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backend module service v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BackendModuleServiceV1GetModulesGet Returns an array of enabled modules
*/
func (a *Client) BackendModuleServiceV1GetModulesGet(params *BackendModuleServiceV1GetModulesGetParams) (*BackendModuleServiceV1GetModulesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackendModuleServiceV1GetModulesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backendModuleServiceV1GetModulesGet",
		Method:             "GET",
		PathPattern:        "/V1/modules",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackendModuleServiceV1GetModulesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackendModuleServiceV1GetModulesGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}