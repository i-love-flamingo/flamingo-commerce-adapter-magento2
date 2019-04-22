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

// NewSalesInvoiceManagementV1GetCommentsListGetParams creates a new SalesInvoiceManagementV1GetCommentsListGetParams object
// with the default values initialized.
func NewSalesInvoiceManagementV1GetCommentsListGetParams() *SalesInvoiceManagementV1GetCommentsListGetParams {
	var ()
	return &SalesInvoiceManagementV1GetCommentsListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesInvoiceManagementV1GetCommentsListGetParamsWithTimeout creates a new SalesInvoiceManagementV1GetCommentsListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesInvoiceManagementV1GetCommentsListGetParamsWithTimeout(timeout time.Duration) *SalesInvoiceManagementV1GetCommentsListGetParams {
	var ()
	return &SalesInvoiceManagementV1GetCommentsListGetParams{

		timeout: timeout,
	}
}

// NewSalesInvoiceManagementV1GetCommentsListGetParamsWithContext creates a new SalesInvoiceManagementV1GetCommentsListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesInvoiceManagementV1GetCommentsListGetParamsWithContext(ctx context.Context) *SalesInvoiceManagementV1GetCommentsListGetParams {
	var ()
	return &SalesInvoiceManagementV1GetCommentsListGetParams{

		Context: ctx,
	}
}

// NewSalesInvoiceManagementV1GetCommentsListGetParamsWithHTTPClient creates a new SalesInvoiceManagementV1GetCommentsListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesInvoiceManagementV1GetCommentsListGetParamsWithHTTPClient(client *http.Client) *SalesInvoiceManagementV1GetCommentsListGetParams {
	var ()
	return &SalesInvoiceManagementV1GetCommentsListGetParams{
		HTTPClient: client,
	}
}

/*SalesInvoiceManagementV1GetCommentsListGetParams contains all the parameters to send to the API endpoint
for the sales invoice management v1 get comments list get operation typically these are written to a http.Request
*/
type SalesInvoiceManagementV1GetCommentsListGetParams struct {

	/*ID
	  The invoice ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) WithTimeout(timeout time.Duration) *SalesInvoiceManagementV1GetCommentsListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) WithContext(ctx context.Context) *SalesInvoiceManagementV1GetCommentsListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) WithHTTPClient(client *http.Client) *SalesInvoiceManagementV1GetCommentsListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) WithID(id int64) *SalesInvoiceManagementV1GetCommentsListGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales invoice management v1 get comments list get params
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesInvoiceManagementV1GetCommentsListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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