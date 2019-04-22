// Code generated by go-swagger; DO NOT EDIT.

package quote_shipping_method_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteShippingMethodManagementV1GetListGetReader is a Reader for the QuoteShippingMethodManagementV1GetListGet structure.
type QuoteShippingMethodManagementV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteShippingMethodManagementV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteShippingMethodManagementV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteShippingMethodManagementV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteShippingMethodManagementV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteShippingMethodManagementV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteShippingMethodManagementV1GetListGetOK creates a QuoteShippingMethodManagementV1GetListGetOK with default headers values
func NewQuoteShippingMethodManagementV1GetListGetOK() *QuoteShippingMethodManagementV1GetListGetOK {
	return &QuoteShippingMethodManagementV1GetListGetOK{}
}

/*QuoteShippingMethodManagementV1GetListGetOK handles this case with default header values.

200 Success.
*/
type QuoteShippingMethodManagementV1GetListGetOK struct {
	Payload []*models.QuoteDataShippingMethodInterface
}

func (o *QuoteShippingMethodManagementV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/shipping-methods][%d] quoteShippingMethodManagementV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *QuoteShippingMethodManagementV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteShippingMethodManagementV1GetListGetBadRequest creates a QuoteShippingMethodManagementV1GetListGetBadRequest with default headers values
func NewQuoteShippingMethodManagementV1GetListGetBadRequest() *QuoteShippingMethodManagementV1GetListGetBadRequest {
	return &QuoteShippingMethodManagementV1GetListGetBadRequest{}
}

/*QuoteShippingMethodManagementV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteShippingMethodManagementV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteShippingMethodManagementV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/shipping-methods][%d] quoteShippingMethodManagementV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteShippingMethodManagementV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteShippingMethodManagementV1GetListGetUnauthorized creates a QuoteShippingMethodManagementV1GetListGetUnauthorized with default headers values
func NewQuoteShippingMethodManagementV1GetListGetUnauthorized() *QuoteShippingMethodManagementV1GetListGetUnauthorized {
	return &QuoteShippingMethodManagementV1GetListGetUnauthorized{}
}

/*QuoteShippingMethodManagementV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteShippingMethodManagementV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteShippingMethodManagementV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/shipping-methods][%d] quoteShippingMethodManagementV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteShippingMethodManagementV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteShippingMethodManagementV1GetListGetDefault creates a QuoteShippingMethodManagementV1GetListGetDefault with default headers values
func NewQuoteShippingMethodManagementV1GetListGetDefault(code int) *QuoteShippingMethodManagementV1GetListGetDefault {
	return &QuoteShippingMethodManagementV1GetListGetDefault{
		_statusCode: code,
	}
}

/*QuoteShippingMethodManagementV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type QuoteShippingMethodManagementV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote shipping method management v1 get list get default response
func (o *QuoteShippingMethodManagementV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteShippingMethodManagementV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/shipping-methods][%d] quoteShippingMethodManagementV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteShippingMethodManagementV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
