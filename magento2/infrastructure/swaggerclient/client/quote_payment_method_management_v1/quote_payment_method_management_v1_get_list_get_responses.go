// Code generated by go-swagger; DO NOT EDIT.

package quote_payment_method_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuotePaymentMethodManagementV1GetListGetReader is a Reader for the QuotePaymentMethodManagementV1GetListGet structure.
type QuotePaymentMethodManagementV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotePaymentMethodManagementV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuotePaymentMethodManagementV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuotePaymentMethodManagementV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuotePaymentMethodManagementV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuotePaymentMethodManagementV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotePaymentMethodManagementV1GetListGetOK creates a QuotePaymentMethodManagementV1GetListGetOK with default headers values
func NewQuotePaymentMethodManagementV1GetListGetOK() *QuotePaymentMethodManagementV1GetListGetOK {
	return &QuotePaymentMethodManagementV1GetListGetOK{}
}

/*QuotePaymentMethodManagementV1GetListGetOK handles this case with default header values.

200 Success.
*/
type QuotePaymentMethodManagementV1GetListGetOK struct {
	Payload []*models.QuoteDataPaymentMethodInterface
}

func (o *QuotePaymentMethodManagementV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/payment-methods][%d] quotePaymentMethodManagementV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetListGetBadRequest creates a QuotePaymentMethodManagementV1GetListGetBadRequest with default headers values
func NewQuotePaymentMethodManagementV1GetListGetBadRequest() *QuotePaymentMethodManagementV1GetListGetBadRequest {
	return &QuotePaymentMethodManagementV1GetListGetBadRequest{}
}

/*QuotePaymentMethodManagementV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type QuotePaymentMethodManagementV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/payment-methods][%d] quotePaymentMethodManagementV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetListGetUnauthorized creates a QuotePaymentMethodManagementV1GetListGetUnauthorized with default headers values
func NewQuotePaymentMethodManagementV1GetListGetUnauthorized() *QuotePaymentMethodManagementV1GetListGetUnauthorized {
	return &QuotePaymentMethodManagementV1GetListGetUnauthorized{}
}

/*QuotePaymentMethodManagementV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuotePaymentMethodManagementV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/payment-methods][%d] quotePaymentMethodManagementV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetListGetDefault creates a QuotePaymentMethodManagementV1GetListGetDefault with default headers values
func NewQuotePaymentMethodManagementV1GetListGetDefault(code int) *QuotePaymentMethodManagementV1GetListGetDefault {
	return &QuotePaymentMethodManagementV1GetListGetDefault{
		_statusCode: code,
	}
}

/*QuotePaymentMethodManagementV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type QuotePaymentMethodManagementV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote payment method management v1 get list get default response
func (o *QuotePaymentMethodManagementV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *QuotePaymentMethodManagementV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/payment-methods][%d] quotePaymentMethodManagementV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
