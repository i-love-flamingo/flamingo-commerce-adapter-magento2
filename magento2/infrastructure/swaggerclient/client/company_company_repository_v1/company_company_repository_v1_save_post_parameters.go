// Code generated by go-swagger; DO NOT EDIT.

package company_company_repository_v1

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

// NewCompanyCompanyRepositoryV1SavePostParams creates a new CompanyCompanyRepositoryV1SavePostParams object
// with the default values initialized.
func NewCompanyCompanyRepositoryV1SavePostParams() *CompanyCompanyRepositoryV1SavePostParams {
	var ()
	return &CompanyCompanyRepositoryV1SavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompanyCompanyRepositoryV1SavePostParamsWithTimeout creates a new CompanyCompanyRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompanyCompanyRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *CompanyCompanyRepositoryV1SavePostParams {
	var ()
	return &CompanyCompanyRepositoryV1SavePostParams{

		timeout: timeout,
	}
}

// NewCompanyCompanyRepositoryV1SavePostParamsWithContext creates a new CompanyCompanyRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompanyCompanyRepositoryV1SavePostParamsWithContext(ctx context.Context) *CompanyCompanyRepositoryV1SavePostParams {
	var ()
	return &CompanyCompanyRepositoryV1SavePostParams{

		Context: ctx,
	}
}

// NewCompanyCompanyRepositoryV1SavePostParamsWithHTTPClient creates a new CompanyCompanyRepositoryV1SavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompanyCompanyRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *CompanyCompanyRepositoryV1SavePostParams {
	var ()
	return &CompanyCompanyRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/*CompanyCompanyRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
for the company company repository v1 save post operation typically these are written to a http.Request
*/
type CompanyCompanyRepositoryV1SavePostParams struct {

	/*CompanyCompanyRepositoryV1SavePostBody*/
	CompanyCompanyRepositoryV1SavePostBody CompanyCompanyRepositoryV1SavePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *CompanyCompanyRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) WithContext(ctx context.Context) *CompanyCompanyRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *CompanyCompanyRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompanyCompanyRepositoryV1SavePostBody adds the companyCompanyRepositoryV1SavePostBody to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) WithCompanyCompanyRepositoryV1SavePostBody(companyCompanyRepositoryV1SavePostBody CompanyCompanyRepositoryV1SavePostBody) *CompanyCompanyRepositoryV1SavePostParams {
	o.SetCompanyCompanyRepositoryV1SavePostBody(companyCompanyRepositoryV1SavePostBody)
	return o
}

// SetCompanyCompanyRepositoryV1SavePostBody adds the companyCompanyRepositoryV1SavePostBody to the company company repository v1 save post params
func (o *CompanyCompanyRepositoryV1SavePostParams) SetCompanyCompanyRepositoryV1SavePostBody(companyCompanyRepositoryV1SavePostBody CompanyCompanyRepositoryV1SavePostBody) {
	o.CompanyCompanyRepositoryV1SavePostBody = companyCompanyRepositoryV1SavePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CompanyCompanyRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CompanyCompanyRepositoryV1SavePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
