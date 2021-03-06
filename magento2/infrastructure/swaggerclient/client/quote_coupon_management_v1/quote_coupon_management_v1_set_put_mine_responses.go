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

// QuoteCouponManagementV1SetPutMineReader is a Reader for the QuoteCouponManagementV1SetPutMine structure.
type QuoteCouponManagementV1SetPutMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCouponManagementV1SetPutMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCouponManagementV1SetPutMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCouponManagementV1SetPutMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCouponManagementV1SetPutMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCouponManagementV1SetPutMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCouponManagementV1SetPutMineOK creates a QuoteCouponManagementV1SetPutMineOK with default headers values
func NewQuoteCouponManagementV1SetPutMineOK() *QuoteCouponManagementV1SetPutMineOK {
	return &QuoteCouponManagementV1SetPutMineOK{}
}

/*QuoteCouponManagementV1SetPutMineOK handles this case with default header values.

200 Success.
*/
type QuoteCouponManagementV1SetPutMineOK struct {
	Payload bool
}

func (o *QuoteCouponManagementV1SetPutMineOK) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/coupons/{couponCode}][%d] quoteCouponManagementV1SetPutMineOK  %+v", 200, o.Payload)
}

func (o *QuoteCouponManagementV1SetPutMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1SetPutMineBadRequest creates a QuoteCouponManagementV1SetPutMineBadRequest with default headers values
func NewQuoteCouponManagementV1SetPutMineBadRequest() *QuoteCouponManagementV1SetPutMineBadRequest {
	return &QuoteCouponManagementV1SetPutMineBadRequest{}
}

/*QuoteCouponManagementV1SetPutMineBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCouponManagementV1SetPutMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1SetPutMineBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/coupons/{couponCode}][%d] quoteCouponManagementV1SetPutMineBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCouponManagementV1SetPutMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1SetPutMineUnauthorized creates a QuoteCouponManagementV1SetPutMineUnauthorized with default headers values
func NewQuoteCouponManagementV1SetPutMineUnauthorized() *QuoteCouponManagementV1SetPutMineUnauthorized {
	return &QuoteCouponManagementV1SetPutMineUnauthorized{}
}

/*QuoteCouponManagementV1SetPutMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCouponManagementV1SetPutMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1SetPutMineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/coupons/{couponCode}][%d] quoteCouponManagementV1SetPutMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCouponManagementV1SetPutMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1SetPutMineDefault creates a QuoteCouponManagementV1SetPutMineDefault with default headers values
func NewQuoteCouponManagementV1SetPutMineDefault(code int) *QuoteCouponManagementV1SetPutMineDefault {
	return &QuoteCouponManagementV1SetPutMineDefault{
		_statusCode: code,
	}
}

/*QuoteCouponManagementV1SetPutMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteCouponManagementV1SetPutMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote coupon management v1 set put mine default response
func (o *QuoteCouponManagementV1SetPutMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCouponManagementV1SetPutMineDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/coupons/{couponCode}][%d] quoteCouponManagementV1SetPutMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCouponManagementV1SetPutMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
