// Code generated by go-swagger; DO NOT EDIT.

package cms_block_repository_v1

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

// NewCmsBlockRepositoryV1SavePutParams creates a new CmsBlockRepositoryV1SavePutParams object
// with the default values initialized.
func NewCmsBlockRepositoryV1SavePutParams() *CmsBlockRepositoryV1SavePutParams {
	var ()
	return &CmsBlockRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCmsBlockRepositoryV1SavePutParamsWithTimeout creates a new CmsBlockRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCmsBlockRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *CmsBlockRepositoryV1SavePutParams {
	var ()
	return &CmsBlockRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewCmsBlockRepositoryV1SavePutParamsWithContext creates a new CmsBlockRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCmsBlockRepositoryV1SavePutParamsWithContext(ctx context.Context) *CmsBlockRepositoryV1SavePutParams {
	var ()
	return &CmsBlockRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewCmsBlockRepositoryV1SavePutParamsWithHTTPClient creates a new CmsBlockRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCmsBlockRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *CmsBlockRepositoryV1SavePutParams {
	var ()
	return &CmsBlockRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*CmsBlockRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the cms block repository v1 save put operation typically these are written to a http.Request
*/
type CmsBlockRepositoryV1SavePutParams struct {

	/*CmsBlockRepositoryV1SavePutBody*/
	CmsBlockRepositoryV1SavePutBody CmsBlockRepositoryV1SavePutBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *CmsBlockRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) WithContext(ctx context.Context) *CmsBlockRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *CmsBlockRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCmsBlockRepositoryV1SavePutBody adds the cmsBlockRepositoryV1SavePutBody to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) WithCmsBlockRepositoryV1SavePutBody(cmsBlockRepositoryV1SavePutBody CmsBlockRepositoryV1SavePutBody) *CmsBlockRepositoryV1SavePutParams {
	o.SetCmsBlockRepositoryV1SavePutBody(cmsBlockRepositoryV1SavePutBody)
	return o
}

// SetCmsBlockRepositoryV1SavePutBody adds the cmsBlockRepositoryV1SavePutBody to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) SetCmsBlockRepositoryV1SavePutBody(cmsBlockRepositoryV1SavePutBody CmsBlockRepositoryV1SavePutBody) {
	o.CmsBlockRepositoryV1SavePutBody = cmsBlockRepositoryV1SavePutBody
}

// WithID adds the id to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) WithID(id string) *CmsBlockRepositoryV1SavePutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cms block repository v1 save put params
func (o *CmsBlockRepositoryV1SavePutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CmsBlockRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CmsBlockRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
