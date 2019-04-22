// Code generated by go-swagger; DO NOT EDIT.

package customer_address_metadata_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CustomerAddressMetadataV1GetAllAttributesMetadataGetReader is a Reader for the CustomerAddressMetadataV1GetAllAttributesMetadataGet structure.
type CustomerAddressMetadataV1GetAllAttributesMetadataGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAddressMetadataV1GetAllAttributesMetadataGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAddressMetadataV1GetAllAttributesMetadataGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetOK creates a CustomerAddressMetadataV1GetAllAttributesMetadataGetOK with default headers values
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetOK() *CustomerAddressMetadataV1GetAllAttributesMetadataGetOK {
	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetOK{}
}

/*CustomerAddressMetadataV1GetAllAttributesMetadataGetOK handles this case with default header values.

200 Success.
*/
type CustomerAddressMetadataV1GetAllAttributesMetadataGetOK struct {
	Payload []*models.CustomerDataAttributeMetadataInterface
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress][%d] customerAddressMetadataV1GetAllAttributesMetadataGetOK  %+v", 200, o.Payload)
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized creates a CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized with default headers values
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized() *CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized {
	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized{}
}

/*CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress][%d] customerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError creates a CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError with default headers values
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError() *CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError {
	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError{}
}

/*CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress][%d] customerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetAllAttributesMetadataGetDefault creates a CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault with default headers values
func NewCustomerAddressMetadataV1GetAllAttributesMetadataGetDefault(code int) *CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault {
	return &CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault{
		_statusCode: code,
	}
}

/*CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer address metadata v1 get all attributes metadata get default response
func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress][%d] customerAddressMetadataV1GetAllAttributesMetadataGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAddressMetadataV1GetAllAttributesMetadataGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
