// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_total_management_v1

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

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutParams creates a new QuoteGuestCartTotalManagementV1CollectTotalsPutParams object
// with the default values initialized.
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutParams() *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	var ()
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithTimeout creates a new QuoteGuestCartTotalManagementV1CollectTotalsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithTimeout(timeout time.Duration) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	var ()
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutParams{

		timeout: timeout,
	}
}

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithContext creates a new QuoteGuestCartTotalManagementV1CollectTotalsPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithContext(ctx context.Context) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	var ()
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutParams{

		Context: ctx,
	}
}

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithHTTPClient creates a new QuoteGuestCartTotalManagementV1CollectTotalsPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutParamsWithHTTPClient(client *http.Client) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	var ()
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutParams{
		HTTPClient: client,
	}
}

/*QuoteGuestCartTotalManagementV1CollectTotalsPutParams contains all the parameters to send to the API endpoint
for the quote guest cart total management v1 collect totals put operation typically these are written to a http.Request
*/
type QuoteGuestCartTotalManagementV1CollectTotalsPutParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string
	/*QuoteGuestCartTotalManagementV1CollectTotalsPutBody*/
	QuoteGuestCartTotalManagementV1CollectTotalsPutBody QuoteGuestCartTotalManagementV1CollectTotalsPutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WithTimeout(timeout time.Duration) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WithContext(ctx context.Context) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WithHTTPClient(client *http.Client) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WithCartID(cartID string) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithQuoteGuestCartTotalManagementV1CollectTotalsPutBody adds the quoteGuestCartTotalManagementV1CollectTotalsPutBody to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WithQuoteGuestCartTotalManagementV1CollectTotalsPutBody(quoteGuestCartTotalManagementV1CollectTotalsPutBody QuoteGuestCartTotalManagementV1CollectTotalsPutBody) *QuoteGuestCartTotalManagementV1CollectTotalsPutParams {
	o.SetQuoteGuestCartTotalManagementV1CollectTotalsPutBody(quoteGuestCartTotalManagementV1CollectTotalsPutBody)
	return o
}

// SetQuoteGuestCartTotalManagementV1CollectTotalsPutBody adds the quoteGuestCartTotalManagementV1CollectTotalsPutBody to the quote guest cart total management v1 collect totals put params
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) SetQuoteGuestCartTotalManagementV1CollectTotalsPutBody(quoteGuestCartTotalManagementV1CollectTotalsPutBody QuoteGuestCartTotalManagementV1CollectTotalsPutBody) {
	o.QuoteGuestCartTotalManagementV1CollectTotalsPutBody = quoteGuestCartTotalManagementV1CollectTotalsPutBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteGuestCartTotalManagementV1CollectTotalsPutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}