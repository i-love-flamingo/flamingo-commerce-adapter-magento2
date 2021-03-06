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

// NewSalesOrderManagementV1HoldPostParams creates a new SalesOrderManagementV1HoldPostParams object
// with the default values initialized.
func NewSalesOrderManagementV1HoldPostParams() *SalesOrderManagementV1HoldPostParams {
	var ()
	return &SalesOrderManagementV1HoldPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesOrderManagementV1HoldPostParamsWithTimeout creates a new SalesOrderManagementV1HoldPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesOrderManagementV1HoldPostParamsWithTimeout(timeout time.Duration) *SalesOrderManagementV1HoldPostParams {
	var ()
	return &SalesOrderManagementV1HoldPostParams{

		timeout: timeout,
	}
}

// NewSalesOrderManagementV1HoldPostParamsWithContext creates a new SalesOrderManagementV1HoldPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesOrderManagementV1HoldPostParamsWithContext(ctx context.Context) *SalesOrderManagementV1HoldPostParams {
	var ()
	return &SalesOrderManagementV1HoldPostParams{

		Context: ctx,
	}
}

// NewSalesOrderManagementV1HoldPostParamsWithHTTPClient creates a new SalesOrderManagementV1HoldPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesOrderManagementV1HoldPostParamsWithHTTPClient(client *http.Client) *SalesOrderManagementV1HoldPostParams {
	var ()
	return &SalesOrderManagementV1HoldPostParams{
		HTTPClient: client,
	}
}

/*SalesOrderManagementV1HoldPostParams contains all the parameters to send to the API endpoint
for the sales order management v1 hold post operation typically these are written to a http.Request
*/
type SalesOrderManagementV1HoldPostParams struct {

	/*ID
	  The order ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) WithTimeout(timeout time.Duration) *SalesOrderManagementV1HoldPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) WithContext(ctx context.Context) *SalesOrderManagementV1HoldPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) WithHTTPClient(client *http.Client) *SalesOrderManagementV1HoldPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) WithID(id int64) *SalesOrderManagementV1HoldPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales order management v1 hold post params
func (o *SalesOrderManagementV1HoldPostParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesOrderManagementV1HoldPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
