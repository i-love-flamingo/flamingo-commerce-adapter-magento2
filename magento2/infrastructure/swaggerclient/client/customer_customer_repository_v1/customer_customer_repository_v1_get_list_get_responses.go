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

// CustomerCustomerRepositoryV1GetListGetReader is a Reader for the CustomerCustomerRepositoryV1GetListGet structure.
type CustomerCustomerRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCustomerRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerCustomerRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCustomerCustomerRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerCustomerRepositoryV1GetListGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerCustomerRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerCustomerRepositoryV1GetListGetOK creates a CustomerCustomerRepositoryV1GetListGetOK with default headers values
func NewCustomerCustomerRepositoryV1GetListGetOK() *CustomerCustomerRepositoryV1GetListGetOK {
	return &CustomerCustomerRepositoryV1GetListGetOK{}
}

/*CustomerCustomerRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CustomerCustomerRepositoryV1GetListGetOK struct {
	Payload *models.CustomerDataCustomerSearchResultsInterface
}

func (o *CustomerCustomerRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/customers/search][%d] customerCustomerRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerDataCustomerSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetListGetUnauthorized creates a CustomerCustomerRepositoryV1GetListGetUnauthorized with default headers values
func NewCustomerCustomerRepositoryV1GetListGetUnauthorized() *CustomerCustomerRepositoryV1GetListGetUnauthorized {
	return &CustomerCustomerRepositoryV1GetListGetUnauthorized{}
}

/*CustomerCustomerRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerCustomerRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/customers/search][%d] customerCustomerRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetListGetInternalServerError creates a CustomerCustomerRepositoryV1GetListGetInternalServerError with default headers values
func NewCustomerCustomerRepositoryV1GetListGetInternalServerError() *CustomerCustomerRepositoryV1GetListGetInternalServerError {
	return &CustomerCustomerRepositoryV1GetListGetInternalServerError{}
}

/*CustomerCustomerRepositoryV1GetListGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerCustomerRepositoryV1GetListGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1GetListGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/customers/search][%d] customerCustomerRepositoryV1GetListGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetListGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1GetListGetDefault creates a CustomerCustomerRepositoryV1GetListGetDefault with default headers values
func NewCustomerCustomerRepositoryV1GetListGetDefault(code int) *CustomerCustomerRepositoryV1GetListGetDefault {
	return &CustomerCustomerRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CustomerCustomerRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerCustomerRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer customer repository v1 get list get default response
func (o *CustomerCustomerRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerCustomerRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/customers/search][%d] customerCustomerRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCustomerRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
