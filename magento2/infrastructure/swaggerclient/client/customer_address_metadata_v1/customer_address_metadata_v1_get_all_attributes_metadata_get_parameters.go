// Code generated by go-swagger; DO NOT EDIT.

package customer_address_metadata_v1

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

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParams creates a new CustomerAddressMetadataV1GetAllAttributesMetadataGetParams object
// with the default values initialized.
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParams() *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {

	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithTimeout creates a new CustomerAddressMetadataV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {

	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetParams{

		timeout: timeout,
	}
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithContext creates a new CustomerAddressMetadataV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithContext(ctx context.Context) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {

	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetParams{

		Context: ctx,
	}
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithHTTPClient creates a new CustomerAddressMetadataV1GetAllAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetParamsWithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {

	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetParams{
		HTTPClient: client,
	}
}

/*CustomerAddressMetadataV1GetAllAttributesMetadataGetParams contains all the parameters to send to the API endpoint
for the customer address metadata v1 get all attributes metadata get operation typically these are written to a http.Request
*/
type CustomerAddressMetadataV1GetAllAttributesMetadataGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) WithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) WithContext(ctx context.Context) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) WithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer address metadata v1 get all attributes metadata get params
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}