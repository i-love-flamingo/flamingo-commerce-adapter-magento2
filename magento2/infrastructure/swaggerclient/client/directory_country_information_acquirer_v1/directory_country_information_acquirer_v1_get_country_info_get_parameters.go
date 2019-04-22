// Code generated by go-swagger; DO NOT EDIT.

package directory_country_information_acquirer_v1

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

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParams creates a new DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams object
// with the default values initialized.
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParams() *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	var ()
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithTimeout creates a new DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithTimeout(timeout time.Duration) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	var ()
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams{

		timeout: timeout,
	}
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithContext creates a new DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithContext(ctx context.Context) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	var ()
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams{

		Context: ctx,
	}
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithHTTPClient creates a new DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetParamsWithHTTPClient(client *http.Client) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	var ()
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams{
		HTTPClient: client,
	}
}

/*DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams contains all the parameters to send to the API endpoint
for the directory country information acquirer v1 get country info get operation typically these are written to a http.Request
*/
type DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams struct {

	/*CountryID*/
	CountryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) WithTimeout(timeout time.Duration) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) WithContext(ctx context.Context) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) WithHTTPClient(client *http.Client) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountryID adds the countryID to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) WithCountryID(countryID string) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the directory country information acquirer v1 get country info get params
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WriteToRequest writes these params to a swagger request
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param countryId
	if err := r.SetPathParam("countryId", o.CountryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
