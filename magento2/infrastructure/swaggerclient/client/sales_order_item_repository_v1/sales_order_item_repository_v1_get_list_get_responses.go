// Code generated by go-swagger; DO NOT EDIT.

package sales_order_item_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesOrderItemRepositoryV1GetListGetReader is a Reader for the SalesOrderItemRepositoryV1GetListGet structure.
type SalesOrderItemRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesOrderItemRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesOrderItemRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesOrderItemRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesOrderItemRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesOrderItemRepositoryV1GetListGetOK creates a SalesOrderItemRepositoryV1GetListGetOK with default headers values
func NewSalesOrderItemRepositoryV1GetListGetOK() *SalesOrderItemRepositoryV1GetListGetOK {
	return &SalesOrderItemRepositoryV1GetListGetOK{}
}

/*SalesOrderItemRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type SalesOrderItemRepositoryV1GetListGetOK struct {
	Payload *models.SalesDataOrderItemSearchResultInterface
}

func (o *SalesOrderItemRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/orders/items][%d] salesOrderItemRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *SalesOrderItemRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataOrderItemSearchResultInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderItemRepositoryV1GetListGetUnauthorized creates a SalesOrderItemRepositoryV1GetListGetUnauthorized with default headers values
func NewSalesOrderItemRepositoryV1GetListGetUnauthorized() *SalesOrderItemRepositoryV1GetListGetUnauthorized {
	return &SalesOrderItemRepositoryV1GetListGetUnauthorized{}
}

/*SalesOrderItemRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesOrderItemRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesOrderItemRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/orders/items][%d] salesOrderItemRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesOrderItemRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderItemRepositoryV1GetListGetDefault creates a SalesOrderItemRepositoryV1GetListGetDefault with default headers values
func NewSalesOrderItemRepositoryV1GetListGetDefault(code int) *SalesOrderItemRepositoryV1GetListGetDefault {
	return &SalesOrderItemRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*SalesOrderItemRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type SalesOrderItemRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales order item repository v1 get list get default response
func (o *SalesOrderItemRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesOrderItemRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/orders/items][%d] salesOrderItemRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesOrderItemRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
