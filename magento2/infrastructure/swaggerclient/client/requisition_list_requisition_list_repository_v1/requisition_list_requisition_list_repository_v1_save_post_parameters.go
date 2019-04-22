// Code generated by go-swagger; DO NOT EDIT.

package requisition_list_requisition_list_repository_v1

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

// NewRequisitionListRequisitionListRepositoryV1SavePostParams creates a new RequisitionListRequisitionListRepositoryV1SavePostParams object
// with the default values initialized.
func NewRequisitionListRequisitionListRepositoryV1SavePostParams() *RequisitionListRequisitionListRepositoryV1SavePostParams {
	var ()
	return &RequisitionListRequisitionListRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithTimeout creates a new RequisitionListRequisitionListRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	var ()
	return &RequisitionListRequisitionListRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithContext creates a new RequisitionListRequisitionListRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithContext(ctx context.Context) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	var ()
	return &RequisitionListRequisitionListRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithHTTPClient creates a new RequisitionListRequisitionListRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequisitionListRequisitionListRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	var ()
	return &RequisitionListRequisitionListRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*RequisitionListRequisitionListRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the requisition list requisition list repository v1 save post operation typically these are written to a http.Request
*/
type RequisitionListRequisitionListRepositoryV1SavePostParams struct {

	/*RequisitionListRequisitionListRepositoryV1SavePostBody*/
	RequisitionListRequisitionListRepositoryV1SavePostBody RequisitionListRequisitionListRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) WithContext(ctx context.Context) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequisitionListRequisitionListRepositoryV1SavePostBody adds the requisitionListRequisitionListRepositoryV1SavePostBody to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) WithRequisitionListRequisitionListRepositoryV1SavePostBody(requisitionListRequisitionListRepositoryV1SavePostBody RequisitionListRequisitionListRepositoryV1SavePostBody) *RequisitionListRequisitionListRepositoryV1SavePostParams {
	o.SetRequisitionListRequisitionListRepositoryV1SavePostBody(requisitionListRequisitionListRepositoryV1SavePostBody)
	return o
}

// SetRequisitionListRequisitionListRepositoryV1SavePostBody adds the requisitionListRequisitionListRepositoryV1SavePostBody to the requisition list requisition list repository v1 save post params
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) SetRequisitionListRequisitionListRepositoryV1SavePostBody(requisitionListRequisitionListRepositoryV1SavePostBody RequisitionListRequisitionListRepositoryV1SavePostBody) {
	o.RequisitionListRequisitionListRepositoryV1SavePostBody = requisitionListRequisitionListRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *RequisitionListRequisitionListRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.RequisitionListRequisitionListRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
