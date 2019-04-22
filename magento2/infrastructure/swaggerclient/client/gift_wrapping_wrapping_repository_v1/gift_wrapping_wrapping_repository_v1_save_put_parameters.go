// Code generated by go-swagger; DO NOT EDIT.

package gift_wrapping_wrapping_repository_v1

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

// NewGiftWrappingWrappingRepositoryV1SavePutParams creates a new GiftWrappingWrappingRepositoryV1SavePutParams object
// with the default values initialized.
func NewGiftWrappingWrappingRepositoryV1SavePutParams() *GiftWrappingWrappingRepositoryV1SavePutParams {
	var ()
	return &GiftWrappingWrappingRepositoryV1SavePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGiftWrappingWrappingRepositoryV1SavePutParamsWithTimeout creates a new GiftWrappingWrappingRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGiftWrappingWrappingRepositoryV1SavePutParamsWithTimeout(timeout time.Duration) *GiftWrappingWrappingRepositoryV1SavePutParams {
	var ()
	return &GiftWrappingWrappingRepositoryV1SavePutParams{

		timeout: timeout,
	}
}

// NewGiftWrappingWrappingRepositoryV1SavePutParamsWithContext creates a new GiftWrappingWrappingRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewGiftWrappingWrappingRepositoryV1SavePutParamsWithContext(ctx context.Context) *GiftWrappingWrappingRepositoryV1SavePutParams {
	var ()
	return &GiftWrappingWrappingRepositoryV1SavePutParams{

		Context: ctx,
	}
}

// NewGiftWrappingWrappingRepositoryV1SavePutParamsWithHTTPClient creates a new GiftWrappingWrappingRepositoryV1SavePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGiftWrappingWrappingRepositoryV1SavePutParamsWithHTTPClient(client *http.Client) *GiftWrappingWrappingRepositoryV1SavePutParams {
	var ()
	return &GiftWrappingWrappingRepositoryV1SavePutParams{
		HTTPClient: client,
	}
}

/*GiftWrappingWrappingRepositoryV1SavePutParams contains all the parameters to send to the API endpoint
for the gift wrapping wrapping repository v1 save put operation typically these are written to a http.Request
*/
type GiftWrappingWrappingRepositoryV1SavePutParams struct {

	/*GiftWrappingWrappingRepositoryV1SavePutBody*/
	GiftWrappingWrappingRepositoryV1SavePutBody GiftWrappingWrappingRepositoryV1SavePutBody
	/*WrappingID*/
	WrappingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WithTimeout(timeout time.Duration) *GiftWrappingWrappingRepositoryV1SavePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WithContext(ctx context.Context) *GiftWrappingWrappingRepositoryV1SavePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WithHTTPClient(client *http.Client) *GiftWrappingWrappingRepositoryV1SavePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGiftWrappingWrappingRepositoryV1SavePutBody adds the giftWrappingWrappingRepositoryV1SavePutBody to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WithGiftWrappingWrappingRepositoryV1SavePutBody(giftWrappingWrappingRepositoryV1SavePutBody GiftWrappingWrappingRepositoryV1SavePutBody) *GiftWrappingWrappingRepositoryV1SavePutParams {
	o.SetGiftWrappingWrappingRepositoryV1SavePutBody(giftWrappingWrappingRepositoryV1SavePutBody)
	return o
}

// SetGiftWrappingWrappingRepositoryV1SavePutBody adds the giftWrappingWrappingRepositoryV1SavePutBody to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) SetGiftWrappingWrappingRepositoryV1SavePutBody(giftWrappingWrappingRepositoryV1SavePutBody GiftWrappingWrappingRepositoryV1SavePutBody) {
	o.GiftWrappingWrappingRepositoryV1SavePutBody = giftWrappingWrappingRepositoryV1SavePutBody
}

// WithWrappingID adds the wrappingID to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WithWrappingID(wrappingID string) *GiftWrappingWrappingRepositoryV1SavePutParams {
	o.SetWrappingID(wrappingID)
	return o
}

// SetWrappingID adds the wrappingId to the gift wrapping wrapping repository v1 save put params
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) SetWrappingID(wrappingID string) {
	o.WrappingID = wrappingID
}

// WriteToRequest writes these params to a swagger request
func (o *GiftWrappingWrappingRepositoryV1SavePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.GiftWrappingWrappingRepositoryV1SavePutBody); err != nil {
		return err
	}

	// path param wrappingId
	if err := r.SetPathParam("wrappingId", o.WrappingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
