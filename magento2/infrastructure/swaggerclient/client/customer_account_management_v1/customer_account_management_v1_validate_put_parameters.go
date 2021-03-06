// Code generated by go-swagger; DO NOT EDIT.

package customer_account_management_v1

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

// NewCustomerAccountManagementV1ValidatePutParams creates a new CustomerAccountManagementV1ValidatePutParams object
// with the default values initialized.
func NewCustomerAccountManagementV1ValidatePutParams() *CustomerAccountManagementV1ValidatePutParams {
	var ()
	return &CustomerAccountManagementV1ValidatePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAccountManagementV1ValidatePutParamsWithTimeout creates a new CustomerAccountManagementV1ValidatePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAccountManagementV1ValidatePutParamsWithTimeout(timeout time.Duration) *CustomerAccountManagementV1ValidatePutParams {
	var ()
	return &CustomerAccountManagementV1ValidatePutParams{

		timeout: timeout,
	}
}

// NewCustomerAccountManagementV1ValidatePutParamsWithContext creates a new CustomerAccountManagementV1ValidatePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAccountManagementV1ValidatePutParamsWithContext(ctx context.Context) *CustomerAccountManagementV1ValidatePutParams {
	var ()
	return &CustomerAccountManagementV1ValidatePutParams{

		Context: ctx,
	}
}

// NewCustomerAccountManagementV1ValidatePutParamsWithHTTPClient creates a new CustomerAccountManagementV1ValidatePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAccountManagementV1ValidatePutParamsWithHTTPClient(client *http.Client) *CustomerAccountManagementV1ValidatePutParams {
	var ()
	return &CustomerAccountManagementV1ValidatePutParams{
		HTTPClient: client,
	}
}

/*CustomerAccountManagementV1ValidatePutParams contains all the parameters to send to the API endpoint
for the customer account management v1 validate put operation typically these are written to a http.Request
*/
type CustomerAccountManagementV1ValidatePutParams struct {

	/*CustomerAccountManagementV1ValidatePutBody*/
	CustomerAccountManagementV1ValidatePutBody CustomerAccountManagementV1ValidatePutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) WithTimeout(timeout time.Duration) *CustomerAccountManagementV1ValidatePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) WithContext(ctx context.Context) *CustomerAccountManagementV1ValidatePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) WithHTTPClient(client *http.Client) *CustomerAccountManagementV1ValidatePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerAccountManagementV1ValidatePutBody adds the customerAccountManagementV1ValidatePutBody to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) WithCustomerAccountManagementV1ValidatePutBody(customerAccountManagementV1ValidatePutBody CustomerAccountManagementV1ValidatePutBody) *CustomerAccountManagementV1ValidatePutParams {
	o.SetCustomerAccountManagementV1ValidatePutBody(customerAccountManagementV1ValidatePutBody)
	return o
}

// SetCustomerAccountManagementV1ValidatePutBody adds the customerAccountManagementV1ValidatePutBody to the customer account management v1 validate put params
func (o *CustomerAccountManagementV1ValidatePutParams) SetCustomerAccountManagementV1ValidatePutBody(customerAccountManagementV1ValidatePutBody CustomerAccountManagementV1ValidatePutBody) {
	o.CustomerAccountManagementV1ValidatePutBody = customerAccountManagementV1ValidatePutBody
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAccountManagementV1ValidatePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CustomerAccountManagementV1ValidatePutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
