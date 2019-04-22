// Code generated by go-swagger; DO NOT EDIT.

package customer_balance_balance_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new customer balance balance management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer balance balance management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomerBalanceBalanceManagementV1ApplyPost Apply store credit
*/
func (a *Client) CustomerBalanceBalanceManagementV1ApplyPost(params *CustomerBalanceBalanceManagementV1ApplyPostParams) (*CustomerBalanceBalanceManagementV1ApplyPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerBalanceBalanceManagementV1ApplyPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerBalanceBalanceManagementV1ApplyPost",
		Method:             "POST",
		PathPattern:        "/V1/carts/mine/balance/apply",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerBalanceBalanceManagementV1ApplyPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerBalanceBalanceManagementV1ApplyPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
