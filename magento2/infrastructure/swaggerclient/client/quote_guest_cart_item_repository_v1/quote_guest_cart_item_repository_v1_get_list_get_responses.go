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

// QuoteGuestCartItemRepositoryV1GetListGetReader is a Reader for the QuoteGuestCartItemRepositoryV1GetListGet structure.
type QuoteGuestCartItemRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCartItemRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteGuestCartItemRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteGuestCartItemRepositoryV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteGuestCartItemRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCartItemRepositoryV1GetListGetOK creates a QuoteGuestCartItemRepositoryV1GetListGetOK with default headers values
func NewQuoteGuestCartItemRepositoryV1GetListGetOK() *QuoteGuestCartItemRepositoryV1GetListGetOK {
	return &QuoteGuestCartItemRepositoryV1GetListGetOK{}
}

/*QuoteGuestCartItemRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type QuoteGuestCartItemRepositoryV1GetListGetOK struct {
	Payload []*models.QuoteDataCartItemInterface
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1GetListGetBadRequest creates a QuoteGuestCartItemRepositoryV1GetListGetBadRequest with default headers values
func NewQuoteGuestCartItemRepositoryV1GetListGetBadRequest() *QuoteGuestCartItemRepositoryV1GetListGetBadRequest {
	return &QuoteGuestCartItemRepositoryV1GetListGetBadRequest{}
}

/*QuoteGuestCartItemRepositoryV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteGuestCartItemRepositoryV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1GetListGetDefault creates a QuoteGuestCartItemRepositoryV1GetListGetDefault with default headers values
func NewQuoteGuestCartItemRepositoryV1GetListGetDefault(code int) *QuoteGuestCartItemRepositoryV1GetListGetDefault {
	return &QuoteGuestCartItemRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*QuoteGuestCartItemRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type QuoteGuestCartItemRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest cart item repository v1 get list get default response
func (o *QuoteGuestCartItemRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteGuestCartItemRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
