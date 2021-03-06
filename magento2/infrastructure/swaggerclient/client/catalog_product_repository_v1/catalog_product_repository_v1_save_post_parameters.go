// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

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

// NewCatalogProductRepositoryV1SavePostParams creates a new CatalogProductRepositoryV1SavePostParams object
// with the default values initialized.
func NewCatalogProductRepositoryV1SavePostParams() *CatalogProductRepositoryV1SavePostParams {
	var ()
	return &CatalogProductRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductRepositoryV1SavePostParamsWithTimeout creates a new CatalogProductRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CatalogProductRepositoryV1SavePostParams {
	var ()
	return &CatalogProductRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCatalogProductRepositoryV1SavePostParamsWithContext creates a new CatalogProductRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductRepositoryV1SavePostParamsWithContext(ctx context.Context) *CatalogProductRepositoryV1SavePostParams {
	var ()
	return &CatalogProductRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCatalogProductRepositoryV1SavePostParamsWithHTTPClient creates a new CatalogProductRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CatalogProductRepositoryV1SavePostParams {
	var ()
	return &CatalogProductRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CatalogProductRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the catalog product repository v1 save post operation typically these are written to a http.Request
*/
type CatalogProductRepositoryV1SavePostParams struct {

	/*CatalogProductRepositoryV1SavePostBody*/
	CatalogProductRepositoryV1SavePostBody CatalogProductRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CatalogProductRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) WithContext(ctx context.Context) *CatalogProductRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CatalogProductRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogProductRepositoryV1SavePostBody adds the catalogProductRepositoryV1SavePostBody to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) WithCatalogProductRepositoryV1SavePostBody(catalogProductRepositoryV1SavePostBody CatalogProductRepositoryV1SavePostBody) *CatalogProductRepositoryV1SavePostParams {
	o.SetCatalogProductRepositoryV1SavePostBody(catalogProductRepositoryV1SavePostBody)
	return o
}

// SetCatalogProductRepositoryV1SavePostBody adds the catalogProductRepositoryV1SavePostBody to the catalog product repository v1 save post params
func (o *CatalogProductRepositoryV1SavePostParams) SetCatalogProductRepositoryV1SavePostBody(catalogProductRepositoryV1SavePostBody CatalogProductRepositoryV1SavePostBody) {
	o.CatalogProductRepositoryV1SavePostBody = catalogProductRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogProductRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
