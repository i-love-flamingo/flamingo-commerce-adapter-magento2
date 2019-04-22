// Code generated by go-swagger; DO NOT EDIT.

package customer_customer_repository_v1

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

// CustomerCustomerRepositoryV1SavePutReader is a Reader for the CustomerCustomerRepositoryV1SavePut structure.
type CustomerCustomerRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCustomerRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomerCustomerRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomerCustomerRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCustomerCustomerRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomerCustomerRepositoryV1SavePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCustomerCustomerRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerCustomerRepositoryV1SavePutOK creates a CustomerCustomerRepositoryV1SavePutOK with default headers values
func NewCustomerCustomerRepositoryV1SavePutOK() *CustomerCustomerRepositoryV1SavePutOK {
	return &CustomerCustomerRepositoryV1SavePutOK{}
}

/*CustomerCustomerRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CustomerCustomerRepositoryV1SavePutOK struct {
	Payload *models.CustomerDataCustomerInterface
}

func (o *CustomerCustomerRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/{customerId}][%d] customerCustomerRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CustomerCustomerRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerDataCustomerInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1SavePutBadRequest creates a CustomerCustomerRepositoryV1SavePutBadRequest with default headers values
func NewCustomerCustomerRepositoryV1SavePutBadRequest() *CustomerCustomerRepositoryV1SavePutBadRequest {
	return &CustomerCustomerRepositoryV1SavePutBadRequest{}
}

/*CustomerCustomerRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CustomerCustomerRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/{customerId}][%d] customerCustomerRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerCustomerRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1SavePutUnauthorized creates a CustomerCustomerRepositoryV1SavePutUnauthorized with default headers values
func NewCustomerCustomerRepositoryV1SavePutUnauthorized() *CustomerCustomerRepositoryV1SavePutUnauthorized {
	return &CustomerCustomerRepositoryV1SavePutUnauthorized{}
}

/*CustomerCustomerRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CustomerCustomerRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/{customerId}][%d] customerCustomerRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerCustomerRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1SavePutInternalServerError creates a CustomerCustomerRepositoryV1SavePutInternalServerError with default headers values
func NewCustomerCustomerRepositoryV1SavePutInternalServerError() *CustomerCustomerRepositoryV1SavePutInternalServerError {
	return &CustomerCustomerRepositoryV1SavePutInternalServerError{}
}

/*CustomerCustomerRepositoryV1SavePutInternalServerError handles this case with default header values.

Internal Server error
*/
type CustomerCustomerRepositoryV1SavePutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerCustomerRepositoryV1SavePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/{customerId}][%d] customerCustomerRepositoryV1SavePutInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerCustomerRepositoryV1SavePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCustomerRepositoryV1SavePutDefault creates a CustomerCustomerRepositoryV1SavePutDefault with default headers values
func NewCustomerCustomerRepositoryV1SavePutDefault(code int) *CustomerCustomerRepositoryV1SavePutDefault {
	return &CustomerCustomerRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CustomerCustomerRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CustomerCustomerRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer customer repository v1 save put default response
func (o *CustomerCustomerRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CustomerCustomerRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/{customerId}][%d] customerCustomerRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCustomerRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerCustomerRepositoryV1SavePutBody customer customer repository v1 save put body
swagger:model CustomerCustomerRepositoryV1SavePutBody
*/
type CustomerCustomerRepositoryV1SavePutBody struct {

	// customer
	// Required: true
	Customer *models.CustomerDataCustomerInterface `json:"customer"`

	// password hash
	PasswordHash string `json:"passwordHash,omitempty"`
}

// Validate validates this customer customer repository v1 save put body
func (o *CustomerCustomerRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerCustomerRepositoryV1SavePutBody) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customerCustomerRepositoryV1SavePutBody"+"."+"customer", "body", o.Customer); err != nil {
		return err
	}

	if o.Customer != nil {
		if err := o.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerCustomerRepositoryV1SavePutBody" + "." + "customer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerCustomerRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerCustomerRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CustomerCustomerRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
