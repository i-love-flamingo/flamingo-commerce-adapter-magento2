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

// NewGiftMessageGuestCartRepositoryV1SavePostParams creates a new GiftMessageGuestCartRepositoryV1SavePostParams object
// with the default values initialized.
func NewGiftMessageGuestCartRepositoryV1SavePostParams() *GiftMessageGuestCartRepositoryV1SavePostParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftMessageGuestCartRepositoryV1SavePostParamsWithTimeout creates a new GiftMessageGuestCartRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftMessageGuestCartRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *GiftMessageGuestCartRepositoryV1SavePostParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewGiftMessageGuestCartRepositoryV1SavePostParamsWithContext creates a new GiftMessageGuestCartRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftMessageGuestCartRepositoryV1SavePostParamsWithContext(ctx context.Context) *GiftMessageGuestCartRepositoryV1SavePostParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewGiftMessageGuestCartRepositoryV1SavePostParamsWithHTTPClient creates a new GiftMessageGuestCartRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftMessageGuestCartRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *GiftMessageGuestCartRepositoryV1SavePostParams {
	var ()
	return &GiftMessageGuestCartRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*GiftMessageGuestCartRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the gift message guest cart repository v1 save post operation typically these are written to a http.Request
*/
type GiftMessageGuestCartRepositoryV1SavePostParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID string
	/*GiftMessageGuestCartRepositoryV1SavePostBody*/
	GiftMessageGuestCartRepositoryV1SavePostBody GiftMessageGuestCartRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *GiftMessageGuestCartRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WithContext(ctx context.Context) *GiftMessageGuestCartRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *GiftMessageGuestCartRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WithCartID(cartID string) *GiftMessageGuestCartRepositoryV1SavePostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithGiftMessageGuestCartRepositoryV1SavePostBody adds the giftMessageGuestCartRepositoryV1SavePostBody to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WithGiftMessageGuestCartRepositoryV1SavePostBody(giftMessageGuestCartRepositoryV1SavePostBody GiftMessageGuestCartRepositoryV1SavePostBody) *GiftMessageGuestCartRepositoryV1SavePostParams {
	o.SetGiftMessageGuestCartRepositoryV1SavePostBody(giftMessageGuestCartRepositoryV1SavePostBody)
	return o
}

// SetGiftMessageGuestCartRepositoryV1SavePostBody adds the giftMessageGuestCartRepositoryV1SavePostBody to the gift message guest cart repository v1 save post params
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) SetGiftMessageGuestCartRepositoryV1SavePostBody(giftMessageGuestCartRepositoryV1SavePostBody GiftMessageGuestCartRepositoryV1SavePostBody) {
	o.GiftMessageGuestCartRepositoryV1SavePostBody = giftMessageGuestCartRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *GiftMessageGuestCartRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.GiftMessageGuestCartRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
