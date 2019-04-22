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

// NewCatalogProductAttributeRepositoryV1SavePostParams creates a new CatalogProductAttributeRepositoryV1SavePostParams object
// with the default values initialized.
func NewCatalogProductAttributeRepositoryV1SavePostParams() *CatalogProductAttributeRepositoryV1SavePostParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePostParamsWithTimeout creates a new CatalogProductAttributeRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductAttributeRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CatalogProductAttributeRepositoryV1SavePostParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePostParamsWithContext creates a new CatalogProductAttributeRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductAttributeRepositoryV1SavePostParamsWithContext(ctx context.Context) *CatalogProductAttributeRepositoryV1SavePostParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCatalogProductAttributeRepositoryV1SavePostParamsWithHTTPClient creates a new CatalogProductAttributeRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductAttributeRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CatalogProductAttributeRepositoryV1SavePostParams {
	var ()
	return &CatalogProductAttributeRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CatalogProductAttributeRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the catalog product attribute repository v1 save post operation typically these are written to a http.Request
*/
type CatalogProductAttributeRepositoryV1SavePostParams struct {

	/*CatalogProductAttributeRepositoryV1SavePostBody*/
	CatalogProductAttributeRepositoryV1SavePostBody CatalogProductAttributeRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CatalogProductAttributeRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) WithContext(ctx context.Context) *CatalogProductAttributeRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CatalogProductAttributeRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogProductAttributeRepositoryV1SavePostBody adds the catalogProductAttributeRepositoryV1SavePostBody to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) WithCatalogProductAttributeRepositoryV1SavePostBody(catalogProductAttributeRepositoryV1SavePostBody CatalogProductAttributeRepositoryV1SavePostBody) *CatalogProductAttributeRepositoryV1SavePostParams {
	o.SetCatalogProductAttributeRepositoryV1SavePostBody(catalogProductAttributeRepositoryV1SavePostBody)
	return o
}

// SetCatalogProductAttributeRepositoryV1SavePostBody adds the catalogProductAttributeRepositoryV1SavePostBody to the catalog product attribute repository v1 save post params
func (o *CatalogProductAttributeRepositoryV1SavePostParams) SetCatalogProductAttributeRepositoryV1SavePostBody(catalogProductAttributeRepositoryV1SavePostBody CatalogProductAttributeRepositoryV1SavePostBody) {
	o.CatalogProductAttributeRepositoryV1SavePostBody = catalogProductAttributeRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductAttributeRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogProductAttributeRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
