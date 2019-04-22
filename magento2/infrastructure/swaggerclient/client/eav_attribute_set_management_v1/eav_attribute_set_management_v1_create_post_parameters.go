// Code generated by go-swagger; DO NOT EDIT.

package eav_attribute_set_management_v1

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

// NewEavAttributeSetManagementV1CreatePostParams creates a new EavAttributeSetManagementV1CreatePostParams object
// with the default values initialized.
func NewEavAttributeSetManagementV1CreatePostParams() *EavAttributeSetManagementV1CreatePostParams {
	var ()
	return &EavAttributeSetManagementV1CreatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEavAttributeSetManagementV1CreatePostParamsWithTimeout creates a new EavAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEavAttributeSetManagementV1CreatePostParamsWithTimeout(timeout time.Duration) *EavAttributeSetManagementV1CreatePostParams {
	var ()
	return &EavAttributeSetManagementV1CreatePostParams{

		timeout: timeout,
	}
}

// NewEavAttributeSetManagementV1CreatePostParamsWithContext creates a new EavAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEavAttributeSetManagementV1CreatePostParamsWithContext(ctx context.Context) *EavAttributeSetManagementV1CreatePostParams {
	var ()
	return &EavAttributeSetManagementV1CreatePostParams{

		Context: ctx,
	}
}

// NewEavAttributeSetManagementV1CreatePostParamsWithHTTPClient creates a new EavAttributeSetManagementV1CreatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEavAttributeSetManagementV1CreatePostParamsWithHTTPClient(client *http.Client) *EavAttributeSetManagementV1CreatePostParams {
	var ()
	return &EavAttributeSetManagementV1CreatePostParams{
		HTTPClient: client,
	}
}

/*EavAttributeSetManagementV1CreatePostParams contains all the parameters to send to the API endpoint
for the eav attribute set management v1 create post operation typically these are written to a http.Request
*/
type EavAttributeSetManagementV1CreatePostParams struct {

	/*EavAttributeSetManagementV1CreatePostBody*/
	EavAttributeSetManagementV1CreatePostBody EavAttributeSetManagementV1CreatePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) WithTimeout(timeout time.Duration) *EavAttributeSetManagementV1CreatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) WithContext(ctx context.Context) *EavAttributeSetManagementV1CreatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) WithHTTPClient(client *http.Client) *EavAttributeSetManagementV1CreatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEavAttributeSetManagementV1CreatePostBody adds the eavAttributeSetManagementV1CreatePostBody to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) WithEavAttributeSetManagementV1CreatePostBody(eavAttributeSetManagementV1CreatePostBody EavAttributeSetManagementV1CreatePostBody) *EavAttributeSetManagementV1CreatePostParams {
	o.SetEavAttributeSetManagementV1CreatePostBody(eavAttributeSetManagementV1CreatePostBody)
	return o
}

// SetEavAttributeSetManagementV1CreatePostBody adds the eavAttributeSetManagementV1CreatePostBody to the eav attribute set management v1 create post params
func (o *EavAttributeSetManagementV1CreatePostParams) SetEavAttributeSetManagementV1CreatePostBody(eavAttributeSetManagementV1CreatePostBody EavAttributeSetManagementV1CreatePostBody) {
	o.EavAttributeSetManagementV1CreatePostBody = eavAttributeSetManagementV1CreatePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *EavAttributeSetManagementV1CreatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.EavAttributeSetManagementV1CreatePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}