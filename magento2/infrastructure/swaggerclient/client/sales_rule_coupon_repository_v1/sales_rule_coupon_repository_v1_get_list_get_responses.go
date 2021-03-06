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

// SalesRuleCouponRepositoryV1GetListGetReader is a Reader for the SalesRuleCouponRepositoryV1GetListGet structure.
type SalesRuleCouponRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleCouponRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleCouponRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesRuleCouponRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleCouponRepositoryV1GetListGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleCouponRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleCouponRepositoryV1GetListGetOK creates a SalesRuleCouponRepositoryV1GetListGetOK with default headers values
func NewSalesRuleCouponRepositoryV1GetListGetOK() *SalesRuleCouponRepositoryV1GetListGetOK {
	return &SalesRuleCouponRepositoryV1GetListGetOK{}
}

/*SalesRuleCouponRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type SalesRuleCouponRepositoryV1GetListGetOK struct {
	Payload *models.SalesRuleDataCouponSearchResultInterface
}

func (o *SalesRuleCouponRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/coupons/search][%d] salesRuleCouponRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesRuleDataCouponSearchResultInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1GetListGetUnauthorized creates a SalesRuleCouponRepositoryV1GetListGetUnauthorized with default headers values
func NewSalesRuleCouponRepositoryV1GetListGetUnauthorized() *SalesRuleCouponRepositoryV1GetListGetUnauthorized {
	return &SalesRuleCouponRepositoryV1GetListGetUnauthorized{}
}

/*SalesRuleCouponRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleCouponRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/coupons/search][%d] salesRuleCouponRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1GetListGetInternalServerError creates a SalesRuleCouponRepositoryV1GetListGetInternalServerError with default headers values
func NewSalesRuleCouponRepositoryV1GetListGetInternalServerError() *SalesRuleCouponRepositoryV1GetListGetInternalServerError {
	return &SalesRuleCouponRepositoryV1GetListGetInternalServerError{}
}

/*SalesRuleCouponRepositoryV1GetListGetInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleCouponRepositoryV1GetListGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponRepositoryV1GetListGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/coupons/search][%d] salesRuleCouponRepositoryV1GetListGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1GetListGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponRepositoryV1GetListGetDefault creates a SalesRuleCouponRepositoryV1GetListGetDefault with default headers values
func NewSalesRuleCouponRepositoryV1GetListGetDefault(code int) *SalesRuleCouponRepositoryV1GetListGetDefault {
	return &SalesRuleCouponRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*SalesRuleCouponRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleCouponRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule coupon repository v1 get list get default response
func (o *SalesRuleCouponRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleCouponRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/coupons/search][%d] salesRuleCouponRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleCouponRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
