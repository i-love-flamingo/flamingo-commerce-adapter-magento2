// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_coupon_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteGuestCouponManagementV1RemoveDeleteReader is a Reader for the QuoteGuestCouponManagementV1RemoveDelete structure.
type QuoteGuestCouponManagementV1RemoveDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCouponManagementV1RemoveDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteGuestCouponManagementV1RemoveDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteGuestCouponManagementV1RemoveDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteGuestCouponManagementV1RemoveDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCouponManagementV1RemoveDeleteOK creates a QuoteGuestCouponManagementV1RemoveDeleteOK with default headers values
func NewQuoteGuestCouponManagementV1RemoveDeleteOK() *QuoteGuestCouponManagementV1RemoveDeleteOK {
	return &QuoteGuestCouponManagementV1RemoveDeleteOK{}
}

/*QuoteGuestCouponManagementV1RemoveDeleteOK handles this case with default header values.

200 Success.
*/
type QuoteGuestCouponManagementV1RemoveDeleteOK struct {
	Payload bool
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1RemoveDeleteOK  %+v", 200, o.Payload)
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCouponManagementV1RemoveDeleteBadRequest creates a QuoteGuestCouponManagementV1RemoveDeleteBadRequest with default headers values
func NewQuoteGuestCouponManagementV1RemoveDeleteBadRequest() *QuoteGuestCouponManagementV1RemoveDeleteBadRequest {
	return &QuoteGuestCouponManagementV1RemoveDeleteBadRequest{}
}

/*QuoteGuestCouponManagementV1RemoveDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteGuestCouponManagementV1RemoveDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1RemoveDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCouponManagementV1RemoveDeleteDefault creates a QuoteGuestCouponManagementV1RemoveDeleteDefault with default headers values
func NewQuoteGuestCouponManagementV1RemoveDeleteDefault(code int) *QuoteGuestCouponManagementV1RemoveDeleteDefault {
	return &QuoteGuestCouponManagementV1RemoveDeleteDefault{
		_statusCode: code,
	}
}

/*QuoteGuestCouponManagementV1RemoveDeleteDefault handles this case with default header values.

Unexpected error
*/
type QuoteGuestCouponManagementV1RemoveDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest coupon management v1 remove delete default response
func (o *QuoteGuestCouponManagementV1RemoveDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1RemoveDelete default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteGuestCouponManagementV1RemoveDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
