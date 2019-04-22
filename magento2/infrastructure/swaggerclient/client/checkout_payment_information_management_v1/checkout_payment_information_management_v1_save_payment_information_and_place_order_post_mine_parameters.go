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

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams object
// with the default values initialized.
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams() *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithTimeout creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams{

		timeout: timeout,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithContext creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams{

		Context: ctx,
	}
}

// NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithHTTPClient creates a new CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParamsWithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	var ()
	return &CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams{
		HTTPClient: client,
	}
}

/*CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams contains all the parameters to send to the API endpoint
for the checkout payment information management v1 save payment information and place order post mine operation typically these are written to a http.Request
*/
type CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams struct {

	/*CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody*/
	CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) WithTimeout(timeout time.Duration) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) WithContext(ctx context.Context) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) WithHTTPClient(client *http.Client) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody adds the checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) WithCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineBody) *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams {
	o.SetCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody)
	return o
}

// SetCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody adds the checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody to the checkout payment information management v1 save payment information and place order post mine params
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) SetCheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineBody) {
	o.CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody = checkoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostMineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CheckoutPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
