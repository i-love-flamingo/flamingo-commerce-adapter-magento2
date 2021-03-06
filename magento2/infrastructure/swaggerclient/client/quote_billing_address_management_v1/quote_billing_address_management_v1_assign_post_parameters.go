// Code generated by go-swagger; DO NOT EDIT.

package quote_billing_address_management_v1

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

// NewQuoteBillingAddressManagementV1AssignPostParams creates a new QuoteBillingAddressManagementV1AssignPostParams object
// with the default values initialized.
func NewQuoteBillingAddressManagementV1AssignPostParams() *QuoteBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteBillingAddressManagementV1AssignPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteBillingAddressManagementV1AssignPostParamsWithTimeout creates a new QuoteBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteBillingAddressManagementV1AssignPostParamsWithTimeout(timeout time.Duration) *QuoteBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteBillingAddressManagementV1AssignPostParams{

		timeout: timeout,
	}
}

// NewQuoteBillingAddressManagementV1AssignPostParamsWithContext creates a new QuoteBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteBillingAddressManagementV1AssignPostParamsWithContext(ctx context.Context) *QuoteBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteBillingAddressManagementV1AssignPostParams{

		Context: ctx,
	}
}

// NewQuoteBillingAddressManagementV1AssignPostParamsWithHTTPClient creates a new QuoteBillingAddressManagementV1AssignPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteBillingAddressManagementV1AssignPostParamsWithHTTPClient(client *http.Client) *QuoteBillingAddressManagementV1AssignPostParams {
	var ()
	return &QuoteBillingAddressManagementV1AssignPostParams{
		HTTPClient: client,
	}
}

/*QuoteBillingAddressManagementV1AssignPostParams contains all the parameters to send to the API endpoint
for the quote billing address management v1 assign post operation typically these are written to a http.Request
*/
type QuoteBillingAddressManagementV1AssignPostParams struct {

	/*CartID
	  The cart ID.

	*/
	CartID int64
	/*QuoteBillingAddressManagementV1AssignPostBody*/
	QuoteBillingAddressManagementV1AssignPostBody QuoteBillingAddressManagementV1AssignPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) WithTimeout(timeout time.Duration) *QuoteBillingAddressManagementV1AssignPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) WithContext(ctx context.Context) *QuoteBillingAddressManagementV1AssignPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) WithHTTPClient(client *http.Client) *QuoteBillingAddressManagementV1AssignPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) WithCartID(cartID int64) *QuoteBillingAddressManagementV1AssignPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) SetCartID(cartID int64) {
	o.CartID = cartID
}

// WithQuoteBillingAddressManagementV1AssignPostBody adds the quoteBillingAddressManagementV1AssignPostBody to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) WithQuoteBillingAddressManagementV1AssignPostBody(quoteBillingAddressManagementV1AssignPostBody QuoteBillingAddressManagementV1AssignPostBody) *QuoteBillingAddressManagementV1AssignPostParams {
	o.SetQuoteBillingAddressManagementV1AssignPostBody(quoteBillingAddressManagementV1AssignPostBody)
	return o
}

// SetQuoteBillingAddressManagementV1AssignPostBody adds the quoteBillingAddressManagementV1AssignPostBody to the quote billing address management v1 assign post params
func (o *QuoteBillingAddressManagementV1AssignPostParams) SetQuoteBillingAddressManagementV1AssignPostBody(quoteBillingAddressManagementV1AssignPostBody QuoteBillingAddressManagementV1AssignPostBody) {
	o.QuoteBillingAddressManagementV1AssignPostBody = quoteBillingAddressManagementV1AssignPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteBillingAddressManagementV1AssignPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", swag.FormatInt64(o.CartID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.QuoteBillingAddressManagementV1AssignPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
