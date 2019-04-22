// Code generated by go-swagger; DO NOT EDIT.

package company_credit_credit_history_management_v1

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

// NewCompanyCreditCreditHistoryManagementV1UpdatePutParams creates a new CompanyCreditCreditHistoryManagementV1UpdatePutParams object
// with the default values initialized.
func NewCompanyCreditCreditHistoryManagementV1UpdatePutParams() *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	var ()
	return &CompanyCreditCreditHistoryManagementV1UpdatePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithTimeout creates a new CompanyCreditCreditHistoryManagementV1UpdatePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithTimeout(timeout time.Duration) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	var ()
	return &CompanyCreditCreditHistoryManagementV1UpdatePutParams{

		timeout: timeout,
	}
}

// NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithContext creates a new CompanyCreditCreditHistoryManagementV1UpdatePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithContext(ctx context.Context) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	var ()
	return &CompanyCreditCreditHistoryManagementV1UpdatePutParams{

		Context: ctx,
	}
}

// NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithHTTPClient creates a new CompanyCreditCreditHistoryManagementV1UpdatePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompanyCreditCreditHistoryManagementV1UpdatePutParamsWithHTTPClient(client *http.Client) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	var ()
	return &CompanyCreditCreditHistoryManagementV1UpdatePutParams{
		HTTPClient: client,
	}
}

/*CompanyCreditCreditHistoryManagementV1UpdatePutParams contains all the parameters to send to the API endpoint
for the company credit credit history management v1 update put operation typically these are written to a http.Request
*/
type CompanyCreditCreditHistoryManagementV1UpdatePutParams struct {

	/*CompanyCreditCreditHistoryManagementV1UpdatePutBody*/
	CompanyCreditCreditHistoryManagementV1UpdatePutBody CompanyCreditCreditHistoryManagementV1UpdatePutBody
	/*HistoryID*/
	HistoryID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WithTimeout(timeout time.Duration) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WithContext(ctx context.Context) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WithHTTPClient(client *http.Client) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompanyCreditCreditHistoryManagementV1UpdatePutBody adds the companyCreditCreditHistoryManagementV1UpdatePutBody to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WithCompanyCreditCreditHistoryManagementV1UpdatePutBody(companyCreditCreditHistoryManagementV1UpdatePutBody CompanyCreditCreditHistoryManagementV1UpdatePutBody) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	o.SetCompanyCreditCreditHistoryManagementV1UpdatePutBody(companyCreditCreditHistoryManagementV1UpdatePutBody)
	return o
}

// SetCompanyCreditCreditHistoryManagementV1UpdatePutBody adds the companyCreditCreditHistoryManagementV1UpdatePutBody to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) SetCompanyCreditCreditHistoryManagementV1UpdatePutBody(companyCreditCreditHistoryManagementV1UpdatePutBody CompanyCreditCreditHistoryManagementV1UpdatePutBody) {
	o.CompanyCreditCreditHistoryManagementV1UpdatePutBody = companyCreditCreditHistoryManagementV1UpdatePutBody
}

// WithHistoryID adds the historyID to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WithHistoryID(historyID int64) *CompanyCreditCreditHistoryManagementV1UpdatePutParams {
	o.SetHistoryID(historyID)
	return o
}

// SetHistoryID adds the historyId to the company credit credit history management v1 update put params
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) SetHistoryID(historyID int64) {
	o.HistoryID = historyID
}

// WriteToRequest writes these params to a swagger request
func (o *CompanyCreditCreditHistoryManagementV1UpdatePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CompanyCreditCreditHistoryManagementV1UpdatePutBody); err != nil {
		return err
	}

	// path param historyId
	if err := r.SetPathParam("historyId", swag.FormatInt64(o.HistoryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
