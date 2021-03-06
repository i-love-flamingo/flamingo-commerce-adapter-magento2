// Code generated by go-swagger; DO NOT EDIT.

package customer_group_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CustomerGroupManagementV1GetDefaultGroupGetReader is a Reader for the CustomerGroupManagementV1GetDefaultGroupGet structure.
type CustomerGroupManagementV1GetDefaultGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGroupManagementV1GetDefaultGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerGroupManagementV1GetDefaultGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomerGroupManagementV1GetDefaultGroupGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCustomerGroupManagementV1GetDefaultGroupGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerGroupManagementV1GetDefaultGroupGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerGroupManagementV1GetDefaultGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGroupManagementV1GetDefaultGroupGetOK creates a CustomerGroupManagementV1GetDefaultGroupGetOK with default headers values
func NewCustomerGroupManagementV1GetDefaultGroupGetOK() *CustomerGroupManagementV1GetDefaultGroupGetOK {
	return &CustomerGroupManagementV1GetDefaultGroupGetOK{}
}

/*CustomerGroupManagementV1GetDefaultGroupGetOK handles this case with default header values.

200 Success.
*/
type CustomerGroupManagementV1GetDefaultGroupGetOK struct {
	Payload *models.CustomerDataGroupInterface
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/customerGroups/default][%d] customerGroupManagementV1GetDefaultGroupGetOK  %+v", 200, o.Payload)
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerDataGroupInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGroupManagementV1GetDefaultGroupGetBadRequest creates a CustomerGroupManagementV1GetDefaultGroupGetBadRequest with default headers values
func NewCustomerGroupManagementV1GetDefaultGroupGetBadRequest() *CustomerGroupManagementV1GetDefaultGroupGetBadRequest {
	return &CustomerGroupManagementV1GetDefaultGroupGetBadRequest{}
}

/*CustomerGroupManagementV1GetDefaultGroupGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CustomerGroupManagementV1GetDefaultGroupGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/customerGroups/default][%d] customerGroupManagementV1GetDefaultGroupGetBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGroupManagementV1GetDefaultGroupGetUnauthorized creates a CustomerGroupManagementV1GetDefaultGroupGetUnauthorized with default headers values
func NewCustomerGroupManagementV1GetDefaultGroupGetUnauthorized() *CustomerGroupManagementV1GetDefaultGroupGetUnauthorized {
	return &CustomerGroupManagementV1GetDefaultGroupGetUnauthorized{}
}

/*CustomerGroupManagementV1GetDefaultGroupGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerGroupManagementV1GetDefaultGroupGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/customerGroups/default][%d] customerGroupManagementV1GetDefaultGroupGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGroupManagementV1GetDefaultGroupGetInternalServerError creates a CustomerGroupManagementV1GetDefaultGroupGetInternalServerError with default headers values
func NewCustomerGroupManagementV1GetDefaultGroupGetInternalServerError() *CustomerGroupManagementV1GetDefaultGroupGetInternalServerError {
	return &CustomerGroupManagementV1GetDefaultGroupGetInternalServerError{}
}

/*CustomerGroupManagementV1GetDefaultGroupGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerGroupManagementV1GetDefaultGroupGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/customerGroups/default][%d] customerGroupManagementV1GetDefaultGroupGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGroupManagementV1GetDefaultGroupGetDefault creates a CustomerGroupManagementV1GetDefaultGroupGetDefault with default headers values
func NewCustomerGroupManagementV1GetDefaultGroupGetDefault(code int) *CustomerGroupManagementV1GetDefaultGroupGetDefault {
	return &CustomerGroupManagementV1GetDefaultGroupGetDefault{
		_statusCode: code,
	}
}

/*CustomerGroupManagementV1GetDefaultGroupGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerGroupManagementV1GetDefaultGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer group management v1 get default group get default response
func (o *CustomerGroupManagementV1GetDefaultGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/customerGroups/default][%d] customerGroupManagementV1GetDefaultGroupGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGroupManagementV1GetDefaultGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
