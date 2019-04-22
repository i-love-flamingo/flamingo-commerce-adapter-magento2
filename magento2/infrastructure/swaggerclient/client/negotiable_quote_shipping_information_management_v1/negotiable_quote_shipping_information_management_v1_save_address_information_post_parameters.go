// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_shipping_information_management_v1

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

// NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams creates a new NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams object
// with the default values initialized.
func NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams() *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	var ()
	return &NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithTimeout creates a new NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithTimeout(timeout time.Duration) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	var ()
	return &NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams{

		timeout: timeout,
	}
}

// NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithContext creates a new NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithContext(ctx context.Context) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	var ()
	return &NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams{

		Context: ctx,
	}
}

// NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithHTTPClient creates a new NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParamsWithHTTPClient(client *http.Client) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	var ()
	return &NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams{
		HTTPClient: client,
	}
}

/*NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams contains all the parameters to send to the API endpoint
for the negotiable quote shipping information management v1 save address information post operation typically these are written to a http.Request
*/
type NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams struct {

	/*CartID*/
	CartID int64
	/*NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody*/
	NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WithTimeout(timeout time.Duration) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WithContext(ctx context.Context) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WithHTTPClient(client *http.Client) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WithCartID(cartID int64) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) SetCartID(cartID int64) {
	o.CartID = cartID
}

// WithNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody adds the negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WithNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody(negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody) *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody(negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody)
	return o
}

// SetNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody adds the negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody to the negotiable quote shipping information management v1 save address information post params
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) SetNegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody(negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody) {
	o.NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody = negotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", swag.FormatInt64(o.CartID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.NegotiableQuoteShippingInformationManagementV1SaveAddressInformationPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}