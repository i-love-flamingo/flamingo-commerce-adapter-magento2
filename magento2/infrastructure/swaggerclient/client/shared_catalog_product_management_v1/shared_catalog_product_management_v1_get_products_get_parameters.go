// Code generated by go-swagger; DO NOT EDIT.

package shared_catalog_product_management_v1

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

// NewSharedCatalogProductManagementV1GetProductsGetParams creates a new SharedCatalogProductManagementV1GetProductsGetParams object
// with the default values initialized.
func NewSharedCatalogProductManagementV1GetProductsGetParams() *SharedCatalogProductManagementV1GetProductsGetParams {
	var ()
	return &SharedCatalogProductManagementV1GetProductsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSharedCatalogProductManagementV1GetProductsGetParamsWithTimeout creates a new SharedCatalogProductManagementV1GetProductsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSharedCatalogProductManagementV1GetProductsGetParamsWithTimeout(timeout time.Duration) *SharedCatalogProductManagementV1GetProductsGetParams {
	var ()
	return &SharedCatalogProductManagementV1GetProductsGetParams{

		timeout: timeout,
	}
}

// NewSharedCatalogProductManagementV1GetProductsGetParamsWithContext creates a new SharedCatalogProductManagementV1GetProductsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSharedCatalogProductManagementV1GetProductsGetParamsWithContext(ctx context.Context) *SharedCatalogProductManagementV1GetProductsGetParams {
	var ()
	return &SharedCatalogProductManagementV1GetProductsGetParams{

		Context: ctx,
	}
}

// NewSharedCatalogProductManagementV1GetProductsGetParamsWithHTTPClient creates a new SharedCatalogProductManagementV1GetProductsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSharedCatalogProductManagementV1GetProductsGetParamsWithHTTPClient(client *http.Client) *SharedCatalogProductManagementV1GetProductsGetParams {
	var ()
	return &SharedCatalogProductManagementV1GetProductsGetParams{
		HTTPClient: client,
	}
}

/*SharedCatalogProductManagementV1GetProductsGetParams contains all the parameters to send to the API endpoint
for the shared catalog product management v1 get products get operation typically these are written to a http.Request
*/
type SharedCatalogProductManagementV1GetProductsGetParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) WithTimeout(timeout time.Duration) *SharedCatalogProductManagementV1GetProductsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) WithContext(ctx context.Context) *SharedCatalogProductManagementV1GetProductsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) WithHTTPClient(client *http.Client) *SharedCatalogProductManagementV1GetProductsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) WithID(id int64) *SharedCatalogProductManagementV1GetProductsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the shared catalog product management v1 get products get params
func (o *SharedCatalogProductManagementV1GetProductsGetParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SharedCatalogProductManagementV1GetProductsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
