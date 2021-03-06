// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_repository_v1

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

// NewCatalogProductAttributeRepositoryV1SavePutParams creates a new CatalogProductAttributeRepositoryV1SavePutParams object
// with the default values initialized.
func NewCatalogProductAttributeRepositoryV1SavePutParams() *CatalogProductAttributeRepositoryV1SavePutParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePutParamsWithTimeout creates a new CatalogProductAttributeRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductAttributeRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CatalogProductAttributeRepositoryV1SavePutParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePutParamsWithContext creates a new CatalogProductAttributeRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductAttributeRepositoryV1SavePutParamsWithContext(ctx context.Context) *CatalogProductAttributeRepositoryV1SavePutParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePutParamsWithHTTPClient creates a new CatalogProductAttributeRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductAttributeRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CatalogProductAttributeRepositoryV1SavePutParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CatalogProductAttributeRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the catalog product attribute repository v1 save put operation typically these are written to a http.Request
*/
type CatalogProductAttributeRepositoryV1SavePutParams struct {

	/*AttributeCode*/
	AttributeCode string
	/*CatalogProductAttributeRepositoryV1SavePutBody*/
	CatalogProductAttributeRepositoryV1SavePutBody CatalogProductAttributeRepositoryV1SavePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CatalogProductAttributeRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WithContext(ctx context.Context) *CatalogProductAttributeRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CatalogProductAttributeRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCode adds the attributeCode to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WithAttributeCode(attributeCode string) *CatalogProductAttributeRepositoryV1SavePutParams {
	o.SetAttributeCode(attributeCode)
	return o
}

// SetAttributeCode adds the attributeCode to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) SetAttributeCode(attributeCode string) {
	o.AttributeCode = attributeCode
}

// WithCatalogProductAttributeRepositoryV1SavePutBody adds the catalogProductAttributeRepositoryV1SavePutBody to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WithCatalogProductAttributeRepositoryV1SavePutBody(catalogProductAttributeRepositoryV1SavePutBody CatalogProductAttributeRepositoryV1SavePutBody) *CatalogProductAttributeRepositoryV1SavePutParams {
	o.SetCatalogProductAttributeRepositoryV1SavePutBody(catalogProductAttributeRepositoryV1SavePutBody)
	return o
}

// SetCatalogProductAttributeRepositoryV1SavePutBody adds the catalogProductAttributeRepositoryV1SavePutBody to the catalog product attribute repository v1 save put params
func (o *CatalogProductAttributeRepositoryV1SavePutParams) SetCatalogProductAttributeRepositoryV1SavePutBody(catalogProductAttributeRepositoryV1SavePutBody CatalogProductAttributeRepositoryV1SavePutBody) {
	o.CatalogProductAttributeRepositoryV1SavePutBody = catalogProductAttributeRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductAttributeRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeCode
	if err := r.SetPathParam("attributeCode", o.AttributeCode); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CatalogProductAttributeRepositoryV1SavePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
