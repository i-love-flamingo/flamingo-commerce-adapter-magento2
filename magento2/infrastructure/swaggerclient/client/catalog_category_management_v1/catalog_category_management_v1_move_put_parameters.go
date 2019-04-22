// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCatalogCategoryManagementV1MovePutParams creates a new CatalogCategoryManagementV1MovePutParams object
// with the default values initialized.
func NewCatalogCategoryManagementV1MovePutParams() *CatalogCategoryManagementV1MovePutParams {
	var ()
	return &CatalogCategoryManagementV1MovePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCategoryManagementV1MovePutParamsWithTimeout creates a new CatalogCategoryManagementV1MovePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogCategoryManagementV1MovePutParamsWithTimeout(timeout time.Duration) *CatalogCategoryManagementV1MovePutParams {
	var ()
	return &CatalogCategoryManagementV1MovePutParams{

		timeout: timeout,
	}
}

// NewCatalogCategoryManagementV1MovePutParamsWithContext creates a new CatalogCategoryManagementV1MovePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogCategoryManagementV1MovePutParamsWithContext(ctx context.Context) *CatalogCategoryManagementV1MovePutParams {
	var ()
	return &CatalogCategoryManagementV1MovePutParams{

		Context: ctx,
	}
}

// NewCatalogCategoryManagementV1MovePutParamsWithHTTPClient creates a new CatalogCategoryManagementV1MovePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogCategoryManagementV1MovePutParamsWithHTTPClient(client *http.Client) *CatalogCategoryManagementV1MovePutParams {
	var ()
	return &CatalogCategoryManagementV1MovePutParams{
		HTTPClient: client,
	}
}

/*CatalogCategoryManagementV1MovePutParams contains all the parameters to send to the API endpoint
for the catalog category management v1 move put operation typically these are written to a http.Request
*/
type CatalogCategoryManagementV1MovePutParams struct {

	/*CatalogCategoryManagementV1MovePutBody*/
	CatalogCategoryManagementV1MovePutBody CatalogCategoryManagementV1MovePutBody
	/*CategoryID*/
	CategoryID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) WithTimeout(timeout time.Duration) *CatalogCategoryManagementV1MovePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) WithContext(ctx context.Context) *CatalogCategoryManagementV1MovePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) WithHTTPClient(client *http.Client) *CatalogCategoryManagementV1MovePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogCategoryManagementV1MovePutBody adds the catalogCategoryManagementV1MovePutBody to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) WithCatalogCategoryManagementV1MovePutBody(catalogCategoryManagementV1MovePutBody CatalogCategoryManagementV1MovePutBody) *CatalogCategoryManagementV1MovePutParams {
	o.SetCatalogCategoryManagementV1MovePutBody(catalogCategoryManagementV1MovePutBody)
	return o
}

// SetCatalogCategoryManagementV1MovePutBody adds the catalogCategoryManagementV1MovePutBody to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) SetCatalogCategoryManagementV1MovePutBody(catalogCategoryManagementV1MovePutBody CatalogCategoryManagementV1MovePutBody) {
	o.CatalogCategoryManagementV1MovePutBody = catalogCategoryManagementV1MovePutBody
}

// WithCategoryID adds the categoryID to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) WithCategoryID(categoryID int64) *CatalogCategoryManagementV1MovePutParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the catalog category management v1 move put params
func (o *CatalogCategoryManagementV1MovePutParams) SetCategoryID(categoryID int64) {
	o.CategoryID = categoryID
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCategoryManagementV1MovePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogCategoryManagementV1MovePutBody); err != nil {
		return err
	}

	// path param categoryId
	if err := r.SetPathParam("categoryId", swag.FormatInt64(o.CategoryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
