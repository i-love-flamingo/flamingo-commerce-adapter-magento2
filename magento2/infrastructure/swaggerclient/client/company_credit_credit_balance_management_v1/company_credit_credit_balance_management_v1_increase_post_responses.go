// Code generated by go-swagger; DO NOT EDIT.

package company_credit_credit_balance_management_v1

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

// CompanyCreditCreditBalanceManagementV1IncreasePostReader is a Reader for the CompanyCreditCreditBalanceManagementV1IncreasePost structure.
type CompanyCreditCreditBalanceManagementV1IncreasePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyCreditCreditBalanceManagementV1IncreasePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyCreditCreditBalanceManagementV1IncreasePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyCreditCreditBalanceManagementV1IncreasePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyCreditCreditBalanceManagementV1IncreasePostOK creates a CompanyCreditCreditBalanceManagementV1IncreasePostOK with default headers values
func NewCompanyCreditCreditBalanceManagementV1IncreasePostOK() *CompanyCreditCreditBalanceManagementV1IncreasePostOK {
	return &CompanyCreditCreditBalanceManagementV1IncreasePostOK{}
}

/*CompanyCreditCreditBalanceManagementV1IncreasePostOK handles this case with default header values.

200 Success.
*/
type CompanyCreditCreditBalanceManagementV1IncreasePostOK struct {
	Payload bool
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/companyCredits/{creditId}/increaseBalance][%d] companyCreditCreditBalanceManagementV1IncreasePostOK  %+v", 200, o.Payload)
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized creates a CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized with default headers values
func NewCompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized() *CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized {
	return &CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized{}
}

/*CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/companyCredits/{creditId}/increaseBalance][%d] companyCreditCreditBalanceManagementV1IncreasePostUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCreditCreditBalanceManagementV1IncreasePostDefault creates a CompanyCreditCreditBalanceManagementV1IncreasePostDefault with default headers values
func NewCompanyCreditCreditBalanceManagementV1IncreasePostDefault(code int) *CompanyCreditCreditBalanceManagementV1IncreasePostDefault {
	return &CompanyCreditCreditBalanceManagementV1IncreasePostDefault{
		_statusCode: code,
	}
}

/*CompanyCreditCreditBalanceManagementV1IncreasePostDefault handles this case with default header values.

Unexpected error
*/
type CompanyCreditCreditBalanceManagementV1IncreasePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company credit credit balance management v1 increase post default response
func (o *CompanyCreditCreditBalanceManagementV1IncreasePostDefault) Code() int {
	return o._statusCode
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/companyCredits/{creditId}/increaseBalance][%d] companyCreditCreditBalanceManagementV1IncreasePost default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompanyCreditCreditBalanceManagementV1IncreasePostBody company credit credit balance management v1 increase post body
swagger:model CompanyCreditCreditBalanceManagementV1IncreasePostBody
*/
type CompanyCreditCreditBalanceManagementV1IncreasePostBody struct {

	// [optional]
	Comment string `json:"comment,omitempty"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// operation type
	// Required: true
	OperationType *int64 `json:"operationType"`

	// options
	Options *models.CompanyCreditDataCreditBalanceOptionsInterface `json:"options,omitempty"`

	// value
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this company credit credit balance management v1 increase post body
func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("companyCreditCreditBalanceManagementV1IncreasePostBody"+"."+"currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) validateOperationType(formats strfmt.Registry) error {

	if err := validate.Required("companyCreditCreditBalanceManagementV1IncreasePostBody"+"."+"operationType", "body", o.OperationType); err != nil {
		return err
	}

	return nil
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(o.Options) { // not required
		return nil
	}

	if o.Options != nil {
		if err := o.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("companyCreditCreditBalanceManagementV1IncreasePostBody" + "." + "options")
			}
			return err
		}
	}

	return nil
}

func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("companyCreditCreditBalanceManagementV1IncreasePostBody"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompanyCreditCreditBalanceManagementV1IncreasePostBody) UnmarshalBinary(b []byte) error {
	var res CompanyCreditCreditBalanceManagementV1IncreasePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
