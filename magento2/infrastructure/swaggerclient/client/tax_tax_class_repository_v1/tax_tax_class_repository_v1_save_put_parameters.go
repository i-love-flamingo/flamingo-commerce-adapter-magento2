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

	strfmt "github.com/go-openapi/strfmt"
)

// NewTaxTaxClassRepositoryV1SavePutParams creates a new TaxTaxClassRepositoryV1SavePutParams object
// with the default values initialized.
func NewTaxTaxClassRepositoryV1SavePutParams() *TaxTaxClassRepositoryV1SavePutParams {
	var ()
	return &TaxTaxClassRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTaxTaxClassRepositoryV1SavePutParamsWithTimeout creates a new TaxTaxClassRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaxTaxClassRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *TaxTaxClassRepositoryV1SavePutParams {
	var ()
	return &TaxTaxClassRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewTaxTaxClassRepositoryV1SavePutParamsWithContext creates a new TaxTaxClassRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaxTaxClassRepositoryV1SavePutParamsWithContext(ctx context.Context) *TaxTaxClassRepositoryV1SavePutParams {
	var ()
	return &TaxTaxClassRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewTaxTaxClassRepositoryV1SavePutParamsWithHTTPClient creates a new TaxTaxClassRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaxTaxClassRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *TaxTaxClassRepositoryV1SavePutParams {
	var ()
	return &TaxTaxClassRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*TaxTaxClassRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the tax tax class repository v1 save put operation typically these are written to a http.Request
*/
type TaxTaxClassRepositoryV1SavePutParams struct {

	/*ClassID*/
	ClassID string
	/*TaxTaxClassRepositoryV1SavePutBody*/
	TaxTaxClassRepositoryV1SavePutBody TaxTaxClassRepositoryV1SavePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *TaxTaxClassRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) WithContext(ctx context.Context) *TaxTaxClassRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *TaxTaxClassRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassID adds the classID to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) WithClassID(classID string) *TaxTaxClassRepositoryV1SavePutParams {
	o.SetClassID(classID)
	return o
}

// SetClassID adds the classId to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) SetClassID(classID string) {
	o.ClassID = classID
}

// WithTaxTaxClassRepositoryV1SavePutBody adds the taxTaxClassRepositoryV1SavePutBody to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) WithTaxTaxClassRepositoryV1SavePutBody(taxTaxClassRepositoryV1SavePutBody TaxTaxClassRepositoryV1SavePutBody) *TaxTaxClassRepositoryV1SavePutParams {
	o.SetTaxTaxClassRepositoryV1SavePutBody(taxTaxClassRepositoryV1SavePutBody)
	return o
}

// SetTaxTaxClassRepositoryV1SavePutBody adds the taxTaxClassRepositoryV1SavePutBody to the tax tax class repository v1 save put params
func (o *TaxTaxClassRepositoryV1SavePutParams) SetTaxTaxClassRepositoryV1SavePutBody(taxTaxClassRepositoryV1SavePutBody TaxTaxClassRepositoryV1SavePutBody) {
	o.TaxTaxClassRepositoryV1SavePutBody = taxTaxClassRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *TaxTaxClassRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param classId
	if err := r.SetPathParam("classId", o.ClassID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.TaxTaxClassRepositoryV1SavePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
