// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_link_management_v1

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

// NewCatalogProductLinkManagementV1SetProductLinksPostParams creates a new CatalogProductLinkManagementV1SetProductLinksPostParams object
// with the default values initialized.
func NewCatalogProductLinkManagementV1SetProductLinksPostParams() *CatalogProductLinkManagementV1SetProductLinksPostParams {
	var ()
	return &CatalogProductLinkManagementV1SetProductLinksPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithTimeout creates a new CatalogProductLinkManagementV1SetProductLinksPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithTimeout(timeout time.Duration) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	var ()
	return &CatalogProductLinkManagementV1SetProductLinksPostParams{

		timeout: timeout,
	}
}

// NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithContext creates a new CatalogProductLinkManagementV1SetProductLinksPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithContext(ctx context.Context) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	var ()
	return &CatalogProductLinkManagementV1SetProductLinksPostParams{

		Context: ctx,
	}
}

// NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithHTTPClient creates a new CatalogProductLinkManagementV1SetProductLinksPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductLinkManagementV1SetProductLinksPostParamsWithHTTPClient(client *http.Client) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	var ()
	return &CatalogProductLinkManagementV1SetProductLinksPostParams{
		HTTPClient: client,
	}
}

/*CatalogProductLinkManagementV1SetProductLinksPostParams contains all the parameters to send to the API endpoint
for the catalog product link management v1 set product links post operation typically these are written to a http.Request
*/
type CatalogProductLinkManagementV1SetProductLinksPostParams struct {

	/*CatalogProductLinkManagementV1SetProductLinksPostBody*/
	CatalogProductLinkManagementV1SetProductLinksPostBody CatalogProductLinkManagementV1SetProductLinksPostBody
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WithTimeout(timeout time.Duration) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WithContext(ctx context.Context) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WithHTTPClient(client *http.Client) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogProductLinkManagementV1SetProductLinksPostBody adds the catalogProductLinkManagementV1SetProductLinksPostBody to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WithCatalogProductLinkManagementV1SetProductLinksPostBody(catalogProductLinkManagementV1SetProductLinksPostBody CatalogProductLinkManagementV1SetProductLinksPostBody) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	o.SetCatalogProductLinkManagementV1SetProductLinksPostBody(catalogProductLinkManagementV1SetProductLinksPostBody)
	return o
}

// SetCatalogProductLinkManagementV1SetProductLinksPostBody adds the catalogProductLinkManagementV1SetProductLinksPostBody to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) SetCatalogProductLinkManagementV1SetProductLinksPostBody(catalogProductLinkManagementV1SetProductLinksPostBody CatalogProductLinkManagementV1SetProductLinksPostBody) {
	o.CatalogProductLinkManagementV1SetProductLinksPostBody = catalogProductLinkManagementV1SetProductLinksPostBody
}

// WithSku adds the sku to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WithSku(sku string) *CatalogProductLinkManagementV1SetProductLinksPostParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the catalog product link management v1 set product links post params
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductLinkManagementV1SetProductLinksPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogProductLinkManagementV1SetProductLinksPostBody); err != nil {
		return err
	}

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
