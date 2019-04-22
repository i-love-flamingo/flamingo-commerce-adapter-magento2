// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_negotiable_quote_shipping_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new negotiable quote negotiable quote shipping management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for negotiable quote negotiable quote shipping management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPut Updates the shipping method on a negotiable quote.
*/
func (a *Client) NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPut(params *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) (*NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPut",
		Method:             "PUT",
		PathPattern:        "/V1/negotiableQuote/{quoteId}/shippingMethod",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}