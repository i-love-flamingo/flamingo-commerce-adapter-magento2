// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_total_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteCartTotalRepositoryV1GetGetMineReader is a Reader for the QuoteCartTotalRepositoryV1GetGetMine structure.
type QuoteCartTotalRepositoryV1GetGetMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartTotalRepositoryV1GetGetMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartTotalRepositoryV1GetGetMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteCartTotalRepositoryV1GetGetMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteCartTotalRepositoryV1GetGetMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartTotalRepositoryV1GetGetMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartTotalRepositoryV1GetGetMineOK creates a QuoteCartTotalRepositoryV1GetGetMineOK with default headers values
func NewQuoteCartTotalRepositoryV1GetGetMineOK() *QuoteCartTotalRepositoryV1GetGetMineOK {
	return &QuoteCartTotalRepositoryV1GetGetMineOK{}
}

/*QuoteCartTotalRepositoryV1GetGetMineOK handles this case with default header values.

200 Success.
*/
type QuoteCartTotalRepositoryV1GetGetMineOK struct {
	Payload *models.QuoteDataTotalsInterface
}

func (o *QuoteCartTotalRepositoryV1GetGetMineOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/totals][%d] quoteCartTotalRepositoryV1GetGetMineOK  %+v", 200, o.Payload)
}

func (o *QuoteCartTotalRepositoryV1GetGetMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataTotalsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartTotalRepositoryV1GetGetMineBadRequest creates a QuoteCartTotalRepositoryV1GetGetMineBadRequest with default headers values
func NewQuoteCartTotalRepositoryV1GetGetMineBadRequest() *QuoteCartTotalRepositoryV1GetGetMineBadRequest {
	return &QuoteCartTotalRepositoryV1GetGetMineBadRequest{}
}

/*QuoteCartTotalRepositoryV1GetGetMineBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteCartTotalRepositoryV1GetGetMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartTotalRepositoryV1GetGetMineBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/totals][%d] quoteCartTotalRepositoryV1GetGetMineBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteCartTotalRepositoryV1GetGetMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartTotalRepositoryV1GetGetMineUnauthorized creates a QuoteCartTotalRepositoryV1GetGetMineUnauthorized with default headers values
func NewQuoteCartTotalRepositoryV1GetGetMineUnauthorized() *QuoteCartTotalRepositoryV1GetGetMineUnauthorized {
	return &QuoteCartTotalRepositoryV1GetGetMineUnauthorized{}
}

/*QuoteCartTotalRepositoryV1GetGetMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartTotalRepositoryV1GetGetMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartTotalRepositoryV1GetGetMineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/totals][%d] quoteCartTotalRepositoryV1GetGetMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartTotalRepositoryV1GetGetMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartTotalRepositoryV1GetGetMineDefault creates a QuoteCartTotalRepositoryV1GetGetMineDefault with default headers values
func NewQuoteCartTotalRepositoryV1GetGetMineDefault(code int) *QuoteCartTotalRepositoryV1GetGetMineDefault {
	return &QuoteCartTotalRepositoryV1GetGetMineDefault{
		_statusCode: code,
	}
}

/*QuoteCartTotalRepositoryV1GetGetMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartTotalRepositoryV1GetGetMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart total repository v1 get get mine default response
func (o *QuoteCartTotalRepositoryV1GetGetMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartTotalRepositoryV1GetGetMineDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/totals][%d] quoteCartTotalRepositoryV1GetGetMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartTotalRepositoryV1GetGetMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
