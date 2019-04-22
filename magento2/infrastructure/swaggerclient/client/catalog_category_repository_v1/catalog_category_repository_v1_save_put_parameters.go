// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_repository_v1

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

// NewCatalogCategoryRepositoryV1SavePutParams creates a new CatalogCategoryRepositoryV1SavePutParams object
// with the default values initialized.
func NewCatalogCategoryRepositoryV1SavePutParams() *CatalogCategoryRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCategoryRepositoryV1SavePutParamsWithTimeout creates a new CatalogCategoryRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogCategoryRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CatalogCategoryRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCatalogCategoryRepositoryV1SavePutParamsWithContext creates a new CatalogCategoryRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogCategoryRepositoryV1SavePutParamsWithContext(ctx context.Context) *CatalogCategoryRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCatalogCategoryRepositoryV1SavePutParamsWithHTTPClient creates a new CatalogCategoryRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogCategoryRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CatalogCategoryRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CatalogCategoryRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the catalog category repository v1 save put operation typically these are written to a http.Request
*/
type CatalogCategoryRepositoryV1SavePutParams struct {

	/*CatalogCategoryRepositoryV1SavePutBody*/
	CatalogCategoryRepositoryV1SavePutBody CatalogCategoryRepositoryV1SavePutBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CatalogCategoryRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) WithContext(ctx context.Context) *CatalogCategoryRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CatalogCategoryRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogCategoryRepositoryV1SavePutBody adds the catalogCategoryRepositoryV1SavePutBody to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) WithCatalogCategoryRepositoryV1SavePutBody(catalogCategoryRepositoryV1SavePutBody CatalogCategoryRepositoryV1SavePutBody) *CatalogCategoryRepositoryV1SavePutParams {
	o.SetCatalogCategoryRepositoryV1SavePutBody(catalogCategoryRepositoryV1SavePutBody)
	return o
}

// SetCatalogCategoryRepositoryV1SavePutBody adds the catalogCategoryRepositoryV1SavePutBody to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) SetCatalogCategoryRepositoryV1SavePutBody(catalogCategoryRepositoryV1SavePutBody CatalogCategoryRepositoryV1SavePutBody) {
	o.CatalogCategoryRepositoryV1SavePutBody = catalogCategoryRepositoryV1SavePutBody
}

// WithID adds the id to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) WithID(id string) *CatalogCategoryRepositoryV1SavePutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the catalog category repository v1 save put params
func (o *CatalogCategoryRepositoryV1SavePutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCategoryRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogCategoryRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}