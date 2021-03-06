// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_rate_repository_v1

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

// NewTaxTaxRateRepositoryV1GetGetParams creates a new TaxTaxRateRepositoryV1GetGetParams object
// with the default values initialized.
func NewTaxTaxRateRepositoryV1GetGetParams() *TaxTaxRateRepositoryV1GetGetParams {
	var ()
	return &TaxTaxRateRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTaxTaxRateRepositoryV1GetGetParamsWithTimeout creates a new TaxTaxRateRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaxTaxRateRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *TaxTaxRateRepositoryV1GetGetParams {
	var ()
	return &TaxTaxRateRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewTaxTaxRateRepositoryV1GetGetParamsWithContext creates a new TaxTaxRateRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaxTaxRateRepositoryV1GetGetParamsWithContext(ctx context.Context) *TaxTaxRateRepositoryV1GetGetParams {
	var ()
	return &TaxTaxRateRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewTaxTaxRateRepositoryV1GetGetParamsWithHTTPClient creates a new TaxTaxRateRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaxTaxRateRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *TaxTaxRateRepositoryV1GetGetParams {
	var ()
	return &TaxTaxRateRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*TaxTaxRateRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the tax tax rate repository v1 get get operation typically these are written to a http.Request
*/
type TaxTaxRateRepositoryV1GetGetParams struct {

	/*RateID*/
	RateID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *TaxTaxRateRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) WithContext(ctx context.Context) *TaxTaxRateRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *TaxTaxRateRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRateID adds the rateID to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) WithRateID(rateID int64) *TaxTaxRateRepositoryV1GetGetParams {
	o.SetRateID(rateID)
	return o
}

// SetRateID adds the rateId to the tax tax rate repository v1 get get params
func (o *TaxTaxRateRepositoryV1GetGetParams) SetRateID(rateID int64) {
	o.RateID = rateID
}

// WriteToRequest writes these params to a swagger request
func (o *TaxTaxRateRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rateId
	if err := r.SetPathParam("rateId", swag.FormatInt64(o.RateID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
