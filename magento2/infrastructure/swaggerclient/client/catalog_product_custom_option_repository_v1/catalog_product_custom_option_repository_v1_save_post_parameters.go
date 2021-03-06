// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_custom_option_repository_v1

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

// NewCatalogProductCustomOptionRepositoryV1SavePostParams creates a new CatalogProductCustomOptionRepositoryV1SavePostParams object
// with the default values initialized.
func NewCatalogProductCustomOptionRepositoryV1SavePostParams() *CatalogProductCustomOptionRepositoryV1SavePostParams {
	var ()
	return &CatalogProductCustomOptionRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithTimeout creates a new CatalogProductCustomOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	var ()
	return &CatalogProductCustomOptionRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithContext creates a new CatalogProductCustomOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithContext(ctx context.Context) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	var ()
	return &CatalogProductCustomOptionRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithHTTPClient creates a new CatalogProductCustomOptionRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductCustomOptionRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	var ()
	return &CatalogProductCustomOptionRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CatalogProductCustomOptionRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the catalog product custom option repository v1 save post operation typically these are written to a http.Request
*/
type CatalogProductCustomOptionRepositoryV1SavePostParams struct {

	/*CatalogProductCustomOptionRepositoryV1SavePostBody*/
	CatalogProductCustomOptionRepositoryV1SavePostBody CatalogProductCustomOptionRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) WithContext(ctx context.Context) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogProductCustomOptionRepositoryV1SavePostBody adds the catalogProductCustomOptionRepositoryV1SavePostBody to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) WithCatalogProductCustomOptionRepositoryV1SavePostBody(catalogProductCustomOptionRepositoryV1SavePostBody CatalogProductCustomOptionRepositoryV1SavePostBody) *CatalogProductCustomOptionRepositoryV1SavePostParams {
	o.SetCatalogProductCustomOptionRepositoryV1SavePostBody(catalogProductCustomOptionRepositoryV1SavePostBody)
	return o
}

// SetCatalogProductCustomOptionRepositoryV1SavePostBody adds the catalogProductCustomOptionRepositoryV1SavePostBody to the catalog product custom option repository v1 save post params
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) SetCatalogProductCustomOptionRepositoryV1SavePostBody(catalogProductCustomOptionRepositoryV1SavePostBody CatalogProductCustomOptionRepositoryV1SavePostBody) {
	o.CatalogProductCustomOptionRepositoryV1SavePostBody = catalogProductCustomOptionRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductCustomOptionRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogProductCustomOptionRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
