// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_negotiable_quote_shipping_management_v1

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

// NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams creates a new NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams object
// with the default values initialized.
func NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams() *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	var ()
	return &NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithTimeout creates a new NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithTimeout(timeout time.Duration) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	var ()
	return &NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithContext creates a new NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithContext(ctx context.Context) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	var ()
	return &NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithHTTPClient creates a new NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParamsWithHTTPClient(client *http.Client) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	var ()
	return &NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams contains all the parameters to send to the API endpoint
for the negotiable quote negotiable quote shipping management v1 set shipping method put operation typically these are written to a http.Request
*/
type NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams struct {

	/*NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody*/
	NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody
	/*QuoteID
	  Negotiable Quote id

	*/
	QuoteID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WithTimeout(timeout time.Duration) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WithContext(ctx context.Context) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WithHTTPClient(client *http.Client) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody adds the negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WithNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody(negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	o.SetNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody(negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody)
	return o
}

// SetNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody adds the negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) SetNegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody(negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody) {
	o.NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody = negotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody
}

// WithQuoteID adds the quoteID to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WithQuoteID(quoteID int64) *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams {
	o.SetQuoteID(quoteID)
	return o
}

// SetQuoteID adds the quoteId to the negotiable quote negotiable quote shipping management v1 set shipping method put params
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) SetQuoteID(quoteID int64) {
	o.QuoteID = quoteID
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.NegotiableQuoteNegotiableQuoteShippingManagementV1SetShippingMethodPutBody); err != nil {
		return err
	}

	// path param quoteId
	if err := r.SetPathParam("quoteId", swag.FormatInt64(o.QuoteID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
