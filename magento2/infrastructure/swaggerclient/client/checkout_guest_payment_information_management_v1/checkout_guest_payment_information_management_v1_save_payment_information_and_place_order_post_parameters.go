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

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams object
// with the default values initialized.
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams() *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithTimeout creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams{

		timeout: timeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithContext creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams{

		Context: ctx,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithHTTPClient creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParamsWithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams{
		HTTPClient: client,
	}
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams contains all the parameters to send to the API endpoint
for the checkout guest payment information management v1 save payment information and place order post operation typically these are written to a http.Request
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams struct {

	/*CartID*/
	CartID string
	/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody*/
	CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WithCartID(cartID string) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody adds the checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WithCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams {
	o.SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody)
	return o
}

// SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody adds the checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody to the checkout guest payment information management v1 save payment information and place order post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody) {
	o.CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody = checkoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CheckoutGuestPaymentInformationManagementV1SavePaymentInformationAndPlaceOrderPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
