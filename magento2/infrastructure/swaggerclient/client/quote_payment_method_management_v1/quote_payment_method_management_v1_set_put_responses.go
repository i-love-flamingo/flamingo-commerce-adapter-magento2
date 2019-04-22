// Code generated by go-swagger; DO NOT EDIT.

package quote_payment_method_management_v1

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

// QuotePaymentMethodManagementV1SetPutReader is a Reader for the QuotePaymentMethodManagementV1SetPut structure.
type QuotePaymentMethodManagementV1SetPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotePaymentMethodManagementV1SetPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuotePaymentMethodManagementV1SetPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuotePaymentMethodManagementV1SetPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuotePaymentMethodManagementV1SetPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuotePaymentMethodManagementV1SetPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotePaymentMethodManagementV1SetPutOK creates a QuotePaymentMethodManagementV1SetPutOK with default headers values
func NewQuotePaymentMethodManagementV1SetPutOK() *QuotePaymentMethodManagementV1SetPutOK {
	return &QuotePaymentMethodManagementV1SetPutOK{}
}

/*QuotePaymentMethodManagementV1SetPutOK handles this case with default header values.

200 Success.
*/
type QuotePaymentMethodManagementV1SetPutOK struct {
	Payload string
}

func (o *QuotePaymentMethodManagementV1SetPutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/selected-payment-method][%d] quotePaymentMethodManagementV1SetPutOK  %+v", 200, o.Payload)
}

func (o *QuotePaymentMethodManagementV1SetPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1SetPutBadRequest creates a QuotePaymentMethodManagementV1SetPutBadRequest with default headers values
func NewQuotePaymentMethodManagementV1SetPutBadRequest() *QuotePaymentMethodManagementV1SetPutBadRequest {
	return &QuotePaymentMethodManagementV1SetPutBadRequest{}
}

/*QuotePaymentMethodManagementV1SetPutBadRequest handles this case with default header values.

400 Bad Request
*/
type QuotePaymentMethodManagementV1SetPutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1SetPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/selected-payment-method][%d] quotePaymentMethodManagementV1SetPutBadRequest  %+v", 400, o.Payload)
}

func (o *QuotePaymentMethodManagementV1SetPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1SetPutUnauthorized creates a QuotePaymentMethodManagementV1SetPutUnauthorized with default headers values
func NewQuotePaymentMethodManagementV1SetPutUnauthorized() *QuotePaymentMethodManagementV1SetPutUnauthorized {
	return &QuotePaymentMethodManagementV1SetPutUnauthorized{}
}

/*QuotePaymentMethodManagementV1SetPutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuotePaymentMethodManagementV1SetPutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1SetPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/selected-payment-method][%d] quotePaymentMethodManagementV1SetPutUnauthorized  %+v", 401, o.Payload)
}

func (o *QuotePaymentMethodManagementV1SetPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1SetPutDefault creates a QuotePaymentMethodManagementV1SetPutDefault with default headers values
func NewQuotePaymentMethodManagementV1SetPutDefault(code int) *QuotePaymentMethodManagementV1SetPutDefault {
	return &QuotePaymentMethodManagementV1SetPutDefault{
		_statusCode: code,
	}
}

/*QuotePaymentMethodManagementV1SetPutDefault handles this case with default header values.

Unexpected error
*/
type QuotePaymentMethodManagementV1SetPutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote payment method management v1 set put default response
func (o *QuotePaymentMethodManagementV1SetPutDefault) Code() int {
	return o._statusCode
}

func (o *QuotePaymentMethodManagementV1SetPutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/selected-payment-method][%d] quotePaymentMethodManagementV1SetPut default  %+v", o._statusCode, o.Payload)
}

func (o *QuotePaymentMethodManagementV1SetPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuotePaymentMethodManagementV1SetPutBody quote payment method management v1 set put body
swagger:model QuotePaymentMethodManagementV1SetPutBody
*/
type QuotePaymentMethodManagementV1SetPutBody struct {

	// method
	// Required: true
	Method *models.QuoteDataPaymentInterface `json:"method"`
}

// Validate validates this quote payment method management v1 set put body
func (o *QuotePaymentMethodManagementV1SetPutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuotePaymentMethodManagementV1SetPutBody) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("quotePaymentMethodManagementV1SetPutBody"+"."+"method", "body", o.Method); err != nil {
		return err
	}

	if o.Method != nil {
		if err := o.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quotePaymentMethodManagementV1SetPutBody" + "." + "method")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuotePaymentMethodManagementV1SetPutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuotePaymentMethodManagementV1SetPutBody) UnmarshalBinary(b []byte) error {
	var res QuotePaymentMethodManagementV1SetPutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
