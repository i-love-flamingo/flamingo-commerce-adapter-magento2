// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_metadata_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CustomerCustomerMetadataV1GetAttributesGetReader is a Reader for the CustomerCustomerMetadataV1GetAttributesGet structure.
type CustomerCustomerMetadataV1GetAttributesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCustomerMetadataV1GetAttributesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerCustomerMetadataV1GetAttributesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCustomerCustomerMetadataV1GetAttributesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerCustomerMetadataV1GetAttributesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerCustomerMetadataV1GetAttributesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerCustomerMetadataV1GetAttributesGetOK creates a CustomerCustomerMetadataV1GetAttributesGetOK with default headers values
func NewCustomerCustomerMetadataV1GetAttributesGetOK() *CustomerCustomerMetadataV1GetAttributesGetOK {
	return &CustomerCustomerMetadataV1GetAttributesGetOK{}
}

/*CustomerCustomerMetadataV1GetAttributesGetOK handles this case with default header values.

200 Success.
*/
type CustomerCustomerMetadataV1GetAttributesGetOK struct {
	Payload []*models.CustomerDataAttributeMetadataInterface
}

func (o *CustomerCustomerMetadataV1GetAttributesGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customer/form/{formCode}][%d] customerCustomerMetadataV1GetAttributesGetOK  %+v", 200, o.Payload)
}

func (o *CustomerCustomerMetadataV1GetAttributesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerMetadataV1GetAttributesGetUnauthorized creates a CustomerCustomerMetadataV1GetAttributesGetUnauthorized with default headers values
func NewCustomerCustomerMetadataV1GetAttributesGetUnauthorized() *CustomerCustomerMetadataV1GetAttributesGetUnauthorized {
	return &CustomerCustomerMetadataV1GetAttributesGetUnauthorized{}
}

/*CustomerCustomerMetadataV1GetAttributesGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerCustomerMetadataV1GetAttributesGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerMetadataV1GetAttributesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customer/form/{formCode}][%d] customerCustomerMetadataV1GetAttributesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerCustomerMetadataV1GetAttributesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerMetadataV1GetAttributesGetInternalServerError creates a CustomerCustomerMetadataV1GetAttributesGetInternalServerError with default headers values
func NewCustomerCustomerMetadataV1GetAttributesGetInternalServerError() *CustomerCustomerMetadataV1GetAttributesGetInternalServerError {
	return &CustomerCustomerMetadataV1GetAttributesGetInternalServerError{}
}

/*CustomerCustomerMetadataV1GetAttributesGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerCustomerMetadataV1GetAttributesGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerMetadataV1GetAttributesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customer/form/{formCode}][%d] customerCustomerMetadataV1GetAttributesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerCustomerMetadataV1GetAttributesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerMetadataV1GetAttributesGetDefault creates a CustomerCustomerMetadataV1GetAttributesGetDefault with default headers values
func NewCustomerCustomerMetadataV1GetAttributesGetDefault(code int) *CustomerCustomerMetadataV1GetAttributesGetDefault {
	return &CustomerCustomerMetadataV1GetAttributesGetDefault{
		_statusCode: code,
	}
}

/*CustomerCustomerMetadataV1GetAttributesGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerCustomerMetadataV1GetAttributesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer customer metadata v1 get attributes get default response
func (o *CustomerCustomerMetadataV1GetAttributesGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerCustomerMetadataV1GetAttributesGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customer/form/{formCode}][%d] customerCustomerMetadataV1GetAttributesGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCustomerMetadataV1GetAttributesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
