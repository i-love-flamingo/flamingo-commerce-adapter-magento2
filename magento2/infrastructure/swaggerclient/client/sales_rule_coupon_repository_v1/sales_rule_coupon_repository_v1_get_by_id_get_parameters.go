// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_coupon_repository_v1

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

// NewSalesRuleCouponRepositoryV1GetByIDGetParams creates a new SalesRuleCouponRepositoryV1GetByIDGetParams object
// with the default values initialized.
func NewSalesRuleCouponRepositoryV1GetByIDGetParams() *SalesRuleCouponRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleCouponRepositoryV1GetByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithTimeout creates a new SalesRuleCouponRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithTimeout(timeout time.Duration) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleCouponRepositoryV1GetByIDGetParams{

		timeout: timeout,
	}
}

// NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithContext creates a new SalesRuleCouponRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithContext(ctx context.Context) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleCouponRepositoryV1GetByIDGetParams{

		Context: ctx,
	}
}

// NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithHTTPClient creates a new SalesRuleCouponRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesRuleCouponRepositoryV1GetByIDGetParamsWithHTTPClient(client *http.Client) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleCouponRepositoryV1GetByIDGetParams{
		HTTPClient: client,
	}
}

/*SalesRuleCouponRepositoryV1GetByIDGetParams contains all the parameters to send to the API endpoint
for the sales rule coupon repository v1 get by Id get operation typically these are written to a http.Request
*/
type SalesRuleCouponRepositoryV1GetByIDGetParams struct {

	/*CouponID*/
	CouponID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) WithTimeout(timeout time.Duration) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) WithContext(ctx context.Context) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) WithHTTPClient(client *http.Client) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouponID adds the couponID to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) WithCouponID(couponID int64) *SalesRuleCouponRepositoryV1GetByIDGetParams {
	o.SetCouponID(couponID)
	return o
}

// SetCouponID adds the couponId to the sales rule coupon repository v1 get by Id get params
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) SetCouponID(couponID int64) {
	o.CouponID = couponID
}

// WriteToRequest writes these params to a swagger request
func (o *SalesRuleCouponRepositoryV1GetByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param couponId
	if err := r.SetPathParam("couponId", swag.FormatInt64(o.CouponID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
