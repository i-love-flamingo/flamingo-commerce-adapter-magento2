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

// NewCatalogCategoryManagementV1GetTreeGetParams creates a new CatalogCategoryManagementV1GetTreeGetParams object
// with the default values initialized.
func NewCatalogCategoryManagementV1GetTreeGetParams() *CatalogCategoryManagementV1GetTreeGetParams {
	var ()
	return &CatalogCategoryManagementV1GetTreeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCategoryManagementV1GetTreeGetParamsWithTimeout creates a new CatalogCategoryManagementV1GetTreeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogCategoryManagementV1GetTreeGetParamsWithTimeout(timeout time.Duration) *CatalogCategoryManagementV1GetTreeGetParams {
	var ()
	return &CatalogCategoryManagementV1GetTreeGetParams{

		timeout: timeout,
	}
}

// NewCatalogCategoryManagementV1GetTreeGetParamsWithContext creates a new CatalogCategoryManagementV1GetTreeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogCategoryManagementV1GetTreeGetParamsWithContext(ctx context.Context) *CatalogCategoryManagementV1GetTreeGetParams {
	var ()
	return &CatalogCategoryManagementV1GetTreeGetParams{

		Context: ctx,
	}
}

// NewCatalogCategoryManagementV1GetTreeGetParamsWithHTTPClient creates a new CatalogCategoryManagementV1GetTreeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogCategoryManagementV1GetTreeGetParamsWithHTTPClient(client *http.Client) *CatalogCategoryManagementV1GetTreeGetParams {
	var ()
	return &CatalogCategoryManagementV1GetTreeGetParams{
		HTTPClient: client,
	}
}

/*CatalogCategoryManagementV1GetTreeGetParams contains all the parameters to send to the API endpoint
for the catalog category management v1 get tree get operation typically these are written to a http.Request
*/
type CatalogCategoryManagementV1GetTreeGetParams struct {

	/*Depth*/
	Depth *int64
	/*RootCategoryID*/
	RootCategoryID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) WithTimeout(timeout time.Duration) *CatalogCategoryManagementV1GetTreeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) WithContext(ctx context.Context) *CatalogCategoryManagementV1GetTreeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) WithHTTPClient(client *http.Client) *CatalogCategoryManagementV1GetTreeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDepth adds the depth to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) WithDepth(depth *int64) *CatalogCategoryManagementV1GetTreeGetParams {
	o.SetDepth(depth)
	return o
}

// SetDepth adds the depth to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) SetDepth(depth *int64) {
	o.Depth = depth
}

// WithRootCategoryID adds the rootCategoryID to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) WithRootCategoryID(rootCategoryID *int64) *CatalogCategoryManagementV1GetTreeGetParams {
	o.SetRootCategoryID(rootCategoryID)
	return o
}

// SetRootCategoryID adds the rootCategoryId to the catalog category management v1 get tree get params
func (o *CatalogCategoryManagementV1GetTreeGetParams) SetRootCategoryID(rootCategoryID *int64) {
	o.RootCategoryID = rootCategoryID
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCategoryManagementV1GetTreeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Depth != nil {

		// query param depth
		var qrDepth int64
		if o.Depth != nil {
			qrDepth = *o.Depth
		}
		qDepth := swag.FormatInt64(qrDepth)
		if qDepth != "" {
			if err := r.SetQueryParam("depth", qDepth); err != nil {
				return err
			}
		}

	}

	if o.RootCategoryID != nil {

		// query param rootCategoryId
		var qrRootCategoryID int64
		if o.RootCategoryID != nil {
			qrRootCategoryID = *o.RootCategoryID
		}
		qRootCategoryID := swag.FormatInt64(qrRootCategoryID)
		if qRootCategoryID != "" {
			if err := r.SetQueryParam("rootCategoryId", qRootCategoryID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
