// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_item_repository_v1

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

// QuoteCartItemRepositoryV1SavePutReader is a Reader for the QuoteCartItemRepositoryV1SavePut structure.
type QuoteCartItemRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartItemRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartItemRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCartItemRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCartItemRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartItemRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartItemRepositoryV1SavePutOK creates a QuoteCartItemRepositoryV1SavePutOK with default headers values
func NewQuoteCartItemRepositoryV1SavePutOK() *QuoteCartItemRepositoryV1SavePutOK {
	return &QuoteCartItemRepositoryV1SavePutOK{}
}

/*QuoteCartItemRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type QuoteCartItemRepositoryV1SavePutOK struct {
	Payload *models.QuoteDataCartItemInterface
}

func (o *QuoteCartItemRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataCartItemInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePutBadRequest creates a QuoteCartItemRepositoryV1SavePutBadRequest with default headers values
func NewQuoteCartItemRepositoryV1SavePutBadRequest() *QuoteCartItemRepositoryV1SavePutBadRequest {
	return &QuoteCartItemRepositoryV1SavePutBadRequest{}
}

/*QuoteCartItemRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCartItemRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePutUnauthorized creates a QuoteCartItemRepositoryV1SavePutUnauthorized with default headers values
func NewQuoteCartItemRepositoryV1SavePutUnauthorized() *QuoteCartItemRepositoryV1SavePutUnauthorized {
	return &QuoteCartItemRepositoryV1SavePutUnauthorized{}
}

/*QuoteCartItemRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartItemRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePutDefault creates a QuoteCartItemRepositoryV1SavePutDefault with default headers values
func NewQuoteCartItemRepositoryV1SavePutDefault(code int) *QuoteCartItemRepositoryV1SavePutDefault {
	return &QuoteCartItemRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*QuoteCartItemRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartItemRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart item repository v1 save put default response
func (o *QuoteCartItemRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartItemRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteCartItemRepositoryV1SavePutBody quote cart item repository v1 save put body
swagger:model QuoteCartItemRepositoryV1SavePutBody
*/
type QuoteCartItemRepositoryV1SavePutBody struct {

	// cart item
	// Required: true
	CartItem *models.QuoteDataCartItemInterface `json:"cartItem"`
}

// Validate validates this quote cart item repository v1 save put body
func (o *QuoteCartItemRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCartItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteCartItemRepositoryV1SavePutBody) validateCartItem(formats strfmt.Registry) error {

	if err := validate.Required("quoteCartItemRepositoryV1SavePutBody"+"."+"cartItem", "body", o.CartItem); err != nil {
		return err
	}

	if o.CartItem != nil {
		if err := o.CartItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteCartItemRepositoryV1SavePutBody" + "." + "cartItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteCartItemRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteCartItemRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res QuoteCartItemRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
