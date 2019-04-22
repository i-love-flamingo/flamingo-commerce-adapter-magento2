// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_negotiable_cart_repository_v1

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

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParams creates a new NegotiableQuoteNegotiableCartRepositoryV1SavePutParams object
// with the default values initialized.
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParams() *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	var ()
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithTimeout creates a new NegotiableQuoteNegotiableCartRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	var ()
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithContext creates a new NegotiableQuoteNegotiableCartRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithContext(ctx context.Context) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	var ()
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithHTTPClient creates a new NegotiableQuoteNegotiableCartRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	var ()
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteNegotiableCartRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the negotiable quote negotiable cart repository v1 save put operation typically these are written to a http.Request
*/
type NegotiableQuoteNegotiableCartRepositoryV1SavePutParams struct {

	/*NegotiableQuoteNegotiableCartRepositoryV1SavePutBody*/
	NegotiableQuoteNegotiableCartRepositoryV1SavePutBody NegotiableQuoteNegotiableCartRepositoryV1SavePutBody
	/*QuoteID*/
	QuoteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WithContext(ctx context.Context) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNegotiableQuoteNegotiableCartRepositoryV1SavePutBody adds the negotiableQuoteNegotiableCartRepositoryV1SavePutBody to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WithNegotiableQuoteNegotiableCartRepositoryV1SavePutBody(negotiableQuoteNegotiableCartRepositoryV1SavePutBody NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	o.SetNegotiableQuoteNegotiableCartRepositoryV1SavePutBody(negotiableQuoteNegotiableCartRepositoryV1SavePutBody)
	return o
}

// SetNegotiableQuoteNegotiableCartRepositoryV1SavePutBody adds the negotiableQuoteNegotiableCartRepositoryV1SavePutBody to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) SetNegotiableQuoteNegotiableCartRepositoryV1SavePutBody(negotiableQuoteNegotiableCartRepositoryV1SavePutBody NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) {
	o.NegotiableQuoteNegotiableCartRepositoryV1SavePutBody = negotiableQuoteNegotiableCartRepositoryV1SavePutBody
}

// WithQuoteID adds the quoteID to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WithQuoteID(quoteID string) *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams {
	o.SetQuoteID(quoteID)
	return o
}

// SetQuoteID adds the quoteId to the negotiable quote negotiable cart repository v1 save put params
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) SetQuoteID(quoteID string) {
	o.QuoteID = quoteID
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.NegotiableQuoteNegotiableCartRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param quoteId
	if err := r.SetPathParam("quoteId", o.QuoteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
