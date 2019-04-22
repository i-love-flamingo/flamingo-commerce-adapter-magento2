// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_comment_locator_v1

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

// NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParams creates a new NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams object
// with the default values initialized.
func NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParams() *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	var ()
	return &NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithTimeout creates a new NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithTimeout(timeout time.Duration) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	var ()
	return &NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithContext creates a new NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithContext(ctx context.Context) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	var ()
	return &NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithHTTPClient creates a new NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteCommentLocatorV1GetListForQuoteGetParamsWithHTTPClient(client *http.Client) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	var ()
	return &NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams contains all the parameters to send to the API endpoint
for the negotiable quote comment locator v1 get list for quote get operation typically these are written to a http.Request
*/
type NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams struct {

	/*QuoteID
	  Negotiable Quote ID.

	*/
	QuoteID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) WithTimeout(timeout time.Duration) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) WithContext(ctx context.Context) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) WithHTTPClient(client *http.Client) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteID adds the quoteID to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) WithQuoteID(quoteID int64) *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams {
	o.SetQuoteID(quoteID)
	return o
}

// SetQuoteID adds the quoteId to the negotiable quote comment locator v1 get list for quote get params
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) SetQuoteID(quoteID int64) {
	o.QuoteID = quoteID
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteCommentLocatorV1GetListForQuoteGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param quoteId
	if err := r.SetPathParam("quoteId", swag.FormatInt64(o.QuoteID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}