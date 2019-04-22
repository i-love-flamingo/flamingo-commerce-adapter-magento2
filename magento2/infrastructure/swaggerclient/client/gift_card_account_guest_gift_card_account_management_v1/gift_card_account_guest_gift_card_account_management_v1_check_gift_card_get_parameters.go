// Code generated by go-swagger; DO NOT EDIT.

package gift_card_account_guest_gift_card_account_management_v1

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

// NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams creates a new GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams object
// with the default values initialized.
func NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams() *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithTimeout creates a new GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithTimeout(timeout time.Duration) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams{

		timeout: timeout,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithContext creates a new GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithContext(ctx context.Context) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams{

		Context: ctx,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithHTTPClient creates a new GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParamsWithHTTPClient(client *http.Client) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams{
		HTTPClient: client,
	}
}

/*GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams contains all the parameters to send to the API endpoint
for the gift card account guest gift card account management v1 check gift card get operation typically these are written to a http.Request
*/
type GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams struct {

	/*CartID*/
	CartID string
	/*GiftCardCode*/
	GiftCardCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WithTimeout(timeout time.Duration) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WithContext(ctx context.Context) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WithHTTPClient(client *http.Client) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WithCartID(cartID string) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithGiftCardCode adds the giftCardCode to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WithGiftCardCode(giftCardCode string) *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams {
	o.SetGiftCardCode(giftCardCode)
	return o
}

// SetGiftCardCode adds the giftCardCode to the gift card account guest gift card account management v1 check gift card get params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) SetGiftCardCode(giftCardCode string) {
	o.GiftCardCode = giftCardCode
}

// WriteToRequest writes these params to a swagger request
func (o *GiftCardAccountGuestGiftCardAccountManagementV1CheckGiftCardGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	// path param giftCardCode
	if err := r.SetPathParam("giftCardCode", o.GiftCardCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
