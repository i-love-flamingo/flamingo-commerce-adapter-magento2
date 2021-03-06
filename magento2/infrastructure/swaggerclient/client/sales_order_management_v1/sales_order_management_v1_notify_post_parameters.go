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

// NewSalesOrderManagementV1NotifyPostParams creates a new SalesOrderManagementV1NotifyPostParams object
// with the default values initialized.
func NewSalesOrderManagementV1NotifyPostParams() *SalesOrderManagementV1NotifyPostParams {
	var ()
	return &SalesOrderManagementV1NotifyPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesOrderManagementV1NotifyPostParamsWithTimeout creates a new SalesOrderManagementV1NotifyPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesOrderManagementV1NotifyPostParamsWithTimeout(timeout time.Duration) *SalesOrderManagementV1NotifyPostParams {
	var ()
	return &SalesOrderManagementV1NotifyPostParams{

		timeout: timeout,
	}
}

// NewSalesOrderManagementV1NotifyPostParamsWithContext creates a new SalesOrderManagementV1NotifyPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesOrderManagementV1NotifyPostParamsWithContext(ctx context.Context) *SalesOrderManagementV1NotifyPostParams {
	var ()
	return &SalesOrderManagementV1NotifyPostParams{

		Context: ctx,
	}
}

// NewSalesOrderManagementV1NotifyPostParamsWithHTTPClient creates a new SalesOrderManagementV1NotifyPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesOrderManagementV1NotifyPostParamsWithHTTPClient(client *http.Client) *SalesOrderManagementV1NotifyPostParams {
	var ()
	return &SalesOrderManagementV1NotifyPostParams{
		HTTPClient: client,
	}
}

/*SalesOrderManagementV1NotifyPostParams contains all the parameters to send to the API endpoint
for the sales order management v1 notify post operation typically these are written to a http.Request
*/
type SalesOrderManagementV1NotifyPostParams struct {

	/*ID
	  The order ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) WithTimeout(timeout time.Duration) *SalesOrderManagementV1NotifyPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) WithContext(ctx context.Context) *SalesOrderManagementV1NotifyPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) WithHTTPClient(client *http.Client) *SalesOrderManagementV1NotifyPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) WithID(id int64) *SalesOrderManagementV1NotifyPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales order management v1 notify post params
func (o *SalesOrderManagementV1NotifyPostParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesOrderManagementV1NotifyPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
