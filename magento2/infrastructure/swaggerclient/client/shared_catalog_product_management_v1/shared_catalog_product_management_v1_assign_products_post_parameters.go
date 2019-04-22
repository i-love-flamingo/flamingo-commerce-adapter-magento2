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

// NewSharedCatalogProductManagementV1AssignProductsPostParams creates a new SharedCatalogProductManagementV1AssignProductsPostParams object
// with the default values initialized.
func NewSharedCatalogProductManagementV1AssignProductsPostParams() *SharedCatalogProductManagementV1AssignProductsPostParams {
	var ()
	return &SharedCatalogProductManagementV1AssignProductsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSharedCatalogProductManagementV1AssignProductsPostParamsWithTimeout creates a new SharedCatalogProductManagementV1AssignProductsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSharedCatalogProductManagementV1AssignProductsPostParamsWithTimeout(timeout time.Duration) *SharedCatalogProductManagementV1AssignProductsPostParams {
	var ()
	return &SharedCatalogProductManagementV1AssignProductsPostParams{

		timeout: timeout,
	}
}

// NewSharedCatalogProductManagementV1AssignProductsPostParamsWithContext creates a new SharedCatalogProductManagementV1AssignProductsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSharedCatalogProductManagementV1AssignProductsPostParamsWithContext(ctx context.Context) *SharedCatalogProductManagementV1AssignProductsPostParams {
	var ()
	return &SharedCatalogProductManagementV1AssignProductsPostParams{

		Context: ctx,
	}
}

// NewSharedCatalogProductManagementV1AssignProductsPostParamsWithHTTPClient creates a new SharedCatalogProductManagementV1AssignProductsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSharedCatalogProductManagementV1AssignProductsPostParamsWithHTTPClient(client *http.Client) *SharedCatalogProductManagementV1AssignProductsPostParams {
	var ()
	return &SharedCatalogProductManagementV1AssignProductsPostParams{
		HTTPClient: client,
	}
}

/*SharedCatalogProductManagementV1AssignProductsPostParams contains all the parameters to send to the API endpoint
for the shared catalog product management v1 assign products post operation typically these are written to a http.Request
*/
type SharedCatalogProductManagementV1AssignProductsPostParams struct {

	/*ID*/
	ID int64
	/*SharedCatalogProductManagementV1AssignProductsPostBody*/
	SharedCatalogProductManagementV1AssignProductsPostBody SharedCatalogProductManagementV1AssignProductsPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WithTimeout(timeout time.Duration) *SharedCatalogProductManagementV1AssignProductsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WithContext(ctx context.Context) *SharedCatalogProductManagementV1AssignProductsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WithHTTPClient(client *http.Client) *SharedCatalogProductManagementV1AssignProductsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WithID(id int64) *SharedCatalogProductManagementV1AssignProductsPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) SetID(id int64) {
	o.ID = id
}

// WithSharedCatalogProductManagementV1AssignProductsPostBody adds the sharedCatalogProductManagementV1AssignProductsPostBody to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WithSharedCatalogProductManagementV1AssignProductsPostBody(sharedCatalogProductManagementV1AssignProductsPostBody SharedCatalogProductManagementV1AssignProductsPostBody) *SharedCatalogProductManagementV1AssignProductsPostParams {
	o.SetSharedCatalogProductManagementV1AssignProductsPostBody(sharedCatalogProductManagementV1AssignProductsPostBody)
	return o
}

// SetSharedCatalogProductManagementV1AssignProductsPostBody adds the sharedCatalogProductManagementV1AssignProductsPostBody to the shared catalog product management v1 assign products post params
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) SetSharedCatalogProductManagementV1AssignProductsPostBody(sharedCatalogProductManagementV1AssignProductsPostBody SharedCatalogProductManagementV1AssignProductsPostBody) {
	o.SharedCatalogProductManagementV1AssignProductsPostBody = sharedCatalogProductManagementV1AssignProductsPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *SharedCatalogProductManagementV1AssignProductsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SharedCatalogProductManagementV1AssignProductsPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
