// Code generated by go-swagger; DO NOT EDIT.

package gift_message_guest_cart_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new gift message guest cart repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gift message guest cart repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GiftMessageGuestCartRepositoryV1GetGet Return the gift message for a specified order.
*/
func (a *Client) GiftMessageGuestCartRepositoryV1GetGet(params *GiftMessageGuestCartRepositoryV1GetGetParams) (*GiftMessageGuestCartRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftMessageGuestCartRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftMessageGuestCartRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}/gift-message",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftMessageGuestCartRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftMessageGuestCartRepositoryV1GetGetOK), nil

}

/*
GiftMessageGuestCartRepositoryV1SavePost Set the gift message for an entire order.
*/
func (a *Client) GiftMessageGuestCartRepositoryV1SavePost(params *GiftMessageGuestCartRepositoryV1SavePostParams) (*GiftMessageGuestCartRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftMessageGuestCartRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftMessageGuestCartRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/guest-carts/{cartId}/gift-message",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftMessageGuestCartRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftMessageGuestCartRepositoryV1SavePostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
