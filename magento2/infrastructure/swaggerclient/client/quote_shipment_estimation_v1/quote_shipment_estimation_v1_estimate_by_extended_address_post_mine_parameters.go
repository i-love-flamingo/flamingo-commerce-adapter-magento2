// Code generated by go-swagger; DO NOT EDIT.

package quote_shipment_estimation_v1

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

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams creates a new QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams object
// with the default values initialized.
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams() *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	var ()
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithTimeout creates a new QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithTimeout(timeout time.Duration) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	var ()
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams{

		timeout: timeout,
	}
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithContext creates a new QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithContext(ctx context.Context) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	var ()
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams{

		Context: ctx,
	}
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithHTTPClient creates a new QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParamsWithHTTPClient(client *http.Client) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	var ()
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams{
		HTTPClient: client,
	}
}

/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams contains all the parameters to send to the API endpoint
for the quote shipment estimation v1 estimate by extended address post mine operation typically these are written to a http.Request
*/
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams struct {

	/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostBody*/
	QuoteShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) WithTimeout(timeout time.Duration) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) WithContext(ctx context.Context) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) WithHTTPClient(client *http.Client) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody adds the quoteShipmentEstimationV1EstimateByExtendedAddressPostBody to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) WithQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams {
	o.SetQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteShipmentEstimationV1EstimateByExtendedAddressPostBody)
	return o
}

// SetQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody adds the quoteShipmentEstimationV1EstimateByExtendedAddressPostBody to the quote shipment estimation v1 estimate by extended address post mine params
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) SetQuoteShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) {
	o.QuoteShipmentEstimationV1EstimateByExtendedAddressPostBody = quoteShipmentEstimationV1EstimateByExtendedAddressPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QuoteShipmentEstimationV1EstimateByExtendedAddressPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
