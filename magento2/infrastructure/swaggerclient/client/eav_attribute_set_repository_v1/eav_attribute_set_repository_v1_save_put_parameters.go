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

	strfmt "github.com/go-openapi/strfmt"
)

// NewEavAttributeSetRepositoryV1SavePutParams creates a new EavAttributeSetRepositoryV1SavePutParams object
// with the default values initialized.
func NewEavAttributeSetRepositoryV1SavePutParams() *EavAttributeSetRepositoryV1SavePutParams {
	var ()
	return &EavAttributeSetRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEavAttributeSetRepositoryV1SavePutParamsWithTimeout creates a new EavAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEavAttributeSetRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *EavAttributeSetRepositoryV1SavePutParams {
	var ()
	return &EavAttributeSetRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewEavAttributeSetRepositoryV1SavePutParamsWithContext creates a new EavAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewEavAttributeSetRepositoryV1SavePutParamsWithContext(ctx context.Context) *EavAttributeSetRepositoryV1SavePutParams {
	var ()
	return &EavAttributeSetRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewEavAttributeSetRepositoryV1SavePutParamsWithHTTPClient creates a new EavAttributeSetRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEavAttributeSetRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *EavAttributeSetRepositoryV1SavePutParams {
	var ()
	return &EavAttributeSetRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*EavAttributeSetRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the eav attribute set repository v1 save put operation typically these are written to a http.Request
*/
type EavAttributeSetRepositoryV1SavePutParams struct {

	/*AttributeSetID*/
	AttributeSetID string
	/*EavAttributeSetRepositoryV1SavePutBody*/
	EavAttributeSetRepositoryV1SavePutBody EavAttributeSetRepositoryV1SavePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *EavAttributeSetRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) WithContext(ctx context.Context) *EavAttributeSetRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *EavAttributeSetRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeSetID adds the attributeSetID to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) WithAttributeSetID(attributeSetID string) *EavAttributeSetRepositoryV1SavePutParams {
	o.SetAttributeSetID(attributeSetID)
	return o
}

// SetAttributeSetID adds the attributeSetId to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) SetAttributeSetID(attributeSetID string) {
	o.AttributeSetID = attributeSetID
}

// WithEavAttributeSetRepositoryV1SavePutBody adds the eavAttributeSetRepositoryV1SavePutBody to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) WithEavAttributeSetRepositoryV1SavePutBody(eavAttributeSetRepositoryV1SavePutBody EavAttributeSetRepositoryV1SavePutBody) *EavAttributeSetRepositoryV1SavePutParams {
	o.SetEavAttributeSetRepositoryV1SavePutBody(eavAttributeSetRepositoryV1SavePutBody)
	return o
}

// SetEavAttributeSetRepositoryV1SavePutBody adds the eavAttributeSetRepositoryV1SavePutBody to the eav attribute set repository v1 save put params
func (o *EavAttributeSetRepositoryV1SavePutParams) SetEavAttributeSetRepositoryV1SavePutBody(eavAttributeSetRepositoryV1SavePutBody EavAttributeSetRepositoryV1SavePutBody) {
	o.EavAttributeSetRepositoryV1SavePutBody = eavAttributeSetRepositoryV1SavePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *EavAttributeSetRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeSetId
	if err := r.SetPathParam("attributeSetId", o.AttributeSetID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.EavAttributeSetRepositoryV1SavePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
