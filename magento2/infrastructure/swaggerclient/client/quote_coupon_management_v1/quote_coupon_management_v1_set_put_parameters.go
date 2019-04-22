// Code generated by go-swagger; DO NOT EDIT.

package quote_coupon_management_v1

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

// NewQuoteCouponManagementV1SetPutParams creates a new QuoteCouponManagementV1SetPutParams object
// with the default values initialized.
func NewQuoteCouponManagementV1SetPutParams() *QuoteCouponManagementV1SetPutParams {
	var ()
	return &QuoteCouponManagementV1SetPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCouponManagementV1SetPutParamsWithTimeout creates a new QuoteCouponManagementV1SetPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCouponManagementV1SetPutParamsWithTimeout(timeout time.Duration) *QuoteCouponManagementV1SetPutParams {
	var ()
	return &QuoteCouponManagementV1SetPutParams{

		timeout: timeout,
	}
}

// NewQuoteCouponManagementV1SetPutParamsWithContext creates a new QuoteCouponManagementV1SetPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCouponManagementV1SetPutParamsWithContext(ctx context.Context) *QuoteCouponManagementV1SetPutParams {
	var ()
	return &QuoteCouponManagementV1SetPutParams{

		Context: ctx,
	}
}

// NewQuoteCouponManagementV1SetPutParamsWithHTTPClient creates a new QuoteCouponManagementV1SetPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCouponManagementV1SetPutParamsWithHTTPClient(client *http.Client) *QuoteCouponManagementV1SetPutParams {
	var ()
	return &QuoteCouponManagementV1SetPutParams{
		HTTPClient: client,
	}
}

/*QuoteCouponManagementV1SetPutParams contains all the parameters to send to the API endpoint
for the quote coupon management v1 set put operation typically these are written to a http.Request
*/
type QuoteCouponManagementV1SetPutParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID int64
	/*CouponCode
	  The coupon code data.

	*/
	CouponCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) WithTimeout(timeout time.Duration) *QuoteCouponManagementV1SetPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) WithContext(ctx context.Context) *QuoteCouponManagementV1SetPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) WithHTTPClient(client *http.Client) *QuoteCouponManagementV1SetPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) WithCartID(cartID int64) *QuoteCouponManagementV1SetPutParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) SetCartID(cartID int64) {
	o.CartID = cartID
}

// WithCouponCode adds the couponCode to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) WithCouponCode(couponCode string) *QuoteCouponManagementV1SetPutParams {
	o.SetCouponCode(couponCode)
	return o
}

// SetCouponCode adds the couponCode to the quote coupon management v1 set put params
func (o *QuoteCouponManagementV1SetPutParams) SetCouponCode(couponCode string) {
	o.CouponCode = couponCode
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCouponManagementV1SetPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", swag.FormatInt64(o.CartID)); err != nil {
		return err
	}

	// path param couponCode
	if err := r.SetPathParam("couponCode", o.CouponCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}