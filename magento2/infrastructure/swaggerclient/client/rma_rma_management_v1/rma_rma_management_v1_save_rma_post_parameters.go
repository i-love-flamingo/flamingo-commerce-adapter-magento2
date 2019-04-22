// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRmaRmaManagementV1SaveRmaPostParams creates a new RmaRmaManagementV1SaveRmaPostParams object
// with the default values initialized.
func NewRmaRmaManagementV1SaveRmaPostParams() *RmaRmaManagementV1SaveRmaPostParams {
	var ()
	return &RmaRmaManagementV1SaveRmaPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRmaRmaManagementV1SaveRmaPostParamsWithTimeout creates a new RmaRmaManagementV1SaveRmaPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRmaRmaManagementV1SaveRmaPostParamsWithTimeout(timeout time.Duration) *RmaRmaManagementV1SaveRmaPostParams {
	var ()
	return &RmaRmaManagementV1SaveRmaPostParams{

		timeout: timeout,
	}
}

// NewRmaRmaManagementV1SaveRmaPostParamsWithContext creates a new RmaRmaManagementV1SaveRmaPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewRmaRmaManagementV1SaveRmaPostParamsWithContext(ctx context.Context) *RmaRmaManagementV1SaveRmaPostParams {
	var ()
	return &RmaRmaManagementV1SaveRmaPostParams{

		Context: ctx,
	}
}

// NewRmaRmaManagementV1SaveRmaPostParamsWithHTTPClient creates a new RmaRmaManagementV1SaveRmaPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRmaRmaManagementV1SaveRmaPostParamsWithHTTPClient(client *http.Client) *RmaRmaManagementV1SaveRmaPostParams {
	var ()
	return &RmaRmaManagementV1SaveRmaPostParams{
		HTTPClient: client,
	}
}

/*RmaRmaManagementV1SaveRmaPostParams contains all the parameters to send to the API endpoint
for the rma rma management v1 save rma post operation typically these are written to a http.Request
*/
type RmaRmaManagementV1SaveRmaPostParams struct {

	/*RmaRmaManagementV1SaveRmaPostBody*/
	RmaRmaManagementV1SaveRmaPostBody RmaRmaManagementV1SaveRmaPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) WithTimeout(timeout time.Duration) *RmaRmaManagementV1SaveRmaPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) WithContext(ctx context.Context) *RmaRmaManagementV1SaveRmaPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) WithHTTPClient(client *http.Client) *RmaRmaManagementV1SaveRmaPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRmaRmaManagementV1SaveRmaPostBody adds the rmaRmaManagementV1SaveRmaPostBody to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) WithRmaRmaManagementV1SaveRmaPostBody(rmaRmaManagementV1SaveRmaPostBody RmaRmaManagementV1SaveRmaPostBody) *RmaRmaManagementV1SaveRmaPostParams {
	o.SetRmaRmaManagementV1SaveRmaPostBody(rmaRmaManagementV1SaveRmaPostBody)
	return o
}

// SetRmaRmaManagementV1SaveRmaPostBody adds the rmaRmaManagementV1SaveRmaPostBody to the rma rma management v1 save rma post params
func (o *RmaRmaManagementV1SaveRmaPostParams) SetRmaRmaManagementV1SaveRmaPostBody(rmaRmaManagementV1SaveRmaPostBody RmaRmaManagementV1SaveRmaPostBody) {
	o.RmaRmaManagementV1SaveRmaPostBody = rmaRmaManagementV1SaveRmaPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *RmaRmaManagementV1SaveRmaPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.RmaRmaManagementV1SaveRmaPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}