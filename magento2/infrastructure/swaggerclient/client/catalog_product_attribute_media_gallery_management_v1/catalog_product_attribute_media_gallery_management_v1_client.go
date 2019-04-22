// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_media_gallery_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute media gallery management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute media gallery management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductAttributeMediaGalleryManagementV1CreatePost Create new gallery entry
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1CreatePost(params *CatalogProductAttributeMediaGalleryManagementV1CreatePostParams) (*CatalogProductAttributeMediaGalleryManagementV1CreatePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1CreatePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1CreatePost",
		Method:             "POST",
		PathPattern:        "/V1/products/{sku}/media",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1CreatePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeMediaGalleryManagementV1CreatePostOK), nil

}

/*
CatalogProductAttributeMediaGalleryManagementV1GetGet Return information about gallery entry
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1GetGet(params *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) (*CatalogProductAttributeMediaGalleryManagementV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}/media/{entryId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeMediaGalleryManagementV1GetGetOK), nil

}

/*
CatalogProductAttributeMediaGalleryManagementV1GetListGet Retrieve the list of gallery entries associated with given product
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1GetListGet(params *CatalogProductAttributeMediaGalleryManagementV1GetListGetParams) (*CatalogProductAttributeMediaGalleryManagementV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}/media",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeMediaGalleryManagementV1GetListGetOK), nil

}

/*
CatalogProductAttributeMediaGalleryManagementV1RemoveDelete Remove gallery entry
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1RemoveDelete(params *CatalogProductAttributeMediaGalleryManagementV1RemoveDeleteParams) (*CatalogProductAttributeMediaGalleryManagementV1RemoveDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1RemoveDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1RemoveDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/{sku}/media/{entryId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1RemoveDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeMediaGalleryManagementV1RemoveDeleteOK), nil

}

/*
CatalogProductAttributeMediaGalleryManagementV1UpdatePut Update gallery entry
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1UpdatePut(params *CatalogProductAttributeMediaGalleryManagementV1UpdatePutParams) (*CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1UpdatePut",
		Method:             "PUT",
		PathPattern:        "/V1/products/{sku}/media/{entryId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1UpdatePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}