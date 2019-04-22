// Code generated by go-swagger; DO NOT EDIT.

package catalog_attribute_set_management_v1

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

// NewCatalogAttributeSetManagementV1CreatePostParams creates a new CatalogAttributeSetManagementV1CreatePostParams object
// with the default values initialized.
func NewCatalogAttributeSetManagementV1CreatePostParams() *CatalogAttributeSetManagementV1CreatePostParams {
	var ()
	return &CatalogAttributeSetManagementV1CreatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogAttributeSetManagementV1CreatePostParamsWithTimeout creates a new CatalogAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogAttributeSetManagementV1CreatePostParamsWithTimeout(timeout time.Duration) *CatalogAttributeSetManagementV1CreatePostParams {
	var ()
	return &CatalogAttributeSetManagementV1CreatePostParams{

		timeout: timeout,
	}
}

// NewCatalogAttributeSetManagementV1CreatePostParamsWithContext creates a new CatalogAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogAttributeSetManagementV1CreatePostParamsWithContext(ctx context.Context) *CatalogAttributeSetManagementV1CreatePostParams {
	var ()
	return &CatalogAttributeSetManagementV1CreatePostParams{

		Context: ctx,
	}
}

// NewCatalogAttributeSetManagementV1CreatePostParamsWithHTTPClient creates a new CatalogAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogAttributeSetManagementV1CreatePostParamsWithHTTPClient(client *http.Client) *CatalogAttributeSetManagementV1CreatePostParams {
	var ()
	return &CatalogAttributeSetManagementV1CreatePostParams{
		HTTPClient: client,
	}
}

/*CatalogAttributeSetManagementV1CreatePostParams contains all the parameters to send to the API endpoint
for the catalog attribute set management v1 create post operation typically these are written to a http.Request
*/
type CatalogAttributeSetManagementV1CreatePostParams struct {

	/*CatalogAttributeSetManagementV1CreatePostBody*/
	CatalogAttributeSetManagementV1CreatePostBody CatalogAttributeSetManagementV1CreatePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) WithTimeout(timeout time.Duration) *CatalogAttributeSetManagementV1CreatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) WithContext(ctx context.Context) *CatalogAttributeSetManagementV1CreatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) WithHTTPClient(client *http.Client) *CatalogAttributeSetManagementV1CreatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogAttributeSetManagementV1CreatePostBody adds the catalogAttributeSetManagementV1CreatePostBody to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) WithCatalogAttributeSetManagementV1CreatePostBody(catalogAttributeSetManagementV1CreatePostBody CatalogAttributeSetManagementV1CreatePostBody) *CatalogAttributeSetManagementV1CreatePostParams {
	o.SetCatalogAttributeSetManagementV1CreatePostBody(catalogAttributeSetManagementV1CreatePostBody)
	return o
}

// SetCatalogAttributeSetManagementV1CreatePostBody adds the catalogAttributeSetManagementV1CreatePostBody to the catalog attribute set management v1 create post params
func (o *CatalogAttributeSetManagementV1CreatePostParams) SetCatalogAttributeSetManagementV1CreatePostBody(catalogAttributeSetManagementV1CreatePostBody CatalogAttributeSetManagementV1CreatePostBody) {
	o.CatalogAttributeSetManagementV1CreatePostBody = catalogAttributeSetManagementV1CreatePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogAttributeSetManagementV1CreatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogAttributeSetManagementV1CreatePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}