// Code generated by go-swagger; DO NOT EDIT.

package sales_ship_order_v1

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

// NewSalesShipOrderV1ExecutePostParams creates a new SalesShipOrderV1ExecutePostParams object
// with the default values initialized.
func NewSalesShipOrderV1ExecutePostParams() *SalesShipOrderV1ExecutePostParams {
	var ()
	return &SalesShipOrderV1ExecutePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesShipOrderV1ExecutePostParamsWithTimeout creates a new SalesShipOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesShipOrderV1ExecutePostParamsWithTimeout(timeout time.Duration) *SalesShipOrderV1ExecutePostParams {
	var ()
	return &SalesShipOrderV1ExecutePostParams{

		timeout: timeout,
	}
}

// NewSalesShipOrderV1ExecutePostParamsWithContext creates a new SalesShipOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesShipOrderV1ExecutePostParamsWithContext(ctx context.Context) *SalesShipOrderV1ExecutePostParams {
	var ()
	return &SalesShipOrderV1ExecutePostParams{

		Context: ctx,
	}
}

// NewSalesShipOrderV1ExecutePostParamsWithHTTPClient creates a new SalesShipOrderV1ExecutePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesShipOrderV1ExecutePostParamsWithHTTPClient(client *http.Client) *SalesShipOrderV1ExecutePostParams {
	var ()
	return &SalesShipOrderV1ExecutePostParams{
		HTTPClient: client,
	}
}

/*SalesShipOrderV1ExecutePostParams contains all the parameters to send to the API endpoint
for the sales ship order v1 execute post operation typically these are written to a http.Request
*/
type SalesShipOrderV1ExecutePostParams struct {

	/*OrderID*/
	OrderID int64
	/*SalesShipOrderV1ExecutePostBody*/
	SalesShipOrderV1ExecutePostBody SalesShipOrderV1ExecutePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) WithTimeout(timeout time.Duration) *SalesShipOrderV1ExecutePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) WithContext(ctx context.Context) *SalesShipOrderV1ExecutePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) WithHTTPClient(client *http.Client) *SalesShipOrderV1ExecutePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) WithOrderID(orderID int64) *SalesShipOrderV1ExecutePostParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WithSalesShipOrderV1ExecutePostBody adds the salesShipOrderV1ExecutePostBody to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) WithSalesShipOrderV1ExecutePostBody(salesShipOrderV1ExecutePostBody SalesShipOrderV1ExecutePostBody) *SalesShipOrderV1ExecutePostParams {
	o.SetSalesShipOrderV1ExecutePostBody(salesShipOrderV1ExecutePostBody)
	return o
}

// SetSalesShipOrderV1ExecutePostBody adds the salesShipOrderV1ExecutePostBody to the sales ship order v1 execute post params
func (o *SalesShipOrderV1ExecutePostParams) SetSalesShipOrderV1ExecutePostBody(salesShipOrderV1ExecutePostBody SalesShipOrderV1ExecutePostBody) {
	o.SalesShipOrderV1ExecutePostBody = salesShipOrderV1ExecutePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *SalesShipOrderV1ExecutePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SalesShipOrderV1ExecutePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}