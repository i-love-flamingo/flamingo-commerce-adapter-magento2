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

// TaxTaxRateRepositoryV1SavePutReader is a Reader for the TaxTaxRateRepositoryV1SavePut structure.
type TaxTaxRateRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaxTaxRateRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTaxTaxRateRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTaxTaxRateRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTaxTaxRateRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTaxTaxRateRepositoryV1SavePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTaxTaxRateRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaxTaxRateRepositoryV1SavePutOK creates a TaxTaxRateRepositoryV1SavePutOK with default headers values
func NewTaxTaxRateRepositoryV1SavePutOK() *TaxTaxRateRepositoryV1SavePutOK {
	return &TaxTaxRateRepositoryV1SavePutOK{}
}

/*TaxTaxRateRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type TaxTaxRateRepositoryV1SavePutOK struct {
	Payload *models.TaxDataTaxRateInterface
}

func (o *TaxTaxRateRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/taxRates][%d] taxTaxRateRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxDataTaxRateInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePutBadRequest creates a TaxTaxRateRepositoryV1SavePutBadRequest with default headers values
func NewTaxTaxRateRepositoryV1SavePutBadRequest() *TaxTaxRateRepositoryV1SavePutBadRequest {
	return &TaxTaxRateRepositoryV1SavePutBadRequest{}
}

/*TaxTaxRateRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type TaxTaxRateRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/taxRates][%d] taxTaxRateRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePutUnauthorized creates a TaxTaxRateRepositoryV1SavePutUnauthorized with default headers values
func NewTaxTaxRateRepositoryV1SavePutUnauthorized() *TaxTaxRateRepositoryV1SavePutUnauthorized {
	return &TaxTaxRateRepositoryV1SavePutUnauthorized{}
}

/*TaxTaxRateRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type TaxTaxRateRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/taxRates][%d] taxTaxRateRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePutInternalServerError creates a TaxTaxRateRepositoryV1SavePutInternalServerError with default headers values
func NewTaxTaxRateRepositoryV1SavePutInternalServerError() *TaxTaxRateRepositoryV1SavePutInternalServerError {
	return &TaxTaxRateRepositoryV1SavePutInternalServerError{}
}

/*TaxTaxRateRepositoryV1SavePutInternalServerError handles this case with default header values.

Internal Server error
*/
type TaxTaxRateRepositoryV1SavePutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1SavePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/taxRates][%d] taxTaxRateRepositoryV1SavePutInternalServerError  %+v", 500, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1SavePutDefault creates a TaxTaxRateRepositoryV1SavePutDefault with default headers values
func NewTaxTaxRateRepositoryV1SavePutDefault(code int) *TaxTaxRateRepositoryV1SavePutDefault {
	return &TaxTaxRateRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*TaxTaxRateRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type TaxTaxRateRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the tax tax rate repository v1 save put default response
func (o *TaxTaxRateRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *TaxTaxRateRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/taxRates][%d] taxTaxRateRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *TaxTaxRateRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TaxTaxRateRepositoryV1SavePutBody tax tax rate repository v1 save put body
swagger:model TaxTaxRateRepositoryV1SavePutBody
*/
type TaxTaxRateRepositoryV1SavePutBody struct {

	// tax rate
	// Required: true
	TaxRate *models.TaxDataTaxRateInterface `json:"taxRate"`
}

// Validate validates this tax tax rate repository v1 save put body
func (o *TaxTaxRateRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTaxRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TaxTaxRateRepositoryV1SavePutBody) validateTaxRate(formats strfmt.Registry) error {

	if err := validate.Required("taxTaxRateRepositoryV1SavePutBody"+"."+"taxRate", "body", o.TaxRate); err != nil {
		return err
	}

	if o.TaxRate != nil {
		if err := o.TaxRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxTaxRateRepositoryV1SavePutBody" + "." + "taxRate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TaxTaxRateRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxTaxRateRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res TaxTaxRateRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}