// Code generated by go-swagger; DO NOT EDIT.

package shared_catalog_category_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new shared catalog category management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shared catalog category management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SharedCatalogCategoryManagementV1AssignCategoriesPost Add categories into the shared catalog.
*/
func (a *Client) SharedCatalogCategoryManagementV1AssignCategoriesPost(params *SharedCatalogCategoryManagementV1AssignCategoriesPostParams) (*SharedCatalogCategoryManagementV1AssignCategoriesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogCategoryManagementV1AssignCategoriesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogCategoryManagementV1AssignCategoriesPost",
		Method:             "POST",
		PathPattern:        "/V1/sharedCatalog/{id}/assignCategories",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogCategoryManagementV1AssignCategoriesPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogCategoryManagementV1AssignCategoriesPostOK), nil

}

/*
SharedCatalogCategoryManagementV1GetCategoriesGet Return the list of categories in the selected shared catalog.
*/
func (a *Client) SharedCatalogCategoryManagementV1GetCategoriesGet(params *SharedCatalogCategoryManagementV1GetCategoriesGetParams) (*SharedCatalogCategoryManagementV1GetCategoriesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogCategoryManagementV1GetCategoriesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogCategoryManagementV1GetCategoriesGet",
		Method:             "GET",
		PathPattern:        "/V1/sharedCatalog/{id}/categories",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogCategoryManagementV1GetCategoriesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogCategoryManagementV1GetCategoriesGetOK), nil

}

/*
SharedCatalogCategoryManagementV1UnassignCategoriesPost Remove the specified categories from the shared catalog.
*/
func (a *Client) SharedCatalogCategoryManagementV1UnassignCategoriesPost(params *SharedCatalogCategoryManagementV1UnassignCategoriesPostParams) (*SharedCatalogCategoryManagementV1UnassignCategoriesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedCatalogCategoryManagementV1UnassignCategoriesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sharedCatalogCategoryManagementV1UnassignCategoriesPost",
		Method:             "POST",
		PathPattern:        "/V1/sharedCatalog/{id}/unassignCategories",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SharedCatalogCategoryManagementV1UnassignCategoriesPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedCatalogCategoryManagementV1UnassignCategoriesPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
