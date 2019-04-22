// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_shipping_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new checkout guest shipping information management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for checkout guest shipping information management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckoutGuestShippingInformationManagementV1SaveAddressInformationPost checkout guest shipping information management v1 save address information post API
*/
func (a *Client) CheckoutGuestShippingInformationManagementV1SaveAddressInformationPost(params *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) (*CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkoutGuestShippingInformationManagementV1SaveAddressInformationPost",
		Method:             "POST",
		PathPattern:        "/V1/guest-carts/{cartId}/shipping-information",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
