// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteCartRepositoryV1GetListGetReader is a Reader for the QuoteCartRepositoryV1GetListGet structure.
type QuoteCartRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewQuoteCartRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartRepositoryV1GetListGetOK creates a QuoteCartRepositoryV1GetListGetOK with default headers values
func NewQuoteCartRepositoryV1GetListGetOK() *QuoteCartRepositoryV1GetListGetOK {
	return &QuoteCartRepositoryV1GetListGetOK{}
}

/*QuoteCartRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type QuoteCartRepositoryV1GetListGetOK struct {
	Payload *models.QuoteDataCartSearchResultsInterface
}

func (o *QuoteCartRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/search][%d] quoteCartRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *QuoteCartRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataCartSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartRepositoryV1GetListGetUnauthorized creates a QuoteCartRepositoryV1GetListGetUnauthorized with default headers values
func NewQuoteCartRepositoryV1GetListGetUnauthorized() *QuoteCartRepositoryV1GetListGetUnauthorized {
	return &QuoteCartRepositoryV1GetListGetUnauthorized{}
}

/*QuoteCartRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/search][%d] quoteCartRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartRepositoryV1GetListGetDefault creates a QuoteCartRepositoryV1GetListGetDefault with default headers values
func NewQuoteCartRepositoryV1GetListGetDefault(code int) *QuoteCartRepositoryV1GetListGetDefault {
	return &QuoteCartRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*QuoteCartRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart repository v1 get list get default response
func (o *QuoteCartRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/search][%d] quoteCartRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
