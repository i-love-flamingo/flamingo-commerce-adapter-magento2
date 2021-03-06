// Code generated by go-swagger; DO NOT EDIT.

package quote_billing_address_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// QuoteBillingAddressManagementV1GetGetReader is a Reader for the QuoteBillingAddressManagementV1GetGet structure.
type QuoteBillingAddressManagementV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteBillingAddressManagementV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteBillingAddressManagementV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteBillingAddressManagementV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteBillingAddressManagementV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteBillingAddressManagementV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteBillingAddressManagementV1GetGetOK creates a QuoteBillingAddressManagementV1GetGetOK with default headers values
func NewQuoteBillingAddressManagementV1GetGetOK() *QuoteBillingAddressManagementV1GetGetOK {
	return &QuoteBillingAddressManagementV1GetGetOK{}
}

/*QuoteBillingAddressManagementV1GetGetOK handles this case with default header values.

200 Success.
*/
type QuoteBillingAddressManagementV1GetGetOK struct {
	Payload *models.QuoteDataAddressInterface
}

func (o *QuoteBillingAddressManagementV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1GetGetOK  %+v", 200, o.Payload)
}

func (o *QuoteBillingAddressManagementV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataAddressInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1GetGetBadRequest creates a QuoteBillingAddressManagementV1GetGetBadRequest with default headers values
func NewQuoteBillingAddressManagementV1GetGetBadRequest() *QuoteBillingAddressManagementV1GetGetBadRequest {
	return &QuoteBillingAddressManagementV1GetGetBadRequest{}
}

/*QuoteBillingAddressManagementV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type QuoteBillingAddressManagementV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteBillingAddressManagementV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteBillingAddressManagementV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1GetGetUnauthorized creates a QuoteBillingAddressManagementV1GetGetUnauthorized with default headers values
func NewQuoteBillingAddressManagementV1GetGetUnauthorized() *QuoteBillingAddressManagementV1GetGetUnauthorized {
	return &QuoteBillingAddressManagementV1GetGetUnauthorized{}
}

/*QuoteBillingAddressManagementV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteBillingAddressManagementV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteBillingAddressManagementV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteBillingAddressManagementV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteBillingAddressManagementV1GetGetDefault creates a QuoteBillingAddressManagementV1GetGetDefault with default headers values
func NewQuoteBillingAddressManagementV1GetGetDefault(code int) *QuoteBillingAddressManagementV1GetGetDefault {
	return &QuoteBillingAddressManagementV1GetGetDefault{
		_statusCode: code,
	}
}

/*QuoteBillingAddressManagementV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type QuoteBillingAddressManagementV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote billing address management v1 get get default response
func (o *QuoteBillingAddressManagementV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteBillingAddressManagementV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/{cartId}/billing-address][%d] quoteBillingAddressManagementV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteBillingAddressManagementV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
