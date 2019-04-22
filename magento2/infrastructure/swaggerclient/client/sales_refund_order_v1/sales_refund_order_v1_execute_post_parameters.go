// Code generated by go-swagger; DO NOT EDIT.

package sales_refund_order_v1

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

// NewSalesRefundOrderV1ExecutePostParams creates a new SalesRefundOrderV1ExecutePostParams object
// with the default values initialized.
func NewSalesRefundOrderV1ExecutePostParams() *SalesRefundOrderV1ExecutePostParams {
	var ()
	return &SalesRefundOrderV1ExecutePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesRefundOrderV1ExecutePostParamsWithTimeout creates a new SalesRefundOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesRefundOrderV1ExecutePostParamsWithTimeout(timeout time.Duration) *SalesRefundOrderV1ExecutePostParams {
	var ()
	return &SalesRefundOrderV1ExecutePostParams{

		timeout: timeout,
	}
}

// NewSalesRefundOrderV1ExecutePostParamsWithContext creates a new SalesRefundOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesRefundOrderV1ExecutePostParamsWithContext(ctx context.Context) *SalesRefundOrderV1ExecutePostParams {
	var ()
	return &SalesRefundOrderV1ExecutePostParams{

		Context: ctx,
	}
}

// NewSalesRefundOrderV1ExecutePostParamsWithHTTPClient creates a new SalesRefundOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesRefundOrderV1ExecutePostParamsWithHTTPClient(client *http.Client) *SalesRefundOrderV1ExecutePostParams {
	var ()
	return &SalesRefundOrderV1ExecutePostParams{
		HTTPClient: client,
	}
}

/*SalesRefundOrderV1ExecutePostParams contains all the parameters to send to the API endpoint
for the sales refund order v1 execute post operation typically these are written to a http.Request
*/
type SalesRefundOrderV1ExecutePostParams struct {

	/*OrderID*/
	OrderID int64
	/*SalesRefundOrderV1ExecutePostBody*/
	SalesRefundOrderV1ExecutePostBody SalesRefundOrderV1ExecutePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) WithTimeout(timeout time.Duration) *SalesRefundOrderV1ExecutePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) WithContext(ctx context.Context) *SalesRefundOrderV1ExecutePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) WithHTTPClient(client *http.Client) *SalesRefundOrderV1ExecutePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) WithOrderID(orderID int64) *SalesRefundOrderV1ExecutePostParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WithSalesRefundOrderV1ExecutePostBody adds the salesRefundOrderV1ExecutePostBody to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) WithSalesRefundOrderV1ExecutePostBody(salesRefundOrderV1ExecutePostBody SalesRefundOrderV1ExecutePostBody) *SalesRefundOrderV1ExecutePostParams {
	o.SetSalesRefundOrderV1ExecutePostBody(salesRefundOrderV1ExecutePostBody)
	return o
}

// SetSalesRefundOrderV1ExecutePostBody adds the salesRefundOrderV1ExecutePostBody to the sales refund order v1 execute post params
func (o *SalesRefundOrderV1ExecutePostParams) SetSalesRefundOrderV1ExecutePostBody(salesRefundOrderV1ExecutePostBody SalesRefundOrderV1ExecutePostBody) {
	o.SalesRefundOrderV1ExecutePostBody = salesRefundOrderV1ExecutePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *SalesRefundOrderV1ExecutePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SalesRefundOrderV1ExecutePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
