// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_link_type_list_v1

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

// NewCatalogProductLinkTypeListV1GetItemAttributesGetParams creates a new CatalogProductLinkTypeListV1GetItemAttributesGetParams object
// with the default values initialized.
func NewCatalogProductLinkTypeListV1GetItemAttributesGetParams() *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	var ()
	return &CatalogProductLinkTypeListV1GetItemAttributesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithTimeout creates a new CatalogProductLinkTypeListV1GetItemAttributesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithTimeout(timeout time.Duration) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	var ()
	return &CatalogProductLinkTypeListV1GetItemAttributesGetParams{

		timeout: timeout,
	}
}

// NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithContext creates a new CatalogProductLinkTypeListV1GetItemAttributesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithContext(ctx context.Context) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	var ()
	return &CatalogProductLinkTypeListV1GetItemAttributesGetParams{

		Context: ctx,
	}
}

// NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithHTTPClient creates a new CatalogProductLinkTypeListV1GetItemAttributesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogProductLinkTypeListV1GetItemAttributesGetParamsWithHTTPClient(client *http.Client) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	var ()
	return &CatalogProductLinkTypeListV1GetItemAttributesGetParams{
		HTTPClient: client,
	}
}

/*CatalogProductLinkTypeListV1GetItemAttributesGetParams contains all the parameters to send to the API endpoint
for the catalog product link type list v1 get item attributes get operation typically these are written to a http.Request
*/
type CatalogProductLinkTypeListV1GetItemAttributesGetParams struct {

	/*Type*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) WithTimeout(timeout time.Duration) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) WithContext(ctx context.Context) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) WithHTTPClient(client *http.Client) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) WithType(typeVar string) *CatalogProductLinkTypeListV1GetItemAttributesGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the catalog product link type list v1 get item attributes get params
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductLinkTypeListV1GetItemAttributesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
