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

// NewCustomerAccountManagementV1ChangePasswordByIDPutParams creates a new CustomerAccountManagementV1ChangePasswordByIDPutParams object
// with the default values initialized.
func NewCustomerAccountManagementV1ChangePasswordByIDPutParams() *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ChangePasswordByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithTimeout creates a new CustomerAccountManagementV1ChangePasswordByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithTimeout(timeout time.Duration) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ChangePasswordByIDPutParams{

		timeout: timeout,
	}
}

// NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithContext creates a new CustomerAccountManagementV1ChangePasswordByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithContext(ctx context.Context) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ChangePasswordByIDPutParams{

		Context: ctx,
	}
}

// NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithHTTPClient creates a new CustomerAccountManagementV1ChangePasswordByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerAccountManagementV1ChangePasswordByIDPutParamsWithHTTPClient(client *http.Client) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	var ()
	return &CustomerAccountManagementV1ChangePasswordByIDPutParams{
		HTTPClient: client,
	}
}

/*CustomerAccountManagementV1ChangePasswordByIDPutParams contains all the parameters to send to the API endpoint
for the customer account management v1 change password by Id put operation typically these are written to a http.Request
*/
type CustomerAccountManagementV1ChangePasswordByIDPutParams struct {

	/*CustomerAccountManagementV1ChangePasswordByIDPutBody*/
	CustomerAccountManagementV1ChangePasswordByIDPutBody CustomerAccountManagementV1ChangePasswordByIDPutBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) WithTimeout(timeout time.Duration) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) WithContext(ctx context.Context) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) WithHTTPClient(client *http.Client) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerAccountManagementV1ChangePasswordByIDPutBody adds the customerAccountManagementV1ChangePasswordByIDPutBody to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) WithCustomerAccountManagementV1ChangePasswordByIDPutBody(customerAccountManagementV1ChangePasswordByIDPutBody CustomerAccountManagementV1ChangePasswordByIDPutBody) *CustomerAccountManagementV1ChangePasswordByIDPutParams {
	o.SetCustomerAccountManagementV1ChangePasswordByIDPutBody(customerAccountManagementV1ChangePasswordByIDPutBody)
	return o
}

// SetCustomerAccountManagementV1ChangePasswordByIDPutBody adds the customerAccountManagementV1ChangePasswordByIdPutBody to the customer account management v1 change password by Id put params
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) SetCustomerAccountManagementV1ChangePasswordByIDPutBody(customerAccountManagementV1ChangePasswordByIDPutBody CustomerAccountManagementV1ChangePasswordByIDPutBody) {
	o.CustomerAccountManagementV1ChangePasswordByIDPutBody = customerAccountManagementV1ChangePasswordByIDPutBody
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerAccountManagementV1ChangePasswordByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CustomerAccountManagementV1ChangePasswordByIDPutBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}