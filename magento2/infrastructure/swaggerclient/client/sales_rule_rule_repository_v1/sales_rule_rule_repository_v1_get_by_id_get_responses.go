// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_rule_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesRuleRuleRepositoryV1GetByIDGetReader is a Reader for the SalesRuleRuleRepositoryV1GetByIDGet structure.
type SalesRuleRuleRepositoryV1GetByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleRuleRepositoryV1GetByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleRuleRepositoryV1GetByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesRuleRuleRepositoryV1GetByIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesRuleRuleRepositoryV1GetByIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleRuleRepositoryV1GetByIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleRuleRepositoryV1GetByIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleRuleRepositoryV1GetByIDGetOK creates a SalesRuleRuleRepositoryV1GetByIDGetOK with default headers values
func NewSalesRuleRuleRepositoryV1GetByIDGetOK() *SalesRuleRuleRepositoryV1GetByIDGetOK {
	return &SalesRuleRuleRepositoryV1GetByIDGetOK{}
}

/*SalesRuleRuleRepositoryV1GetByIDGetOK handles this case with default header values.

200 Success.
*/
type SalesRuleRuleRepositoryV1GetByIDGetOK struct {
	Payload *models.SalesRuleDataRuleInterface
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/salesRules/{ruleId}][%d] salesRuleRuleRepositoryV1GetByIdGetOK  %+v", 200, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesRuleDataRuleInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1GetByIDGetBadRequest creates a SalesRuleRuleRepositoryV1GetByIDGetBadRequest with default headers values
func NewSalesRuleRuleRepositoryV1GetByIDGetBadRequest() *SalesRuleRuleRepositoryV1GetByIDGetBadRequest {
	return &SalesRuleRuleRepositoryV1GetByIDGetBadRequest{}
}

/*SalesRuleRuleRepositoryV1GetByIDGetBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesRuleRuleRepositoryV1GetByIDGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/salesRules/{ruleId}][%d] salesRuleRuleRepositoryV1GetByIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1GetByIDGetUnauthorized creates a SalesRuleRuleRepositoryV1GetByIDGetUnauthorized with default headers values
func NewSalesRuleRuleRepositoryV1GetByIDGetUnauthorized() *SalesRuleRuleRepositoryV1GetByIDGetUnauthorized {
	return &SalesRuleRuleRepositoryV1GetByIDGetUnauthorized{}
}

/*SalesRuleRuleRepositoryV1GetByIDGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleRuleRepositoryV1GetByIDGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/salesRules/{ruleId}][%d] salesRuleRuleRepositoryV1GetByIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1GetByIDGetInternalServerError creates a SalesRuleRuleRepositoryV1GetByIDGetInternalServerError with default headers values
func NewSalesRuleRuleRepositoryV1GetByIDGetInternalServerError() *SalesRuleRuleRepositoryV1GetByIDGetInternalServerError {
	return &SalesRuleRuleRepositoryV1GetByIDGetInternalServerError{}
}

/*SalesRuleRuleRepositoryV1GetByIDGetInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleRuleRepositoryV1GetByIDGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/salesRules/{ruleId}][%d] salesRuleRuleRepositoryV1GetByIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1GetByIDGetDefault creates a SalesRuleRuleRepositoryV1GetByIDGetDefault with default headers values
func NewSalesRuleRuleRepositoryV1GetByIDGetDefault(code int) *SalesRuleRuleRepositoryV1GetByIDGetDefault {
	return &SalesRuleRuleRepositoryV1GetByIDGetDefault{
		_statusCode: code,
	}
}

/*SalesRuleRuleRepositoryV1GetByIDGetDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleRuleRepositoryV1GetByIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule rule repository v1 get by Id get default response
func (o *SalesRuleRuleRepositoryV1GetByIDGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/salesRules/{ruleId}][%d] salesRuleRuleRepositoryV1GetByIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1GetByIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
