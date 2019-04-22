// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_link_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product link management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product link management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductLinkManagementV1GetLinkedItemsByTypeGet Provide the list of links for a specific product
*/
func (a *Client) CatalogProductLinkManagementV1GetLinkedItemsByTypeGet(params *CatalogProductLinkManagementV1GetLinkedItemsByTypeGetParams) (*CatalogProductLinkManagementV1GetLinkedItemsByTypeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductLinkManagementV1GetLinkedItemsByTypeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductLinkManagementV1GetLinkedItemsByTypeGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}/links/{type}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductLinkManagementV1GetLinkedItemsByTypeGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductLinkManagementV1GetLinkedItemsByTypeGetOK), nil

}

/*
CatalogProductLinkManagementV1SetProductLinksPost Assign a product link to another product
*/
func (a *Client) CatalogProductLinkManagementV1SetProductLinksPost(params *CatalogProductLinkManagementV1SetProductLinksPostParams) (*CatalogProductLinkManagementV1SetProductLinksPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductLinkManagementV1SetProductLinksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductLinkManagementV1SetProductLinksPost",
		Method:             "POST",
		PathPattern:        "/V1/products/{sku}/links",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductLinkManagementV1SetProductLinksPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductLinkManagementV1SetProductLinksPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}