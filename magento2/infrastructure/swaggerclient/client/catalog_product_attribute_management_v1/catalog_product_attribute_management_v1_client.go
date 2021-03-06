// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductAttributeManagementV1AssignPost Assign attribute to attribute set
*/
func (a *Client) CatalogProductAttributeManagementV1AssignPost(params *CatalogProductAttributeManagementV1AssignPostParams) (*CatalogProductAttributeManagementV1AssignPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeManagementV1AssignPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeManagementV1AssignPost",
		Method:             "POST",
		PathPattern:        "/V1/products/attribute-sets/attributes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeManagementV1AssignPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeManagementV1AssignPostOK), nil

}

/*
CatalogProductAttributeManagementV1GetAttributesGet Retrieve related attributes based on given attribute set ID
*/
func (a *Client) CatalogProductAttributeManagementV1GetAttributesGet(params *CatalogProductAttributeManagementV1GetAttributesGetParams) (*CatalogProductAttributeManagementV1GetAttributesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeManagementV1GetAttributesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeManagementV1GetAttributesGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attribute-sets/{attributeSetId}/attributes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeManagementV1GetAttributesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeManagementV1GetAttributesGetOK), nil

}

/*
CatalogProductAttributeManagementV1UnassignDelete Remove attribute from attribute set
*/
func (a *Client) CatalogProductAttributeManagementV1UnassignDelete(params *CatalogProductAttributeManagementV1UnassignDeleteParams) (*CatalogProductAttributeManagementV1UnassignDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeManagementV1UnassignDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeManagementV1UnassignDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/attribute-sets/{attributeSetId}/attributes/{attributeCode}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeManagementV1UnassignDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeManagementV1UnassignDeleteOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
