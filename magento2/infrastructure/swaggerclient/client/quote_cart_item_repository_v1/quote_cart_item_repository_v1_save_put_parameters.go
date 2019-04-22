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

// NewQuoteCartItemRepositoryV1SavePutParams creates a new QuoteCartItemRepositoryV1SavePutParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1SavePutParams() *QuoteCartItemRepositoryV1SavePutParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePutParamsWithTimeout creates a new QuoteCartItemRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePutParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1SavePutParamsWithContext creates a new QuoteCartItemRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1SavePutParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePutParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1SavePutParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePutParams {
	var ()
	return &QuoteCartItemRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 save put operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1SavePutParams struct {

	/*CartID*/
	CartID string
	/*ItemID*/
	ItemID string
	/*QuoteCartItemRepositoryV1SavePutBody*/
	QuoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithCartID(cartID string) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithItemID adds the itemID to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithItemID(itemID string) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithQuoteCartItemRepositoryV1SavePutBody adds the quoteCartItemRepositoryV1SavePutBody to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) WithQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutBody) *QuoteCartItemRepositoryV1SavePutParams {
	o.SetQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody)
	return o
}

// SetQuoteCartItemRepositoryV1SavePutBody adds the quoteCartItemRepositoryV1SavePutBody to the quote cart item repository v1 save put params
func (o *QuoteCartItemRepositoryV1SavePutParams) SetQuoteCartItemRepositoryV1SavePutBody(quoteCartItemRepositoryV1SavePutBody QuoteCartItemRepositoryV1SavePutBody) {
	o.QuoteCartItemRepositoryV1SavePutBody = quoteCartItemRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

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
