// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_metadata_v1

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

// NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams creates a new CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams object
// with the default values initialized.
func NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams() *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	var ()
	return &CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithTimeout creates a new CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithTimeout(timeout time.Duration) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	var ()
	return &CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams{

		timeout: timeout,
	}
}

// NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithContext creates a new CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithContext(ctx context.Context) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	var ()
	return &CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams{

		Context: ctx,
	}
}

// NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithHTTPClient creates a new CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerCustomerMetadataV1GetCustomAttributesMetadataGetParamsWithHTTPClient(client *http.Client) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	var ()
	return &CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams{
		HTTPClient: client,
	}
}

/*CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams contains all the parameters to send to the API endpoint
for the customer customer metadata v1 get custom attributes metadata get operation typically these are written to a http.Request
*/
type CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams struct {

	/*DataInterfaceName*/
	DataInterfaceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) WithTimeout(timeout time.Duration) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) WithContext(ctx context.Context) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) WithHTTPClient(client *http.Client) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataInterfaceName adds the dataInterfaceName to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) WithDataInterfaceName(dataInterfaceName *string) *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams {
	o.SetDataInterfaceName(dataInterfaceName)
	return o
}

// SetDataInterfaceName adds the dataInterfaceName to the customer customer metadata v1 get custom attributes metadata get params
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) SetDataInterfaceName(dataInterfaceName *string) {
	o.DataInterfaceName = dataInterfaceName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerCustomerMetadataV1GetCustomAttributesMetadataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataInterfaceName != nil {

		// query param dataInterfaceName
		var qrDataInterfaceName string
		if o.DataInterfaceName != nil {
			qrDataInterfaceName = *o.DataInterfaceName
		}
		qDataInterfaceName := qrDataInterfaceName
		if qDataInterfaceName != "" {
			if err := r.SetQueryParam("dataInterfaceName", qDataInterfaceName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}