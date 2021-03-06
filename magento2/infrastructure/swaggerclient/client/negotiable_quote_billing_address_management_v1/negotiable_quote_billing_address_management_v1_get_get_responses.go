// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_billing_address_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// NegotiableQuoteBillingAddressManagementV1GetGetReader is a Reader for the NegotiableQuoteBillingAddressManagementV1GetGet structure.
type NegotiableQuoteBillingAddressManagementV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NegotiableQuoteBillingAddressManagementV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNegotiableQuoteBillingAddressManagementV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNegotiableQuoteBillingAddressManagementV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewNegotiableQuoteBillingAddressManagementV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNegotiableQuoteBillingAddressManagementV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNegotiableQuoteBillingAddressManagementV1GetGetOK creates a NegotiableQuoteBillingAddressManagementV1GetGetOK with default headers values
func NewNegotiableQuoteBillingAddressManagementV1GetGetOK() *NegotiableQuoteBillingAddressManagementV1GetGetOK {
	return &NegotiableQuoteBillingAddressManagementV1GetGetOK{}
}

/*NegotiableQuoteBillingAddressManagementV1GetGetOK handles this case with default header values.

200 Success.
*/
type NegotiableQuoteBillingAddressManagementV1GetGetOK struct {
	Payload *models.QuoteDataAddressInterface
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/negotiable-carts/{cartId}/billing-address][%d] negotiableQuoteBillingAddressManagementV1GetGetOK  %+v", 200, o.Payload)
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataAddressInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteBillingAddressManagementV1GetGetBadRequest creates a NegotiableQuoteBillingAddressManagementV1GetGetBadRequest with default headers values
func NewNegotiableQuoteBillingAddressManagementV1GetGetBadRequest() *NegotiableQuoteBillingAddressManagementV1GetGetBadRequest {
	return &NegotiableQuoteBillingAddressManagementV1GetGetBadRequest{}
}

/*NegotiableQuoteBillingAddressManagementV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type NegotiableQuoteBillingAddressManagementV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/negotiable-carts/{cartId}/billing-address][%d] negotiableQuoteBillingAddressManagementV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteBillingAddressManagementV1GetGetUnauthorized creates a NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized with default headers values
func NewNegotiableQuoteBillingAddressManagementV1GetGetUnauthorized() *NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized {
	return &NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized{}
}

/*NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/negotiable-carts/{cartId}/billing-address][%d] negotiableQuoteBillingAddressManagementV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteBillingAddressManagementV1GetGetDefault creates a NegotiableQuoteBillingAddressManagementV1GetGetDefault with default headers values
func NewNegotiableQuoteBillingAddressManagementV1GetGetDefault(code int) *NegotiableQuoteBillingAddressManagementV1GetGetDefault {
	return &NegotiableQuoteBillingAddressManagementV1GetGetDefault{
		_statusCode: code,
	}
}

/*NegotiableQuoteBillingAddressManagementV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type NegotiableQuoteBillingAddressManagementV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the negotiable quote billing address management v1 get get default response
func (o *NegotiableQuoteBillingAddressManagementV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/negotiable-carts/{cartId}/billing-address][%d] negotiableQuoteBillingAddressManagementV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *NegotiableQuoteBillingAddressManagementV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
