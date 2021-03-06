// Code generated by go-swagger; DO NOT EDIT.

package quote_shipping_method_management_v1

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

// NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams creates a new QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams object
// with the default values initialized.
func NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams() *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	var ()
	return &QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithTimeout creates a new QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithTimeout(timeout time.Duration) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	var ()
	return &QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams{

		timeout: timeout,
	}
}

// NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithContext creates a new QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithContext(ctx context.Context) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	var ()
	return &QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams{

		Context: ctx,
	}
}

// NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithHTTPClient creates a new QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteShippingMethodManagementV1EstimateByAddressIDPostMineParamsWithHTTPClient(client *http.Client) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	var ()
	return &QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams{
		HTTPClient: client,
	}
}

/*QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams contains all the parameters to send to the API endpoint
for the quote shipping method management v1 estimate by address Id post mine operation typically these are written to a http.Request
*/
type QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams struct {

	/*QuoteShippingMethodManagementV1EstimateByAddressIDPostBody*/
	QuoteShippingMethodManagementV1EstimateByAddressIDPostBody QuoteShippingMethodManagementV1EstimateByAddressIDPostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) WithTimeout(timeout time.Duration) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) WithContext(ctx context.Context) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) WithHTTPClient(client *http.Client) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteShippingMethodManagementV1EstimateByAddressIDPostBody adds the quoteShippingMethodManagementV1EstimateByAddressIDPostBody to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) WithQuoteShippingMethodManagementV1EstimateByAddressIDPostBody(quoteShippingMethodManagementV1EstimateByAddressIDPostBody QuoteShippingMethodManagementV1EstimateByAddressIDPostMineBody) *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams {
	o.SetQuoteShippingMethodManagementV1EstimateByAddressIDPostBody(quoteShippingMethodManagementV1EstimateByAddressIDPostBody)
	return o
}

// SetQuoteShippingMethodManagementV1EstimateByAddressIDPostBody adds the quoteShippingMethodManagementV1EstimateByAddressIdPostBody to the quote shipping method management v1 estimate by address Id post mine params
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) SetQuoteShippingMethodManagementV1EstimateByAddressIDPostBody(quoteShippingMethodManagementV1EstimateByAddressIDPostBody QuoteShippingMethodManagementV1EstimateByAddressIDPostMineBody) {
	o.QuoteShippingMethodManagementV1EstimateByAddressIDPostBody = quoteShippingMethodManagementV1EstimateByAddressIDPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteShippingMethodManagementV1EstimateByAddressIDPostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QuoteShippingMethodManagementV1EstimateByAddressIDPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
