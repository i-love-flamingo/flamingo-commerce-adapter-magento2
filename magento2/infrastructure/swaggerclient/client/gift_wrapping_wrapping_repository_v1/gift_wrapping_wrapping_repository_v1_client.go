// Code generated by go-swagger; DO NOT EDIT.

package gift_wrapping_wrapping_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new gift wrapping wrapping repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gift wrapping wrapping repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GiftWrappingWrappingRepositoryV1DeleteByIDDelete Delete gift wrapping
*/
func (a *Client) GiftWrappingWrappingRepositoryV1DeleteByIDDelete(params *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteParams) (*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftWrappingWrappingRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/gift-wrappings/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftWrappingWrappingRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK), nil

}

/*
GiftWrappingWrappingRepositoryV1GetGet Return data object for specified wrapping ID and store.
*/
func (a *Client) GiftWrappingWrappingRepositoryV1GetGet(params *GiftWrappingWrappingRepositoryV1GetGetParams) (*GiftWrappingWrappingRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftWrappingWrappingRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftWrappingWrappingRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/gift-wrappings/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftWrappingWrappingRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftWrappingWrappingRepositoryV1GetGetOK), nil

}

/*
GiftWrappingWrappingRepositoryV1GetListGet Return list of gift wrapping data objects based on search criteria
*/
func (a *Client) GiftWrappingWrappingRepositoryV1GetListGet(params *GiftWrappingWrappingRepositoryV1GetListGetParams) (*GiftWrappingWrappingRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftWrappingWrappingRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftWrappingWrappingRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/gift-wrappings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftWrappingWrappingRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftWrappingWrappingRepositoryV1GetListGetOK), nil

}

/*
GiftWrappingWrappingRepositoryV1SavePost Create/Update new gift wrapping with data object values
*/
func (a *Client) GiftWrappingWrappingRepositoryV1SavePost(params *GiftWrappingWrappingRepositoryV1SavePostParams) (*GiftWrappingWrappingRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftWrappingWrappingRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftWrappingWrappingRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/gift-wrappings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftWrappingWrappingRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftWrappingWrappingRepositoryV1SavePostOK), nil

}

/*
GiftWrappingWrappingRepositoryV1SavePut Create/Update new gift wrapping with data object values
*/
func (a *Client) GiftWrappingWrappingRepositoryV1SavePut(params *GiftWrappingWrappingRepositoryV1SavePutParams) (*GiftWrappingWrappingRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGiftWrappingWrappingRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "giftWrappingWrappingRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/gift-wrappings/{wrappingId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GiftWrappingWrappingRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GiftWrappingWrappingRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}