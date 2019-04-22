// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_totals_information_management_v1

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

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostParams creates a new CheckoutGuestTotalsInformationManagementV1CalculatePostParams object
// with the default values initialized.
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostParams() *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	var ()
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithTimeout creates a new CheckoutGuestTotalsInformationManagementV1CalculatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithTimeout(timeout time.Duration) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	var ()
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostParams{

		timeout: timeout,
	}
}

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithContext creates a new CheckoutGuestTotalsInformationManagementV1CalculatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithContext(ctx context.Context) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	var ()
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostParams{

		Context: ctx,
	}
}

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithHTTPClient creates a new CheckoutGuestTotalsInformationManagementV1CalculatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostParamsWithHTTPClient(client *http.Client) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	var ()
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostParams{
		HTTPClient: client,
	}
}

/*CheckoutGuestTotalsInformationManagementV1CalculatePostParams contains all the parameters to send to the API endpoint
for the checkout guest totals information management v1 calculate post operation typically these are written to a http.Request
*/
type CheckoutGuestTotalsInformationManagementV1CalculatePostParams struct {

	/*CartID*/
	CartID string
	/*CheckoutGuestTotalsInformationManagementV1CalculatePostBody*/
	CheckoutGuestTotalsInformationManagementV1CalculatePostBody CheckoutGuestTotalsInformationManagementV1CalculatePostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WithTimeout(timeout time.Duration) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WithContext(ctx context.Context) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WithHTTPClient(client *http.Client) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WithCartID(cartID string) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithCheckoutGuestTotalsInformationManagementV1CalculatePostBody adds the checkoutGuestTotalsInformationManagementV1CalculatePostBody to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WithCheckoutGuestTotalsInformationManagementV1CalculatePostBody(checkoutGuestTotalsInformationManagementV1CalculatePostBody CheckoutGuestTotalsInformationManagementV1CalculatePostBody) *CheckoutGuestTotalsInformationManagementV1CalculatePostParams {
	o.SetCheckoutGuestTotalsInformationManagementV1CalculatePostBody(checkoutGuestTotalsInformationManagementV1CalculatePostBody)
	return o
}

// SetCheckoutGuestTotalsInformationManagementV1CalculatePostBody adds the checkoutGuestTotalsInformationManagementV1CalculatePostBody to the checkout guest totals information management v1 calculate post params
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) SetCheckoutGuestTotalsInformationManagementV1CalculatePostBody(checkoutGuestTotalsInformationManagementV1CalculatePostBody CheckoutGuestTotalsInformationManagementV1CalculatePostBody) {
	o.CheckoutGuestTotalsInformationManagementV1CalculatePostBody = checkoutGuestTotalsInformationManagementV1CalculatePostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.CheckoutGuestTotalsInformationManagementV1CalculatePostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}