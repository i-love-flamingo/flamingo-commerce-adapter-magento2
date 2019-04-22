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

// NewQuoteCartItemRepositoryV1SavePostMineParams creates a new QuoteCartItemRepositoryV1SavePostMineParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1SavePostMineParams() *QuoteCartItemRepositoryV1SavePostMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePostMineParamsWithTimeout creates a new QuoteCartItemRepositoryV1SavePostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1SavePostMineParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePostMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostMineParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePostMineParamsWithContext creates a new QuoteCartItemRepositoryV1SavePostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1SavePostMineParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePostMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostMineParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1SavePostMineParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1SavePostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1SavePostMineParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePostMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePostMineParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1SavePostMineParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 save post mine operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1SavePostMineParams struct {

	/*QuoteCartItemRepositoryV1SavePostBody*/
	QuoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteCartItemRepositoryV1SavePostBody adds the quoteCartItemRepositoryV1SavePostBody to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) WithQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostMineBody) *QuoteCartItemRepositoryV1SavePostMineParams {
	o.SetQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody)
	return o
}

// SetQuoteCartItemRepositoryV1SavePostBody adds the quoteCartItemRepositoryV1SavePostBody to the quote cart item repository v1 save post mine params
func (o *QuoteCartItemRepositoryV1SavePostMineParams) SetQuoteCartItemRepositoryV1SavePostBody(quoteCartItemRepositoryV1SavePostBody QuoteCartItemRepositoryV1SavePostMineBody) {
	o.QuoteCartItemRepositoryV1SavePostBody = quoteCartItemRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1SavePostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QuoteCartItemRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
