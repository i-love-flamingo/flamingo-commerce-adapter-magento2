// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_item_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteGuestCartItemRepositoryV1DeleteByIDDeleteReader is a Reader for the QuoteGuestCartItemRepositoryV1DeleteByIDDelete structure.
type QuoteGuestCartItemRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK creates a QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK with default headers values
func NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK() *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK {
	return &QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK{}
}

/*QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/items/{itemId}][%d] quoteGuestCartItemRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest creates a QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest() *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest {
	return &QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/items/{itemId}][%d] quoteGuestCartItemRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault creates a QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewQuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault(code int) *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault {
	return &QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest cart item repository v1 delete by Id delete default response
func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/guest-carts/{cartId}/items/{itemId}][%d] quoteGuestCartItemRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
