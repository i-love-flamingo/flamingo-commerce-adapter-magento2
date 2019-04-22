// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_attribute_repository_v1

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

// NewCatalogCategoryAttributeRepositoryV1GetGetParams creates a new CatalogCategoryAttributeRepositoryV1GetGetParams object
// with the default values initialized.
func NewCatalogCategoryAttributeRepositoryV1GetGetParams() *CatalogCategoryAttributeRepositoryV1GetGetParams {
	var ()
	return &CatalogCategoryAttributeRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithTimeout creates a new CatalogCategoryAttributeRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	var ()
	return &CatalogCategoryAttributeRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithContext creates a new CatalogCategoryAttributeRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithContext(ctx context.Context) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	var ()
	return &CatalogCategoryAttributeRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithHTTPClient creates a new CatalogCategoryAttributeRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogCategoryAttributeRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	var ()
	return &CatalogCategoryAttributeRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*CatalogCategoryAttributeRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the catalog category attribute repository v1 get get operation typically these are written to a http.Request
*/
type CatalogCategoryAttributeRepositoryV1GetGetParams struct {

	/*AttributeCode*/
	AttributeCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) WithContext(ctx context.Context) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCode adds the attributeCode to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) WithAttributeCode(attributeCode string) *CatalogCategoryAttributeRepositoryV1GetGetParams {
	o.SetAttributeCode(attributeCode)
	return o
}

// SetAttributeCode adds the attributeCode to the catalog category attribute repository v1 get get params
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) SetAttributeCode(attributeCode string) {
	o.AttributeCode = attributeCode
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCategoryAttributeRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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