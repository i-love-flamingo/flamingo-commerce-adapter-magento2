// Code generated by go-swagger; DO NOT EDIT.

package sales_order_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesOrderRepositoryV1GetGetReader is a Reader for the SalesOrderRepositoryV1GetGet structure.
type SalesOrderRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesOrderRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesOrderRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesOrderRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesOrderRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesOrderRepositoryV1GetGetOK creates a SalesOrderRepositoryV1GetGetOK with default headers values
func NewSalesOrderRepositoryV1GetGetOK() *SalesOrderRepositoryV1GetGetOK {
	return &SalesOrderRepositoryV1GetGetOK{}
}

/*SalesOrderRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type SalesOrderRepositoryV1GetGetOK struct {
	Payload *models.SalesDataOrderInterface
}

func (o *SalesOrderRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/orders/{id}][%d] salesOrderRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *SalesOrderRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataOrderInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderRepositoryV1GetGetUnauthorized creates a SalesOrderRepositoryV1GetGetUnauthorized with default headers values
func NewSalesOrderRepositoryV1GetGetUnauthorized() *SalesOrderRepositoryV1GetGetUnauthorized {
	return &SalesOrderRepositoryV1GetGetUnauthorized{}
}

/*SalesOrderRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesOrderRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesOrderRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/orders/{id}][%d] salesOrderRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesOrderRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderRepositoryV1GetGetDefault creates a SalesOrderRepositoryV1GetGetDefault with default headers values
func NewSalesOrderRepositoryV1GetGetDefault(code int) *SalesOrderRepositoryV1GetGetDefault {
	return &SalesOrderRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*SalesOrderRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type SalesOrderRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales order repository v1 get get default response
func (o *SalesOrderRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesOrderRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/orders/{id}][%d] salesOrderRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesOrderRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
