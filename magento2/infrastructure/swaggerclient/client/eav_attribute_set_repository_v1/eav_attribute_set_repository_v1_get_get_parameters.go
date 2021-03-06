// Code generated by go-swagger; DO NOT EDIT.

package eav_attribute_set_repository_v1

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

// NewEavAttributeSetRepositoryV1GetGetParams creates a new EavAttributeSetRepositoryV1GetGetParams object
// with the default values initialized.
func NewEavAttributeSetRepositoryV1GetGetParams() *EavAttributeSetRepositoryV1GetGetParams {
	var ()
	return &EavAttributeSetRepositoryV1GetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEavAttributeSetRepositoryV1GetGetParamsWithTimeout creates a new EavAttributeSetRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEavAttributeSetRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *EavAttributeSetRepositoryV1GetGetParams {
	var ()
	return &EavAttributeSetRepositoryV1GetGetParams{

		timeout: timeout,
	}
}

// NewEavAttributeSetRepositoryV1GetGetParamsWithContext creates a new EavAttributeSetRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEavAttributeSetRepositoryV1GetGetParamsWithContext(ctx context.Context) *EavAttributeSetRepositoryV1GetGetParams {
	var ()
	return &EavAttributeSetRepositoryV1GetGetParams{

		Context: ctx,
	}
}

// NewEavAttributeSetRepositoryV1GetGetParamsWithHTTPClient creates a new EavAttributeSetRepositoryV1GetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEavAttributeSetRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *EavAttributeSetRepositoryV1GetGetParams {
	var ()
	return &EavAttributeSetRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/*EavAttributeSetRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
for the eav attribute set repository v1 get get operation typically these are written to a http.Request
*/
type EavAttributeSetRepositoryV1GetGetParams struct {

	/*AttributeSetID*/
	AttributeSetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *EavAttributeSetRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) WithContext(ctx context.Context) *EavAttributeSetRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *EavAttributeSetRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeSetID adds the attributeSetID to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) WithAttributeSetID(attributeSetID int64) *EavAttributeSetRepositoryV1GetGetParams {
	o.SetAttributeSetID(attributeSetID)
	return o
}

// SetAttributeSetID adds the attributeSetId to the eav attribute set repository v1 get get params
func (o *EavAttributeSetRepositoryV1GetGetParams) SetAttributeSetID(attributeSetID int64) {
	o.AttributeSetID = attributeSetID
}

// WriteToRequest writes these params to a swagger request
func (o *EavAttributeSetRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeSetId
	if err := r.SetPathParam("attributeSetId", swag.FormatInt64(o.AttributeSetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
