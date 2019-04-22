// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_billing_address_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteGuestBillingAddressManagementV1GetGetReader is a Reader for the QuoteGuestBillingAddressManagementV1GetGet structure.
type QuoteGuestBillingAddressManagementV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestBillingAddressManagementV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteGuestBillingAddressManagementV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteGuestBillingAddressManagementV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteGuestBillingAddressManagementV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestBillingAddressManagementV1GetGetOK creates a QuoteGuestBillingAddressManagementV1GetGetOK with default headers values
func NewQuoteGuestBillingAddressManagementV1GetGetOK() *QuoteGuestBillingAddressManagementV1GetGetOK {
	return &QuoteGuestBillingAddressManagementV1GetGetOK{}
}

/*QuoteGuestBillingAddressManagementV1GetGetOK handles this case with default header values.

200 Success.
*/
type QuoteGuestBillingAddressManagementV1GetGetOK struct {
	Payload *models.QuoteDataAddressInterface
}

func (o *QuoteGuestBillingAddressManagementV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/billing-address][%d] quoteGuestBillingAddressManagementV1GetGetOK  %+v", 200, o.Payload)
}

func (o *QuoteGuestBillingAddressManagementV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataAddressInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestBillingAddressManagementV1GetGetBadRequest creates a QuoteGuestBillingAddressManagementV1GetGetBadRequest with default headers values
func NewQuoteGuestBillingAddressManagementV1GetGetBadRequest() *QuoteGuestBillingAddressManagementV1GetGetBadRequest {
	return &QuoteGuestBillingAddressManagementV1GetGetBadRequest{}
}

/*QuoteGuestBillingAddressManagementV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteGuestBillingAddressManagementV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestBillingAddressManagementV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/billing-address][%d] quoteGuestBillingAddressManagementV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteGuestBillingAddressManagementV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestBillingAddressManagementV1GetGetDefault creates a QuoteGuestBillingAddressManagementV1GetGetDefault with default headers values
func NewQuoteGuestBillingAddressManagementV1GetGetDefault(code int) *QuoteGuestBillingAddressManagementV1GetGetDefault {
	return &QuoteGuestBillingAddressManagementV1GetGetDefault{
		_statusCode: code,
	}
}

/*QuoteGuestBillingAddressManagementV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type QuoteGuestBillingAddressManagementV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest billing address management v1 get get default response
func (o *QuoteGuestBillingAddressManagementV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestBillingAddressManagementV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/billing-address][%d] quoteGuestBillingAddressManagementV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteGuestBillingAddressManagementV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}