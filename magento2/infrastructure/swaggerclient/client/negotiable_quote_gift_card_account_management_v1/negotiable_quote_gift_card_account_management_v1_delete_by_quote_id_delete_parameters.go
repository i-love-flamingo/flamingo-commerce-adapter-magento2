// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_gift_card_account_management_v1

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

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams creates a new NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized.
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams() *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithTimeout creates a new NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithTimeout(timeout time.Duration) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithContext creates a new NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithContext(ctx context.Context) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithHTTPClient creates a new NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParamsWithHTTPClient(client *http.Client) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	var ()
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams contains all the parameters to send to the API endpoint
for the negotiable quote gift card account management v1 delete by quote Id delete operation typically these are written to a http.Request
*/
type NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams struct {

	/*CartID*/
	CartID int64
	/*GiftCardCode*/
	GiftCardCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithTimeout(timeout time.Duration) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithContext(ctx context.Context) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithHTTPClient(client *http.Client) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithCartID(cartID int64) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetCartID(cartID int64) {
	o.CartID = cartID
}

// WithGiftCardCode adds the giftCardCode to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WithGiftCardCode(giftCardCode string) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams {
	o.SetGiftCardCode(giftCardCode)
	return o
}

// SetGiftCardCode adds the giftCardCode to the negotiable quote gift card account management v1 delete by quote Id delete params
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) SetGiftCardCode(giftCardCode string) {
	o.GiftCardCode = giftCardCode
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", swag.FormatInt64(o.CartID)); err != nil {
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
