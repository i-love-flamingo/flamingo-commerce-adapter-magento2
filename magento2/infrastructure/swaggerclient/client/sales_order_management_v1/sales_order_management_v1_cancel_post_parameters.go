// Code generated by go-swagger; DO NOT EDIT.

package sales_order_management_v1

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

// NewSalesOrderManagementV1CancelPostParams creates a new SalesOrderManagementV1CancelPostParams object
// with the default values initialized.
func NewSalesOrderManagementV1CancelPostParams() *SalesOrderManagementV1CancelPostParams {
	var ()
	return &SalesOrderManagementV1CancelPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesOrderManagementV1CancelPostParamsWithTimeout creates a new SalesOrderManagementV1CancelPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesOrderManagementV1CancelPostParamsWithTimeout(timeout time.Duration) *SalesOrderManagementV1CancelPostParams {
	var ()
	return &SalesOrderManagementV1CancelPostParams{

		timeout: timeout,
	}
}

// NewSalesOrderManagementV1CancelPostParamsWithContext creates a new SalesOrderManagementV1CancelPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesOrderManagementV1CancelPostParamsWithContext(ctx context.Context) *SalesOrderManagementV1CancelPostParams {
	var ()
	return &SalesOrderManagementV1CancelPostParams{

		Context: ctx,
	}
}

// NewSalesOrderManagementV1CancelPostParamsWithHTTPClient creates a new SalesOrderManagementV1CancelPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesOrderManagementV1CancelPostParamsWithHTTPClient(client *http.Client) *SalesOrderManagementV1CancelPostParams {
	var ()
	return &SalesOrderManagementV1CancelPostParams{
		HTTPClient: client,
	}
}

/*SalesOrderManagementV1CancelPostParams contains all the parameters to send to the API endpoint
for the sales order management v1 cancel post operation typically these are written to a http.Request
*/
type SalesOrderManagementV1CancelPostParams struct {

	/*ID
	  The order ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) WithTimeout(timeout time.Duration) *SalesOrderManagementV1CancelPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) WithContext(ctx context.Context) *SalesOrderManagementV1CancelPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) WithHTTPClient(client *http.Client) *SalesOrderManagementV1CancelPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) WithID(id int64) *SalesOrderManagementV1CancelPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales order management v1 cancel post params
func (o *SalesOrderManagementV1CancelPostParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesOrderManagementV1CancelPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
