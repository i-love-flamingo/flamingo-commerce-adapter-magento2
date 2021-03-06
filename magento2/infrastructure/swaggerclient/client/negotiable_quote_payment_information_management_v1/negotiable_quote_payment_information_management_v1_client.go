// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_payment_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new negotiable quote payment information management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for negotiable quote payment information management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGet Get payment information
*/
func (a *Client) NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGet(params *NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGetParams) (*NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNegotiableQuotePaymentInformationManagementV1GetPaymentInformationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "negotiableQuotePaymentInformationManagementV1GetPaymentInformationGet",
		Method:             "GET",
		PathPattern:        "/V1/negotiable-carts/{cartId}/payment-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NegotiableQuotePaymentInformationManagementV1GetPaymentInformationGetOK), nil

}

/*
NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPost Set payment information and place order for a specified cart.
*/
func (a *Client) NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPost(params *NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) (*NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "negotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPost",
		Method:             "POST",
		PathPattern:        "/V1/negotiable-carts/{cartId}/payment-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NegotiableQuotePaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostOK), nil

}

/*
NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPost Set payment information for a specified cart.
*/
func (a *Client) NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPost(params *NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPostParams) (*NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNegotiableQuotePaymentInformationManagementV1SavePaymentInformationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "negotiableQuotePaymentInformationManagementV1SavePaymentInformationPost",
		Method:             "POST",
		PathPattern:        "/V1/negotiable-carts/{cartId}/set-payment-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NegotiableQuotePaymentInformationManagementV1SavePaymentInformationPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
