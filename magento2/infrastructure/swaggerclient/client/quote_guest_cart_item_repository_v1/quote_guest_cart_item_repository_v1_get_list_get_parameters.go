// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_item_repository_v1

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

// NewQuoteGuestCartItemRepositoryV1GetListGetParams creates a new QuoteGuestCartItemRepositoryV1GetListGetParams object
// with the default values initialized.
func NewQuoteGuestCartItemRepositoryV1GetListGetParams() *QuoteGuestCartItemRepositoryV1GetListGetParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1GetListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithTimeout creates a new QuoteGuestCartItemRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithTimeout(timeout time.Duration) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1GetListGetParams{

		timeout: timeout,
	}
}

// NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithContext creates a new QuoteGuestCartItemRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithContext(ctx context.Context) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1GetListGetParams{

		Context: ctx,
	}
}

// NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithHTTPClient creates a new QuoteGuestCartItemRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestCartItemRepositoryV1GetListGetParamsWithHTTPClient(client *http.Client) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1GetListGetParams{
		HTTPClient: client,
	}
}

/*QuoteGuestCartItemRepositoryV1GetListGetParams contains all the parameters to send to the API endpoint
for the quote guest cart item repository v1 get list get operation typically these are written to a http.Request
*/
type QuoteGuestCartItemRepositoryV1GetListGetParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) WithTimeout(timeout time.Duration) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) WithContext(ctx context.Context) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) WithHTTPClient(client *http.Client) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) WithCartID(cartID string) *QuoteGuestCartItemRepositoryV1GetListGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest cart item repository v1 get list get params
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCartItemRepositoryV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
