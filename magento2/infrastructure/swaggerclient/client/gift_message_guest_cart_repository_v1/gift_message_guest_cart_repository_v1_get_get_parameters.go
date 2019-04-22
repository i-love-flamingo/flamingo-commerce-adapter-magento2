// Code generated by go-swagger; DO NOT EDIT.

package gift_message_guest_cart_repository_v1

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

// NewGiftMessageGuestCartRepositoryV1GetGetParams creates a new GiftMessageGuestCartRepositoryV1GetGetParams object
// with the default values initialized.
func NewGiftMessageGuestCartRepositoryV1GetGetParams() *GiftMessageGuestCartRepositoryV1GetGetParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftMessageGuestCartRepositoryV1GetGetParamsWithTimeout creates a new GiftMessageGuestCartRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftMessageGuestCartRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *GiftMessageGuestCartRepositoryV1GetGetParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewGiftMessageGuestCartRepositoryV1GetGetParamsWithContext creates a new GiftMessageGuestCartRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftMessageGuestCartRepositoryV1GetGetParamsWithContext(ctx context.Context) *GiftMessageGuestCartRepositoryV1GetGetParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewGiftMessageGuestCartRepositoryV1GetGetParamsWithHTTPClient creates a new GiftMessageGuestCartRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftMessageGuestCartRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *GiftMessageGuestCartRepositoryV1GetGetParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*GiftMessageGuestCartRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the gift message guest cart repository v1 get get operation typically these are written to a http.Request
*/
type GiftMessageGuestCartRepositoryV1GetGetParams struct {

	/*CartID
	  The shopping cart ID.

	*/
	CartID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *GiftMessageGuestCartRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) WithContext(ctx context.Context) *GiftMessageGuestCartRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *GiftMessageGuestCartRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) WithCartID(cartID string) *GiftMessageGuestCartRepositoryV1GetGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift message guest cart repository v1 get get params
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WriteToRequest writes these params to a swagger request
func (o *GiftMessageGuestCartRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
