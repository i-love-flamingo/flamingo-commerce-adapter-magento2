// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_total_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quote guest cart total management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest cart total management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QuoteGuestCartTotalManagementV1CollectTotalsPut Set shipping/billing methods and additional data for cart and collect totals for guest.
*/
func (a *Client) QuoteGuestCartTotalManagementV1CollectTotalsPut(params *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) (*QuoteGuestCartTotalManagementV1CollectTotalsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartTotalManagementV1CollectTotalsPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartTotalManagementV1CollectTotalsPut",
		Method:             "PUT",
		PathPattern:        "/V1/guest-carts/{cartId}/collect-totals",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartTotalManagementV1CollectTotalsPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartTotalManagementV1CollectTotalsPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
