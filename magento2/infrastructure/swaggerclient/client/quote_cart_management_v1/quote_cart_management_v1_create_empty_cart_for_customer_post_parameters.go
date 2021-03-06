// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_management_v1

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

// NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParams creates a new QuoteCartManagementV1CreateEmptyCartForCustomerPostParams object
// with the default values initialized.
func NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParams() *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {

	return &QuoteCartManagementV1CreateEmptyCartForCustomerPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithTimeout creates a new QuoteCartManagementV1CreateEmptyCartForCustomerPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithTimeout(timeout time.Duration) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {

	return &QuoteCartManagementV1CreateEmptyCartForCustomerPostParams{

		timeout: timeout,
	}
}

// NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithContext creates a new QuoteCartManagementV1CreateEmptyCartForCustomerPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithContext(ctx context.Context) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {

	return &QuoteCartManagementV1CreateEmptyCartForCustomerPostParams{

		Context: ctx,
	}
}

// NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithHTTPClient creates a new QuoteCartManagementV1CreateEmptyCartForCustomerPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteCartManagementV1CreateEmptyCartForCustomerPostParamsWithHTTPClient(client *http.Client) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {

	return &QuoteCartManagementV1CreateEmptyCartForCustomerPostParams{
		HTTPClient: client,
	}
}

/*QuoteCartManagementV1CreateEmptyCartForCustomerPostParams contains all the parameters to send to the API endpoint
for the quote cart management v1 create empty cart for customer post operation typically these are written to a http.Request
*/
type QuoteCartManagementV1CreateEmptyCartForCustomerPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) WithTimeout(timeout time.Duration) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) WithContext(ctx context.Context) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) WithHTTPClient(client *http.Client) *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote cart management v1 create empty cart for customer post params
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteCartManagementV1CreateEmptyCartForCustomerPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
