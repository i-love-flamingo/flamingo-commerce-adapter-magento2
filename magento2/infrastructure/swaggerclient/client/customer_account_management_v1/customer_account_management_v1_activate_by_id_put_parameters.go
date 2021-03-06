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

// NewCustomerAccountManagementV1ActivateByIDPutParams creates a new CustomerAccountManagementV1ActivateByIDPutParams object
// with the default values initialized.
func NewCustomerAccountManagementV1ActivateByIDPutParams() *CustomerAccountManagementV1ActivateByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ActivateByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAccountManagementV1ActivateByIDPutParamsWithTimeout creates a new CustomerAccountManagementV1ActivateByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAccountManagementV1ActivateByIDPutParamsWithTimeout(timeout time.Duration) *CustomerAccountManagementV1ActivateByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ActivateByIDPutParams{

		timeout: timeout,
	}
}

// NewCustomerAccountManagementV1ActivateByIDPutParamsWithContext creates a new CustomerAccountManagementV1ActivateByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAccountManagementV1ActivateByIDPutParamsWithContext(ctx context.Context) *CustomerAccountManagementV1ActivateByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ActivateByIDPutParams{

		Context: ctx,
	}
}

// NewCustomerAccountManagementV1ActivateByIDPutParamsWithHTTPClient creates a new CustomerAccountManagementV1ActivateByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAccountManagementV1ActivateByIDPutParamsWithHTTPClient(client *http.Client) *CustomerAccountManagementV1ActivateByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ActivateByIDPutParams{
		HTTPClient: client,
	}
}

/*CustomerAccountManagementV1ActivateByIDPutParams contains all the parameters to send to the API endpoint
for the customer account management v1 activate by Id put operation typically these are written to a http.Request
*/
type CustomerAccountManagementV1ActivateByIDPutParams struct {

	/*CustomerAccountManagementV1ActivateByIDPutBody*/
	CustomerAccountManagementV1ActivateByIDPutBody CustomerAccountManagementV1ActivateByIDPutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) WithTimeout(timeout time.Duration) *CustomerAccountManagementV1ActivateByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) WithContext(ctx context.Context) *CustomerAccountManagementV1ActivateByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) WithHTTPClient(client *http.Client) *CustomerAccountManagementV1ActivateByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerAccountManagementV1ActivateByIDPutBody adds the customerAccountManagementV1ActivateByIDPutBody to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) WithCustomerAccountManagementV1ActivateByIDPutBody(customerAccountManagementV1ActivateByIDPutBody CustomerAccountManagementV1ActivateByIDPutBody) *CustomerAccountManagementV1ActivateByIDPutParams {
	o.SetCustomerAccountManagementV1ActivateByIDPutBody(customerAccountManagementV1ActivateByIDPutBody)
	return o
}

// SetCustomerAccountManagementV1ActivateByIDPutBody adds the customerAccountManagementV1ActivateByIdPutBody to the customer account management v1 activate by Id put params
func (o *CustomerAccountManagementV1ActivateByIDPutParams) SetCustomerAccountManagementV1ActivateByIDPutBody(customerAccountManagementV1ActivateByIDPutBody CustomerAccountManagementV1ActivateByIDPutBody) {
	o.CustomerAccountManagementV1ActivateByIDPutBody = customerAccountManagementV1ActivateByIDPutBody
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAccountManagementV1ActivateByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CustomerAccountManagementV1ActivateByIDPutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
