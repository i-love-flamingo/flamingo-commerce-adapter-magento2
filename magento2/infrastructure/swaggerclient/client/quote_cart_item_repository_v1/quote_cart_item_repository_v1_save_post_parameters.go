// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_item_repository_v1

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

// NewQuoteCartItemRepositoryV1SavePostParams creates a new QuoteCartItemRepositoryV1SavePostParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1SavePostParams() *QuoteCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePostParamsWithTimeout creates a new QuoteCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePostParamsWithContext creates a new QuoteCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1SavePostParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1SavePostParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePostParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 save post operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1SavePostParams struct {

	/*QuoteCartItemRepositoryV1SavePostBody*/
	QuoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostBody
	/*QuoteID*/
	QuoteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteCartItemRepositoryV1SavePostBody adds the quoteCartItemRepositoryV1SavePostBody to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) WithQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostBody) *QuoteCartItemRepositoryV1SavePostParams {
	o.SetQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody)
	return o
}

// SetQuoteCartItemRepositoryV1SavePostBody adds the quoteCartItemRepositoryV1SavePostBody to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) SetQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostBody) {
	o.QuoteCartItemRepositoryV1SavePostBody = quoteCartItemRepositoryV1SavePostBody
}

// WithQuoteID adds the quoteID to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) WithQuoteID(quoteID string) *QuoteCartItemRepositoryV1SavePostParams {
	o.SetQuoteID(quoteID)
	return o
}

// SetQuoteID adds the quoteId to the quote cart item repository v1 save post params
func (o *QuoteCartItemRepositoryV1SavePostParams) SetQuoteID(quoteID string) {
	o.QuoteID = quoteID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QuoteCartItemRepositoryV1SavePostBody); err != nil {
		return err
	}

	// path param quoteId
	if err := r.SetPathParam("quoteId", o.QuoteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}