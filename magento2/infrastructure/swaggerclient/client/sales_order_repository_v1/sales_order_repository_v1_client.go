// Code generated by go-swagger; DO NOT EDIT.

package sales_order_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sales order repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sales order repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SalesOrderRepositoryV1GetGet Loads a specified order.
*/
func (a *Client) SalesOrderRepositoryV1GetGet(params *SalesOrderRepositoryV1GetGetParams) (*SalesOrderRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesOrderRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesOrderRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/orders/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesOrderRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesOrderRepositoryV1GetGetOK), nil

}

/*
SalesOrderRepositoryV1GetListGet Lists orders that match specified search criteria. This call returns an array of objects, but detailed information about each object’s attributes might not be included. See http://devdocs.magento.com/codelinks/attributes.html#OrderRepositoryInterface to determine which call to use to get detailed information about all attributes for an object.
*/
func (a *Client) SalesOrderRepositoryV1GetListGet(params *SalesOrderRepositoryV1GetListGetParams) (*SalesOrderRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesOrderRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesOrderRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/orders",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesOrderRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesOrderRepositoryV1GetListGetOK), nil

}

/*
SalesOrderRepositoryV1SavePost Performs persist operations for a specified order.
*/
func (a *Client) SalesOrderRepositoryV1SavePost(params *SalesOrderRepositoryV1SavePostParams) (*SalesOrderRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesOrderRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesOrderRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/orders/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesOrderRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesOrderRepositoryV1SavePostOK), nil

}

/*
SalesOrderRepositoryV1SavePut Performs persist operations for a specified order.
*/
func (a *Client) SalesOrderRepositoryV1SavePut(params *SalesOrderRepositoryV1SavePutParams) (*SalesOrderRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesOrderRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesOrderRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/orders/create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesOrderRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesOrderRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}