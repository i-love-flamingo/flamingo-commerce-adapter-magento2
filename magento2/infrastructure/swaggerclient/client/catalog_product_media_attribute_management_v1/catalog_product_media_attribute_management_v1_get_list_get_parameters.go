// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_media_attribute_management_v1

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

// NewCatalogProductMediaAttributeManagementV1GetListGetParams creates a new CatalogProductMediaAttributeManagementV1GetListGetParams object
// with the default values initialized.
func NewCatalogProductMediaAttributeManagementV1GetListGetParams() *CatalogProductMediaAttributeManagementV1GetListGetParams {
	var ()
	return &CatalogProductMediaAttributeManagementV1GetListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithTimeout creates a new CatalogProductMediaAttributeManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithTimeout(timeout time.Duration) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	var ()
	return &CatalogProductMediaAttributeManagementV1GetListGetParams{

		timeout: timeout,
	}
}

// NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithContext creates a new CatalogProductMediaAttributeManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithContext(ctx context.Context) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	var ()
	return &CatalogProductMediaAttributeManagementV1GetListGetParams{

		Context: ctx,
	}
}

// NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithHTTPClient creates a new CatalogProductMediaAttributeManagementV1GetListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductMediaAttributeManagementV1GetListGetParamsWithHTTPClient(client *http.Client) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	var ()
	return &CatalogProductMediaAttributeManagementV1GetListGetParams{
		HTTPClient: client,
	}
}

/*CatalogProductMediaAttributeManagementV1GetListGetParams contains all the parameters to send to the API endpoint
for the catalog product media attribute management v1 get list get operation typically these are written to a http.Request
*/
type CatalogProductMediaAttributeManagementV1GetListGetParams struct {

	/*AttributeSetName*/
	AttributeSetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) WithTimeout(timeout time.Duration) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) WithContext(ctx context.Context) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) WithHTTPClient(client *http.Client) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeSetName adds the attributeSetName to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) WithAttributeSetName(attributeSetName string) *CatalogProductMediaAttributeManagementV1GetListGetParams {
	o.SetAttributeSetName(attributeSetName)
	return o
}

// SetAttributeSetName adds the attributeSetName to the catalog product media attribute management v1 get list get params
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) SetAttributeSetName(attributeSetName string) {
	o.AttributeSetName = attributeSetName
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductMediaAttributeManagementV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeSetName
	if err := r.SetPathParam("attributeSetName", o.AttributeSetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}