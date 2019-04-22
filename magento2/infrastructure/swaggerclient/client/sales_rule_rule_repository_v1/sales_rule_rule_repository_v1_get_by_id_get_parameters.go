// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_rule_repository_v1

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

// NewSalesRuleRuleRepositoryV1GetByIDGetParams creates a new SalesRuleRuleRepositoryV1GetByIDGetParams object
// with the default values initialized.
func NewSalesRuleRuleRepositoryV1GetByIDGetParams() *SalesRuleRuleRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleRuleRepositoryV1GetByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithTimeout creates a new SalesRuleRuleRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithTimeout(timeout time.Duration) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleRuleRepositoryV1GetByIDGetParams{

		timeout: timeout,
	}
}

// NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithContext creates a new SalesRuleRuleRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithContext(ctx context.Context) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleRuleRepositoryV1GetByIDGetParams{

		Context: ctx,
	}
}

// NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithHTTPClient creates a new SalesRuleRuleRepositoryV1GetByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesRuleRuleRepositoryV1GetByIDGetParamsWithHTTPClient(client *http.Client) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	var ()
	return &SalesRuleRuleRepositoryV1GetByIDGetParams{
		HTTPClient: client,
	}
}

/*SalesRuleRuleRepositoryV1GetByIDGetParams contains all the parameters to send to the API endpoint
for the sales rule rule repository v1 get by Id get operation typically these are written to a http.Request
*/
type SalesRuleRuleRepositoryV1GetByIDGetParams struct {

	/*RuleID*/
	RuleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) WithTimeout(timeout time.Duration) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) WithContext(ctx context.Context) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) WithHTTPClient(client *http.Client) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) WithRuleID(ruleID int64) *SalesRuleRuleRepositoryV1GetByIDGetParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the sales rule rule repository v1 get by Id get params
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) SetRuleID(ruleID int64) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *SalesRuleRuleRepositoryV1GetByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", swag.FormatInt64(o.RuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}