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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteParams creates a new QuoteCartItemRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized.
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteParams() *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithTimeout creates a new QuoteCartItemRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteParams{

		timeout: timeout,
	}
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithContext creates a new QuoteCartItemRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithContext(ctx context.Context) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteParams{

		Context: ctx,
	}
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithHTTPClient creates a new QuoteCartItemRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteParamsWithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteParams{
		HTTPClient: client,
	}
}

/*QuoteCartItemRepositoryV1DeleteByIDDeleteParams contains all the parameters to send to the API endpoint
for the quote cart item repository v1 delete by Id delete operation typically these are written to a http.Request
*/
type QuoteCartItemRepositoryV1DeleteByIDDeleteParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID int64
	/*ItemID
	  The item ID of the item to be removed.

	*/
	ItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WithTimeout(timeout time.Duration) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WithContext(ctx context.Context) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WithHTTPClient(client *http.Client) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WithCartID(cartID int64) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) SetCartID(cartID int64) {
	o.CartID = cartID
}

// WithItemID adds the itemID to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WithItemID(itemID int64) *QuoteCartItemRepositoryV1DeleteByIDDeleteParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the quote cart item repository v1 delete by Id delete params
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", swag.FormatInt64(o.CartID)); err != nil {
		return err
	}

	// path param itemId
	if err := r.SetPathParam("itemId", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
