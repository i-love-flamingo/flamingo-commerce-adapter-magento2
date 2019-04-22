// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_attributes_management_v1

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

// NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams creates a new RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams object
// with the default values initialized.
func NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams() *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {

	return &RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithTimeout creates a new RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithTimeout(timeout time.Duration) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {

	return &RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams{

		timeout: timeout,
	}
}

// NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithContext creates a new RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithContext(ctx context.Context) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {

	return &RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams{

		Context: ctx,
	}
}

// NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithHTTPClient creates a new RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRmaRmaAttributesManagementV1GetAllAttributesMetadataGetParamsWithHTTPClient(client *http.Client) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {

	return &RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams{
		HTTPClient: client,
	}
}

/*RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams contains all the parameters to send to the API endpoint
for the rma rma attributes management v1 get all attributes metadata get operation typically these are written to a http.Request
*/
type RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) WithTimeout(timeout time.Duration) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) WithContext(ctx context.Context) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) WithHTTPClient(client *http.Client) *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rma rma attributes management v1 get all attributes metadata get params
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RmaRmaAttributesManagementV1GetAllAttributesMetadataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}