// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_item_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteCartItemRepositoryV1DeleteByIDDeleteReader is a Reader for the QuoteCartItemRepositoryV1DeleteByIDDelete structure.
type QuoteCartItemRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartItemRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartItemRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteOK creates a QuoteCartItemRepositoryV1DeleteByIDDeleteOK with default headers values
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteOK() *QuoteCartItemRepositoryV1DeleteByIDDeleteOK {
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteOK{}
}

/*QuoteCartItemRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type QuoteCartItemRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest creates a QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest() *QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest {
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized creates a QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized() *QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized {
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartItemRepositoryV1DeleteByIDDeleteDefault creates a QuoteCartItemRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewQuoteCartItemRepositoryV1DeleteByIDDeleteDefault(code int) *QuoteCartItemRepositoryV1DeleteByIDDeleteDefault {
	return &QuoteCartItemRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*QuoteCartItemRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartItemRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart item repository v1 delete by Id delete default response
func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/carts/{cartId}/items/{itemId}][%d] quoteCartItemRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartItemRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}