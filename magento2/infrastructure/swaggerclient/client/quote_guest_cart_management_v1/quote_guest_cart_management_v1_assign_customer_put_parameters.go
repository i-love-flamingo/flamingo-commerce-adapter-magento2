// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_management_v1

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

// NewQuoteGuestCartManagementV1AssignCustomerPutParams creates a new QuoteGuestCartManagementV1AssignCustomerPutParams object
// with the default values initialized.
func NewQuoteGuestCartManagementV1AssignCustomerPutParams() *QuoteGuestCartManagementV1AssignCustomerPutParams {
	var ()
	return &QuoteGuestCartManagementV1AssignCustomerPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithTimeout creates a new QuoteGuestCartManagementV1AssignCustomerPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithTimeout(timeout time.Duration) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	var ()
	return &QuoteGuestCartManagementV1AssignCustomerPutParams{

		timeout: timeout,
	}
}

// NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithContext creates a new QuoteGuestCartManagementV1AssignCustomerPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithContext(ctx context.Context) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	var ()
	return &QuoteGuestCartManagementV1AssignCustomerPutParams{

		Context: ctx,
	}
}

// NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithHTTPClient creates a new QuoteGuestCartManagementV1AssignCustomerPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestCartManagementV1AssignCustomerPutParamsWithHTTPClient(client *http.Client) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	var ()
	return &QuoteGuestCartManagementV1AssignCustomerPutParams{
		HTTPClient: client,
	}
}

/*QuoteGuestCartManagementV1AssignCustomerPutParams contains all the parameters to send to the API endpoint
for the quote guest cart management v1 assign customer put operation typically these are written to a http.Request
*/
type QuoteGuestCartManagementV1AssignCustomerPutParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string
	/*QuoteGuestCartManagementV1AssignCustomerPutBody*/
	QuoteGuestCartManagementV1AssignCustomerPutBody QuoteGuestCartManagementV1AssignCustomerPutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WithTimeout(timeout time.Duration) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WithContext(ctx context.Context) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WithHTTPClient(client *http.Client) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WithCartID(cartID string) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithQuoteGuestCartManagementV1AssignCustomerPutBody adds the quoteGuestCartManagementV1AssignCustomerPutBody to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WithQuoteGuestCartManagementV1AssignCustomerPutBody(quoteGuestCartManagementV1AssignCustomerPutBody QuoteGuestCartManagementV1AssignCustomerPutBody) *QuoteGuestCartManagementV1AssignCustomerPutParams {
	o.SetQuoteGuestCartManagementV1AssignCustomerPutBody(quoteGuestCartManagementV1AssignCustomerPutBody)
	return o
}

// SetQuoteGuestCartManagementV1AssignCustomerPutBody adds the quoteGuestCartManagementV1AssignCustomerPutBody to the quote guest cart management v1 assign customer put params
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) SetQuoteGuestCartManagementV1AssignCustomerPutBody(quoteGuestCartManagementV1AssignCustomerPutBody QuoteGuestCartManagementV1AssignCustomerPutBody) {
	o.QuoteGuestCartManagementV1AssignCustomerPutBody = quoteGuestCartManagementV1AssignCustomerPutBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCartManagementV1AssignCustomerPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteGuestCartManagementV1AssignCustomerPutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
