// Code generated by go-swagger; DO NOT EDIT.

package sales_creditmemo_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesCreditmemoRepositoryV1GetGetReader is a Reader for the SalesCreditmemoRepositoryV1GetGet structure.
type SalesCreditmemoRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesCreditmemoRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesCreditmemoRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesCreditmemoRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesCreditmemoRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesCreditmemoRepositoryV1GetGetOK creates a SalesCreditmemoRepositoryV1GetGetOK with default headers values
func NewSalesCreditmemoRepositoryV1GetGetOK() *SalesCreditmemoRepositoryV1GetGetOK {
	return &SalesCreditmemoRepositoryV1GetGetOK{}
}

/*SalesCreditmemoRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type SalesCreditmemoRepositoryV1GetGetOK struct {
	Payload *models.SalesDataCreditmemoInterface
}

func (o *SalesCreditmemoRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/creditmemo/{id}][%d] salesCreditmemoRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *SalesCreditmemoRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataCreditmemoInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoRepositoryV1GetGetUnauthorized creates a SalesCreditmemoRepositoryV1GetGetUnauthorized with default headers values
func NewSalesCreditmemoRepositoryV1GetGetUnauthorized() *SalesCreditmemoRepositoryV1GetGetUnauthorized {
	return &SalesCreditmemoRepositoryV1GetGetUnauthorized{}
}

/*SalesCreditmemoRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesCreditmemoRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesCreditmemoRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/creditmemo/{id}][%d] salesCreditmemoRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesCreditmemoRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoRepositoryV1GetGetDefault creates a SalesCreditmemoRepositoryV1GetGetDefault with default headers values
func NewSalesCreditmemoRepositoryV1GetGetDefault(code int) *SalesCreditmemoRepositoryV1GetGetDefault {
	return &SalesCreditmemoRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*SalesCreditmemoRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type SalesCreditmemoRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales creditmemo repository v1 get get default response
func (o *SalesCreditmemoRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesCreditmemoRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/creditmemo/{id}][%d] salesCreditmemoRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesCreditmemoRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
