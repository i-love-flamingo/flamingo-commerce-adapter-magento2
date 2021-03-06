// Code generated by go-swagger; DO NOT EDIT.

package sales_invoice_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesInvoiceManagementV1SetCapturePostReader is a Reader for the SalesInvoiceManagementV1SetCapturePost structure.
type SalesInvoiceManagementV1SetCapturePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesInvoiceManagementV1SetCapturePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesInvoiceManagementV1SetCapturePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesInvoiceManagementV1SetCapturePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesInvoiceManagementV1SetCapturePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesInvoiceManagementV1SetCapturePostOK creates a SalesInvoiceManagementV1SetCapturePostOK with default headers values
func NewSalesInvoiceManagementV1SetCapturePostOK() *SalesInvoiceManagementV1SetCapturePostOK {
	return &SalesInvoiceManagementV1SetCapturePostOK{}
}

/*SalesInvoiceManagementV1SetCapturePostOK handles this case with default header values.

200 Success.
*/
type SalesInvoiceManagementV1SetCapturePostOK struct {
	Payload string
}

func (o *SalesInvoiceManagementV1SetCapturePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/invoices/{id}/capture][%d] salesInvoiceManagementV1SetCapturePostOK  %+v", 200, o.Payload)
}

func (o *SalesInvoiceManagementV1SetCapturePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesInvoiceManagementV1SetCapturePostUnauthorized creates a SalesInvoiceManagementV1SetCapturePostUnauthorized with default headers values
func NewSalesInvoiceManagementV1SetCapturePostUnauthorized() *SalesInvoiceManagementV1SetCapturePostUnauthorized {
	return &SalesInvoiceManagementV1SetCapturePostUnauthorized{}
}

/*SalesInvoiceManagementV1SetCapturePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesInvoiceManagementV1SetCapturePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesInvoiceManagementV1SetCapturePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/invoices/{id}/capture][%d] salesInvoiceManagementV1SetCapturePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesInvoiceManagementV1SetCapturePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesInvoiceManagementV1SetCapturePostDefault creates a SalesInvoiceManagementV1SetCapturePostDefault with default headers values
func NewSalesInvoiceManagementV1SetCapturePostDefault(code int) *SalesInvoiceManagementV1SetCapturePostDefault {
	return &SalesInvoiceManagementV1SetCapturePostDefault{
		_statusCode: code,
	}
}

/*SalesInvoiceManagementV1SetCapturePostDefault handles this case with default header values.

Unexpected error
*/
type SalesInvoiceManagementV1SetCapturePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales invoice management v1 set capture post default response
func (o *SalesInvoiceManagementV1SetCapturePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesInvoiceManagementV1SetCapturePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/invoices/{id}/capture][%d] salesInvoiceManagementV1SetCapturePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesInvoiceManagementV1SetCapturePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
