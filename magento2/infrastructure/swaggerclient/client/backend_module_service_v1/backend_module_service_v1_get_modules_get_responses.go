// Code generated by go-swagger; DO NOT EDIT.

package backend_module_service_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// BackendModuleServiceV1GetModulesGetReader is a Reader for the BackendModuleServiceV1GetModulesGet structure.
type BackendModuleServiceV1GetModulesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackendModuleServiceV1GetModulesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBackendModuleServiceV1GetModulesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewBackendModuleServiceV1GetModulesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewBackendModuleServiceV1GetModulesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBackendModuleServiceV1GetModulesGetOK creates a BackendModuleServiceV1GetModulesGetOK with default headers values
func NewBackendModuleServiceV1GetModulesGetOK() *BackendModuleServiceV1GetModulesGetOK {
	return &BackendModuleServiceV1GetModulesGetOK{}
}

/*BackendModuleServiceV1GetModulesGetOK handles this case with default header values.

200 Success.
*/
type BackendModuleServiceV1GetModulesGetOK struct {
	Payload []string
}

func (o *BackendModuleServiceV1GetModulesGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/modules][%d] backendModuleServiceV1GetModulesGetOK  %+v", 200, o.Payload)
}

func (o *BackendModuleServiceV1GetModulesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackendModuleServiceV1GetModulesGetUnauthorized creates a BackendModuleServiceV1GetModulesGetUnauthorized with default headers values
func NewBackendModuleServiceV1GetModulesGetUnauthorized() *BackendModuleServiceV1GetModulesGetUnauthorized {
	return &BackendModuleServiceV1GetModulesGetUnauthorized{}
}

/*BackendModuleServiceV1GetModulesGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type BackendModuleServiceV1GetModulesGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *BackendModuleServiceV1GetModulesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/modules][%d] backendModuleServiceV1GetModulesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *BackendModuleServiceV1GetModulesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackendModuleServiceV1GetModulesGetDefault creates a BackendModuleServiceV1GetModulesGetDefault with default headers values
func NewBackendModuleServiceV1GetModulesGetDefault(code int) *BackendModuleServiceV1GetModulesGetDefault {
	return &BackendModuleServiceV1GetModulesGetDefault{
		_statusCode: code,
	}
}

/*BackendModuleServiceV1GetModulesGetDefault handles this case with default header values.

Unexpected error
*/
type BackendModuleServiceV1GetModulesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the backend module service v1 get modules get default response
func (o *BackendModuleServiceV1GetModulesGetDefault) Code() int {
	return o._statusCode
}

func (o *BackendModuleServiceV1GetModulesGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/modules][%d] backendModuleServiceV1GetModulesGet default  %+v", o._statusCode, o.Payload)
}

func (o *BackendModuleServiceV1GetModulesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
