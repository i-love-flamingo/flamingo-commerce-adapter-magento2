// Code generated by go-swagger; DO NOT EDIT.

package customer_group_repository_v1

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

// NewCustomerGroupRepositoryV1SavePostParams creates a new CustomerGroupRepositoryV1SavePostParams object
// with the default values initialized.
func NewCustomerGroupRepositoryV1SavePostParams() *CustomerGroupRepositoryV1SavePostParams {
	var ()
	return &CustomerGroupRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGroupRepositoryV1SavePostParamsWithTimeout creates a new CustomerGroupRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGroupRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CustomerGroupRepositoryV1SavePostParams {
	var ()
	return &CustomerGroupRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCustomerGroupRepositoryV1SavePostParamsWithContext creates a new CustomerGroupRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGroupRepositoryV1SavePostParamsWithContext(ctx context.Context) *CustomerGroupRepositoryV1SavePostParams {
	var ()
	return &CustomerGroupRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCustomerGroupRepositoryV1SavePostParamsWithHTTPClient creates a new CustomerGroupRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGroupRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CustomerGroupRepositoryV1SavePostParams {
	var ()
	return &CustomerGroupRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CustomerGroupRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the customer group repository v1 save post operation typically these are written to a http.Request
*/
type CustomerGroupRepositoryV1SavePostParams struct {

	/*CustomerGroupRepositoryV1SavePostBody*/
	CustomerGroupRepositoryV1SavePostBody CustomerGroupRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CustomerGroupRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) WithContext(ctx context.Context) *CustomerGroupRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CustomerGroupRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerGroupRepositoryV1SavePostBody adds the customerGroupRepositoryV1SavePostBody to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) WithCustomerGroupRepositoryV1SavePostBody(customerGroupRepositoryV1SavePostBody CustomerGroupRepositoryV1SavePostBody) *CustomerGroupRepositoryV1SavePostParams {
	o.SetCustomerGroupRepositoryV1SavePostBody(customerGroupRepositoryV1SavePostBody)
	return o
}

// SetCustomerGroupRepositoryV1SavePostBody adds the customerGroupRepositoryV1SavePostBody to the customer group repository v1 save post params
func (o *CustomerGroupRepositoryV1SavePostParams) SetCustomerGroupRepositoryV1SavePostBody(customerGroupRepositoryV1SavePostBody CustomerGroupRepositoryV1SavePostBody) {
	o.CustomerGroupRepositoryV1SavePostBody = customerGroupRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGroupRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CustomerGroupRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}