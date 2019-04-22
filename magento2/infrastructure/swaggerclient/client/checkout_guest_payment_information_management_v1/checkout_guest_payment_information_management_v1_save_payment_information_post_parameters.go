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

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams object
// with the default values initialized.
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams() *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithTimeout creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams{

		timeout: timeout,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithContext creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams{

		Context: ctx,
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithHTTPClient creates a new CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParamsWithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	var ()
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams{
		HTTPClient: client,
	}
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams contains all the parameters to send to the API endpoint
for the checkout guest payment information management v1 save payment information post operation typically these are written to a http.Request
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams struct {

	/*CartID*/
	CartID string
	/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody*/
	CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WithTimeout(timeout time.Duration) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WithContext(ctx context.Context) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WithHTTPClient(client *http.Client) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WithCartID(cartID string) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody adds the checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WithCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams {
	o.SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody)
	return o
}

// SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody adds the checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody to the checkout guest payment information management v1 save payment information post params
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) SetCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody(checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) {
	o.CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody = checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
