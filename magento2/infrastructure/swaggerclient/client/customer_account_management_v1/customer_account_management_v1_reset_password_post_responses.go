// Code generated by go-swagger; DO NOT EDIT.

package customer_account_management_v1

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

// CustomerAccountManagementV1ResetPasswordPostReader is a Reader for the CustomerAccountManagementV1ResetPasswordPost structure.
type CustomerAccountManagementV1ResetPasswordPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountManagementV1ResetPasswordPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAccountManagementV1ResetPasswordPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewCustomerAccountManagementV1ResetPasswordPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAccountManagementV1ResetPasswordPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAccountManagementV1ResetPasswordPostOK creates a CustomerAccountManagementV1ResetPasswordPostOK with default headers values
func NewCustomerAccountManagementV1ResetPasswordPostOK() *CustomerAccountManagementV1ResetPasswordPostOK {
	return &CustomerAccountManagementV1ResetPasswordPostOK{}
}

/*CustomerAccountManagementV1ResetPasswordPostOK handles this case with default header values.

200 Success.
*/
type CustomerAccountManagementV1ResetPasswordPostOK struct {
	Payload bool
}

func (o *CustomerAccountManagementV1ResetPasswordPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/customers/resetPassword][%d] customerAccountManagementV1ResetPasswordPostOK  %+v", 200, o.Payload)
}

func (o *CustomerAccountManagementV1ResetPasswordPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1ResetPasswordPostInternalServerError creates a CustomerAccountManagementV1ResetPasswordPostInternalServerError with default headers values
func NewCustomerAccountManagementV1ResetPasswordPostInternalServerError() *CustomerAccountManagementV1ResetPasswordPostInternalServerError {
	return &CustomerAccountManagementV1ResetPasswordPostInternalServerError{}
}

/*CustomerAccountManagementV1ResetPasswordPostInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAccountManagementV1ResetPasswordPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1ResetPasswordPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/customers/resetPassword][%d] customerAccountManagementV1ResetPasswordPostInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAccountManagementV1ResetPasswordPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1ResetPasswordPostDefault creates a CustomerAccountManagementV1ResetPasswordPostDefault with default headers values
func NewCustomerAccountManagementV1ResetPasswordPostDefault(code int) *CustomerAccountManagementV1ResetPasswordPostDefault {
	return &CustomerAccountManagementV1ResetPasswordPostDefault{
		_statusCode: code,
	}
}

/*CustomerAccountManagementV1ResetPasswordPostDefault handles this case with default header values.

Unexpected error
*/
type CustomerAccountManagementV1ResetPasswordPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer account management v1 reset password post default response
func (o *CustomerAccountManagementV1ResetPasswordPostDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAccountManagementV1ResetPasswordPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/customers/resetPassword][%d] customerAccountManagementV1ResetPasswordPost default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAccountManagementV1ResetPasswordPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerAccountManagementV1ResetPasswordPostBody customer account management v1 reset password post body
swagger:model CustomerAccountManagementV1ResetPasswordPostBody
*/
type CustomerAccountManagementV1ResetPasswordPostBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// new password
	// Required: true
	NewPassword *string `json:"newPassword"`

	// reset token
	// Required: true
	ResetToken *string `json:"resetToken"`
}

// Validate validates this customer account management v1 reset password post body
func (o *CustomerAccountManagementV1ResetPasswordPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResetToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1ResetPasswordPostBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1ResetPasswordPostBody"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CustomerAccountManagementV1ResetPasswordPostBody) validateNewPassword(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1ResetPasswordPostBody"+"."+"newPassword", "body", o.NewPassword); err != nil {
		return err
	}

	return nil
}

func (o *CustomerAccountManagementV1ResetPasswordPostBody) validateResetToken(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1ResetPasswordPostBody"+"."+"resetToken", "body", o.ResetToken); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerAccountManagementV1ResetPasswordPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerAccountManagementV1ResetPasswordPostBody) UnmarshalBinary(b []byte) error {
	var res CustomerAccountManagementV1ResetPasswordPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}