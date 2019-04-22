// Code generated by go-swagger; DO NOT EDIT.

package shared_catalog_product_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new shared catalog product management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shared catalog product management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SharedCatalogProductManagementV1AssignProductsPost Add products into the shared catalog.
*/
func (a *Client) SharedCatalogProductManagementV1AssignProductsPost(params *SharedCatalogProductManagementV1AssignProductsPostParams) (*SharedCatalogProductManagementV1AssignProductsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogProductManagementV1AssignProductsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogProductManagementV1AssignProductsPost",
		Method:             "POST",
		PathPattern:        "/V1/sharedCatalog/{id}/assignProducts",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogProductManagementV1AssignProductsPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogProductManagementV1AssignProductsPostOK), nil

}

/*
SharedCatalogProductManagementV1GetProductsGet Return the list of product SKUs in the selected shared catalog.
*/
func (a *Client) SharedCatalogProductManagementV1GetProductsGet(params *SharedCatalogProductManagementV1GetProductsGetParams) (*SharedCatalogProductManagementV1GetProductsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogProductManagementV1GetProductsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogProductManagementV1GetProductsGet",
		Method:             "GET",
		PathPattern:        "/V1/sharedCatalog/{id}/products",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogProductManagementV1GetProductsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogProductManagementV1GetProductsGetOK), nil

}

/*
SharedCatalogProductManagementV1UnassignProductsPost Remove the specified products from the shared catalog.
*/
func (a *Client) SharedCatalogProductManagementV1UnassignProductsPost(params *SharedCatalogProductManagementV1UnassignProductsPostParams) (*SharedCatalogProductManagementV1UnassignProductsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogProductManagementV1UnassignProductsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogProductManagementV1UnassignProductsPost",
		Method:             "POST",
		PathPattern:        "/V1/sharedCatalog/{id}/unassignProducts",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogProductManagementV1UnassignProductsPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogProductManagementV1UnassignProductsPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}