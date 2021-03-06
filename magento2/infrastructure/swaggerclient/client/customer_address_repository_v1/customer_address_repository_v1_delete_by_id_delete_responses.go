// Code generated by go-swagger; DO NOT EDIT.

package customer_address_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CustomerAddressRepositoryV1DeleteByIDDeleteReader is a Reader for the CustomerAddressRepositoryV1DeleteByIDDelete structure.
type CustomerAddressRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAddressRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAddressRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomerAddressRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAddressRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAddressRepositoryV1DeleteByIDDeleteOK creates a CustomerAddressRepositoryV1DeleteByIDDeleteOK with default headers values
func NewCustomerAddressRepositoryV1DeleteByIDDeleteOK() *CustomerAddressRepositoryV1DeleteByIDDeleteOK {
	return &CustomerAddressRepositoryV1DeleteByIDDeleteOK{}
}

/*CustomerAddressRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type CustomerAddressRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/addresses/{addressId}][%d] customerAddressRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressRepositoryV1DeleteByIDDeleteBadRequest creates a CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewCustomerAddressRepositoryV1DeleteByIDDeleteBadRequest() *CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest {
	return &CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/addresses/{addressId}][%d] customerAddressRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized creates a CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewCustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized() *CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized {
	return &CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/addresses/{addressId}][%d] customerAddressRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError creates a CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError with default headers values
func NewCustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError() *CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError {
	return &CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError{}
}

/*CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /V1/addresses/{addressId}][%d] customerAddressRepositoryV1DeleteByIdDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressRepositoryV1DeleteByIDDeleteDefault creates a CustomerAddressRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewCustomerAddressRepositoryV1DeleteByIDDeleteDefault(code int) *CustomerAddressRepositoryV1DeleteByIDDeleteDefault {
	return &CustomerAddressRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*CustomerAddressRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type CustomerAddressRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer address repository v1 delete by Id delete default response
func (o *CustomerAddressRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/addresses/{addressId}][%d] customerAddressRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAddressRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
