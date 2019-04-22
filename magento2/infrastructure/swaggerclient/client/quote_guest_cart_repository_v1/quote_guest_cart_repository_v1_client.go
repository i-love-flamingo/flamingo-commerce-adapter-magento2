// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quote guest cart repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest cart repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QuoteGuestCartRepositoryV1GetGet Enable a guest user to return information for a specified cart.
*/
func (a *Client) QuoteGuestCartRepositoryV1GetGet(params *QuoteGuestCartRepositoryV1GetGetParams) (*QuoteGuestCartRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartRepositoryV1GetGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}