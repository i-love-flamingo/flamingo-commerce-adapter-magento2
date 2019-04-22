// Code generated by go-swagger; DO NOT EDIT.

package gift_registry_guest_cart_shipping_method_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new gift registry guest cart shipping method management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gift registry guest cart shipping method management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPost Estimate shipping
*/
func (a *Client) GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPost(params *GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPostParams) (*GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIdPost",
		Method:             "POST",
		PathPattern:        "/V1/guest-giftregistry/{cartId}/estimate-shipping-methods",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftRegistryGuestCartShippingMethodManagementV1EstimateByRegistryIDPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
