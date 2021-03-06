// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog product repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CatalogProductRepositoryV1DeleteByIDDelete catalog product repository v1 delete by Id delete API
*/
func (a *Client) CatalogProductRepositoryV1DeleteByIDDelete(params *CatalogProductRepositoryV1DeleteByIDDeleteParams) (*CatalogProductRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/products/{sku}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductRepositoryV1DeleteByIDDeleteOK), nil

}

/*
CatalogProductRepositoryV1GetGet Get info about product by product SKU
*/
func (a *Client) CatalogProductRepositoryV1GetGet(params *CatalogProductRepositoryV1GetGetParams) (*CatalogProductRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductRepositoryV1GetGetOK), nil

}

/*
CatalogProductRepositoryV1GetListGet Get product list
*/
func (a *Client) CatalogProductRepositoryV1GetListGet(params *CatalogProductRepositoryV1GetListGetParams) (*CatalogProductRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductRepositoryV1GetListGetOK), nil

}

/*
CatalogProductRepositoryV1SavePost Create product
*/
func (a *Client) CatalogProductRepositoryV1SavePost(params *CatalogProductRepositoryV1SavePostParams) (*CatalogProductRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/products",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductRepositoryV1SavePostOK), nil

}

/*
CatalogProductRepositoryV1SavePut Create product
*/
func (a *Client) CatalogProductRepositoryV1SavePut(params *CatalogProductRepositoryV1SavePutParams) (*CatalogProductRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/products/{sku}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CatalogProductRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
