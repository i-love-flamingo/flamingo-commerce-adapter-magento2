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

// SalesShipmentManagementV1GetLabelGetReader is a Reader for the SalesShipmentManagementV1GetLabelGet structure.
type SalesShipmentManagementV1GetLabelGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesShipmentManagementV1GetLabelGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesShipmentManagementV1GetLabelGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesShipmentManagementV1GetLabelGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesShipmentManagementV1GetLabelGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesShipmentManagementV1GetLabelGetOK creates a SalesShipmentManagementV1GetLabelGetOK with default headers values
func NewSalesShipmentManagementV1GetLabelGetOK() *SalesShipmentManagementV1GetLabelGetOK {
	return &SalesShipmentManagementV1GetLabelGetOK{}
}

/*SalesShipmentManagementV1GetLabelGetOK handles this case with default header values.

200 Success.
*/
type SalesShipmentManagementV1GetLabelGetOK struct {
	Payload string
}

func (o *SalesShipmentManagementV1GetLabelGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/label][%d] salesShipmentManagementV1GetLabelGetOK  %+v", 200, o.Payload)
}

func (o *SalesShipmentManagementV1GetLabelGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentManagementV1GetLabelGetUnauthorized creates a SalesShipmentManagementV1GetLabelGetUnauthorized with default headers values
func NewSalesShipmentManagementV1GetLabelGetUnauthorized() *SalesShipmentManagementV1GetLabelGetUnauthorized {
	return &SalesShipmentManagementV1GetLabelGetUnauthorized{}
}

/*SalesShipmentManagementV1GetLabelGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesShipmentManagementV1GetLabelGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentManagementV1GetLabelGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/label][%d] salesShipmentManagementV1GetLabelGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesShipmentManagementV1GetLabelGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentManagementV1GetLabelGetDefault creates a SalesShipmentManagementV1GetLabelGetDefault with default headers values
func NewSalesShipmentManagementV1GetLabelGetDefault(code int) *SalesShipmentManagementV1GetLabelGetDefault {
	return &SalesShipmentManagementV1GetLabelGetDefault{
		_statusCode: code,
	}
}

/*SalesShipmentManagementV1GetLabelGetDefault handles this case with default header values.

Unexpected error
*/
type SalesShipmentManagementV1GetLabelGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales shipment management v1 get label get default response
func (o *SalesShipmentManagementV1GetLabelGetDefault) Code() int {
	return o._statusCode
}

func (o *SalesShipmentManagementV1GetLabelGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/shipment/{id}/label][%d] salesShipmentManagementV1GetLabelGet default  %+v", o._statusCode, o.Payload)
}

func (o *SalesShipmentManagementV1GetLabelGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
