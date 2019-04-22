// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteCartManagementV1PlaceOrderPutMineReader is a Reader for the QuoteCartManagementV1PlaceOrderPutMine structure.
type QuoteCartManagementV1PlaceOrderPutMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartManagementV1PlaceOrderPutMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartManagementV1PlaceOrderPutMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCartManagementV1PlaceOrderPutMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCartManagementV1PlaceOrderPutMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartManagementV1PlaceOrderPutMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartManagementV1PlaceOrderPutMineOK creates a QuoteCartManagementV1PlaceOrderPutMineOK with default headers values
func NewQuoteCartManagementV1PlaceOrderPutMineOK() *QuoteCartManagementV1PlaceOrderPutMineOK {
	return &QuoteCartManagementV1PlaceOrderPutMineOK{}
}

/*QuoteCartManagementV1PlaceOrderPutMineOK handles this case with default header values.

200 Success.
*/
type QuoteCartManagementV1PlaceOrderPutMineOK struct {
	Payload int64
}

func (o *QuoteCartManagementV1PlaceOrderPutMineOK) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/order][%d] quoteCartManagementV1PlaceOrderPutMineOK  %+v", 200, o.Payload)
}

func (o *QuoteCartManagementV1PlaceOrderPutMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartManagementV1PlaceOrderPutMineBadRequest creates a QuoteCartManagementV1PlaceOrderPutMineBadRequest with default headers values
func NewQuoteCartManagementV1PlaceOrderPutMineBadRequest() *QuoteCartManagementV1PlaceOrderPutMineBadRequest {
	return &QuoteCartManagementV1PlaceOrderPutMineBadRequest{}
}

/*QuoteCartManagementV1PlaceOrderPutMineBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCartManagementV1PlaceOrderPutMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartManagementV1PlaceOrderPutMineBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/order][%d] quoteCartManagementV1PlaceOrderPutMineBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCartManagementV1PlaceOrderPutMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartManagementV1PlaceOrderPutMineUnauthorized creates a QuoteCartManagementV1PlaceOrderPutMineUnauthorized with default headers values
func NewQuoteCartManagementV1PlaceOrderPutMineUnauthorized() *QuoteCartManagementV1PlaceOrderPutMineUnauthorized {
	return &QuoteCartManagementV1PlaceOrderPutMineUnauthorized{}
}

/*QuoteCartManagementV1PlaceOrderPutMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartManagementV1PlaceOrderPutMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartManagementV1PlaceOrderPutMineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/order][%d] quoteCartManagementV1PlaceOrderPutMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartManagementV1PlaceOrderPutMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartManagementV1PlaceOrderPutMineDefault creates a QuoteCartManagementV1PlaceOrderPutMineDefault with default headers values
func NewQuoteCartManagementV1PlaceOrderPutMineDefault(code int) *QuoteCartManagementV1PlaceOrderPutMineDefault {
	return &QuoteCartManagementV1PlaceOrderPutMineDefault{
		_statusCode: code,
	}
}

/*QuoteCartManagementV1PlaceOrderPutMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartManagementV1PlaceOrderPutMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart management v1 place order put mine default response
func (o *QuoteCartManagementV1PlaceOrderPutMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartManagementV1PlaceOrderPutMineDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/order][%d] quoteCartManagementV1PlaceOrderPutMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartManagementV1PlaceOrderPutMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteCartManagementV1PlaceOrderPutMineBody quote cart management v1 place order put mine body
swagger:model QuoteCartManagementV1PlaceOrderPutMineBody
*/
type QuoteCartManagementV1PlaceOrderPutMineBody struct {

	// payment method
	PaymentMethod *models.QuoteDataPaymentInterface `json:"paymentMethod,omitempty"`
}

// Validate validates this quote cart management v1 place order put mine body
func (o *QuoteCartManagementV1PlaceOrderPutMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteCartManagementV1PlaceOrderPutMineBody) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.PaymentMethod) { // not required
		return nil
	}

	if o.PaymentMethod != nil {
		if err := o.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteCartManagementV1PlaceOrderPutBody" + "." + "paymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteCartManagementV1PlaceOrderPutMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteCartManagementV1PlaceOrderPutMineBody) UnmarshalBinary(b []byte) error {
	var res QuoteCartManagementV1PlaceOrderPutMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}