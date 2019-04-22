// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_option_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute option management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute option management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductAttributeOptionManagementV1AddPost Add option to attribute
*/
func (a *Client) CatalogProductAttributeOptionManagementV1AddPost(params *CatalogProductAttributeOptionManagementV1AddPostParams) (*CatalogProductAttributeOptionManagementV1AddPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeOptionManagementV1AddPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeOptionManagementV1AddPost",
		Method:             "POST",
		PathPattern:        "/V1/products/attributes/{attributeCode}/options",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeOptionManagementV1AddPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeOptionManagementV1AddPostOK), nil

}

/*
CatalogProductAttributeOptionManagementV1DeleteDelete Delete option from attribute
*/
func (a *Client) CatalogProductAttributeOptionManagementV1DeleteDelete(params *CatalogProductAttributeOptionManagementV1DeleteDeleteParams) (*CatalogProductAttributeOptionManagementV1DeleteDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeOptionManagementV1DeleteDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeOptionManagementV1DeleteDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/attributes/{attributeCode}/options/{optionId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeOptionManagementV1DeleteDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeOptionManagementV1DeleteDeleteOK), nil

}

/*
CatalogProductAttributeOptionManagementV1GetItemsGet Retrieve list of attribute options
*/
func (a *Client) CatalogProductAttributeOptionManagementV1GetItemsGet(params *CatalogProductAttributeOptionManagementV1GetItemsGetParams) (*CatalogProductAttributeOptionManagementV1GetItemsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeOptionManagementV1GetItemsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeOptionManagementV1GetItemsGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attributes/{attributeCode}/options",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeOptionManagementV1GetItemsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeOptionManagementV1GetItemsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
