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

// QuotePaymentMethodManagementV1GetGetMineReader is a Reader for the QuotePaymentMethodManagementV1GetGetMine structure.
type QuotePaymentMethodManagementV1GetGetMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotePaymentMethodManagementV1GetGetMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuotePaymentMethodManagementV1GetGetMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuotePaymentMethodManagementV1GetGetMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuotePaymentMethodManagementV1GetGetMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuotePaymentMethodManagementV1GetGetMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotePaymentMethodManagementV1GetGetMineOK creates a QuotePaymentMethodManagementV1GetGetMineOK with default headers values
func NewQuotePaymentMethodManagementV1GetGetMineOK() *QuotePaymentMethodManagementV1GetGetMineOK {
	return &QuotePaymentMethodManagementV1GetGetMineOK{}
}

/*QuotePaymentMethodManagementV1GetGetMineOK handles this case with default header values.

200 Success.
*/
type QuotePaymentMethodManagementV1GetGetMineOK struct {
	Payload *models.QuoteDataPaymentInterface
}

func (o *QuotePaymentMethodManagementV1GetGetMineOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/selected-payment-method][%d] quotePaymentMethodManagementV1GetGetMineOK  %+v", 200, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetGetMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataPaymentInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetGetMineBadRequest creates a QuotePaymentMethodManagementV1GetGetMineBadRequest with default headers values
func NewQuotePaymentMethodManagementV1GetGetMineBadRequest() *QuotePaymentMethodManagementV1GetGetMineBadRequest {
	return &QuotePaymentMethodManagementV1GetGetMineBadRequest{}
}

/*QuotePaymentMethodManagementV1GetGetMineBadRequest handles this case with default header values.

400 Bad Request
*/
type QuotePaymentMethodManagementV1GetGetMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1GetGetMineBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/selected-payment-method][%d] quotePaymentMethodManagementV1GetGetMineBadRequest  %+v", 400, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetGetMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetGetMineUnauthorized creates a QuotePaymentMethodManagementV1GetGetMineUnauthorized with default headers values
func NewQuotePaymentMethodManagementV1GetGetMineUnauthorized() *QuotePaymentMethodManagementV1GetGetMineUnauthorized {
	return &QuotePaymentMethodManagementV1GetGetMineUnauthorized{}
}

/*QuotePaymentMethodManagementV1GetGetMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuotePaymentMethodManagementV1GetGetMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuotePaymentMethodManagementV1GetGetMineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/selected-payment-method][%d] quotePaymentMethodManagementV1GetGetMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetGetMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotePaymentMethodManagementV1GetGetMineDefault creates a QuotePaymentMethodManagementV1GetGetMineDefault with default headers values
func NewQuotePaymentMethodManagementV1GetGetMineDefault(code int) *QuotePaymentMethodManagementV1GetGetMineDefault {
	return &QuotePaymentMethodManagementV1GetGetMineDefault{
		_statusCode: code,
	}
}

/*QuotePaymentMethodManagementV1GetGetMineDefault handles this case with default header values.

Unexpected error
*/
type QuotePaymentMethodManagementV1GetGetMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote payment method management v1 get get mine default response
func (o *QuotePaymentMethodManagementV1GetGetMineDefault) Code() int {
	return o._statusCode
}

func (o *QuotePaymentMethodManagementV1GetGetMineDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/selected-payment-method][%d] quotePaymentMethodManagementV1GetGetMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuotePaymentMethodManagementV1GetGetMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}