// Code generated by go-swagger; DO NOT EDIT.

package gift_card_account_gift_card_account_management_v1

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

// NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams creates a new GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams object
// with the default values initialized.
func NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams() *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	var ()
	return &GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithTimeout creates a new GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithTimeout(timeout time.Duration) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	var ()
	return &GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams{

		timeout: timeout,
	}
}

// NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithContext creates a new GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithContext(ctx context.Context) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	var ()
	return &GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams{

		Context: ctx,
	}
}

// NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithHTTPClient creates a new GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParamsWithHTTPClient(client *http.Client) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	var ()
	return &GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams{
		HTTPClient: client,
	}
}

/*GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams contains all the parameters to send to the API endpoint
for the gift card account gift card account management v1 save by quote Id post mine operation typically these are written to a http.Request
*/
type GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams struct {

	/*GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody*/
	GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) WithTimeout(timeout time.Duration) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) WithContext(ctx context.Context) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) WithHTTPClient(client *http.Client) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody adds the giftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) WithGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody(giftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineBody) *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams {
	o.SetGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody(giftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody)
	return o
}

// SetGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody adds the giftCardAccountGiftCardAccountManagementV1SaveByQuoteIdPostBody to the gift card account gift card account management v1 save by quote Id post mine params
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) SetGiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody(giftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineBody) {
	o.GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody = giftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.GiftCardAccountGiftCardAccountManagementV1SaveByQuoteIDPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}