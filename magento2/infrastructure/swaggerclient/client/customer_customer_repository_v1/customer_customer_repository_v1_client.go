// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new customer customer repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer customer repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomerCustomerRepositoryV1DeleteByIDDelete Delete customer by Customer ID.
*/
func (a *Client) CustomerCustomerRepositoryV1DeleteByIDDelete(params *CustomerCustomerRepositoryV1DeleteByIDDeleteParams) (*CustomerCustomerRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/customers/{customerId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1DeleteByIDDeleteOK), nil

}

/*
CustomerCustomerRepositoryV1GetByIDGet Get customer by Customer ID.
*/
func (a *Client) CustomerCustomerRepositoryV1GetByIDGet(params *CustomerCustomerRepositoryV1GetByIDGetParams) (*CustomerCustomerRepositoryV1GetByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1GetByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1GetByIdGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1GetByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1GetByIDGetOK), nil

}

/*
CustomerCustomerRepositoryV1GetByIDGetMine Get customer by Customer ID.
*/
func (a *Client) CustomerCustomerRepositoryV1GetByIDGetMine(params *CustomerCustomerRepositoryV1GetByIDGetMineParams) (*CustomerCustomerRepositoryV1GetByIDGetMineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1GetByIDGetMineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1GetByIdGetMine",
		Method:             "GET",
		PathPattern:        "/V1/customers/me",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1GetByIDGetMineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1GetByIDGetMineOK), nil

}

/*
CustomerCustomerRepositoryV1GetListGet Retrieve customers which match a specified criteria. This call returns an array of objects, but detailed information about each object’s attributes might not be included. See http://devdocs.magento.com/codelinks/attributes.html#CustomerRepositoryInterface to determine which call to use to get detailed information about all attributes for an object.
*/
func (a *Client) CustomerCustomerRepositoryV1GetListGet(params *CustomerCustomerRepositoryV1GetListGetParams) (*CustomerCustomerRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/search",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1GetListGetOK), nil

}

/*
CustomerCustomerRepositoryV1SavePut Create or update a customer.
*/
func (a *Client) CustomerCustomerRepositoryV1SavePut(params *CustomerCustomerRepositoryV1SavePutParams) (*CustomerCustomerRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/{customerId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1SavePutOK), nil

}

/*
CustomerCustomerRepositoryV1SavePutMe Create or update a customer.
*/
func (a *Client) CustomerCustomerRepositoryV1SavePutMe(params *CustomerCustomerRepositoryV1SavePutMeParams) (*CustomerCustomerRepositoryV1SavePutMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCustomerRepositoryV1SavePutMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerCustomerRepositoryV1SavePutMe",
		Method:             "PUT",
		PathPattern:        "/V1/customers/me",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerCustomerRepositoryV1SavePutMeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerCustomerRepositoryV1SavePutMeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}