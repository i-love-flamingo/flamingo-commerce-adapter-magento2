// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_item_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quote guest cart item repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest cart item repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QuoteGuestCartItemRepositoryV1DeleteByIDDelete Remove the specified item from the specified cart.
*/
func (a *Client) QuoteGuestCartItemRepositoryV1DeleteByIDDelete(params *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteParams) (*QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartItemRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/guest-carts/{cartId}/items/{itemId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartItemRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK), nil

}

/*
QuoteGuestCartItemRepositoryV1GetListGet List items that are assigned to a specified cart.
*/
func (a *Client) QuoteGuestCartItemRepositoryV1GetListGet(params *QuoteGuestCartItemRepositoryV1GetListGetParams) (*QuoteGuestCartItemRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartItemRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartItemRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}/items",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartItemRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartItemRepositoryV1GetListGetOK), nil

}

/*
QuoteGuestCartItemRepositoryV1SavePost Add/update the specified cart item.
*/
func (a *Client) QuoteGuestCartItemRepositoryV1SavePost(params *QuoteGuestCartItemRepositoryV1SavePostParams) (*QuoteGuestCartItemRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartItemRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartItemRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/guest-carts/{cartId}/items",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartItemRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartItemRepositoryV1SavePostOK), nil

}

/*
QuoteGuestCartItemRepositoryV1SavePut Add/update the specified cart item.
*/
func (a *Client) QuoteGuestCartItemRepositoryV1SavePut(params *QuoteGuestCartItemRepositoryV1SavePutParams) (*QuoteGuestCartItemRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartItemRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGuestCartItemRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/guest-carts/{cartId}/items/{itemId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteGuestCartItemRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGuestCartItemRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
