// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_class_repository_v1

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

// NewTaxTaxClassRepositoryV1GetGetParams creates a new TaxTaxClassRepositoryV1GetGetParams object
// with the default values initialized.
func NewTaxTaxClassRepositoryV1GetGetParams() *TaxTaxClassRepositoryV1GetGetParams {
	var ()
	return &TaxTaxClassRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTaxTaxClassRepositoryV1GetGetParamsWithTimeout creates a new TaxTaxClassRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaxTaxClassRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *TaxTaxClassRepositoryV1GetGetParams {
	var ()
	return &TaxTaxClassRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewTaxTaxClassRepositoryV1GetGetParamsWithContext creates a new TaxTaxClassRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaxTaxClassRepositoryV1GetGetParamsWithContext(ctx context.Context) *TaxTaxClassRepositoryV1GetGetParams {
	var ()
	return &TaxTaxClassRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewTaxTaxClassRepositoryV1GetGetParamsWithHTTPClient creates a new TaxTaxClassRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaxTaxClassRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *TaxTaxClassRepositoryV1GetGetParams {
	var ()
	return &TaxTaxClassRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*TaxTaxClassRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the tax tax class repository v1 get get operation typically these are written to a http.Request
*/
type TaxTaxClassRepositoryV1GetGetParams struct {

	/*TaxClassID*/
	TaxClassID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *TaxTaxClassRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) WithContext(ctx context.Context) *TaxTaxClassRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *TaxTaxClassRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaxClassID adds the taxClassID to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) WithTaxClassID(taxClassID int64) *TaxTaxClassRepositoryV1GetGetParams {
	o.SetTaxClassID(taxClassID)
	return o
}

// SetTaxClassID adds the taxClassId to the tax tax class repository v1 get get params
func (o *TaxTaxClassRepositoryV1GetGetParams) SetTaxClassID(taxClassID int64) {
	o.TaxClassID = taxClassID
}

// WriteToRequest writes these params to a swagger request
func (o *TaxTaxClassRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taxClassId
	if err := r.SetPathParam("taxClassId", swag.FormatInt64(o.TaxClassID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
