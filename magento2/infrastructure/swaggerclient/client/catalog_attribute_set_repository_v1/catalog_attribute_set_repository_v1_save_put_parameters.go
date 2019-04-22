// Code generated by go-swagger; DO NOT EDIT.

package catalog_attribute_set_repository_v1

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

// NewCatalogAttributeSetRepositoryV1SavePutParams creates a new CatalogAttributeSetRepositoryV1SavePutParams object
// with the default values initialized.
func NewCatalogAttributeSetRepositoryV1SavePutParams() *CatalogAttributeSetRepositoryV1SavePutParams {
	var ()
	return &CatalogAttributeSetRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogAttributeSetRepositoryV1SavePutParamsWithTimeout creates a new CatalogAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogAttributeSetRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CatalogAttributeSetRepositoryV1SavePutParams {
	var ()
	return &CatalogAttributeSetRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCatalogAttributeSetRepositoryV1SavePutParamsWithContext creates a new CatalogAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogAttributeSetRepositoryV1SavePutParamsWithContext(ctx context.Context) *CatalogAttributeSetRepositoryV1SavePutParams {
	var ()
	return &CatalogAttributeSetRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCatalogAttributeSetRepositoryV1SavePutParamsWithHTTPClient creates a new CatalogAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogAttributeSetRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CatalogAttributeSetRepositoryV1SavePutParams {
	var ()
	return &CatalogAttributeSetRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CatalogAttributeSetRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the catalog attribute set repository v1 save put operation typically these are written to a http.Request
*/
type CatalogAttributeSetRepositoryV1SavePutParams struct {

	/*AttributeSetID*/
	AttributeSetID string
	/*CatalogAttributeSetRepositoryV1SavePutBody*/
	CatalogAttributeSetRepositoryV1SavePutBody CatalogAttributeSetRepositoryV1SavePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CatalogAttributeSetRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WithContext(ctx context.Context) *CatalogAttributeSetRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CatalogAttributeSetRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeSetID adds the attributeSetID to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WithAttributeSetID(attributeSetID string) *CatalogAttributeSetRepositoryV1SavePutParams {
	o.SetAttributeSetID(attributeSetID)
	return o
}

// SetAttributeSetID adds the attributeSetId to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) SetAttributeSetID(attributeSetID string) {
	o.AttributeSetID = attributeSetID
}

// WithCatalogAttributeSetRepositoryV1SavePutBody adds the catalogAttributeSetRepositoryV1SavePutBody to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WithCatalogAttributeSetRepositoryV1SavePutBody(catalogAttributeSetRepositoryV1SavePutBody CatalogAttributeSetRepositoryV1SavePutBody) *CatalogAttributeSetRepositoryV1SavePutParams {
	o.SetCatalogAttributeSetRepositoryV1SavePutBody(catalogAttributeSetRepositoryV1SavePutBody)
	return o
}

// SetCatalogAttributeSetRepositoryV1SavePutBody adds the catalogAttributeSetRepositoryV1SavePutBody to the catalog attribute set repository v1 save put params
func (o *CatalogAttributeSetRepositoryV1SavePutParams) SetCatalogAttributeSetRepositoryV1SavePutBody(catalogAttributeSetRepositoryV1SavePutBody CatalogAttributeSetRepositoryV1SavePutBody) {
	o.CatalogAttributeSetRepositoryV1SavePutBody = catalogAttributeSetRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogAttributeSetRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeSetId
	if err := r.SetPathParam("attributeSetId", o.AttributeSetID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CatalogAttributeSetRepositoryV1SavePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
