// Code generated by go-swagger; DO NOT EDIT.

package checkout_payment_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CheckoutPaymentInformationManagementV1GetPaymentInformationGetReader is a Reader for the CheckoutPaymentInformationManagementV1GetPaymentInformationGet structure.
type CheckoutPaymentInformationManagementV1GetPaymentInformationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetOK creates a CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK with default headers values
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetOK() *CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK {
	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK{}
}

/*CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK handles this case with default header values.

200 Success.
*/
type CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK struct {
	Payload *models.CheckoutDataPaymentDetailsInterface
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/payment-information][%d] checkoutPaymentInformationManagementV1GetPaymentInformationGetOK  %+v", 200, o.Payload)
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CheckoutDataPaymentDetailsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized creates a CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized with default headers values
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized() *CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized {
	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized{}
}

/*CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/payment-information][%d] checkoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault creates a CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault with default headers values
func NewCheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault(code int) *CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault {
	return &CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault{
		_statusCode: code,
	}
}

/*CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault handles this case with default header values.

Unexpected error
*/
type CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the checkout payment information management v1 get payment information get default response
func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/carts/mine/payment-information][%d] checkoutPaymentInformationManagementV1GetPaymentInformationGet default  %+v", o._statusCode, o.Payload)
}

func (o *CheckoutPaymentInformationManagementV1GetPaymentInformationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
