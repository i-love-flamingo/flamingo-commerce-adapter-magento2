// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog category management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog category management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogCategoryManagementV1GetTreeGet Retrieve list of categories
*/
func (a *Client) CatalogCategoryManagementV1GetTreeGet(params *CatalogCategoryManagementV1GetTreeGetParams) (*CatalogCategoryManagementV1GetTreeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCategoryManagementV1GetTreeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogCategoryManagementV1GetTreeGet",
		Method:             "GET",
		PathPattern:        "/V1/categories",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogCategoryManagementV1GetTreeGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogCategoryManagementV1GetTreeGetOK), nil

}

/*
CatalogCategoryManagementV1MovePut Move category
*/
func (a *Client) CatalogCategoryManagementV1MovePut(params *CatalogCategoryManagementV1MovePutParams) (*CatalogCategoryManagementV1MovePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCategoryManagementV1MovePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogCategoryManagementV1MovePut",
		Method:             "PUT",
		PathPattern:        "/V1/categories/{categoryId}/move",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogCategoryManagementV1MovePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogCategoryManagementV1MovePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
