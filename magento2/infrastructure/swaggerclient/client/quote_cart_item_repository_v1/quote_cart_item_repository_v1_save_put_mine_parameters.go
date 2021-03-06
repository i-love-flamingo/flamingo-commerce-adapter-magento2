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

// NewQuoteCartItemRepositoryV1SavePutMineParams creates a new QuoteCartItemRepositoryV1SavePutMineParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1SavePutMineParams() *QuoteCartItemRepositoryV1SavePutMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePutMineParamsWithTimeout creates a new QuoteCartItemRepositoryV1SavePutMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1SavePutMineParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePutMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutMineParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePutMineParamsWithContext creates a new QuoteCartItemRepositoryV1SavePutMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1SavePutMineParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePutMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutMineParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1SavePutMineParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1SavePutMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1SavePutMineParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePutMineParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutMineParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1SavePutMineParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 save put mine operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1SavePutMineParams struct {

	/*ItemID*/
	ItemID string
	/*QuoteCartItemRepositoryV1SavePutBody*/
	QuoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePutMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePutMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePutMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WithItemID(itemID string) *QuoteCartItemRepositoryV1SavePutMineParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithQuoteCartItemRepositoryV1SavePutBody adds the quoteCartItemRepositoryV1SavePutBody to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WithQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutMineBody) *QuoteCartItemRepositoryV1SavePutMineParams {
	o.SetQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody)
	return o
}

// SetQuoteCartItemRepositoryV1SavePutBody adds the quoteCartItemRepositoryV1SavePutBody to the quote cart item repository v1 save put mine params
func (o *QuoteCartItemRepositoryV1SavePutMineParams) SetQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutMineBody) {
	o.QuoteCartItemRepositoryV1SavePutBody = quoteCartItemRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1SavePutMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteCartItemRepositoryV1SavePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
