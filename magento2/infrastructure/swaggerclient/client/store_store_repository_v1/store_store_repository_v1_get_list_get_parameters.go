// Code generated by go-swagger; DO NOT EDIT.

package store_store_repository_v1

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

// NewStoreStoreRepositoryV1GetListGetParams creates a new StoreStoreRepositoryV1GetListGetParams object
// with the default values initialized.
func NewStoreStoreRepositoryV1GetListGetParams() *StoreStoreRepositoryV1GetListGetParams {

	return &StoreStoreRepositoryV1GetListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStoreStoreRepositoryV1GetListGetParamsWithTimeout creates a new StoreStoreRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStoreStoreRepositoryV1GetListGetParamsWithTimeout(timeout time.Duration) *StoreStoreRepositoryV1GetListGetParams {

	return &StoreStoreRepositoryV1GetListGetParams{

		timeout: timeout,
	}
}

// NewStoreStoreRepositoryV1GetListGetParamsWithContext creates a new StoreStoreRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStoreStoreRepositoryV1GetListGetParamsWithContext(ctx context.Context) *StoreStoreRepositoryV1GetListGetParams {

	return &StoreStoreRepositoryV1GetListGetParams{

		Context: ctx,
	}
}

// NewStoreStoreRepositoryV1GetListGetParamsWithHTTPClient creates a new StoreStoreRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStoreStoreRepositoryV1GetListGetParamsWithHTTPClient(client *http.Client) *StoreStoreRepositoryV1GetListGetParams {

	return &StoreStoreRepositoryV1GetListGetParams{
		HTTPClient: client,
	}
}

/*StoreStoreRepositoryV1GetListGetParams contains all the parameters to send to the API endpoint
for the store store repository v1 get list get operation typically these are written to a http.Request
*/
type StoreStoreRepositoryV1GetListGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) WithTimeout(timeout time.Duration) *StoreStoreRepositoryV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) WithContext(ctx context.Context) *StoreStoreRepositoryV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) WithHTTPClient(client *http.Client) *StoreStoreRepositoryV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the store store repository v1 get list get params
func (o *StoreStoreRepositoryV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StoreStoreRepositoryV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
