// Code generated by go-swagger; DO NOT EDIT.

package company_role_repository_v1

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

// NewCompanyRoleRepositoryV1DeleteDeleteParams creates a new CompanyRoleRepositoryV1DeleteDeleteParams object
// with the default values initialized.
func NewCompanyRoleRepositoryV1DeleteDeleteParams() *CompanyRoleRepositoryV1DeleteDeleteParams {
	var ()
	return &CompanyRoleRepositoryV1DeleteDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompanyRoleRepositoryV1DeleteDeleteParamsWithTimeout creates a new CompanyRoleRepositoryV1DeleteDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompanyRoleRepositoryV1DeleteDeleteParamsWithTimeout(timeout time.Duration) *CompanyRoleRepositoryV1DeleteDeleteParams {
	var ()
	return &CompanyRoleRepositoryV1DeleteDeleteParams{

		timeout: timeout,
	}
}

// NewCompanyRoleRepositoryV1DeleteDeleteParamsWithContext creates a new CompanyRoleRepositoryV1DeleteDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompanyRoleRepositoryV1DeleteDeleteParamsWithContext(ctx context.Context) *CompanyRoleRepositoryV1DeleteDeleteParams {
	var ()
	return &CompanyRoleRepositoryV1DeleteDeleteParams{

		Context: ctx,
	}
}

// NewCompanyRoleRepositoryV1DeleteDeleteParamsWithHTTPClient creates a new CompanyRoleRepositoryV1DeleteDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompanyRoleRepositoryV1DeleteDeleteParamsWithHTTPClient(client *http.Client) *CompanyRoleRepositoryV1DeleteDeleteParams {
	var ()
	return &CompanyRoleRepositoryV1DeleteDeleteParams{
		HTTPClient: client,
	}
}

/*CompanyRoleRepositoryV1DeleteDeleteParams contains all the parameters to send to the API endpoint
for the company role repository v1 delete delete operation typically these are written to a http.Request
*/
type CompanyRoleRepositoryV1DeleteDeleteParams struct {

	/*RoleID*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) WithTimeout(timeout time.Duration) *CompanyRoleRepositoryV1DeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) WithContext(ctx context.Context) *CompanyRoleRepositoryV1DeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) WithHTTPClient(client *http.Client) *CompanyRoleRepositoryV1DeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) WithRoleID(roleID int64) *CompanyRoleRepositoryV1DeleteDeleteParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the company role repository v1 delete delete params
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *CompanyRoleRepositoryV1DeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
