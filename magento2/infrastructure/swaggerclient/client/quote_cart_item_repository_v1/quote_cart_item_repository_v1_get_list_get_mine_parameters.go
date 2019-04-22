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

// NewQuoteCartItemRepositoryV1GetListGetMineParams creates a new QuoteCartItemRepositoryV1GetListGetMineParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1GetListGetMineParams() *QuoteCartItemRepositoryV1GetListGetMineParams {

	return &QuoteCartItemRepositoryV1GetListGetMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1GetListGetMineParamsWithTimeout creates a new QuoteCartItemRepositoryV1GetListGetMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1GetListGetMineParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1GetListGetMineParams {

	return &QuoteCartItemRepositoryV1GetListGetMineParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1GetListGetMineParamsWithContext creates a new QuoteCartItemRepositoryV1GetListGetMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1GetListGetMineParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1GetListGetMineParams {

	return &QuoteCartItemRepositoryV1GetListGetMineParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1GetListGetMineParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1GetListGetMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1GetListGetMineParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1GetListGetMineParams {

	return &QuoteCartItemRepositoryV1GetListGetMineParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1GetListGetMineParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 get list get mine operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1GetListGetMineParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1GetListGetMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1GetListGetMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1GetListGetMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 get list get mine params
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1GetListGetMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}