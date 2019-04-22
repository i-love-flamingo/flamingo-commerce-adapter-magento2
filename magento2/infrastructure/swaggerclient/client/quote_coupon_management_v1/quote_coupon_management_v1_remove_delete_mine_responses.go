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

// QuoteCouponManagementV1RemoveDeleteMineReader is a Reader for the QuoteCouponManagementV1RemoveDeleteMine structure.
type QuoteCouponManagementV1RemoveDeleteMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCouponManagementV1RemoveDeleteMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCouponManagementV1RemoveDeleteMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCouponManagementV1RemoveDeleteMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCouponManagementV1RemoveDeleteMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCouponManagementV1RemoveDeleteMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCouponManagementV1RemoveDeleteMineOK creates a QuoteCouponManagementV1RemoveDeleteMineOK with default headers values
func NewQuoteCouponManagementV1RemoveDeleteMineOK() *QuoteCouponManagementV1RemoveDeleteMineOK {
	return &QuoteCouponManagementV1RemoveDeleteMineOK{}
}

/*QuoteCouponManagementV1RemoveDeleteMineOK handles this case with default header values.

200 Success.
*/
type QuoteCouponManagementV1RemoveDeleteMineOK struct {
	Payload bool
}

func (o *QuoteCouponManagementV1RemoveDeleteMineOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/mine/coupons][%d] quoteCouponManagementV1RemoveDeleteMineOK  %+v", 200, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteMineBadRequest creates a QuoteCouponManagementV1RemoveDeleteMineBadRequest with default headers values
func NewQuoteCouponManagementV1RemoveDeleteMineBadRequest() *QuoteCouponManagementV1RemoveDeleteMineBadRequest {
	return &QuoteCouponManagementV1RemoveDeleteMineBadRequest{}
}

/*QuoteCouponManagementV1RemoveDeleteMineBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCouponManagementV1RemoveDeleteMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1RemoveDeleteMineBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/mine/coupons][%d] quoteCouponManagementV1RemoveDeleteMineBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteMineUnauthorized creates a QuoteCouponManagementV1RemoveDeleteMineUnauthorized with default headers values
func NewQuoteCouponManagementV1RemoveDeleteMineUnauthorized() *QuoteCouponManagementV1RemoveDeleteMineUnauthorized {
	return &QuoteCouponManagementV1RemoveDeleteMineUnauthorized{}
}

/*QuoteCouponManagementV1RemoveDeleteMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCouponManagementV1RemoveDeleteMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCouponManagementV1RemoveDeleteMineUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/mine/coupons][%d] quoteCouponManagementV1RemoveDeleteMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCouponManagementV1RemoveDeleteMineDefault creates a QuoteCouponManagementV1RemoveDeleteMineDefault with default headers values
func NewQuoteCouponManagementV1RemoveDeleteMineDefault(code int) *QuoteCouponManagementV1RemoveDeleteMineDefault {
	return &QuoteCouponManagementV1RemoveDeleteMineDefault{
		_statusCode: code,
	}
}

/*QuoteCouponManagementV1RemoveDeleteMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteCouponManagementV1RemoveDeleteMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote coupon management v1 remove delete mine default response
func (o *QuoteCouponManagementV1RemoveDeleteMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCouponManagementV1RemoveDeleteMineDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/mine/coupons][%d] quoteCouponManagementV1RemoveDeleteMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCouponManagementV1RemoveDeleteMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
