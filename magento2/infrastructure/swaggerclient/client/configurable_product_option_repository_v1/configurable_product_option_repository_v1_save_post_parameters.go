// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_option_repository_v1

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

// NewConfigurableProductOptionRepositoryV1SavePostParams creates a new ConfigurableProductOptionRepositoryV1SavePostParams object
// with the default values initialized.
func NewConfigurableProductOptionRepositoryV1SavePostParams() *ConfigurableProductOptionRepositoryV1SavePostParams {
	var ()
	return &ConfigurableProductOptionRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigurableProductOptionRepositoryV1SavePostParamsWithTimeout creates a new ConfigurableProductOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigurableProductOptionRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *ConfigurableProductOptionRepositoryV1SavePostParams {
	var ()
	return &ConfigurableProductOptionRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewConfigurableProductOptionRepositoryV1SavePostParamsWithContext creates a new ConfigurableProductOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigurableProductOptionRepositoryV1SavePostParamsWithContext(ctx context.Context) *ConfigurableProductOptionRepositoryV1SavePostParams {
	var ()
	return &ConfigurableProductOptionRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewConfigurableProductOptionRepositoryV1SavePostParamsWithHTTPClient creates a new ConfigurableProductOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigurableProductOptionRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *ConfigurableProductOptionRepositoryV1SavePostParams {
	var ()
	return &ConfigurableProductOptionRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*ConfigurableProductOptionRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the configurable product option repository v1 save post operation typically these are written to a http.Request
*/
type ConfigurableProductOptionRepositoryV1SavePostParams struct {

	/*ConfigurableProductOptionRepositoryV1SavePostBody*/
	ConfigurableProductOptionRepositoryV1SavePostBody ConfigurableProductOptionRepositoryV1SavePostBody
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *ConfigurableProductOptionRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WithContext(ctx context.Context) *ConfigurableProductOptionRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *ConfigurableProductOptionRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurableProductOptionRepositoryV1SavePostBody adds the configurableProductOptionRepositoryV1SavePostBody to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WithConfigurableProductOptionRepositoryV1SavePostBody(configurableProductOptionRepositoryV1SavePostBody ConfigurableProductOptionRepositoryV1SavePostBody) *ConfigurableProductOptionRepositoryV1SavePostParams {
	o.SetConfigurableProductOptionRepositoryV1SavePostBody(configurableProductOptionRepositoryV1SavePostBody)
	return o
}

// SetConfigurableProductOptionRepositoryV1SavePostBody adds the configurableProductOptionRepositoryV1SavePostBody to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) SetConfigurableProductOptionRepositoryV1SavePostBody(configurableProductOptionRepositoryV1SavePostBody ConfigurableProductOptionRepositoryV1SavePostBody) {
	o.ConfigurableProductOptionRepositoryV1SavePostBody = configurableProductOptionRepositoryV1SavePostBody
}

// WithSku adds the sku to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WithSku(sku string) *ConfigurableProductOptionRepositoryV1SavePostParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the configurable product option repository v1 save post params
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigurableProductOptionRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.ConfigurableProductOptionRepositoryV1SavePostBody); err != nil {
		return err
	}

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
