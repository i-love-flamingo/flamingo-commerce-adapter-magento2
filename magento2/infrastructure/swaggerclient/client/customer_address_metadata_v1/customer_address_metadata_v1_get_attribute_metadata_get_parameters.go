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

// NewCustomerAddressMetadataV1GetAttributeMetadataGetParams creates a new CustomerAddressMetadataV1GetAttributeMetadataGetParams object
// with the default values initialized.
func NewCustomerAddressMetadataV1GetAttributeMetadataGetParams() *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributeMetadataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithTimeout creates a new CustomerAddressMetadataV1GetAttributeMetadataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributeMetadataGetParams{

		timeout: timeout,
	}
}

// NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithContext creates a new CustomerAddressMetadataV1GetAttributeMetadataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithContext(ctx context.Context) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributeMetadataGetParams{

		Context: ctx,
	}
}

// NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithHTTPClient creates a new CustomerAddressMetadataV1GetAttributeMetadataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAddressMetadataV1GetAttributeMetadataGetParamsWithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributeMetadataGetParams{
		HTTPClient: client,
	}
}

/*CustomerAddressMetadataV1GetAttributeMetadataGetParams contains all the parameters to send to the API endpoint
for the customer address metadata v1 get attribute metadata get operation typically these are written to a http.Request
*/
type CustomerAddressMetadataV1GetAttributeMetadataGetParams struct {

	/*AttributeCode*/
	AttributeCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) WithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) WithContext(ctx context.Context) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) WithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCode adds the attributeCode to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) WithAttributeCode(attributeCode string) *CustomerAddressMetadataV1GetAttributeMetadataGetParams {
	o.SetAttributeCode(attributeCode)
	return o
}

// SetAttributeCode adds the attributeCode to the customer address metadata v1 get attribute metadata get params
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) SetAttributeCode(attributeCode string) {
	o.AttributeCode = attributeCode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAddressMetadataV1GetAttributeMetadataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeCode
	if err := r.SetPathParam("attributeCode", o.AttributeCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
