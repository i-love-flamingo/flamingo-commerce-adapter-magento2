// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_payment_information_management_v1

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

// NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams creates a new CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized.
func NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams() *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithTimeout creates a new CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams{

		timeout: timeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithContext creates a new CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams{

		Context: ctx,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithHTTPClient creates a new CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParamsWithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams{
		HTTPClient: client,
	}
}

/*CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams contains all the parameters to send to the API endpoint
for the checkout guest payment information management v1 get payment information get operation typically these are written to a http.Request
*/
type CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams struct {

	/*CartID*/
	CartID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) WithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) WithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) WithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) WithCartID(cartID string) *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the checkout guest payment information management v1 get payment information get params
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutGuestPaymentInformationManagementV1GetPaymentInformationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}