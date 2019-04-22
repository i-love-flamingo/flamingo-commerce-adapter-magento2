// Code generated by go-swagger; DO NOT EDIT.

package sales_order_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesOrderManagementV1UnHoldPostReader is a Reader for the SalesOrderManagementV1UnHoldPost structure.
type SalesOrderManagementV1UnHoldPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesOrderManagementV1UnHoldPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesOrderManagementV1UnHoldPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesOrderManagementV1UnHoldPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesOrderManagementV1UnHoldPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesOrderManagementV1UnHoldPostOK creates a SalesOrderManagementV1UnHoldPostOK with default headers values
func NewSalesOrderManagementV1UnHoldPostOK() *SalesOrderManagementV1UnHoldPostOK {
	return &SalesOrderManagementV1UnHoldPostOK{}
}

/*SalesOrderManagementV1UnHoldPostOK handles this case with default header values.

200 Success.
*/
type SalesOrderManagementV1UnHoldPostOK struct {
	Payload bool
}

func (o *SalesOrderManagementV1UnHoldPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/unhold][%d] salesOrderManagementV1UnHoldPostOK  %+v", 200, o.Payload)
}

func (o *SalesOrderManagementV1UnHoldPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderManagementV1UnHoldPostUnauthorized creates a SalesOrderManagementV1UnHoldPostUnauthorized with default headers values
func NewSalesOrderManagementV1UnHoldPostUnauthorized() *SalesOrderManagementV1UnHoldPostUnauthorized {
	return &SalesOrderManagementV1UnHoldPostUnauthorized{}
}

/*SalesOrderManagementV1UnHoldPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesOrderManagementV1UnHoldPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesOrderManagementV1UnHoldPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/unhold][%d] salesOrderManagementV1UnHoldPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesOrderManagementV1UnHoldPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderManagementV1UnHoldPostDefault creates a SalesOrderManagementV1UnHoldPostDefault with default headers values
func NewSalesOrderManagementV1UnHoldPostDefault(code int) *SalesOrderManagementV1UnHoldPostDefault {
	return &SalesOrderManagementV1UnHoldPostDefault{
		_statusCode: code,
	}
}

/*SalesOrderManagementV1UnHoldPostDefault handles this case with default header values.

Unexpected error
*/
type SalesOrderManagementV1UnHoldPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales order management v1 un hold post default response
func (o *SalesOrderManagementV1UnHoldPostDefault) Code() int {
	return o._statusCode
}

func (o *SalesOrderManagementV1UnHoldPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/unhold][%d] salesOrderManagementV1UnHoldPost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesOrderManagementV1UnHoldPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}