// Code generated by go-swagger; DO NOT EDIT.

package catalog_special_price_storage_v1

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

// NewCatalogSpecialPriceStorageV1DeletePostParams creates a new CatalogSpecialPriceStorageV1DeletePostParams object
// with the default values initialized.
func NewCatalogSpecialPriceStorageV1DeletePostParams() *CatalogSpecialPriceStorageV1DeletePostParams {
	var ()
	return &CatalogSpecialPriceStorageV1DeletePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogSpecialPriceStorageV1DeletePostParamsWithTimeout creates a new CatalogSpecialPriceStorageV1DeletePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogSpecialPriceStorageV1DeletePostParamsWithTimeout(timeout time.Duration) *CatalogSpecialPriceStorageV1DeletePostParams {
	var ()
	return &CatalogSpecialPriceStorageV1DeletePostParams{

		timeout: timeout,
	}
}

// NewCatalogSpecialPriceStorageV1DeletePostParamsWithContext creates a new CatalogSpecialPriceStorageV1DeletePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogSpecialPriceStorageV1DeletePostParamsWithContext(ctx context.Context) *CatalogSpecialPriceStorageV1DeletePostParams {
	var ()
	return &CatalogSpecialPriceStorageV1DeletePostParams{

		Context: ctx,
	}
}

// NewCatalogSpecialPriceStorageV1DeletePostParamsWithHTTPClient creates a new CatalogSpecialPriceStorageV1DeletePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogSpecialPriceStorageV1DeletePostParamsWithHTTPClient(client *http.Client) *CatalogSpecialPriceStorageV1DeletePostParams {
	var ()
	return &CatalogSpecialPriceStorageV1DeletePostParams{
		HTTPClient: client,
	}
}

/*CatalogSpecialPriceStorageV1DeletePostParams contains all the parameters to send to the API endpoint
for the catalog special price storage v1 delete post operation typically these are written to a http.Request
*/
type CatalogSpecialPriceStorageV1DeletePostParams struct {

	/*CatalogSpecialPriceStorageV1DeletePostBody*/
	CatalogSpecialPriceStorageV1DeletePostBody CatalogSpecialPriceStorageV1DeletePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) WithTimeout(timeout time.Duration) *CatalogSpecialPriceStorageV1DeletePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) WithContext(ctx context.Context) *CatalogSpecialPriceStorageV1DeletePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) WithHTTPClient(client *http.Client) *CatalogSpecialPriceStorageV1DeletePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogSpecialPriceStorageV1DeletePostBody adds the catalogSpecialPriceStorageV1DeletePostBody to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) WithCatalogSpecialPriceStorageV1DeletePostBody(catalogSpecialPriceStorageV1DeletePostBody CatalogSpecialPriceStorageV1DeletePostBody) *CatalogSpecialPriceStorageV1DeletePostParams {
	o.SetCatalogSpecialPriceStorageV1DeletePostBody(catalogSpecialPriceStorageV1DeletePostBody)
	return o
}

// SetCatalogSpecialPriceStorageV1DeletePostBody adds the catalogSpecialPriceStorageV1DeletePostBody to the catalog special price storage v1 delete post params
func (o *CatalogSpecialPriceStorageV1DeletePostParams) SetCatalogSpecialPriceStorageV1DeletePostBody(catalogSpecialPriceStorageV1DeletePostBody CatalogSpecialPriceStorageV1DeletePostBody) {
	o.CatalogSpecialPriceStorageV1DeletePostBody = catalogSpecialPriceStorageV1DeletePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogSpecialPriceStorageV1DeletePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogSpecialPriceStorageV1DeletePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
