// Code generated by go-swagger; DO NOT EDIT.

package sales_shipment_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesShipmentManagementV1GetCommentsListGetReader is a Reader for the SalesShipmentManagementV1GetCommentsListGet structure.
type SalesShipmentManagementV1GetCommentsListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesShipmentManagementV1GetCommentsListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesShipmentManagementV1GetCommentsListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesShipmentManagementV1GetCommentsListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesShipmentManagementV1GetCommentsListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesShipmentManagementV1GetCommentsListGetOK creates a SalesShipmentManagementV1GetCommentsListGetOK with default headers values
func NewSalesShipmentManagementV1GetCommentsListGetOK() *SalesShipmentManagementV1GetCommentsListGetOK {
	return &SalesShipmentManagementV1GetCommentsListGetOK{}
}

/*SalesShipmentManagementV1GetCommentsListGetOK handles this case with default header values.

200 Success.
*/
type SalesShipmentManagementV1GetCommentsListGetOK struct {
	Payload *models.SalesDataShipmentCommentSearchResultInterface
}

func (o *SalesShipmentManagementV1GetCommentsListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/comments][%d] salesShipmentManagementV1GetCommentsListGetOK  %+v", 200, o.Payload)
}

func (o *SalesShipmentManagementV1GetCommentsListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataShipmentCommentSearchResultInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentManagementV1GetCommentsListGetUnauthorized creates a SalesShipmentManagementV1GetCommentsListGetUnauthorized with default headers values
func NewSalesShipmentManagementV1GetCommentsListGetUnauthorized() *SalesShipmentManagementV1GetCommentsListGetUnauthorized {
	return &SalesShipmentManagementV1GetCommentsListGetUnauthorized{}
}

/*SalesShipmentManagementV1GetCommentsListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesShipmentManagementV1GetCommentsListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentManagementV1GetCommentsListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/comments][%d] salesShipmentManagementV1GetCommentsListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesShipmentManagementV1GetCommentsListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentManagementV1GetCommentsListGetDefault creates a SalesShipmentManagementV1GetCommentsListGetDefault with default headers values
func NewSalesShipmentManagementV1GetCommentsListGetDefault(code int) *SalesShipmentManagementV1GetCommentsListGetDefault {
	return &SalesShipmentManagementV1GetCommentsListGetDefault{
		_statusCode: code,
	}
}

/*SalesShipmentManagementV1GetCommentsListGetDefault handles this case with default header values.

Unexpected error
*/
type SalesShipmentManagementV1GetCommentsListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales shipment management v1 get comments list get default response
func (o *SalesShipmentManagementV1GetCommentsListGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesShipmentManagementV1GetCommentsListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/comments][%d] salesShipmentManagementV1GetCommentsListGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesShipmentManagementV1GetCommentsListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
