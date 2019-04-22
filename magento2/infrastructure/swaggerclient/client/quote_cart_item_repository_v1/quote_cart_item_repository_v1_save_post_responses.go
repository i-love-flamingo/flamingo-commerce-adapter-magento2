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

// QuoteCartItemRepositoryV1SavePostReader is a Reader for the QuoteCartItemRepositoryV1SavePost structure.
type QuoteCartItemRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartItemRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartItemRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCartItemRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCartItemRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartItemRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartItemRepositoryV1SavePostOK creates a QuoteCartItemRepositoryV1SavePostOK with default headers values
func NewQuoteCartItemRepositoryV1SavePostOK() *QuoteCartItemRepositoryV1SavePostOK {
	return &QuoteCartItemRepositoryV1SavePostOK{}
}

/*QuoteCartItemRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type QuoteCartItemRepositoryV1SavePostOK struct {
	Payload *models.QuoteDataCartItemInterface
}

func (o *QuoteCartItemRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{quoteId}/items][%d] quoteCartItemRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataCartItemInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePostBadRequest creates a QuoteCartItemRepositoryV1SavePostBadRequest with default headers values
func NewQuoteCartItemRepositoryV1SavePostBadRequest() *QuoteCartItemRepositoryV1SavePostBadRequest {
	return &QuoteCartItemRepositoryV1SavePostBadRequest{}
}

/*QuoteCartItemRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCartItemRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{quoteId}/items][%d] quoteCartItemRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePostUnauthorized creates a QuoteCartItemRepositoryV1SavePostUnauthorized with default headers values
func NewQuoteCartItemRepositoryV1SavePostUnauthorized() *QuoteCartItemRepositoryV1SavePostUnauthorized {
	return &QuoteCartItemRepositoryV1SavePostUnauthorized{}
}

/*QuoteCartItemRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartItemRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{quoteId}/items][%d] quoteCartItemRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1SavePostDefault creates a QuoteCartItemRepositoryV1SavePostDefault with default headers values
func NewQuoteCartItemRepositoryV1SavePostDefault(code int) *QuoteCartItemRepositoryV1SavePostDefault {
	return &QuoteCartItemRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*QuoteCartItemRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartItemRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart item repository v1 save post default response
func (o *QuoteCartItemRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartItemRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/{quoteId}/items][%d] quoteCartItemRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartItemRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteCartItemRepositoryV1SavePostBody quote cart item repository v1 save post body
swagger:model QuoteCartItemRepositoryV1SavePostBody
*/
type QuoteCartItemRepositoryV1SavePostBody struct {

	// cart item
	// Required: true
	CartItem *models.QuoteDataCartItemInterface `json:"cartItem"`
}

// Validate validates this quote cart item repository v1 save post body
func (o *QuoteCartItemRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCartItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteCartItemRepositoryV1SavePostBody) validateCartItem(formats strfmt.Registry) error {

	if err := validate.Required("quoteCartItemRepositoryV1SavePostBody"+"."+"cartItem", "body", o.CartItem); err != nil {
		return err
	}

	if o.CartItem != nil {
		if err := o.CartItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteCartItemRepositoryV1SavePostBody" + "." + "cartItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteCartItemRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteCartItemRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res QuoteCartItemRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
