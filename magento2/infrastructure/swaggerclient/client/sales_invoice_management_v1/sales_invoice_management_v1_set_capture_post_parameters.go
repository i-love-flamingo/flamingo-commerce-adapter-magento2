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

// NewSalesInvoiceManagementV1SetCapturePostParams creates a new SalesInvoiceManagementV1SetCapturePostParams object
// with the default values initialized.
func NewSalesInvoiceManagementV1SetCapturePostParams() *SalesInvoiceManagementV1SetCapturePostParams {
	var ()
	return &SalesInvoiceManagementV1SetCapturePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesInvoiceManagementV1SetCapturePostParamsWithTimeout creates a new SalesInvoiceManagementV1SetCapturePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesInvoiceManagementV1SetCapturePostParamsWithTimeout(timeout time.Duration) *SalesInvoiceManagementV1SetCapturePostParams {
	var ()
	return &SalesInvoiceManagementV1SetCapturePostParams{

		timeout: timeout,
	}
}

// NewSalesInvoiceManagementV1SetCapturePostParamsWithContext creates a new SalesInvoiceManagementV1SetCapturePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesInvoiceManagementV1SetCapturePostParamsWithContext(ctx context.Context) *SalesInvoiceManagementV1SetCapturePostParams {
	var ()
	return &SalesInvoiceManagementV1SetCapturePostParams{

		Context: ctx,
	}
}

// NewSalesInvoiceManagementV1SetCapturePostParamsWithHTTPClient creates a new SalesInvoiceManagementV1SetCapturePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesInvoiceManagementV1SetCapturePostParamsWithHTTPClient(client *http.Client) *SalesInvoiceManagementV1SetCapturePostParams {
	var ()
	return &SalesInvoiceManagementV1SetCapturePostParams{
		HTTPClient: client,
	}
}

/*SalesInvoiceManagementV1SetCapturePostParams contains all the parameters to send to the API endpoint
for the sales invoice management v1 set capture post operation typically these are written to a http.Request
*/
type SalesInvoiceManagementV1SetCapturePostParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) WithTimeout(timeout time.Duration) *SalesInvoiceManagementV1SetCapturePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) WithContext(ctx context.Context) *SalesInvoiceManagementV1SetCapturePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) WithHTTPClient(client *http.Client) *SalesInvoiceManagementV1SetCapturePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) WithID(id int64) *SalesInvoiceManagementV1SetCapturePostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales invoice management v1 set capture post params
func (o *SalesInvoiceManagementV1SetCapturePostParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesInvoiceManagementV1SetCapturePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
