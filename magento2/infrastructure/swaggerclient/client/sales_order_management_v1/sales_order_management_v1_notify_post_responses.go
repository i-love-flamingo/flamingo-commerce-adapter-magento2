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

// SalesOrderManagementV1NotifyPostReader is a Reader for the SalesOrderManagementV1NotifyPost structure.
type SalesOrderManagementV1NotifyPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesOrderManagementV1NotifyPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesOrderManagementV1NotifyPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesOrderManagementV1NotifyPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesOrderManagementV1NotifyPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesOrderManagementV1NotifyPostOK creates a SalesOrderManagementV1NotifyPostOK with default headers values
func NewSalesOrderManagementV1NotifyPostOK() *SalesOrderManagementV1NotifyPostOK {
	return &SalesOrderManagementV1NotifyPostOK{}
}

/*SalesOrderManagementV1NotifyPostOK handles this case with default header values.

200 Success.
*/
type SalesOrderManagementV1NotifyPostOK struct {
	Payload bool
}

func (o *SalesOrderManagementV1NotifyPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/emails][%d] salesOrderManagementV1NotifyPostOK  %+v", 200, o.Payload)
}

func (o *SalesOrderManagementV1NotifyPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderManagementV1NotifyPostUnauthorized creates a SalesOrderManagementV1NotifyPostUnauthorized with default headers values
func NewSalesOrderManagementV1NotifyPostUnauthorized() *SalesOrderManagementV1NotifyPostUnauthorized {
	return &SalesOrderManagementV1NotifyPostUnauthorized{}
}

/*SalesOrderManagementV1NotifyPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesOrderManagementV1NotifyPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesOrderManagementV1NotifyPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/emails][%d] salesOrderManagementV1NotifyPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesOrderManagementV1NotifyPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderManagementV1NotifyPostDefault creates a SalesOrderManagementV1NotifyPostDefault with default headers values
func NewSalesOrderManagementV1NotifyPostDefault(code int) *SalesOrderManagementV1NotifyPostDefault {
	return &SalesOrderManagementV1NotifyPostDefault{
		_statusCode: code,
	}
}

/*SalesOrderManagementV1NotifyPostDefault handles this case with default header values.

Unexpected error
*/
type SalesOrderManagementV1NotifyPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales order management v1 notify post default response
func (o *SalesOrderManagementV1NotifyPostDefault) Code() int {
	return o._statusCode
}

func (o *SalesOrderManagementV1NotifyPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/orders/{id}/emails][%d] salesOrderManagementV1NotifyPost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesOrderManagementV1NotifyPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
