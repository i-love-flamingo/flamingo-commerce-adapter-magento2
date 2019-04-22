// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_rate_repository_v1

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

// TaxTaxRateRepositoryV1SavePostReader is a Reader for the TaxTaxRateRepositoryV1SavePost structure.
type TaxTaxRateRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaxTaxRateRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTaxTaxRateRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTaxTaxRateRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTaxTaxRateRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTaxTaxRateRepositoryV1SavePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTaxTaxRateRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaxTaxRateRepositoryV1SavePostOK creates a TaxTaxRateRepositoryV1SavePostOK with default headers values
func NewTaxTaxRateRepositoryV1SavePostOK() *TaxTaxRateRepositoryV1SavePostOK {
	return &TaxTaxRateRepositoryV1SavePostOK{}
}

/*TaxTaxRateRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type TaxTaxRateRepositoryV1SavePostOK struct {
	Payload *models.TaxDataTaxRateInterface
}

func (o *TaxTaxRateRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/taxRates][%d] taxTaxRateRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxDataTaxRateInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePostBadRequest creates a TaxTaxRateRepositoryV1SavePostBadRequest with default headers values
func NewTaxTaxRateRepositoryV1SavePostBadRequest() *TaxTaxRateRepositoryV1SavePostBadRequest {
	return &TaxTaxRateRepositoryV1SavePostBadRequest{}
}

/*TaxTaxRateRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type TaxTaxRateRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/taxRates][%d] taxTaxRateRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePostUnauthorized creates a TaxTaxRateRepositoryV1SavePostUnauthorized with default headers values
func NewTaxTaxRateRepositoryV1SavePostUnauthorized() *TaxTaxRateRepositoryV1SavePostUnauthorized {
	return &TaxTaxRateRepositoryV1SavePostUnauthorized{}
}

/*TaxTaxRateRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type TaxTaxRateRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/taxRates][%d] taxTaxRateRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePostInternalServerError creates a TaxTaxRateRepositoryV1SavePostInternalServerError with default headers values
func NewTaxTaxRateRepositoryV1SavePostInternalServerError() *TaxTaxRateRepositoryV1SavePostInternalServerError {
	return &TaxTaxRateRepositoryV1SavePostInternalServerError{}
}

/*TaxTaxRateRepositoryV1SavePostInternalServerError handles this case with default header values.

Internal Server error
*/
type TaxTaxRateRepositoryV1SavePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/taxRates][%d] taxTaxRateRepositoryV1SavePostInternalServerError  %+v", 500, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePostDefault creates a TaxTaxRateRepositoryV1SavePostDefault with default headers values
func NewTaxTaxRateRepositoryV1SavePostDefault(code int) *TaxTaxRateRepositoryV1SavePostDefault {
	return &TaxTaxRateRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*TaxTaxRateRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type TaxTaxRateRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the tax tax rate repository v1 save post default response
func (o *TaxTaxRateRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *TaxTaxRateRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/taxRates][%d] taxTaxRateRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TaxTaxRateRepositoryV1SavePostBody tax tax rate repository v1 save post body
swagger:model TaxTaxRateRepositoryV1SavePostBody
*/
type TaxTaxRateRepositoryV1SavePostBody struct {

	// tax rate
	// Required: true
	TaxRate *models.TaxDataTaxRateInterface `json:"taxRate"`
}

// Validate validates this tax tax rate repository v1 save post body
func (o *TaxTaxRateRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTaxRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TaxTaxRateRepositoryV1SavePostBody) validateTaxRate(formats strfmt.Registry) error {

	if err := validate.Required("taxTaxRateRepositoryV1SavePostBody"+"."+"taxRate", "body", o.TaxRate); err != nil {
		return err
	}

	if o.TaxRate != nil {
		if err := o.TaxRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxTaxRateRepositoryV1SavePostBody" + "." + "taxRate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TaxTaxRateRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxTaxRateRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res TaxTaxRateRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}