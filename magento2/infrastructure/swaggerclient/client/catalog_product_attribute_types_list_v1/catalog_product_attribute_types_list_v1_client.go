// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_types_list_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute types list v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute types list v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductAttributeTypesListV1GetItemsGet Retrieve list of product attribute types
*/
func (a *Client) CatalogProductAttributeTypesListV1GetItemsGet(params *CatalogProductAttributeTypesListV1GetItemsGetParams) (*CatalogProductAttributeTypesListV1GetItemsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeTypesListV1GetItemsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeTypesListV1GetItemsGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attributes/types",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeTypesListV1GetItemsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeTypesListV1GetItemsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
