// Code generated by go-swagger; DO NOT EDIT.

package checkout_agreements_checkout_agreements_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new checkout agreements checkout agreements repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for checkout agreements checkout agreements repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGet Lists active checkout agreements.
*/
func (a *Client) CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGet(params *CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGetParams) (*CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkoutAgreementsCheckoutAgreementsRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/carts/licence",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckoutAgreementsCheckoutAgreementsRepositoryV1GetListGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
