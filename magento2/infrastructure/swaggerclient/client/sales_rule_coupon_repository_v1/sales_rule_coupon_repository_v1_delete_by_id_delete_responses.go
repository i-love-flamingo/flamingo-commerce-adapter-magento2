// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_coupon_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesRuleCouponRepositoryV1DeleteByIDDeleteReader is a Reader for the SalesRuleCouponRepositoryV1DeleteByIDDelete structure.
type SalesRuleCouponRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleCouponRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleCouponRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleCouponRepositoryV1DeleteByIDDeleteOK creates a SalesRuleCouponRepositoryV1DeleteByIDDeleteOK with default headers values
func NewSalesRuleCouponRepositoryV1DeleteByIDDeleteOK() *SalesRuleCouponRepositoryV1DeleteByIDDeleteOK {
	return &SalesRuleCouponRepositoryV1DeleteByIDDeleteOK{}
}

/*SalesRuleCouponRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type SalesRuleCouponRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/coupons/{couponId}][%d] salesRuleCouponRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest creates a SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewSalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest() *SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest {
	return &SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/coupons/{couponId}][%d] salesRuleCouponRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized creates a SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewSalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized() *SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized {
	return &SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/coupons/{couponId}][%d] salesRuleCouponRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError creates a SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError with default headers values
func NewSalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError() *SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError {
	return &SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError{}
}

/*SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /V1/coupons/{couponId}][%d] salesRuleCouponRepositoryV1DeleteByIdDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1DeleteByIDDeleteDefault creates a SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewSalesRuleCouponRepositoryV1DeleteByIDDeleteDefault(code int) *SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault {
	return &SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule coupon repository v1 delete by Id delete default response
func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/coupons/{couponId}][%d] salesRuleCouponRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
