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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCmsPageRepositoryV1DeleteByIDDeleteParams creates a new CmsPageRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized.
func NewCmsPageRepositoryV1DeleteByIDDeleteParams() *CmsPageRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &CmsPageRepositoryV1DeleteByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithTimeout creates a new CmsPageRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithTimeout(timeout time.Duration) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &CmsPageRepositoryV1DeleteByIDDeleteParams{

		timeout: timeout,
	}
}

// NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithContext creates a new CmsPageRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithContext(ctx context.Context) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &CmsPageRepositoryV1DeleteByIDDeleteParams{

		Context: ctx,
	}
}

// NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithHTTPClient creates a new CmsPageRepositoryV1DeleteByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCmsPageRepositoryV1DeleteByIDDeleteParamsWithHTTPClient(client *http.Client) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	var ()
	return &CmsPageRepositoryV1DeleteByIDDeleteParams{
		HTTPClient: client,
	}
}

/*CmsPageRepositoryV1DeleteByIDDeleteParams contains all the parameters to send to the API endpoint
for the cms page repository v1 delete by Id delete operation typically these are written to a http.Request
*/
type CmsPageRepositoryV1DeleteByIDDeleteParams struct {

	/*PageID*/
	PageID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) WithTimeout(timeout time.Duration) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) WithContext(ctx context.Context) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) WithHTTPClient(client *http.Client) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageID adds the pageID to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) WithPageID(pageID int64) *CmsPageRepositoryV1DeleteByIDDeleteParams {
	o.SetPageID(pageID)
	return o
}

// SetPageID adds the pageId to the cms page repository v1 delete by Id delete params
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) SetPageID(pageID int64) {
	o.PageID = pageID
}

// WriteToRequest writes these params to a swagger request
func (o *CmsPageRepositoryV1DeleteByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pageId
	if err := r.SetPathParam("pageId", swag.FormatInt64(o.PageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}