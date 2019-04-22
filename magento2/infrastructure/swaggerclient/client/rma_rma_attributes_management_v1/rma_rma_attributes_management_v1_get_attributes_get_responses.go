// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_attributes_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// RmaRmaAttributesManagementV1GetAttributesGetReader is a Reader for the RmaRmaAttributesManagementV1GetAttributesGet structure.
type RmaRmaAttributesManagementV1GetAttributesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RmaRmaAttributesManagementV1GetAttributesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRmaRmaAttributesManagementV1GetAttributesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRmaRmaAttributesManagementV1GetAttributesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRmaRmaAttributesManagementV1GetAttributesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRmaRmaAttributesManagementV1GetAttributesGetOK creates a RmaRmaAttributesManagementV1GetAttributesGetOK with default headers values
func NewRmaRmaAttributesManagementV1GetAttributesGetOK() *RmaRmaAttributesManagementV1GetAttributesGetOK {
	return &RmaRmaAttributesManagementV1GetAttributesGetOK{}
}

/*RmaRmaAttributesManagementV1GetAttributesGetOK handles this case with default header values.

200 Success.
*/
type RmaRmaAttributesManagementV1GetAttributesGetOK struct {
	Payload []*models.CustomerDataAttributeMetadataInterface
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/returnsAttributeMetadata/form/{formCode}][%d] rmaRmaAttributesManagementV1GetAttributesGetOK  %+v", 200, o.Payload)
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaRmaAttributesManagementV1GetAttributesGetUnauthorized creates a RmaRmaAttributesManagementV1GetAttributesGetUnauthorized with default headers values
func NewRmaRmaAttributesManagementV1GetAttributesGetUnauthorized() *RmaRmaAttributesManagementV1GetAttributesGetUnauthorized {
	return &RmaRmaAttributesManagementV1GetAttributesGetUnauthorized{}
}

/*RmaRmaAttributesManagementV1GetAttributesGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type RmaRmaAttributesManagementV1GetAttributesGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/returnsAttributeMetadata/form/{formCode}][%d] rmaRmaAttributesManagementV1GetAttributesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaRmaAttributesManagementV1GetAttributesGetDefault creates a RmaRmaAttributesManagementV1GetAttributesGetDefault with default headers values
func NewRmaRmaAttributesManagementV1GetAttributesGetDefault(code int) *RmaRmaAttributesManagementV1GetAttributesGetDefault {
	return &RmaRmaAttributesManagementV1GetAttributesGetDefault{
		_statusCode: code,
	}
}

/*RmaRmaAttributesManagementV1GetAttributesGetDefault handles this case with default header values.

Unexpected error
*/
type RmaRmaAttributesManagementV1GetAttributesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the rma rma attributes management v1 get attributes get default response
func (o *RmaRmaAttributesManagementV1GetAttributesGetDefault) Code() int {
	return o._statusCode
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/returnsAttributeMetadata/form/{formCode}][%d] rmaRmaAttributesManagementV1GetAttributesGet default  %+v", o._statusCode, o.Payload)
}

func (o *RmaRmaAttributesManagementV1GetAttributesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
