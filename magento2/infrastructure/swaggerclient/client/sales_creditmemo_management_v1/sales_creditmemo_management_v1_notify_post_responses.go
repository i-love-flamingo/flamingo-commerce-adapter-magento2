// Code generated by go-swagger; DO NOT EDIT.

package sales_creditmemo_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesCreditmemoManagementV1NotifyPostReader is a Reader for the SalesCreditmemoManagementV1NotifyPost structure.
type SalesCreditmemoManagementV1NotifyPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesCreditmemoManagementV1NotifyPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesCreditmemoManagementV1NotifyPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesCreditmemoManagementV1NotifyPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesCreditmemoManagementV1NotifyPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesCreditmemoManagementV1NotifyPostOK creates a SalesCreditmemoManagementV1NotifyPostOK with default headers values
func NewSalesCreditmemoManagementV1NotifyPostOK() *SalesCreditmemoManagementV1NotifyPostOK {
	return &SalesCreditmemoManagementV1NotifyPostOK{}
}

/*SalesCreditmemoManagementV1NotifyPostOK handles this case with default header values.

200 Success.
*/
type SalesCreditmemoManagementV1NotifyPostOK struct {
	Payload bool
}

func (o *SalesCreditmemoManagementV1NotifyPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/{id}/emails][%d] salesCreditmemoManagementV1NotifyPostOK  %+v", 200, o.Payload)
}

func (o *SalesCreditmemoManagementV1NotifyPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoManagementV1NotifyPostUnauthorized creates a SalesCreditmemoManagementV1NotifyPostUnauthorized with default headers values
func NewSalesCreditmemoManagementV1NotifyPostUnauthorized() *SalesCreditmemoManagementV1NotifyPostUnauthorized {
	return &SalesCreditmemoManagementV1NotifyPostUnauthorized{}
}

/*SalesCreditmemoManagementV1NotifyPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesCreditmemoManagementV1NotifyPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesCreditmemoManagementV1NotifyPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/{id}/emails][%d] salesCreditmemoManagementV1NotifyPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesCreditmemoManagementV1NotifyPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoManagementV1NotifyPostDefault creates a SalesCreditmemoManagementV1NotifyPostDefault with default headers values
func NewSalesCreditmemoManagementV1NotifyPostDefault(code int) *SalesCreditmemoManagementV1NotifyPostDefault {
	return &SalesCreditmemoManagementV1NotifyPostDefault{
		_statusCode: code,
	}
}

/*SalesCreditmemoManagementV1NotifyPostDefault handles this case with default header values.

Unexpected error
*/
type SalesCreditmemoManagementV1NotifyPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales creditmemo management v1 notify post default response
func (o *SalesCreditmemoManagementV1NotifyPostDefault) Code() int {
	return o._statusCode
}

func (o *SalesCreditmemoManagementV1NotifyPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/{id}/emails][%d] salesCreditmemoManagementV1NotifyPost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesCreditmemoManagementV1NotifyPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}