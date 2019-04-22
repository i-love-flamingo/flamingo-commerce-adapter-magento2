// Code generated by go-swagger; DO NOT EDIT.

package sales_invoice_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesInvoiceRepositoryV1GetGetReader is a Reader for the SalesInvoiceRepositoryV1GetGet structure.
type SalesInvoiceRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesInvoiceRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesInvoiceRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesInvoiceRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesInvoiceRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesInvoiceRepositoryV1GetGetOK creates a SalesInvoiceRepositoryV1GetGetOK with default headers values
func NewSalesInvoiceRepositoryV1GetGetOK() *SalesInvoiceRepositoryV1GetGetOK {
	return &SalesInvoiceRepositoryV1GetGetOK{}
}

/*SalesInvoiceRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type SalesInvoiceRepositoryV1GetGetOK struct {
	Payload *models.SalesDataInvoiceInterface
}

func (o *SalesInvoiceRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/invoices/{id}][%d] salesInvoiceRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *SalesInvoiceRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataInvoiceInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesInvoiceRepositoryV1GetGetUnauthorized creates a SalesInvoiceRepositoryV1GetGetUnauthorized with default headers values
func NewSalesInvoiceRepositoryV1GetGetUnauthorized() *SalesInvoiceRepositoryV1GetGetUnauthorized {
	return &SalesInvoiceRepositoryV1GetGetUnauthorized{}
}

/*SalesInvoiceRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesInvoiceRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesInvoiceRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/invoices/{id}][%d] salesInvoiceRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesInvoiceRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesInvoiceRepositoryV1GetGetDefault creates a SalesInvoiceRepositoryV1GetGetDefault with default headers values
func NewSalesInvoiceRepositoryV1GetGetDefault(code int) *SalesInvoiceRepositoryV1GetGetDefault {
	return &SalesInvoiceRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*SalesInvoiceRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type SalesInvoiceRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales invoice repository v1 get get default response
func (o *SalesInvoiceRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesInvoiceRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/invoices/{id}][%d] salesInvoiceRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesInvoiceRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}