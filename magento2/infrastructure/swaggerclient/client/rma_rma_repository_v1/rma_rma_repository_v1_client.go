// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rma rma repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rma rma repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RmaRmaRepositoryV1DeleteDelete Delete RMA
*/
func (a *Client) RmaRmaRepositoryV1DeleteDelete(params *RmaRmaRepositoryV1DeleteDeleteParams) (*RmaRmaRepositoryV1DeleteDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRmaRmaRepositoryV1DeleteDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rmaRmaRepositoryV1DeleteDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/returns/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RmaRmaRepositoryV1DeleteDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RmaRmaRepositoryV1DeleteDeleteOK), nil

}

/*
RmaRmaRepositoryV1GetGet Return data object for specified RMA id
*/
func (a *Client) RmaRmaRepositoryV1GetGet(params *RmaRmaRepositoryV1GetGetParams) (*RmaRmaRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRmaRmaRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rmaRmaRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/returns/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RmaRmaRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RmaRmaRepositoryV1GetGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}