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

// CustomerAddressMetadataV1GetCustomAttributesMetadataGetReader is a Reader for the CustomerAddressMetadataV1GetCustomAttributesMetadataGet structure.
type CustomerAddressMetadataV1GetCustomAttributesMetadataGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetOK creates a CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK with default headers values
func NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetOK() *CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK {
	return &CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK{}
}

/*CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK handles this case with default header values.

200 Success.
*/
type CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK struct {
	Payload []*models.CustomerDataAttributeMetadataInterface
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress/custom][%d] customerAddressMetadataV1GetCustomAttributesMetadataGetOK  %+v", 200, o.Payload)
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized creates a CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized with default headers values
func NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized() *CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized {
	return &CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized{}
}

/*CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress/custom][%d] customerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError creates a CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError with default headers values
func NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError() *CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError {
	return &CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError{}
}

/*CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress/custom][%d] customerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault creates a CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault with default headers values
func NewCustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault(code int) *CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault {
	return &CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault{
		_statusCode: code,
	}
}

/*CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault handles this case with default header values.

Unexpected error
*/
type CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer address metadata v1 get custom attributes metadata get default response
func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/attributeMetadata/customerAddress/custom][%d] customerAddressMetadataV1GetCustomAttributesMetadataGet default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAddressMetadataV1GetCustomAttributesMetadataGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
