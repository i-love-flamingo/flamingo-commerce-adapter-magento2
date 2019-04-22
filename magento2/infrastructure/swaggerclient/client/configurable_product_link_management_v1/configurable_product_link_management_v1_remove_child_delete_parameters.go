// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_link_management_v1

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

// NewConfigurableProductLinkManagementV1RemoveChildDeleteParams creates a new ConfigurableProductLinkManagementV1RemoveChildDeleteParams object
// with the default values initialized.
func NewConfigurableProductLinkManagementV1RemoveChildDeleteParams() *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	var ()
	return &ConfigurableProductLinkManagementV1RemoveChildDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithTimeout creates a new ConfigurableProductLinkManagementV1RemoveChildDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithTimeout(timeout time.Duration) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	var ()
	return &ConfigurableProductLinkManagementV1RemoveChildDeleteParams{

		timeout: timeout,
	}
}

// NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithContext creates a new ConfigurableProductLinkManagementV1RemoveChildDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithContext(ctx context.Context) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	var ()
	return &ConfigurableProductLinkManagementV1RemoveChildDeleteParams{

		Context: ctx,
	}
}

// NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithHTTPClient creates a new ConfigurableProductLinkManagementV1RemoveChildDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigurableProductLinkManagementV1RemoveChildDeleteParamsWithHTTPClient(client *http.Client) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	var ()
	return &ConfigurableProductLinkManagementV1RemoveChildDeleteParams{
		HTTPClient: client,
	}
}

/*ConfigurableProductLinkManagementV1RemoveChildDeleteParams contains all the parameters to send to the API endpoint
for the configurable product link management v1 remove child delete operation typically these are written to a http.Request
*/
type ConfigurableProductLinkManagementV1RemoveChildDeleteParams struct {

	/*ChildSku*/
	ChildSku string
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WithTimeout(timeout time.Duration) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WithContext(ctx context.Context) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WithHTTPClient(client *http.Client) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildSku adds the childSku to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WithChildSku(childSku string) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	o.SetChildSku(childSku)
	return o
}

// SetChildSku adds the childSku to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) SetChildSku(childSku string) {
	o.ChildSku = childSku
}

// WithSku adds the sku to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WithSku(sku string) *ConfigurableProductLinkManagementV1RemoveChildDeleteParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the configurable product link management v1 remove child delete params
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigurableProductLinkManagementV1RemoveChildDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childSku
	if err := r.SetPathParam("childSku", o.ChildSku); err != nil {
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