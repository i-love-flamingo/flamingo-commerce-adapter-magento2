// Code generated by go-swagger; DO NOT EDIT.

package quote_billing_address_management_v1

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

// QuoteBillingAddressManagementV1AssignPostReader is a Reader for the QuoteBillingAddressManagementV1AssignPost structure.
type QuoteBillingAddressManagementV1AssignPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteBillingAddressManagementV1AssignPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteBillingAddressManagementV1AssignPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteBillingAddressManagementV1AssignPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteBillingAddressManagementV1AssignPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteBillingAddressManagementV1AssignPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteBillingAddressManagementV1AssignPostOK creates a QuoteBillingAddressManagementV1AssignPostOK with default headers values
func NewQuoteBillingAddressManagementV1AssignPostOK() *QuoteBillingAddressManagementV1AssignPostOK {
	return &QuoteBillingAddressManagementV1AssignPostOK{}
}

/*QuoteBillingAddressManagementV1AssignPostOK handles this case with default header values.

200 Success.
*/
type QuoteBillingAddressManagementV1AssignPostOK struct {
	Payload int64
}

func (o *QuoteBillingAddressManagementV1AssignPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1AssignPostOK  %+v", 200, o.Payload)
}

func (o *QuoteBillingAddressManagementV1AssignPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1AssignPostBadRequest creates a QuoteBillingAddressManagementV1AssignPostBadRequest with default headers values
func NewQuoteBillingAddressManagementV1AssignPostBadRequest() *QuoteBillingAddressManagementV1AssignPostBadRequest {
	return &QuoteBillingAddressManagementV1AssignPostBadRequest{}
}

/*QuoteBillingAddressManagementV1AssignPostBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteBillingAddressManagementV1AssignPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteBillingAddressManagementV1AssignPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1AssignPostBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteBillingAddressManagementV1AssignPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1AssignPostUnauthorized creates a QuoteBillingAddressManagementV1AssignPostUnauthorized with default headers values
func NewQuoteBillingAddressManagementV1AssignPostUnauthorized() *QuoteBillingAddressManagementV1AssignPostUnauthorized {
	return &QuoteBillingAddressManagementV1AssignPostUnauthorized{}
}

/*QuoteBillingAddressManagementV1AssignPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteBillingAddressManagementV1AssignPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteBillingAddressManagementV1AssignPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1AssignPostUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteBillingAddressManagementV1AssignPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1AssignPostDefault creates a QuoteBillingAddressManagementV1AssignPostDefault with default headers values
func NewQuoteBillingAddressManagementV1AssignPostDefault(code int) *QuoteBillingAddressManagementV1AssignPostDefault {
	return &QuoteBillingAddressManagementV1AssignPostDefault{
		_statusCode: code,
	}
}

/*QuoteBillingAddressManagementV1AssignPostDefault handles this case with default header values.

Unexpected error
*/
type QuoteBillingAddressManagementV1AssignPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote billing address management v1 assign post default response
func (o *QuoteBillingAddressManagementV1AssignPostDefault) Code() int {
	return o._statusCode
}

func (o *QuoteBillingAddressManagementV1AssignPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1AssignPost default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteBillingAddressManagementV1AssignPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteBillingAddressManagementV1AssignPostBody quote billing address management v1 assign post body
swagger:model QuoteBillingAddressManagementV1AssignPostBody
*/
type QuoteBillingAddressManagementV1AssignPostBody struct {

	// address
	// Required: true
	Address *models.QuoteDataAddressInterface `json:"address"`

	// use for shipping
	UseForShipping bool `json:"useForShipping,omitempty"`
}

// Validate validates this quote billing address management v1 assign post body
func (o *QuoteBillingAddressManagementV1AssignPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteBillingAddressManagementV1AssignPostBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("quoteBillingAddressManagementV1AssignPostBody"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	if o.Address != nil {
		if err := o.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteBillingAddressManagementV1AssignPostBody" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteBillingAddressManagementV1AssignPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteBillingAddressManagementV1AssignPostBody) UnmarshalBinary(b []byte) error {
	var res QuoteBillingAddressManagementV1AssignPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
