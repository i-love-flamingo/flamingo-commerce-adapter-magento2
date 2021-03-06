// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_link_repository_v1

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

// NewCatalogProductLinkRepositoryV1SavePutParams creates a new CatalogProductLinkRepositoryV1SavePutParams object
// with the default values initialized.
func NewCatalogProductLinkRepositoryV1SavePutParams() *CatalogProductLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogProductLinkRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductLinkRepositoryV1SavePutParamsWithTimeout creates a new CatalogProductLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductLinkRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CatalogProductLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogProductLinkRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCatalogProductLinkRepositoryV1SavePutParamsWithContext creates a new CatalogProductLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductLinkRepositoryV1SavePutParamsWithContext(ctx context.Context) *CatalogProductLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogProductLinkRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCatalogProductLinkRepositoryV1SavePutParamsWithHTTPClient creates a new CatalogProductLinkRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductLinkRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CatalogProductLinkRepositoryV1SavePutParams {
	var ()
	return &CatalogProductLinkRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CatalogProductLinkRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the catalog product link repository v1 save put operation typically these are written to a http.Request
*/
type CatalogProductLinkRepositoryV1SavePutParams struct {

	/*CatalogProductLinkRepositoryV1SavePutBody*/
	CatalogProductLinkRepositoryV1SavePutBody CatalogProductLinkRepositoryV1SavePutBody
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CatalogProductLinkRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) WithContext(ctx context.Context) *CatalogProductLinkRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CatalogProductLinkRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogProductLinkRepositoryV1SavePutBody adds the catalogProductLinkRepositoryV1SavePutBody to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) WithCatalogProductLinkRepositoryV1SavePutBody(catalogProductLinkRepositoryV1SavePutBody CatalogProductLinkRepositoryV1SavePutBody) *CatalogProductLinkRepositoryV1SavePutParams {
	o.SetCatalogProductLinkRepositoryV1SavePutBody(catalogProductLinkRepositoryV1SavePutBody)
	return o
}

// SetCatalogProductLinkRepositoryV1SavePutBody adds the catalogProductLinkRepositoryV1SavePutBody to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) SetCatalogProductLinkRepositoryV1SavePutBody(catalogProductLinkRepositoryV1SavePutBody CatalogProductLinkRepositoryV1SavePutBody) {
	o.CatalogProductLinkRepositoryV1SavePutBody = catalogProductLinkRepositoryV1SavePutBody
}

// WithSku adds the sku to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) WithSku(sku string) *CatalogProductLinkRepositoryV1SavePutParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the catalog product link repository v1 save put params
func (o *CatalogProductLinkRepositoryV1SavePutParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductLinkRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogProductLinkRepositoryV1SavePutBody); err != nil {
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
