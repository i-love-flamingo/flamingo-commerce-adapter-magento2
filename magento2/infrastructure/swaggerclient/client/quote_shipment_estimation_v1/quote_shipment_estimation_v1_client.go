// Code generated by go-swagger; DO NOT EDIT.

package quote_shipment_estimation_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quote shipment estimation v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote shipment estimation v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QuoteShipmentEstimationV1EstimateByExtendedAddressPost Estimate shipping by address and return list of available shipping methods
*/
func (a *Client) QuoteShipmentEstimationV1EstimateByExtendedAddressPost(params *QuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) (*QuoteShipmentEstimationV1EstimateByExtendedAddressPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteShipmentEstimationV1EstimateByExtendedAddressPost",
		Method:             "POST",
		PathPattern:        "/V1/carts/{cartId}/estimate-shipping-methods",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteShipmentEstimationV1EstimateByExtendedAddressPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteShipmentEstimationV1EstimateByExtendedAddressPostOK), nil

}

/*
QuoteShipmentEstimationV1EstimateByExtendedAddressPostMine Estimate shipping by address and return list of available shipping methods
*/
func (a *Client) QuoteShipmentEstimationV1EstimateByExtendedAddressPostMine(params *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) (*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteShipmentEstimationV1EstimateByExtendedAddressPostMine",
		Method:             "POST",
		PathPattern:        "/V1/carts/mine/estimate-shipping-methods",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}