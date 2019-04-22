// Code generated by go-swagger; DO NOT EDIT.

package sales_shipment_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sales shipment repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sales shipment repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SalesShipmentRepositoryV1GetGet Loads a specified shipment.
*/
func (a *Client) SalesShipmentRepositoryV1GetGet(params *SalesShipmentRepositoryV1GetGetParams) (*SalesShipmentRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesShipmentRepositoryV1GetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesShipmentRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/shipment/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesShipmentRepositoryV1GetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesShipmentRepositoryV1GetGetOK), nil

}

/*
SalesShipmentRepositoryV1GetListGet Lists shipments that match specified search criteria. This call returns an array of objects, but detailed information about each object’s attributes might not be included. See http://devdocs.magento.com/codelinks/attributes.html#ShipmentRepositoryInterface to determine which call to use to get detailed information about all attributes for an object.
*/
func (a *Client) SalesShipmentRepositoryV1GetListGet(params *SalesShipmentRepositoryV1GetListGetParams) (*SalesShipmentRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesShipmentRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesShipmentRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/shipments",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesShipmentRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesShipmentRepositoryV1GetListGetOK), nil

}

/*
SalesShipmentRepositoryV1SavePost Performs persist operations for a specified shipment.
*/
func (a *Client) SalesShipmentRepositoryV1SavePost(params *SalesShipmentRepositoryV1SavePostParams) (*SalesShipmentRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalesShipmentRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "salesShipmentRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/shipment/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SalesShipmentRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SalesShipmentRepositoryV1SavePostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
