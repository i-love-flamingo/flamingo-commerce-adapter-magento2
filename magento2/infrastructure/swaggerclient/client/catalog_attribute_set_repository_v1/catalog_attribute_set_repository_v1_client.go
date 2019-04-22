// Code generated by go-swagger; DO NOT EDIT.

package catalog_attribute_set_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog attribute set repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog attribute set repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogAttributeSetRepositoryV1DeleteByIDDelete Remove attribute set by given ID
*/
func (a *Client) CatalogAttributeSetRepositoryV1DeleteByIDDelete(params *CatalogAttributeSetRepositoryV1DeleteByIDDeleteParams) (*CatalogAttributeSetRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAttributeSetRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogAttributeSetRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/attribute-sets/{attributeSetId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogAttributeSetRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogAttributeSetRepositoryV1DeleteByIDDeleteOK), nil

}

/*
CatalogAttributeSetRepositoryV1GetGet Retrieve attribute set information based on given ID
*/
func (a *Client) CatalogAttributeSetRepositoryV1GetGet(params *CatalogAttributeSetRepositoryV1GetGetParams) (*CatalogAttributeSetRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAttributeSetRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogAttributeSetRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attribute-sets/{attributeSetId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogAttributeSetRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogAttributeSetRepositoryV1GetGetOK), nil

}

/*
CatalogAttributeSetRepositoryV1GetListGet Retrieve list of Attribute Sets
*/
func (a *Client) CatalogAttributeSetRepositoryV1GetListGet(params *CatalogAttributeSetRepositoryV1GetListGetParams) (*CatalogAttributeSetRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAttributeSetRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogAttributeSetRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attribute-sets/sets/list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogAttributeSetRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogAttributeSetRepositoryV1GetListGetOK), nil

}

/*
CatalogAttributeSetRepositoryV1SavePut Save attribute set data
*/
func (a *Client) CatalogAttributeSetRepositoryV1SavePut(params *CatalogAttributeSetRepositoryV1SavePutParams) (*CatalogAttributeSetRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAttributeSetRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogAttributeSetRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/products/attribute-sets/{attributeSetId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogAttributeSetRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogAttributeSetRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
