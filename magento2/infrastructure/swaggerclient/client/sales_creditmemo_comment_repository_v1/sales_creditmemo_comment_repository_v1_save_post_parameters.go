// Code generated by go-swagger; DO NOT EDIT.

package sales_creditmemo_comment_repository_v1

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

// NewSalesCreditmemoCommentRepositoryV1SavePostParams creates a new SalesCreditmemoCommentRepositoryV1SavePostParams object
// with the default values initialized.
func NewSalesCreditmemoCommentRepositoryV1SavePostParams() *SalesCreditmemoCommentRepositoryV1SavePostParams {
	var ()
	return &SalesCreditmemoCommentRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithTimeout creates a new SalesCreditmemoCommentRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	var ()
	return &SalesCreditmemoCommentRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithContext creates a new SalesCreditmemoCommentRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithContext(ctx context.Context) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	var ()
	return &SalesCreditmemoCommentRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithHTTPClient creates a new SalesCreditmemoCommentRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesCreditmemoCommentRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	var ()
	return &SalesCreditmemoCommentRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*SalesCreditmemoCommentRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the sales creditmemo comment repository v1 save post operation typically these are written to a http.Request
*/
type SalesCreditmemoCommentRepositoryV1SavePostParams struct {

	/*ID*/
	ID string
	/*SalesCreditmemoCommentRepositoryV1SavePostBody*/
	SalesCreditmemoCommentRepositoryV1SavePostBody SalesCreditmemoCommentRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WithContext(ctx context.Context) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WithID(id string) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) SetID(id string) {
	o.ID = id
}

// WithSalesCreditmemoCommentRepositoryV1SavePostBody adds the salesCreditmemoCommentRepositoryV1SavePostBody to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WithSalesCreditmemoCommentRepositoryV1SavePostBody(salesCreditmemoCommentRepositoryV1SavePostBody SalesCreditmemoCommentRepositoryV1SavePostBody) *SalesCreditmemoCommentRepositoryV1SavePostParams {
	o.SetSalesCreditmemoCommentRepositoryV1SavePostBody(salesCreditmemoCommentRepositoryV1SavePostBody)
	return o
}

// SetSalesCreditmemoCommentRepositoryV1SavePostBody adds the salesCreditmemoCommentRepositoryV1SavePostBody to the sales creditmemo comment repository v1 save post params
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) SetSalesCreditmemoCommentRepositoryV1SavePostBody(salesCreditmemoCommentRepositoryV1SavePostBody SalesCreditmemoCommentRepositoryV1SavePostBody) {
	o.SalesCreditmemoCommentRepositoryV1SavePostBody = salesCreditmemoCommentRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *SalesCreditmemoCommentRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SalesCreditmemoCommentRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}