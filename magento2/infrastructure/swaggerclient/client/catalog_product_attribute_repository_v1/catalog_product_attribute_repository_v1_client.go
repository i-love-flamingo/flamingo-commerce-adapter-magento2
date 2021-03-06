// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductAttributeRepositoryV1DeleteByIDDelete Delete Attribute by id
*/
func (a *Client) CatalogProductAttributeRepositoryV1DeleteByIDDelete(params *CatalogProductAttributeRepositoryV1DeleteByIDDeleteParams) (*CatalogProductAttributeRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/attributes/{attributeCode}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeRepositoryV1DeleteByIDDeleteOK), nil

}

/*
CatalogProductAttributeRepositoryV1GetGet Retrieve specific attribute
*/
func (a *Client) CatalogProductAttributeRepositoryV1GetGet(params *CatalogProductAttributeRepositoryV1GetGetParams) (*CatalogProductAttributeRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attributes/{attributeCode}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeRepositoryV1GetGetOK), nil

}

/*
CatalogProductAttributeRepositoryV1GetListGet Retrieve all attributes for entity type
*/
func (a *Client) CatalogProductAttributeRepositoryV1GetListGet(params *CatalogProductAttributeRepositoryV1GetListGetParams) (*CatalogProductAttributeRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attributes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeRepositoryV1GetListGetOK), nil

}

/*
CatalogProductAttributeRepositoryV1SavePost Save attribute data
*/
func (a *Client) CatalogProductAttributeRepositoryV1SavePost(params *CatalogProductAttributeRepositoryV1SavePostParams) (*CatalogProductAttributeRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/products/attributes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeRepositoryV1SavePostOK), nil

}

/*
CatalogProductAttributeRepositoryV1SavePut Save attribute data
*/
func (a *Client) CatalogProductAttributeRepositoryV1SavePut(params *CatalogProductAttributeRepositoryV1SavePutParams) (*CatalogProductAttributeRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductAttributeRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/products/attributes/{attributeCode}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductAttributeRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductAttributeRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
