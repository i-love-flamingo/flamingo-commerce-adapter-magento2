// Code generated by go-swagger; DO NOT EDIT.

package downloadable_sample_repository_v1

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

// NewDownloadableSampleRepositoryV1SavePutParams creates a new DownloadableSampleRepositoryV1SavePutParams object
// with the default values initialized.
func NewDownloadableSampleRepositoryV1SavePutParams() *DownloadableSampleRepositoryV1SavePutParams {
	var ()
	return &DownloadableSampleRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadableSampleRepositoryV1SavePutParamsWithTimeout creates a new DownloadableSampleRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadableSampleRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *DownloadableSampleRepositoryV1SavePutParams {
	var ()
	return &DownloadableSampleRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewDownloadableSampleRepositoryV1SavePutParamsWithContext creates a new DownloadableSampleRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadableSampleRepositoryV1SavePutParamsWithContext(ctx context.Context) *DownloadableSampleRepositoryV1SavePutParams {
	var ()
	return &DownloadableSampleRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewDownloadableSampleRepositoryV1SavePutParamsWithHTTPClient creates a new DownloadableSampleRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadableSampleRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *DownloadableSampleRepositoryV1SavePutParams {
	var ()
	return &DownloadableSampleRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*DownloadableSampleRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the downloadable sample repository v1 save put operation typically these are written to a http.Request
*/
type DownloadableSampleRepositoryV1SavePutParams struct {

	/*DownloadableSampleRepositoryV1SavePutBody*/
	DownloadableSampleRepositoryV1SavePutBody DownloadableSampleRepositoryV1SavePutBody
	/*ID*/
	ID string
	/*Sku*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithContext(ctx context.Context) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDownloadableSampleRepositoryV1SavePutBody adds the downloadableSampleRepositoryV1SavePutBody to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithDownloadableSampleRepositoryV1SavePutBody(downloadableSampleRepositoryV1SavePutBody DownloadableSampleRepositoryV1SavePutBody) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetDownloadableSampleRepositoryV1SavePutBody(downloadableSampleRepositoryV1SavePutBody)
	return o
}

// SetDownloadableSampleRepositoryV1SavePutBody adds the downloadableSampleRepositoryV1SavePutBody to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetDownloadableSampleRepositoryV1SavePutBody(downloadableSampleRepositoryV1SavePutBody DownloadableSampleRepositoryV1SavePutBody) {
	o.DownloadableSampleRepositoryV1SavePutBody = downloadableSampleRepositoryV1SavePutBody
}

// WithID adds the id to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithID(id string) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetID(id string) {
	o.ID = id
}

// WithSku adds the sku to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) WithSku(sku string) *DownloadableSampleRepositoryV1SavePutParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the downloadable sample repository v1 save put params
func (o *DownloadableSampleRepositoryV1SavePutParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadableSampleRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DownloadableSampleRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
