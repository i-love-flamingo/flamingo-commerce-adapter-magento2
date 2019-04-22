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

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams object
// with the default values initialized.
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams() *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithTimeout creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams{

		timeout: timeout,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithContext creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams{

		Context: ctx,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithHTTPClient creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParamsWithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams{
		HTTPClient: client,
	}
}

/*CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams contains all the parameters to send to the API endpoint
for the checkout payment information management v1 save payment information post mine operation typically these are written to a http.Request
*/
type CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams struct {

	/*CheckoutPaymentInformationManagementV1SavePaymentInformationPostBody*/
	CheckoutPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) WithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) WithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) WithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckoutPaymentInformationManagementV1SavePaymentInformationPostBody adds the checkoutPaymentInformationManagementV1SavePaymentInformationPostBody to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) WithCheckoutPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineBody) *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams {
	o.SetCheckoutPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationPostBody)
	return o
}

// SetCheckoutPaymentInformationManagementV1SavePaymentInformationPostBody adds the checkoutPaymentInformationManagementV1SavePaymentInformationPostBody to the checkout payment information management v1 save payment information post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) SetCheckoutPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineBody) {
	o.CheckoutPaymentInformationManagementV1SavePaymentInformationPostBody = checkoutPaymentInformationManagementV1SavePaymentInformationPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationPostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CheckoutPaymentInformationManagementV1SavePaymentInformationPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}