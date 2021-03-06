// Code generated by go-swagger; DO NOT EDIT.

package checkout_payment_information_management_v1

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

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParams creates a new CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized.
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParams() *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {

	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithTimeout creates a new CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {

	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams{

		timeout: timeout,
	}
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithContext creates a new CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {

	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams{

		Context: ctx,
	}
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithHTTPClient creates a new CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetParamsWithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {

	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams{
		HTTPClient: client,
	}
}

/*CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams contains all the parameters to send to the API endpoint
for the checkout payment information management v1 get payment information get operation typically these are written to a http.Request
*/
type CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) WithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) WithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) WithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout payment information management v1 get payment information get params
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
