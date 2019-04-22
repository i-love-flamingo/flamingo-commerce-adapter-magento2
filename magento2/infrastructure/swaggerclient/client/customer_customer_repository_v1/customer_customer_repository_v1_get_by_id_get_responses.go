// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CustomerCustomerRepositoryV1GetByIDGetReader is a Reader for the CustomerCustomerRepositoryV1GetByIDGet structure.
type CustomerCustomerRepositoryV1GetByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCustomerRepositoryV1GetByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerCustomerRepositoryV1GetByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomerCustomerRepositoryV1GetByIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCustomerCustomerRepositoryV1GetByIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerCustomerRepositoryV1GetByIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerCustomerRepositoryV1GetByIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerCustomerRepositoryV1GetByIDGetOK creates a CustomerCustomerRepositoryV1GetByIDGetOK with default headers values
func NewCustomerCustomerRepositoryV1GetByIDGetOK() *CustomerCustomerRepositoryV1GetByIDGetOK {
	return &CustomerCustomerRepositoryV1GetByIDGetOK{}
}

/*CustomerCustomerRepositoryV1GetByIDGetOK handles this case with default header values.

200 Success.
*/
type CustomerCustomerRepositoryV1GetByIDGetOK struct {
	Payload *models.CustomerDataCustomerInterface
}

func (o *CustomerCustomerRepositoryV1GetByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/customers/{customerId}][%d] customerCustomerRepositoryV1GetByIdGetOK  %+v", 200, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerDataCustomerInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetByIDGetBadRequest creates a CustomerCustomerRepositoryV1GetByIDGetBadRequest with default headers values
func NewCustomerCustomerRepositoryV1GetByIDGetBadRequest() *CustomerCustomerRepositoryV1GetByIDGetBadRequest {
	return &CustomerCustomerRepositoryV1GetByIDGetBadRequest{}
}

/*CustomerCustomerRepositoryV1GetByIDGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CustomerCustomerRepositoryV1GetByIDGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1GetByIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/customers/{customerId}][%d] customerCustomerRepositoryV1GetByIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetByIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetByIDGetUnauthorized creates a CustomerCustomerRepositoryV1GetByIDGetUnauthorized with default headers values
func NewCustomerCustomerRepositoryV1GetByIDGetUnauthorized() *CustomerCustomerRepositoryV1GetByIDGetUnauthorized {
	return &CustomerCustomerRepositoryV1GetByIDGetUnauthorized{}
}

/*CustomerCustomerRepositoryV1GetByIDGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerCustomerRepositoryV1GetByIDGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1GetByIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/customers/{customerId}][%d] customerCustomerRepositoryV1GetByIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetByIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetByIDGetInternalServerError creates a CustomerCustomerRepositoryV1GetByIDGetInternalServerError with default headers values
func NewCustomerCustomerRepositoryV1GetByIDGetInternalServerError() *CustomerCustomerRepositoryV1GetByIDGetInternalServerError {
	return &CustomerCustomerRepositoryV1GetByIDGetInternalServerError{}
}

/*CustomerCustomerRepositoryV1GetByIDGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerCustomerRepositoryV1GetByIDGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1GetByIDGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/customers/{customerId}][%d] customerCustomerRepositoryV1GetByIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetByIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetByIDGetDefault creates a CustomerCustomerRepositoryV1GetByIDGetDefault with default headers values
func NewCustomerCustomerRepositoryV1GetByIDGetDefault(code int) *CustomerCustomerRepositoryV1GetByIDGetDefault {
	return &CustomerCustomerRepositoryV1GetByIDGetDefault{
		_statusCode: code,
	}
}

/*CustomerCustomerRepositoryV1GetByIDGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerCustomerRepositoryV1GetByIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer customer repository v1 get by Id get default response
func (o *CustomerCustomerRepositoryV1GetByIDGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerCustomerRepositoryV1GetByIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/customers/{customerId}][%d] customerCustomerRepositoryV1GetByIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetByIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
