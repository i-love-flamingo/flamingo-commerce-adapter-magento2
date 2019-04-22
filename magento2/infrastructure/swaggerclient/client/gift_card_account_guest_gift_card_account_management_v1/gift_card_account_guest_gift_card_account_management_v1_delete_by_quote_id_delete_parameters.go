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

// NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams creates a new GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized.
func NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams() *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithTimeout creates a new GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithTimeout(timeout time.Duration) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		timeout: timeout,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithContext creates a new GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithContext(ctx context.Context) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		Context: ctx,
	}
}

// NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithHTTPClient creates a new GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithHTTPClient(client *http.Client) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{
		HTTPClient: client,
	}
}

/*GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams contains all the parameters to send to the API endpoint
for the gift card account guest gift card account management v1 delete by quote Id delete operation typically these are written to a http.Request
*/
type GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams struct {

	/*CartID*/
	CartID string
	/*GiftCardCode*/
	GiftCardCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithTimeout(timeout time.Duration) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithContext(ctx context.Context) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithHTTPClient(client *http.Client) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithCartID(cartID string) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithGiftCardCode adds the giftCardCode to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithGiftCardCode(giftCardCode string) *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetGiftCardCode(giftCardCode)
	return o
}

// SetGiftCardCode adds the giftCardCode to the gift card account guest gift card account management v1 delete by quote Id delete params
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetGiftCardCode(giftCardCode string) {
	o.GiftCardCode = giftCardCode
}

// WriteToRequest writes these params to a swagger request
func (o *GiftCardAccountGuestGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
