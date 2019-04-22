// Code generated by go-swagger; DO NOT EDIT.

package cms_page_repository_v1

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

// NewCmsPageRepositoryV1SavePostParams creates a new CmsPageRepositoryV1SavePostParams object
// with the default values initialized.
func NewCmsPageRepositoryV1SavePostParams() *CmsPageRepositoryV1SavePostParams {
	var ()
	return &CmsPageRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCmsPageRepositoryV1SavePostParamsWithTimeout creates a new CmsPageRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCmsPageRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CmsPageRepositoryV1SavePostParams {
	var ()
	return &CmsPageRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCmsPageRepositoryV1SavePostParamsWithContext creates a new CmsPageRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCmsPageRepositoryV1SavePostParamsWithContext(ctx context.Context) *CmsPageRepositoryV1SavePostParams {
	var ()
	return &CmsPageRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCmsPageRepositoryV1SavePostParamsWithHTTPClient creates a new CmsPageRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCmsPageRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CmsPageRepositoryV1SavePostParams {
	var ()
	return &CmsPageRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CmsPageRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the cms page repository v1 save post operation typically these are written to a http.Request
*/
type CmsPageRepositoryV1SavePostParams struct {

	/*CmsPageRepositoryV1SavePostBody*/
	CmsPageRepositoryV1SavePostBody CmsPageRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CmsPageRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) WithContext(ctx context.Context) *CmsPageRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CmsPageRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCmsPageRepositoryV1SavePostBody adds the cmsPageRepositoryV1SavePostBody to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) WithCmsPageRepositoryV1SavePostBody(cmsPageRepositoryV1SavePostBody CmsPageRepositoryV1SavePostBody) *CmsPageRepositoryV1SavePostParams {
	o.SetCmsPageRepositoryV1SavePostBody(cmsPageRepositoryV1SavePostBody)
	return o
}

// SetCmsPageRepositoryV1SavePostBody adds the cmsPageRepositoryV1SavePostBody to the cms page repository v1 save post params
func (o *CmsPageRepositoryV1SavePostParams) SetCmsPageRepositoryV1SavePostBody(cmsPageRepositoryV1SavePostBody CmsPageRepositoryV1SavePostBody) {
	o.CmsPageRepositoryV1SavePostBody = cmsPageRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CmsPageRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CmsPageRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}