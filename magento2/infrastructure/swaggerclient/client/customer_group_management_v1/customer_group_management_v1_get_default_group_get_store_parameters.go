// Code generated by go-swagger; DO NOT EDIT.

package customer_group_management_v1

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

// NewCustomerGroupManagementV1GetDefaultGroupGetStoreParams creates a new CustomerGroupManagementV1GetDefaultGroupGetStoreParams object
// with the default values initialized.
func NewCustomerGroupManagementV1GetDefaultGroupGetStoreParams() *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	var ()
	return &CustomerGroupManagementV1GetDefaultGroupGetStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithTimeout creates a new CustomerGroupManagementV1GetDefaultGroupGetStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithTimeout(timeout time.Duration) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	var ()
	return &CustomerGroupManagementV1GetDefaultGroupGetStoreParams{

		timeout: timeout,
	}
}

// NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithContext creates a new CustomerGroupManagementV1GetDefaultGroupGetStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithContext(ctx context.Context) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	var ()
	return &CustomerGroupManagementV1GetDefaultGroupGetStoreParams{

		Context: ctx,
	}
}

// NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithHTTPClient creates a new CustomerGroupManagementV1GetDefaultGroupGetStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGroupManagementV1GetDefaultGroupGetStoreParamsWithHTTPClient(client *http.Client) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	var ()
	return &CustomerGroupManagementV1GetDefaultGroupGetStoreParams{
		HTTPClient: client,
	}
}

/*CustomerGroupManagementV1GetDefaultGroupGetStoreParams contains all the parameters to send to the API endpoint
for the customer group management v1 get default group get store operation typically these are written to a http.Request
*/
type CustomerGroupManagementV1GetDefaultGroupGetStoreParams struct {

	/*StoreID*/
	StoreID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) WithTimeout(timeout time.Duration) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) WithContext(ctx context.Context) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) WithHTTPClient(client *http.Client) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoreID adds the storeID to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) WithStoreID(storeID int64) *CustomerGroupManagementV1GetDefaultGroupGetStoreParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the customer group management v1 get default group get store params
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) SetStoreID(storeID int64) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGroupManagementV1GetDefaultGroupGetStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param storeId
	if err := r.SetPathParam("storeId", swag.FormatInt64(o.StoreID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}