// Code generated by go-swagger; DO NOT EDIT.

package sales_creditmemo_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sales creditmemo management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sales creditmemo management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SalesCreditmemoManagementV1CancelPut Cancels a specified credit memo.
*/
func (a *Client) SalesCreditmemoManagementV1CancelPut(params *SalesCreditmemoManagementV1CancelPutParams) (*SalesCreditmemoManagementV1CancelPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesCreditmemoManagementV1CancelPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesCreditmemoManagementV1CancelPut",
		Method:             "PUT",
		PathPattern:        "/V1/creditmemo/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesCreditmemoManagementV1CancelPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesCreditmemoManagementV1CancelPutOK), nil

}

/*
SalesCreditmemoManagementV1GetCommentsListGet Lists comments for a specified credit memo.
*/
func (a *Client) SalesCreditmemoManagementV1GetCommentsListGet(params *SalesCreditmemoManagementV1GetCommentsListGetParams) (*SalesCreditmemoManagementV1GetCommentsListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesCreditmemoManagementV1GetCommentsListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesCreditmemoManagementV1GetCommentsListGet",
		Method:             "GET",
		PathPattern:        "/V1/creditmemo/{id}/comments",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesCreditmemoManagementV1GetCommentsListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesCreditmemoManagementV1GetCommentsListGetOK), nil

}

/*
SalesCreditmemoManagementV1NotifyPost Emails a user a specified credit memo.
*/
func (a *Client) SalesCreditmemoManagementV1NotifyPost(params *SalesCreditmemoManagementV1NotifyPostParams) (*SalesCreditmemoManagementV1NotifyPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesCreditmemoManagementV1NotifyPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesCreditmemoManagementV1NotifyPost",
		Method:             "POST",
		PathPattern:        "/V1/creditmemo/{id}/emails",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesCreditmemoManagementV1NotifyPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesCreditmemoManagementV1NotifyPostOK), nil

}

/*
SalesCreditmemoManagementV1RefundPost Prepare creditmemo to refund and save it.
*/
func (a *Client) SalesCreditmemoManagementV1RefundPost(params *SalesCreditmemoManagementV1RefundPostParams) (*SalesCreditmemoManagementV1RefundPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesCreditmemoManagementV1RefundPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesCreditmemoManagementV1RefundPost",
		Method:             "POST",
		PathPattern:        "/V1/creditmemo/refund",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesCreditmemoManagementV1RefundPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesCreditmemoManagementV1RefundPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}