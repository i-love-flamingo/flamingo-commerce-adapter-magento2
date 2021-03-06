// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_link_repository_v1

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

// NewCatalogCategoryLinkRepositoryV1SavePutParams creates a new CatalogCategoryLinkRepositoryV1SavePutParams object
// with the default values initialized.
func NewCatalogCategoryLinkRepositoryV1SavePutParams() *CatalogCategoryLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryLinkRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCategoryLinkRepositoryV1SavePutParamsWithTimeout creates a new CatalogCategoryLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogCategoryLinkRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CatalogCategoryLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryLinkRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCatalogCategoryLinkRepositoryV1SavePutParamsWithContext creates a new CatalogCategoryLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogCategoryLinkRepositoryV1SavePutParamsWithContext(ctx context.Context) *CatalogCategoryLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryLinkRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCatalogCategoryLinkRepositoryV1SavePutParamsWithHTTPClient creates a new CatalogCategoryLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogCategoryLinkRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CatalogCategoryLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogCategoryLinkRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CatalogCategoryLinkRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the catalog category link repository v1 save put operation typically these are written to a http.Request
*/
type CatalogCategoryLinkRepositoryV1SavePutParams struct {

	/*CatalogCategoryLinkRepositoryV1SavePutBody*/
	CatalogCategoryLinkRepositoryV1SavePutBody CatalogCategoryLinkRepositoryV1SavePutBody
	/*CategoryID*/
	CategoryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CatalogCategoryLinkRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WithContext(ctx context.Context) *CatalogCategoryLinkRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CatalogCategoryLinkRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogCategoryLinkRepositoryV1SavePutBody adds the catalogCategoryLinkRepositoryV1SavePutBody to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WithCatalogCategoryLinkRepositoryV1SavePutBody(catalogCategoryLinkRepositoryV1SavePutBody CatalogCategoryLinkRepositoryV1SavePutBody) *CatalogCategoryLinkRepositoryV1SavePutParams {
	o.SetCatalogCategoryLinkRepositoryV1SavePutBody(catalogCategoryLinkRepositoryV1SavePutBody)
	return o
}

// SetCatalogCategoryLinkRepositoryV1SavePutBody adds the catalogCategoryLinkRepositoryV1SavePutBody to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) SetCatalogCategoryLinkRepositoryV1SavePutBody(catalogCategoryLinkRepositoryV1SavePutBody CatalogCategoryLinkRepositoryV1SavePutBody) {
	o.CatalogCategoryLinkRepositoryV1SavePutBody = catalogCategoryLinkRepositoryV1SavePutBody
}

// WithCategoryID adds the categoryID to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WithCategoryID(categoryID string) *CatalogCategoryLinkRepositoryV1SavePutParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the catalog category link repository v1 save put params
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCategoryLinkRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogCategoryLinkRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
