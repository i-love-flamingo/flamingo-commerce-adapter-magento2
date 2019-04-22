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

// NewQuoteGuestCartItemRepositoryV1SavePostParams creates a new QuoteGuestCartItemRepositoryV1SavePostParams object
// with the default values initialized.
func NewQuoteGuestCartItemRepositoryV1SavePostParams() *QuoteGuestCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCartItemRepositoryV1SavePostParamsWithTimeout creates a new QuoteGuestCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteGuestCartItemRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *QuoteGuestCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewQuoteGuestCartItemRepositoryV1SavePostParamsWithContext creates a new QuoteGuestCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteGuestCartItemRepositoryV1SavePostParamsWithContext(ctx context.Context) *QuoteGuestCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewQuoteGuestCartItemRepositoryV1SavePostParamsWithHTTPClient creates a new QuoteGuestCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteGuestCartItemRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *QuoteGuestCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteGuestCartItemRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*QuoteGuestCartItemRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the quote guest cart item repository v1 save post operation typically these are written to a http.Request
*/
type QuoteGuestCartItemRepositoryV1SavePostParams struct {

	/*CartID*/
	CartID string
	/*QuoteGuestCartItemRepositoryV1SavePostBody*/
	QuoteGuestCartItemRepositoryV1SavePostBody QuoteGuestCartItemRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *QuoteGuestCartItemRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WithContext(ctx context.Context) *QuoteGuestCartItemRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *QuoteGuestCartItemRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WithCartID(cartID string) *QuoteGuestCartItemRepositoryV1SavePostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithQuoteGuestCartItemRepositoryV1SavePostBody adds the quoteGuestCartItemRepositoryV1SavePostBody to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WithQuoteGuestCartItemRepositoryV1SavePostBody(quoteGuestCartItemRepositoryV1SavePostBody QuoteGuestCartItemRepositoryV1SavePostBody) *QuoteGuestCartItemRepositoryV1SavePostParams {
	o.SetQuoteGuestCartItemRepositoryV1SavePostBody(quoteGuestCartItemRepositoryV1SavePostBody)
	return o
}

// SetQuoteGuestCartItemRepositoryV1SavePostBody adds the quoteGuestCartItemRepositoryV1SavePostBody to the quote guest cart item repository v1 save post params
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) SetQuoteGuestCartItemRepositoryV1SavePostBody(quoteGuestCartItemRepositoryV1SavePostBody QuoteGuestCartItemRepositoryV1SavePostBody) {
	o.QuoteGuestCartItemRepositoryV1SavePostBody = quoteGuestCartItemRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCartItemRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteGuestCartItemRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
