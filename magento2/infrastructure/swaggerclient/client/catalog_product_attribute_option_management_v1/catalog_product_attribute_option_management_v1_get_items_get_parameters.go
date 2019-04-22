// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_option_management_v1

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

// NewCatalogProductAttributeOptionManagementV1GetItemsGetParams creates a new CatalogProductAttributeOptionManagementV1GetItemsGetParams object
// with the default values initialized.
func NewCatalogProductAttributeOptionManagementV1GetItemsGetParams() *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1GetItemsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithTimeout creates a new CatalogProductAttributeOptionManagementV1GetItemsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithTimeout(timeout time.Duration) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1GetItemsGetParams{

		timeout: timeout,
	}
}

// NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithContext creates a new CatalogProductAttributeOptionManagementV1GetItemsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithContext(ctx context.Context) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1GetItemsGetParams{

		Context: ctx,
	}
}

// NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithHTTPClient creates a new CatalogProductAttributeOptionManagementV1GetItemsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductAttributeOptionManagementV1GetItemsGetParamsWithHTTPClient(client *http.Client) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1GetItemsGetParams{
		HTTPClient: client,
	}
}

/*CatalogProductAttributeOptionManagementV1GetItemsGetParams contains all the parameters to send to the API endpoint
for the catalog product attribute option management v1 get items get operation typically these are written to a http.Request
*/
type CatalogProductAttributeOptionManagementV1GetItemsGetParams struct {

	/*AttributeCode*/
	AttributeCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) WithTimeout(timeout time.Duration) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) WithContext(ctx context.Context) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) WithHTTPClient(client *http.Client) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCode adds the attributeCode to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) WithAttributeCode(attributeCode string) *CatalogProductAttributeOptionManagementV1GetItemsGetParams {
	o.SetAttributeCode(attributeCode)
	return o
}

// SetAttributeCode adds the attributeCode to the catalog product attribute option management v1 get items get params
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) SetAttributeCode(attributeCode string) {
	o.AttributeCode = attributeCode
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductAttributeOptionManagementV1GetItemsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeCode
	if err := r.SetPathParam("attributeCode", o.AttributeCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
