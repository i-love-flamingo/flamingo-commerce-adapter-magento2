// Code generated by go-swagger; DO NOT EDIT.

package requisition_list_requisition_list_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new requisition list requisition list repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for requisition list requisition list repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RequisitionListRequisitionListRepositoryV1SavePost Save Requisition List
*/
func (a *Client) RequisitionListRequisitionListRepositoryV1SavePost(params *RequisitionListRequisitionListRepositoryV1SavePostParams) (*RequisitionListRequisitionListRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequisitionListRequisitionListRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requisitionListRequisitionListRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/requisition_lists",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequisitionListRequisitionListRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequisitionListRequisitionListRepositoryV1SavePostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
