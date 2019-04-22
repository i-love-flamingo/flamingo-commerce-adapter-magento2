// Code generated by go-swagger; DO NOT EDIT.

package sales_shipment_track_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sales shipment track repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sales shipment track repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SalesShipmentTrackRepositoryV1DeleteByIDDelete Deletes a specified shipment track by ID.
*/
func (a *Client) SalesShipmentTrackRepositoryV1DeleteByIDDelete(params *SalesShipmentTrackRepositoryV1DeleteByIDDeleteParams) (*SalesShipmentTrackRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesShipmentTrackRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesShipmentTrackRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/shipment/track/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesShipmentTrackRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesShipmentTrackRepositoryV1DeleteByIDDeleteOK), nil

}

/*
SalesShipmentTrackRepositoryV1SavePost Performs persist operations for a specified shipment track.
*/
func (a *Client) SalesShipmentTrackRepositoryV1SavePost(params *SalesShipmentTrackRepositoryV1SavePostParams) (*SalesShipmentTrackRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesShipmentTrackRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesShipmentTrackRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/shipment/track",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesShipmentTrackRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesShipmentTrackRepositoryV1SavePostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
