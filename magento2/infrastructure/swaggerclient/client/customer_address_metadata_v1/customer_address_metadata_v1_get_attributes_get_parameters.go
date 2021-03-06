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

// NewCustomerAddressMetadataV1GetAttributesGetParams creates a new CustomerAddressMetadataV1GetAttributesGetParams object
// with the default values initialized.
func NewCustomerAddressMetadataV1GetAttributesGetParams() *CustomerAddressMetadataV1GetAttributesGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAddressMetadataV1GetAttributesGetParamsWithTimeout creates a new CustomerAddressMetadataV1GetAttributesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAddressMetadataV1GetAttributesGetParamsWithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAttributesGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributesGetParams{

		timeout: timeout,
	}
}

// NewCustomerAddressMetadataV1GetAttributesGetParamsWithContext creates a new CustomerAddressMetadataV1GetAttributesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAddressMetadataV1GetAttributesGetParamsWithContext(ctx context.Context) *CustomerAddressMetadataV1GetAttributesGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributesGetParams{

		Context: ctx,
	}
}

// NewCustomerAddressMetadataV1GetAttributesGetParamsWithHTTPClient creates a new CustomerAddressMetadataV1GetAttributesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAddressMetadataV1GetAttributesGetParamsWithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAttributesGetParams {
	var ()
	return &CustomerAddressMetadataV1GetAttributesGetParams{
		HTTPClient: client,
	}
}

/*CustomerAddressMetadataV1GetAttributesGetParams contains all the parameters to send to the API endpoint
for the customer address metadata v1 get attributes get operation typically these are written to a http.Request
*/
type CustomerAddressMetadataV1GetAttributesGetParams struct {

	/*FormCode*/
	FormCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) WithTimeout(timeout time.Duration) *CustomerAddressMetadataV1GetAttributesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) WithContext(ctx context.Context) *CustomerAddressMetadataV1GetAttributesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) WithHTTPClient(client *http.Client) *CustomerAddressMetadataV1GetAttributesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormCode adds the formCode to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) WithFormCode(formCode string) *CustomerAddressMetadataV1GetAttributesGetParams {
	o.SetFormCode(formCode)
	return o
}

// SetFormCode adds the formCode to the customer address metadata v1 get attributes get params
func (o *CustomerAddressMetadataV1GetAttributesGetParams) SetFormCode(formCode string) {
	o.FormCode = formCode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAddressMetadataV1GetAttributesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param formCode
	if err := r.SetPathParam("formCode", o.FormCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
