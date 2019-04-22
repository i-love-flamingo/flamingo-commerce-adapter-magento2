// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_rule_repository_v1

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

// NewTaxTaxRuleRepositoryV1SavePostParams creates a new TaxTaxRuleRepositoryV1SavePostParams object
// with the default values initialized.
func NewTaxTaxRuleRepositoryV1SavePostParams() *TaxTaxRuleRepositoryV1SavePostParams {
	var ()
	return &TaxTaxRuleRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTaxTaxRuleRepositoryV1SavePostParamsWithTimeout creates a new TaxTaxRuleRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaxTaxRuleRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *TaxTaxRuleRepositoryV1SavePostParams {
	var ()
	return &TaxTaxRuleRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewTaxTaxRuleRepositoryV1SavePostParamsWithContext creates a new TaxTaxRuleRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaxTaxRuleRepositoryV1SavePostParamsWithContext(ctx context.Context) *TaxTaxRuleRepositoryV1SavePostParams {
	var ()
	return &TaxTaxRuleRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewTaxTaxRuleRepositoryV1SavePostParamsWithHTTPClient creates a new TaxTaxRuleRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaxTaxRuleRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *TaxTaxRuleRepositoryV1SavePostParams {
	var ()
	return &TaxTaxRuleRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*TaxTaxRuleRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the tax tax rule repository v1 save post operation typically these are written to a http.Request
*/
type TaxTaxRuleRepositoryV1SavePostParams struct {

	/*TaxTaxRuleRepositoryV1SavePostBody*/
	TaxTaxRuleRepositoryV1SavePostBody TaxTaxRuleRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *TaxTaxRuleRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) WithContext(ctx context.Context) *TaxTaxRuleRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *TaxTaxRuleRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaxTaxRuleRepositoryV1SavePostBody adds the taxTaxRuleRepositoryV1SavePostBody to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) WithTaxTaxRuleRepositoryV1SavePostBody(taxTaxRuleRepositoryV1SavePostBody TaxTaxRuleRepositoryV1SavePostBody) *TaxTaxRuleRepositoryV1SavePostParams {
	o.SetTaxTaxRuleRepositoryV1SavePostBody(taxTaxRuleRepositoryV1SavePostBody)
	return o
}

// SetTaxTaxRuleRepositoryV1SavePostBody adds the taxTaxRuleRepositoryV1SavePostBody to the tax tax rule repository v1 save post params
func (o *TaxTaxRuleRepositoryV1SavePostParams) SetTaxTaxRuleRepositoryV1SavePostBody(taxTaxRuleRepositoryV1SavePostBody TaxTaxRuleRepositoryV1SavePostBody) {
	o.TaxTaxRuleRepositoryV1SavePostBody = taxTaxRuleRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *TaxTaxRuleRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.TaxTaxRuleRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}