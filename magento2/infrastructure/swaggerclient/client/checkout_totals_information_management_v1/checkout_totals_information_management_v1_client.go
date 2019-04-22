// Code generated by go-swagger; DO NOT EDIT.

package checkout_totals_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new checkout totals information management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for checkout totals information management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckoutTotalsInformationManagementV1CalculatePost Calculate quote totals based on address and shipping method.
*/
func (a *Client) CheckoutTotalsInformationManagementV1CalculatePost(params *CheckoutTotalsInformationManagementV1CalculatePostParams) (*CheckoutTotalsInformationManagementV1CalculatePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutTotalsInformationManagementV1CalculatePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkoutTotalsInformationManagementV1CalculatePost",
		Method:             "POST",
		PathPattern:        "/V1/carts/{cartId}/totals-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckoutTotalsInformationManagementV1CalculatePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckoutTotalsInformationManagementV1CalculatePostOK), nil

}

/*
CheckoutTotalsInformationManagementV1CalculatePostMine Calculate quote totals based on address and shipping method.
*/
func (a *Client) CheckoutTotalsInformationManagementV1CalculatePostMine(params *CheckoutTotalsInformationManagementV1CalculatePostMineParams) (*CheckoutTotalsInformationManagementV1CalculatePostMineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutTotalsInformationManagementV1CalculatePostMineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkoutTotalsInformationManagementV1CalculatePostMine",
		Method:             "POST",
		PathPattern:        "/V1/carts/mine/totals-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckoutTotalsInformationManagementV1CalculatePostMineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckoutTotalsInformationManagementV1CalculatePostMineOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}