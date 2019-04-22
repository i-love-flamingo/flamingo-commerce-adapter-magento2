// Code generated by go-swagger; DO NOT EDIT.

package quote_coupon_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteCouponManagementV1RemoveDeleteReader is a Reader for the QuoteCouponManagementV1RemoveDelete structure.
type QuoteCouponManagementV1RemoveDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCouponManagementV1RemoveDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCouponManagementV1RemoveDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCouponManagementV1RemoveDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCouponManagementV1RemoveDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCouponManagementV1RemoveDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCouponManagementV1RemoveDeleteOK creates a QuoteCouponManagementV1RemoveDeleteOK with default headers values
func NewQuoteCouponManagementV1RemoveDeleteOK() *QuoteCouponManagementV1RemoveDeleteOK {
	return &QuoteCouponManagementV1RemoveDeleteOK{}
}

/*QuoteCouponManagementV1RemoveDeleteOK handles this case with default header values.

200 Success.
*/
type QuoteCouponManagementV1RemoveDeleteOK struct {
	Payload bool
}

func (o *QuoteCouponManagementV1RemoveDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/coupons][%d] quoteCouponManagementV1RemoveDeleteOK  %+v", 200, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteBadRequest creates a QuoteCouponManagementV1RemoveDeleteBadRequest with default headers values
func NewQuoteCouponManagementV1RemoveDeleteBadRequest() *QuoteCouponManagementV1RemoveDeleteBadRequest {
	return &QuoteCouponManagementV1RemoveDeleteBadRequest{}
}

/*QuoteCouponManagementV1RemoveDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCouponManagementV1RemoveDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1RemoveDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/coupons][%d] quoteCouponManagementV1RemoveDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteUnauthorized creates a QuoteCouponManagementV1RemoveDeleteUnauthorized with default headers values
func NewQuoteCouponManagementV1RemoveDeleteUnauthorized() *QuoteCouponManagementV1RemoveDeleteUnauthorized {
	return &QuoteCouponManagementV1RemoveDeleteUnauthorized{}
}

/*QuoteCouponManagementV1RemoveDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCouponManagementV1RemoveDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1RemoveDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/coupons][%d] quoteCouponManagementV1RemoveDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteDefault creates a QuoteCouponManagementV1RemoveDeleteDefault with default headers values
func NewQuoteCouponManagementV1RemoveDeleteDefault(code int) *QuoteCouponManagementV1RemoveDeleteDefault {
	return &QuoteCouponManagementV1RemoveDeleteDefault{
		_statusCode: code,
	}
}

/*QuoteCouponManagementV1RemoveDeleteDefault handles this case with default header values.

Unexpected error
*/
type QuoteCouponManagementV1RemoveDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote coupon management v1 remove delete default response
func (o *QuoteCouponManagementV1RemoveDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCouponManagementV1RemoveDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/coupons][%d] quoteCouponManagementV1RemoveDelete default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
