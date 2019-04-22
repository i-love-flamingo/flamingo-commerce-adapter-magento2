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

// NewCatalogProductAttributeOptionManagementV1AddPostParams creates a new CatalogProductAttributeOptionManagementV1AddPostParams object
// with the default values initialized.
func NewCatalogProductAttributeOptionManagementV1AddPostParams() *CatalogProductAttributeOptionManagementV1AddPostParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1AddPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductAttributeOptionManagementV1AddPostParamsWithTimeout creates a new CatalogProductAttributeOptionManagementV1AddPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductAttributeOptionManagementV1AddPostParamsWithTimeout(timeout time.Duration) *CatalogProductAttributeOptionManagementV1AddPostParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1AddPostParams{

		timeout: timeout,
	}
}

// NewCatalogProductAttributeOptionManagementV1AddPostParamsWithContext creates a new CatalogProductAttributeOptionManagementV1AddPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductAttributeOptionManagementV1AddPostParamsWithContext(ctx context.Context) *CatalogProductAttributeOptionManagementV1AddPostParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1AddPostParams{

		Context: ctx,
	}
}

// NewCatalogProductAttributeOptionManagementV1AddPostParamsWithHTTPClient creates a new CatalogProductAttributeOptionManagementV1AddPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductAttributeOptionManagementV1AddPostParamsWithHTTPClient(client *http.Client) *CatalogProductAttributeOptionManagementV1AddPostParams {
	var ()
	return &CatalogProductAttributeOptionManagementV1AddPostParams{
		HTTPClient: client,
	}
}

/*CatalogProductAttributeOptionManagementV1AddPostParams contains all the parameters to send to the API endpoint
for the catalog product attribute option management v1 add post operation typically these are written to a http.Request
*/
type CatalogProductAttributeOptionManagementV1AddPostParams struct {

	/*AttributeCode*/
	AttributeCode string
	/*CatalogProductAttributeOptionManagementV1AddPostBody*/
	CatalogProductAttributeOptionManagementV1AddPostBody CatalogProductAttributeOptionManagementV1AddPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WithTimeout(timeout time.Duration) *CatalogProductAttributeOptionManagementV1AddPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WithContext(ctx context.Context) *CatalogProductAttributeOptionManagementV1AddPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WithHTTPClient(client *http.Client) *CatalogProductAttributeOptionManagementV1AddPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCode adds the attributeCode to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WithAttributeCode(attributeCode string) *CatalogProductAttributeOptionManagementV1AddPostParams {
	o.SetAttributeCode(attributeCode)
	return o
}

// SetAttributeCode adds the attributeCode to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) SetAttributeCode(attributeCode string) {
	o.AttributeCode = attributeCode
}

// WithCatalogProductAttributeOptionManagementV1AddPostBody adds the catalogProductAttributeOptionManagementV1AddPostBody to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WithCatalogProductAttributeOptionManagementV1AddPostBody(catalogProductAttributeOptionManagementV1AddPostBody CatalogProductAttributeOptionManagementV1AddPostBody) *CatalogProductAttributeOptionManagementV1AddPostParams {
	o.SetCatalogProductAttributeOptionManagementV1AddPostBody(catalogProductAttributeOptionManagementV1AddPostBody)
	return o
}

// SetCatalogProductAttributeOptionManagementV1AddPostBody adds the catalogProductAttributeOptionManagementV1AddPostBody to the catalog product attribute option management v1 add post params
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) SetCatalogProductAttributeOptionManagementV1AddPostBody(catalogProductAttributeOptionManagementV1AddPostBody CatalogProductAttributeOptionManagementV1AddPostBody) {
	o.CatalogProductAttributeOptionManagementV1AddPostBody = catalogProductAttributeOptionManagementV1AddPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductAttributeOptionManagementV1AddPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeCode
	if err := r.SetPathParam("attributeCode", o.AttributeCode); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CatalogProductAttributeOptionManagementV1AddPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}