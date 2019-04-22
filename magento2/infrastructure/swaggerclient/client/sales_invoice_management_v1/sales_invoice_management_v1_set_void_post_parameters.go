// Code generated by go-swagger; DO NOT EDIT.

package sales_invoice_management_v1

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

// NewSalesInvoiceManagementV1SetVoidPostParams creates a new SalesInvoiceManagementV1SetVoidPostParams object
// with the default values initialized.
func NewSalesInvoiceManagementV1SetVoidPostParams() *SalesInvoiceManagementV1SetVoidPostParams {
	var ()
	return &SalesInvoiceManagementV1SetVoidPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesInvoiceManagementV1SetVoidPostParamsWithTimeout creates a new SalesInvoiceManagementV1SetVoidPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesInvoiceManagementV1SetVoidPostParamsWithTimeout(timeout time.Duration) *SalesInvoiceManagementV1SetVoidPostParams {
	var ()
	return &SalesInvoiceManagementV1SetVoidPostParams{

		timeout: timeout,
	}
}

// NewSalesInvoiceManagementV1SetVoidPostParamsWithContext creates a new SalesInvoiceManagementV1SetVoidPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesInvoiceManagementV1SetVoidPostParamsWithContext(ctx context.Context) *SalesInvoiceManagementV1SetVoidPostParams {
	var ()
	return &SalesInvoiceManagementV1SetVoidPostParams{

		Context: ctx,
	}
}

// NewSalesInvoiceManagementV1SetVoidPostParamsWithHTTPClient creates a new SalesInvoiceManagementV1SetVoidPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesInvoiceManagementV1SetVoidPostParamsWithHTTPClient(client *http.Client) *SalesInvoiceManagementV1SetVoidPostParams {
	var ()
	return &SalesInvoiceManagementV1SetVoidPostParams{
		HTTPClient: client,
	}
}

/*SalesInvoiceManagementV1SetVoidPostParams contains all the parameters to send to the API endpoint
for the sales invoice management v1 set void post operation typically these are written to a http.Request
*/
type SalesInvoiceManagementV1SetVoidPostParams struct {

	/*ID
	  The invoice ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) WithTimeout(timeout time.Duration) *SalesInvoiceManagementV1SetVoidPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) WithContext(ctx context.Context) *SalesInvoiceManagementV1SetVoidPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) WithHTTPClient(client *http.Client) *SalesInvoiceManagementV1SetVoidPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) WithID(id int64) *SalesInvoiceManagementV1SetVoidPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales invoice management v1 set void post params
func (o *SalesInvoiceManagementV1SetVoidPostParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesInvoiceManagementV1SetVoidPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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