// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_repository_v1

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

// NewCustomerCustomerRepositoryV1GetByIDGetParams creates a new CustomerCustomerRepositoryV1GetByIDGetParams object
// with the default values initialized.
func NewCustomerCustomerRepositoryV1GetByIDGetParams() *CustomerCustomerRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerCustomerRepositoryV1GetByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerCustomerRepositoryV1GetByIDGetParamsWithTimeout creates a new CustomerCustomerRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerCustomerRepositoryV1GetByIDGetParamsWithTimeout(timeout time.Duration) *CustomerCustomerRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerCustomerRepositoryV1GetByIDGetParams{

		timeout: timeout,
	}
}

// NewCustomerCustomerRepositoryV1GetByIDGetParamsWithContext creates a new CustomerCustomerRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerCustomerRepositoryV1GetByIDGetParamsWithContext(ctx context.Context) *CustomerCustomerRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerCustomerRepositoryV1GetByIDGetParams{

		Context: ctx,
	}
}

// NewCustomerCustomerRepositoryV1GetByIDGetParamsWithHTTPClient creates a new CustomerCustomerRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerCustomerRepositoryV1GetByIDGetParamsWithHTTPClient(client *http.Client) *CustomerCustomerRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerCustomerRepositoryV1GetByIDGetParams{
		HTTPClient: client,
	}
}

/*CustomerCustomerRepositoryV1GetByIDGetParams contains all the parameters to send to the API endpoint
for the customer customer repository v1 get by Id get operation typically these are written to a http.Request
*/
type CustomerCustomerRepositoryV1GetByIDGetParams struct {

	/*CustomerID*/
	CustomerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) WithTimeout(timeout time.Duration) *CustomerCustomerRepositoryV1GetByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) WithContext(ctx context.Context) *CustomerCustomerRepositoryV1GetByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) WithHTTPClient(client *http.Client) *CustomerCustomerRepositoryV1GetByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) WithCustomerID(customerID int64) *CustomerCustomerRepositoryV1GetByIDGetParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the customer customer repository v1 get by Id get params
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerCustomerRepositoryV1GetByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
