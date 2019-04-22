// Code generated by go-swagger; DO NOT EDIT.

package company_credit_credit_limit_repository_v1

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

// CompanyCreditCreditLimitRepositoryV1SavePutReader is a Reader for the CompanyCreditCreditLimitRepositoryV1SavePut structure.
type CompanyCreditCreditLimitRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyCreditCreditLimitRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyCreditCreditLimitRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompanyCreditCreditLimitRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCompanyCreditCreditLimitRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCompanyCreditCreditLimitRepositoryV1SavePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyCreditCreditLimitRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyCreditCreditLimitRepositoryV1SavePutOK creates a CompanyCreditCreditLimitRepositoryV1SavePutOK with default headers values
func NewCompanyCreditCreditLimitRepositoryV1SavePutOK() *CompanyCreditCreditLimitRepositoryV1SavePutOK {
	return &CompanyCreditCreditLimitRepositoryV1SavePutOK{}
}

/*CompanyCreditCreditLimitRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CompanyCreditCreditLimitRepositoryV1SavePutOK struct {
	Payload *models.CompanyCreditDataCreditLimitInterface
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/companyCredits/{id}][%d] companyCreditCreditLimitRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyCreditDataCreditLimitInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditLimitRepositoryV1SavePutBadRequest creates a CompanyCreditCreditLimitRepositoryV1SavePutBadRequest with default headers values
func NewCompanyCreditCreditLimitRepositoryV1SavePutBadRequest() *CompanyCreditCreditLimitRepositoryV1SavePutBadRequest {
	return &CompanyCreditCreditLimitRepositoryV1SavePutBadRequest{}
}

/*CompanyCreditCreditLimitRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CompanyCreditCreditLimitRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/companyCredits/{id}][%d] companyCreditCreditLimitRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditLimitRepositoryV1SavePutUnauthorized creates a CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized with default headers values
func NewCompanyCreditCreditLimitRepositoryV1SavePutUnauthorized() *CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized {
	return &CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized{}
}

/*CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/companyCredits/{id}][%d] companyCreditCreditLimitRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditLimitRepositoryV1SavePutInternalServerError creates a CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError with default headers values
func NewCompanyCreditCreditLimitRepositoryV1SavePutInternalServerError() *CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError {
	return &CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError{}
}

/*CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError handles this case with default header values.

Internal Server error
*/
type CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/companyCredits/{id}][%d] companyCreditCreditLimitRepositoryV1SavePutInternalServerError  %+v", 500, o.Payload)
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditLimitRepositoryV1SavePutDefault creates a CompanyCreditCreditLimitRepositoryV1SavePutDefault with default headers values
func NewCompanyCreditCreditLimitRepositoryV1SavePutDefault(code int) *CompanyCreditCreditLimitRepositoryV1SavePutDefault {
	return &CompanyCreditCreditLimitRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CompanyCreditCreditLimitRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CompanyCreditCreditLimitRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company credit credit limit repository v1 save put default response
func (o *CompanyCreditCreditLimitRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/companyCredits/{id}][%d] companyCreditCreditLimitRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompanyCreditCreditLimitRepositoryV1SavePutBody company credit credit limit repository v1 save put body
swagger:model CompanyCreditCreditLimitRepositoryV1SavePutBody
*/
type CompanyCreditCreditLimitRepositoryV1SavePutBody struct {

	// credit limit
	// Required: true
	CreditLimit *models.CompanyCreditDataCreditLimitInterface `json:"creditLimit"`
}

// Validate validates this company credit credit limit repository v1 save put body
func (o *CompanyCreditCreditLimitRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreditLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CompanyCreditCreditLimitRepositoryV1SavePutBody) validateCreditLimit(formats strfmt.Registry) error {

	if err := validate.Required("companyCreditCreditLimitRepositoryV1SavePutBody"+"."+"creditLimit", "body", o.CreditLimit); err != nil {
		return err
	}

	if o.CreditLimit != nil {
		if err := o.CreditLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("companyCreditCreditLimitRepositoryV1SavePutBody" + "." + "creditLimit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CompanyCreditCreditLimitRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompanyCreditCreditLimitRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CompanyCreditCreditLimitRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
