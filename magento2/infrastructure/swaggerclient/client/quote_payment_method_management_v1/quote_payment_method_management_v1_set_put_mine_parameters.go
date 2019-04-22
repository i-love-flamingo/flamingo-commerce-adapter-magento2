// Code generated by go-swagger; DO NOT EDIT.

package quote_payment_method_management_v1

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

// NewQuotePaymentMethodManagementV1SetPutMineParams creates a new QuotePaymentMethodManagementV1SetPutMineParams object
// with the default values initialized.
func NewQuotePaymentMethodManagementV1SetPutMineParams() *QuotePaymentMethodManagementV1SetPutMineParams {
	var ()
	return &QuotePaymentMethodManagementV1SetPutMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuotePaymentMethodManagementV1SetPutMineParamsWithTimeout creates a new QuotePaymentMethodManagementV1SetPutMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuotePaymentMethodManagementV1SetPutMineParamsWithTimeout(timeout time.Duration) *QuotePaymentMethodManagementV1SetPutMineParams {
	var ()
	return &QuotePaymentMethodManagementV1SetPutMineParams{

		timeout: timeout,
	}
}

// NewQuotePaymentMethodManagementV1SetPutMineParamsWithContext creates a new QuotePaymentMethodManagementV1SetPutMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuotePaymentMethodManagementV1SetPutMineParamsWithContext(ctx context.Context) *QuotePaymentMethodManagementV1SetPutMineParams {
	var ()
	return &QuotePaymentMethodManagementV1SetPutMineParams{

		Context: ctx,
	}
}

// NewQuotePaymentMethodManagementV1SetPutMineParamsWithHTTPClient creates a new QuotePaymentMethodManagementV1SetPutMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuotePaymentMethodManagementV1SetPutMineParamsWithHTTPClient(client *http.Client) *QuotePaymentMethodManagementV1SetPutMineParams {
	var ()
	return &QuotePaymentMethodManagementV1SetPutMineParams{
		HTTPClient: client,
	}
}

/*QuotePaymentMethodManagementV1SetPutMineParams contains all the parameters to send to the API endpoint
for the quote payment method management v1 set put mine operation typically these are written to a http.Request
*/
type QuotePaymentMethodManagementV1SetPutMineParams struct {

	/*QuotePaymentMethodManagementV1SetPutBody*/
	QuotePaymentMethodManagementV1SetPutBody QuotePaymentMethodManagementV1SetPutMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) WithTimeout(timeout time.Duration) *QuotePaymentMethodManagementV1SetPutMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) WithContext(ctx context.Context) *QuotePaymentMethodManagementV1SetPutMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) WithHTTPClient(client *http.Client) *QuotePaymentMethodManagementV1SetPutMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuotePaymentMethodManagementV1SetPutBody adds the quotePaymentMethodManagementV1SetPutBody to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) WithQuotePaymentMethodManagementV1SetPutBody(quotePaymentMethodManagementV1SetPutBody QuotePaymentMethodManagementV1SetPutMineBody) *QuotePaymentMethodManagementV1SetPutMineParams {
	o.SetQuotePaymentMethodManagementV1SetPutBody(quotePaymentMethodManagementV1SetPutBody)
	return o
}

// SetQuotePaymentMethodManagementV1SetPutBody adds the quotePaymentMethodManagementV1SetPutBody to the quote payment method management v1 set put mine params
func (o *QuotePaymentMethodManagementV1SetPutMineParams) SetQuotePaymentMethodManagementV1SetPutBody(quotePaymentMethodManagementV1SetPutBody QuotePaymentMethodManagementV1SetPutMineBody) {
	o.QuotePaymentMethodManagementV1SetPutBody = quotePaymentMethodManagementV1SetPutBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuotePaymentMethodManagementV1SetPutMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QuotePaymentMethodManagementV1SetPutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
