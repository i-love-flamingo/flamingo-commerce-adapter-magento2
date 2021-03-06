// Code generated by go-swagger; DO NOT EDIT.

package downloadable_link_repository_v1

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

// NewDownloadableLinkRepositoryV1SavePostParams creates a new DownloadableLinkRepositoryV1SavePostParams object
// with the default values initialized.
func NewDownloadableLinkRepositoryV1SavePostParams() *DownloadableLinkRepositoryV1SavePostParams {
	var ()
	return &DownloadableLinkRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadableLinkRepositoryV1SavePostParamsWithTimeout creates a new DownloadableLinkRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadableLinkRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *DownloadableLinkRepositoryV1SavePostParams {
	var ()
	return &DownloadableLinkRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewDownloadableLinkRepositoryV1SavePostParamsWithContext creates a new DownloadableLinkRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadableLinkRepositoryV1SavePostParamsWithContext(ctx context.Context) *DownloadableLinkRepositoryV1SavePostParams {
	var ()
	return &DownloadableLinkRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewDownloadableLinkRepositoryV1SavePostParamsWithHTTPClient creates a new DownloadableLinkRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadableLinkRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *DownloadableLinkRepositoryV1SavePostParams {
	var ()
	return &DownloadableLinkRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*DownloadableLinkRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the downloadable link repository v1 save post operation typically these are written to a http.Request
*/
type DownloadableLinkRepositoryV1SavePostParams struct {

	/*DownloadableLinkRepositoryV1SavePostBody*/
	DownloadableLinkRepositoryV1SavePostBody DownloadableLinkRepositoryV1SavePostBody
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *DownloadableLinkRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) WithContext(ctx context.Context) *DownloadableLinkRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *DownloadableLinkRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDownloadableLinkRepositoryV1SavePostBody adds the downloadableLinkRepositoryV1SavePostBody to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) WithDownloadableLinkRepositoryV1SavePostBody(downloadableLinkRepositoryV1SavePostBody DownloadableLinkRepositoryV1SavePostBody) *DownloadableLinkRepositoryV1SavePostParams {
	o.SetDownloadableLinkRepositoryV1SavePostBody(downloadableLinkRepositoryV1SavePostBody)
	return o
}

// SetDownloadableLinkRepositoryV1SavePostBody adds the downloadableLinkRepositoryV1SavePostBody to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) SetDownloadableLinkRepositoryV1SavePostBody(downloadableLinkRepositoryV1SavePostBody DownloadableLinkRepositoryV1SavePostBody) {
	o.DownloadableLinkRepositoryV1SavePostBody = downloadableLinkRepositoryV1SavePostBody
}

// WithSku adds the sku to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) WithSku(sku string) *DownloadableLinkRepositoryV1SavePostParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the downloadable link repository v1 save post params
func (o *DownloadableLinkRepositoryV1SavePostParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadableLinkRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DownloadableLinkRepositoryV1SavePostBody); err != nil {
		return err
	}

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
