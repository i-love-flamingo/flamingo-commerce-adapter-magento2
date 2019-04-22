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

// NewQuoteGuestPaymentMethodManagementV1GetGetParams creates a new QuoteGuestPaymentMethodManagementV1GetGetParams object
// with the default values initialized.
func NewQuoteGuestPaymentMethodManagementV1GetGetParams() *QuoteGuestPaymentMethodManagementV1GetGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithTimeout creates a new QuoteGuestPaymentMethodManagementV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithTimeout(timeout time.Duration) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetGetParams{

		timeout: timeout,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithContext creates a new QuoteGuestPaymentMethodManagementV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithContext(ctx context.Context) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetGetParams{

		Context: ctx,
	}
}

// NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithHTTPClient creates a new QuoteGuestPaymentMethodManagementV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestPaymentMethodManagementV1GetGetParamsWithHTTPClient(client *http.Client) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	var ()
	return &QuoteGuestPaymentMethodManagementV1GetGetParams{
		HTTPClient: client,
	}
}

/*QuoteGuestPaymentMethodManagementV1GetGetParams contains all the parameters to send to the API endpoint
for the quote guest payment method management v1 get get operation typically these are written to a http.Request
*/
type QuoteGuestPaymentMethodManagementV1GetGetParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) WithTimeout(timeout time.Duration) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) WithContext(ctx context.Context) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) WithHTTPClient(client *http.Client) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) WithCartID(cartID string) *QuoteGuestPaymentMethodManagementV1GetGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest payment method management v1 get get params
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestPaymentMethodManagementV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
