// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_class_repository_v1

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

// TaxTaxClassRepositoryV1SavePostReader is a Reader for the TaxTaxClassRepositoryV1SavePost structure.
type TaxTaxClassRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaxTaxClassRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTaxTaxClassRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTaxTaxClassRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTaxTaxClassRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTaxTaxClassRepositoryV1SavePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTaxTaxClassRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaxTaxClassRepositoryV1SavePostOK creates a TaxTaxClassRepositoryV1SavePostOK with default headers values
func NewTaxTaxClassRepositoryV1SavePostOK() *TaxTaxClassRepositoryV1SavePostOK {
	return &TaxTaxClassRepositoryV1SavePostOK{}
}

/*TaxTaxClassRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type TaxTaxClassRepositoryV1SavePostOK struct {
	Payload string
}

func (o *TaxTaxClassRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/taxClasses][%d] taxTaxClassRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *TaxTaxClassRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxClassRepositoryV1SavePostBadRequest creates a TaxTaxClassRepositoryV1SavePostBadRequest with default headers values
func NewTaxTaxClassRepositoryV1SavePostBadRequest() *TaxTaxClassRepositoryV1SavePostBadRequest {
	return &TaxTaxClassRepositoryV1SavePostBadRequest{}
}

/*TaxTaxClassRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type TaxTaxClassRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxClassRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/taxClasses][%d] taxTaxClassRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *TaxTaxClassRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxClassRepositoryV1SavePostUnauthorized creates a TaxTaxClassRepositoryV1SavePostUnauthorized with default headers values
func NewTaxTaxClassRepositoryV1SavePostUnauthorized() *TaxTaxClassRepositoryV1SavePostUnauthorized {
	return &TaxTaxClassRepositoryV1SavePostUnauthorized{}
}

/*TaxTaxClassRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type TaxTaxClassRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxClassRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/taxClasses][%d] taxTaxClassRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *TaxTaxClassRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxClassRepositoryV1SavePostInternalServerError creates a TaxTaxClassRepositoryV1SavePostInternalServerError with default headers values
func NewTaxTaxClassRepositoryV1SavePostInternalServerError() *TaxTaxClassRepositoryV1SavePostInternalServerError {
	return &TaxTaxClassRepositoryV1SavePostInternalServerError{}
}

/*TaxTaxClassRepositoryV1SavePostInternalServerError handles this case with default header values.

Internal Server error
*/
type TaxTaxClassRepositoryV1SavePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxClassRepositoryV1SavePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/taxClasses][%d] taxTaxClassRepositoryV1SavePostInternalServerError  %+v", 500, o.Payload)
}

func (o *TaxTaxClassRepositoryV1SavePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxClassRepositoryV1SavePostDefault creates a TaxTaxClassRepositoryV1SavePostDefault with default headers values
func NewTaxTaxClassRepositoryV1SavePostDefault(code int) *TaxTaxClassRepositoryV1SavePostDefault {
	return &TaxTaxClassRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*TaxTaxClassRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type TaxTaxClassRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the tax tax class repository v1 save post default response
func (o *TaxTaxClassRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *TaxTaxClassRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/taxClasses][%d] taxTaxClassRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *TaxTaxClassRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TaxTaxClassRepositoryV1SavePostBody tax tax class repository v1 save post body
swagger:model TaxTaxClassRepositoryV1SavePostBody
*/
type TaxTaxClassRepositoryV1SavePostBody struct {

	// tax class
	// Required: true
	TaxClass *models.TaxDataTaxClassInterface `json:"taxClass"`
}

// Validate validates this tax tax class repository v1 save post body
func (o *TaxTaxClassRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTaxClass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TaxTaxClassRepositoryV1SavePostBody) validateTaxClass(formats strfmt.Registry) error {

	if err := validate.Required("taxTaxClassRepositoryV1SavePostBody"+"."+"taxClass", "body", o.TaxClass); err != nil {
		return err
	}

	if o.TaxClass != nil {
		if err := o.TaxClass.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxTaxClassRepositoryV1SavePostBody" + "." + "taxClass")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TaxTaxClassRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxTaxClassRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res TaxTaxClassRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
