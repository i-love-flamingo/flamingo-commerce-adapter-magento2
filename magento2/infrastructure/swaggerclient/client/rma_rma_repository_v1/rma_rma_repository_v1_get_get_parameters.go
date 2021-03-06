// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_repository_v1

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

// NewRmaRmaRepositoryV1GetGetParams creates a new RmaRmaRepositoryV1GetGetParams object
// with the default values initialized.
func NewRmaRmaRepositoryV1GetGetParams() *RmaRmaRepositoryV1GetGetParams {
	var ()
	return &RmaRmaRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRmaRmaRepositoryV1GetGetParamsWithTimeout creates a new RmaRmaRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRmaRmaRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *RmaRmaRepositoryV1GetGetParams {
	var ()
	return &RmaRmaRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewRmaRmaRepositoryV1GetGetParamsWithContext creates a new RmaRmaRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRmaRmaRepositoryV1GetGetParamsWithContext(ctx context.Context) *RmaRmaRepositoryV1GetGetParams {
	var ()
	return &RmaRmaRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewRmaRmaRepositoryV1GetGetParamsWithHTTPClient creates a new RmaRmaRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRmaRmaRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *RmaRmaRepositoryV1GetGetParams {
	var ()
	return &RmaRmaRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*RmaRmaRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the rma rma repository v1 get get operation typically these are written to a http.Request
*/
type RmaRmaRepositoryV1GetGetParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *RmaRmaRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) WithContext(ctx context.Context) *RmaRmaRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *RmaRmaRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) WithID(id int64) *RmaRmaRepositoryV1GetGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rma rma repository v1 get get params
func (o *RmaRmaRepositoryV1GetGetParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RmaRmaRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
