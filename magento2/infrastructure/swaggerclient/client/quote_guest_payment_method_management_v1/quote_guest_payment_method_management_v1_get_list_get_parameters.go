// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_payment_method_management_v1

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

// NewQuoteGuestPaymentMethodManagementV1GetListGetParams creates a new QuoteGuestPaymentMethodManagementV1GetListGetParams object
// with the default values initialized.
func NewQuoteGuestPaymentMethodManagementV1GetListGetParams() *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithTimeout creates a new QuoteGuestPaymentMethodManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithTimeout(timeout time.Duration) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetListGetParams{

		timeout: timeout,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithContext creates a new QuoteGuestPaymentMethodManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithContext(ctx context.Context) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetListGetParams{

		Context: ctx,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithHTTPClient creates a new QuoteGuestPaymentMethodManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestPaymentMethodManagementV1GetListGetParamsWithHTTPClient(client *http.Client) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetListGetParams{
		HTTPClient: client,
	}
}

/*QuoteGuestPaymentMethodManagementV1GetListGetParams contains all the parameters to send to the API endpoint
for the quote guest payment method management v1 get list get operation typically these are written to a http.Request
*/
type QuoteGuestPaymentMethodManagementV1GetListGetParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) WithTimeout(timeout time.Duration) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) WithContext(ctx context.Context) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) WithHTTPClient(client *http.Client) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) WithCartID(cartID string) *QuoteGuestPaymentMethodManagementV1GetListGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest payment method management v1 get list get params
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestPaymentMethodManagementV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
