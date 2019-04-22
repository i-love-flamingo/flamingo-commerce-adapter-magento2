// Code generated by go-swagger; DO NOT EDIT.

package rma_rma_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// RmaRmaRepositoryV1DeleteDeleteReader is a Reader for the RmaRmaRepositoryV1DeleteDelete structure.
type RmaRmaRepositoryV1DeleteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RmaRmaRepositoryV1DeleteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRmaRmaRepositoryV1DeleteDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRmaRmaRepositoryV1DeleteDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRmaRmaRepositoryV1DeleteDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRmaRmaRepositoryV1DeleteDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRmaRmaRepositoryV1DeleteDeleteOK creates a RmaRmaRepositoryV1DeleteDeleteOK with default headers values
func NewRmaRmaRepositoryV1DeleteDeleteOK() *RmaRmaRepositoryV1DeleteDeleteOK {
	return &RmaRmaRepositoryV1DeleteDeleteOK{}
}

/*RmaRmaRepositoryV1DeleteDeleteOK handles this case with default header values.

200 Success.
*/
type RmaRmaRepositoryV1DeleteDeleteOK struct {
	Payload bool
}

func (o *RmaRmaRepositoryV1DeleteDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/returns/{id}][%d] rmaRmaRepositoryV1DeleteDeleteOK  %+v", 200, o.Payload)
}

func (o *RmaRmaRepositoryV1DeleteDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaRmaRepositoryV1DeleteDeleteBadRequest creates a RmaRmaRepositoryV1DeleteDeleteBadRequest with default headers values
func NewRmaRmaRepositoryV1DeleteDeleteBadRequest() *RmaRmaRepositoryV1DeleteDeleteBadRequest {
	return &RmaRmaRepositoryV1DeleteDeleteBadRequest{}
}

/*RmaRmaRepositoryV1DeleteDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type RmaRmaRepositoryV1DeleteDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *RmaRmaRepositoryV1DeleteDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/returns/{id}][%d] rmaRmaRepositoryV1DeleteDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *RmaRmaRepositoryV1DeleteDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaRmaRepositoryV1DeleteDeleteUnauthorized creates a RmaRmaRepositoryV1DeleteDeleteUnauthorized with default headers values
func NewRmaRmaRepositoryV1DeleteDeleteUnauthorized() *RmaRmaRepositoryV1DeleteDeleteUnauthorized {
	return &RmaRmaRepositoryV1DeleteDeleteUnauthorized{}
}

/*RmaRmaRepositoryV1DeleteDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type RmaRmaRepositoryV1DeleteDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RmaRmaRepositoryV1DeleteDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/returns/{id}][%d] rmaRmaRepositoryV1DeleteDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *RmaRmaRepositoryV1DeleteDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaRmaRepositoryV1DeleteDeleteDefault creates a RmaRmaRepositoryV1DeleteDeleteDefault with default headers values
func NewRmaRmaRepositoryV1DeleteDeleteDefault(code int) *RmaRmaRepositoryV1DeleteDeleteDefault {
	return &RmaRmaRepositoryV1DeleteDeleteDefault{
		_statusCode: code,
	}
}

/*RmaRmaRepositoryV1DeleteDeleteDefault handles this case with default header values.

Unexpected error
*/
type RmaRmaRepositoryV1DeleteDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the rma rma repository v1 delete delete default response
func (o *RmaRmaRepositoryV1DeleteDeleteDefault) Code() int {
	return o._statusCode
}

func (o *RmaRmaRepositoryV1DeleteDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/returns/{id}][%d] rmaRmaRepositoryV1DeleteDelete default  %+v", o._statusCode, o.Payload)
}

func (o *RmaRmaRepositoryV1DeleteDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RmaRmaRepositoryV1DeleteDeleteBody rma rma repository v1 delete delete body
swagger:model RmaRmaRepositoryV1DeleteDeleteBody
*/
type RmaRmaRepositoryV1DeleteDeleteBody struct {

	// rma data object
	// Required: true
	RmaDataObject *models.RmaDataRmaInterface `json:"rmaDataObject"`
}

// Validate validates this rma rma repository v1 delete delete body
func (o *RmaRmaRepositoryV1DeleteDeleteBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRmaDataObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RmaRmaRepositoryV1DeleteDeleteBody) validateRmaDataObject(formats strfmt.Registry) error {

	if err := validate.Required("rmaRmaRepositoryV1DeleteDeleteBody"+"."+"rmaDataObject", "body", o.RmaDataObject); err != nil {
		return err
	}

	if o.RmaDataObject != nil {
		if err := o.RmaDataObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rmaRmaRepositoryV1DeleteDeleteBody" + "." + "rmaDataObject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RmaRmaRepositoryV1DeleteDeleteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RmaRmaRepositoryV1DeleteDeleteBody) UnmarshalBinary(b []byte) error {
	var res RmaRmaRepositoryV1DeleteDeleteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}