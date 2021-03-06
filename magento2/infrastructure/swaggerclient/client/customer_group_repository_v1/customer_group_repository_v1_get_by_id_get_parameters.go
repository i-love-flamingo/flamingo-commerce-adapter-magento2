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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomerGroupRepositoryV1GetByIDGetParams creates a new CustomerGroupRepositoryV1GetByIDGetParams object
// with the default values initialized.
func NewCustomerGroupRepositoryV1GetByIDGetParams() *CustomerGroupRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerGroupRepositoryV1GetByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGroupRepositoryV1GetByIDGetParamsWithTimeout creates a new CustomerGroupRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGroupRepositoryV1GetByIDGetParamsWithTimeout(timeout time.Duration) *CustomerGroupRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerGroupRepositoryV1GetByIDGetParams{

		timeout: timeout,
	}
}

// NewCustomerGroupRepositoryV1GetByIDGetParamsWithContext creates a new CustomerGroupRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGroupRepositoryV1GetByIDGetParamsWithContext(ctx context.Context) *CustomerGroupRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerGroupRepositoryV1GetByIDGetParams{

		Context: ctx,
	}
}

// NewCustomerGroupRepositoryV1GetByIDGetParamsWithHTTPClient creates a new CustomerGroupRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGroupRepositoryV1GetByIDGetParamsWithHTTPClient(client *http.Client) *CustomerGroupRepositoryV1GetByIDGetParams {
	var ()
	return &CustomerGroupRepositoryV1GetByIDGetParams{
		HTTPClient: client,
	}
}

/*CustomerGroupRepositoryV1GetByIDGetParams contains all the parameters to send to the API endpoint
for the customer group repository v1 get by Id get operation typically these are written to a http.Request
*/
type CustomerGroupRepositoryV1GetByIDGetParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) WithTimeout(timeout time.Duration) *CustomerGroupRepositoryV1GetByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) WithContext(ctx context.Context) *CustomerGroupRepositoryV1GetByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) WithHTTPClient(client *http.Client) *CustomerGroupRepositoryV1GetByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) WithID(id int64) *CustomerGroupRepositoryV1GetByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer group repository v1 get by Id get params
func (o *CustomerGroupRepositoryV1GetByIDGetParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGroupRepositoryV1GetByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
