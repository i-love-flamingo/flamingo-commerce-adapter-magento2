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

// CustomerAccountManagementV1ResendConfirmationPostReader is a Reader for the CustomerAccountManagementV1ResendConfirmationPost structure.
type CustomerAccountManagementV1ResendConfirmationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountManagementV1ResendConfirmationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAccountManagementV1ResendConfirmationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCustomerAccountManagementV1ResendConfirmationPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerAccountManagementV1ResendConfirmationPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAccountManagementV1ResendConfirmationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAccountManagementV1ResendConfirmationPostOK creates a CustomerAccountManagementV1ResendConfirmationPostOK with default headers values
func NewCustomerAccountManagementV1ResendConfirmationPostOK() *CustomerAccountManagementV1ResendConfirmationPostOK {
	return &CustomerAccountManagementV1ResendConfirmationPostOK{}
}

/*CustomerAccountManagementV1ResendConfirmationPostOK handles this case with default header values.

200 Success.
*/
type CustomerAccountManagementV1ResendConfirmationPostOK struct {
	Payload bool
}

func (o *CustomerAccountManagementV1ResendConfirmationPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/customers/confirm][%d] customerAccountManagementV1ResendConfirmationPostOK  %+v", 200, o.Payload)
}

func (o *CustomerAccountManagementV1ResendConfirmationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1ResendConfirmationPostUnauthorized creates a CustomerAccountManagementV1ResendConfirmationPostUnauthorized with default headers values
func NewCustomerAccountManagementV1ResendConfirmationPostUnauthorized() *CustomerAccountManagementV1ResendConfirmationPostUnauthorized {
	return &CustomerAccountManagementV1ResendConfirmationPostUnauthorized{}
}

/*CustomerAccountManagementV1ResendConfirmationPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerAccountManagementV1ResendConfirmationPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1ResendConfirmationPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/customers/confirm][%d] customerAccountManagementV1ResendConfirmationPostUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerAccountManagementV1ResendConfirmationPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1ResendConfirmationPostInternalServerError creates a CustomerAccountManagementV1ResendConfirmationPostInternalServerError with default headers values
func NewCustomerAccountManagementV1ResendConfirmationPostInternalServerError() *CustomerAccountManagementV1ResendConfirmationPostInternalServerError {
	return &CustomerAccountManagementV1ResendConfirmationPostInternalServerError{}
}

/*CustomerAccountManagementV1ResendConfirmationPostInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAccountManagementV1ResendConfirmationPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1ResendConfirmationPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/customers/confirm][%d] customerAccountManagementV1ResendConfirmationPostInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAccountManagementV1ResendConfirmationPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1ResendConfirmationPostDefault creates a CustomerAccountManagementV1ResendConfirmationPostDefault with default headers values
func NewCustomerAccountManagementV1ResendConfirmationPostDefault(code int) *CustomerAccountManagementV1ResendConfirmationPostDefault {
	return &CustomerAccountManagementV1ResendConfirmationPostDefault{
		_statusCode: code,
	}
}

/*CustomerAccountManagementV1ResendConfirmationPostDefault handles this case with default header values.

Unexpected error
*/
type CustomerAccountManagementV1ResendConfirmationPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer account management v1 resend confirmation post default response
func (o *CustomerAccountManagementV1ResendConfirmationPostDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAccountManagementV1ResendConfirmationPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/customers/confirm][%d] customerAccountManagementV1ResendConfirmationPost default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAccountManagementV1ResendConfirmationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerAccountManagementV1ResendConfirmationPostBody customer account management v1 resend confirmation post body
swagger:model CustomerAccountManagementV1ResendConfirmationPostBody
*/
type CustomerAccountManagementV1ResendConfirmationPostBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// redirect Url
	RedirectURL string `json:"redirectUrl,omitempty"`

	// website Id
	// Required: true
	WebsiteID *int64 `json:"websiteId"`
}

// Validate validates this customer account management v1 resend confirmation post body
func (o *CustomerAccountManagementV1ResendConfirmationPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebsiteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1ResendConfirmationPostBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1ResendConfirmationPostBody"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CustomerAccountManagementV1ResendConfirmationPostBody) validateWebsiteID(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1ResendConfirmationPostBody"+"."+"websiteId", "body", o.WebsiteID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerAccountManagementV1ResendConfirmationPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerAccountManagementV1ResendConfirmationPostBody) UnmarshalBinary(b []byte) error {
	var res CustomerAccountManagementV1ResendConfirmationPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
