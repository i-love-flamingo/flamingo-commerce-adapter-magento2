// Code generated by go-swagger; DO NOT EDIT.

package analytics_link_provider_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new analytics link provider v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for analytics link provider v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AnalyticsLinkProviderV1GetGet analytics link provider v1 get get API
*/
func (a *Client) AnalyticsLinkProviderV1GetGet(params *AnalyticsLinkProviderV1GetGetParams) (*AnalyticsLinkProviderV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnalyticsLinkProviderV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "analyticsLinkProviderV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/analytics/link",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AnalyticsLinkProviderV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AnalyticsLinkProviderV1GetGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
