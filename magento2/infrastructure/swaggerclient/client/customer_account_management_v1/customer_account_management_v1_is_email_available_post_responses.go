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

// CustomerAccountManagementV1IsEmailAvailablePostReader is a Reader for the CustomerAccountManagementV1IsEmailAvailablePost structure.
type CustomerAccountManagementV1IsEmailAvailablePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountManagementV1IsEmailAvailablePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerAccountManagementV1IsEmailAvailablePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewCustomerAccountManagementV1IsEmailAvailablePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerAccountManagementV1IsEmailAvailablePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAccountManagementV1IsEmailAvailablePostOK creates a CustomerAccountManagementV1IsEmailAvailablePostOK with default headers values
func NewCustomerAccountManagementV1IsEmailAvailablePostOK() *CustomerAccountManagementV1IsEmailAvailablePostOK {
	return &CustomerAccountManagementV1IsEmailAvailablePostOK{}
}

/*CustomerAccountManagementV1IsEmailAvailablePostOK handles this case with default header values.

200 Success.
*/
type CustomerAccountManagementV1IsEmailAvailablePostOK struct {
	Payload bool
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/customers/isEmailAvailable][%d] customerAccountManagementV1IsEmailAvailablePostOK  %+v", 200, o.Payload)
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1IsEmailAvailablePostInternalServerError creates a CustomerAccountManagementV1IsEmailAvailablePostInternalServerError with default headers values
func NewCustomerAccountManagementV1IsEmailAvailablePostInternalServerError() *CustomerAccountManagementV1IsEmailAvailablePostInternalServerError {
	return &CustomerAccountManagementV1IsEmailAvailablePostInternalServerError{}
}

/*CustomerAccountManagementV1IsEmailAvailablePostInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerAccountManagementV1IsEmailAvailablePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/customers/isEmailAvailable][%d] customerAccountManagementV1IsEmailAvailablePostInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1IsEmailAvailablePostDefault creates a CustomerAccountManagementV1IsEmailAvailablePostDefault with default headers values
func NewCustomerAccountManagementV1IsEmailAvailablePostDefault(code int) *CustomerAccountManagementV1IsEmailAvailablePostDefault {
	return &CustomerAccountManagementV1IsEmailAvailablePostDefault{
		_statusCode: code,
	}
}

/*CustomerAccountManagementV1IsEmailAvailablePostDefault handles this case with default header values.

Unexpected error
*/
type CustomerAccountManagementV1IsEmailAvailablePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer account management v1 is email available post default response
func (o *CustomerAccountManagementV1IsEmailAvailablePostDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/customers/isEmailAvailable][%d] customerAccountManagementV1IsEmailAvailablePost default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerAccountManagementV1IsEmailAvailablePostBody customer account management v1 is email available post body
swagger:model CustomerAccountManagementV1IsEmailAvailablePostBody
*/
type CustomerAccountManagementV1IsEmailAvailablePostBody struct {

	// customer email
	// Required: true
	CustomerEmail *string `json:"customerEmail"`

	// If not set, will use the current websiteId
	WebsiteID int64 `json:"websiteId,omitempty"`
}

// Validate validates this customer account management v1 is email available post body
func (o *CustomerAccountManagementV1IsEmailAvailablePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomerEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1IsEmailAvailablePostBody) validateCustomerEmail(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1IsEmailAvailablePostBody"+"."+"customerEmail", "body", o.CustomerEmail); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerAccountManagementV1IsEmailAvailablePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerAccountManagementV1IsEmailAvailablePostBody) UnmarshalBinary(b []byte) error {
	var res CustomerAccountManagementV1IsEmailAvailablePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
