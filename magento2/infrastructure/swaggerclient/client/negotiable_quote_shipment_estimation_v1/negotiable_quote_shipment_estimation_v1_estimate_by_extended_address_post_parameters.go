// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_shipment_estimation_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams creates a new NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the default values initialized.
func NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams() *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	var ()
	return &NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithTimeout creates a new NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithTimeout(timeout time.Duration) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	var ()
	return &NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithContext creates a new NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithContext(ctx context.Context) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	var ()
	return &NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithHTTPClient creates a new NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParamsWithHTTPClient(client *http.Client) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	var ()
	return &NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams contains all the parameters to send to the API endpoint
for the negotiable quote shipment estimation v1 estimate by extended address post operation typically these are written to a http.Request
*/
type NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams struct {

	/*CartID*/
	CartID string
	/*NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody*/
	NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WithTimeout(timeout time.Duration) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WithContext(ctx context.Context) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WithHTTPClient(client *http.Client) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WithCartID(cartID string) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody adds the negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WithNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody) *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody)
	return o
}

// SetNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody adds the negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody to the negotiable quote shipment estimation v1 estimate by extended address post params
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) SetNegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody) {
	o.NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody = negotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.NegotiableQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
