// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_billing_address_management_v1

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

// NewQuoteGuestBillingAddressManagementV1AssignPostParams creates a new QuoteGuestBillingAddressManagementV1AssignPostParams object
// with the default values initialized.
func NewQuoteGuestBillingAddressManagementV1AssignPostParams() *QuoteGuestBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteGuestBillingAddressManagementV1AssignPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithTimeout creates a new QuoteGuestBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithTimeout(timeout time.Duration) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteGuestBillingAddressManagementV1AssignPostParams{

		timeout: timeout,
	}
}

// NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithContext creates a new QuoteGuestBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithContext(ctx context.Context) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteGuestBillingAddressManagementV1AssignPostParams{

		Context: ctx,
	}
}

// NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithHTTPClient creates a new QuoteGuestBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestBillingAddressManagementV1AssignPostParamsWithHTTPClient(client *http.Client) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteGuestBillingAddressManagementV1AssignPostParams{
		HTTPClient: client,
	}
}

/*QuoteGuestBillingAddressManagementV1AssignPostParams contains all the parameters to send to the API endpoint
for the quote guest billing address management v1 assign post operation typically these are written to a http.Request
*/
type QuoteGuestBillingAddressManagementV1AssignPostParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string
	/*QuoteGuestBillingAddressManagementV1AssignPostBody*/
	QuoteGuestBillingAddressManagementV1AssignPostBody QuoteGuestBillingAddressManagementV1AssignPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WithTimeout(timeout time.Duration) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WithContext(ctx context.Context) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WithHTTPClient(client *http.Client) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WithCartID(cartID string) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithQuoteGuestBillingAddressManagementV1AssignPostBody adds the quoteGuestBillingAddressManagementV1AssignPostBody to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WithQuoteGuestBillingAddressManagementV1AssignPostBody(quoteGuestBillingAddressManagementV1AssignPostBody QuoteGuestBillingAddressManagementV1AssignPostBody) *QuoteGuestBillingAddressManagementV1AssignPostParams {
	o.SetQuoteGuestBillingAddressManagementV1AssignPostBody(quoteGuestBillingAddressManagementV1AssignPostBody)
	return o
}

// SetQuoteGuestBillingAddressManagementV1AssignPostBody adds the quoteGuestBillingAddressManagementV1AssignPostBody to the quote guest billing address management v1 assign post params
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) SetQuoteGuestBillingAddressManagementV1AssignPostBody(quoteGuestBillingAddressManagementV1AssignPostBody QuoteGuestBillingAddressManagementV1AssignPostBody) {
	o.QuoteGuestBillingAddressManagementV1AssignPostBody = quoteGuestBillingAddressManagementV1AssignPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestBillingAddressManagementV1AssignPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteGuestBillingAddressManagementV1AssignPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
