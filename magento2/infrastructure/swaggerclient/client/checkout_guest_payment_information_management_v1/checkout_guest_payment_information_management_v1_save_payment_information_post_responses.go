// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_payment_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostReader is a Reader for the CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPost structure.
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK creates a CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK with default headers values
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK() *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK {
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK{}
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK handles this case with default header values.

200 Success.
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK struct {
	Payload int64
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/set-payment-information][%d] checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK  %+v", 200, o.Payload)
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest creates a CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest with default headers values
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest() *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest {
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest{}
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest handles this case with default header values.

400 Bad Request
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/set-payment-information][%d] checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest  %+v", 400, o.Payload)
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault creates a CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault with default headers values
func NewCheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault(code int) *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault {
	return &CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault{
		_statusCode: code,
	}
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault handles this case with default header values.

Unexpected error
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the checkout guest payment information management v1 save payment information post default response
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/set-payment-information][%d] checkoutGuestPaymentInformationManagementV1SavePaymentInformationPost default  %+v", o._statusCode, o.Payload)
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody checkout guest payment information management v1 save payment information post body
swagger:model CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody
*/
type CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody struct {

	// billing address
	BillingAddress *models.QuoteDataAddressInterface `json:"billingAddress,omitempty"`

	// email
	// Required: true
	Email *string `json:"email"`

	// payment method
	// Required: true
	PaymentMethod *models.QuoteDataPaymentInterface `json:"paymentMethod"`
}

// Validate validates this checkout guest payment information management v1 save payment information post body
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) validateBillingAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.BillingAddress) { // not required
		return nil
	}

	if o.BillingAddress != nil {
		if err := o.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody" + "." + "billingAddress")
			}
			return err
		}
	}

	return nil
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody"+"."+"paymentMethod", "body", o.PaymentMethod); err != nil {
		return err
	}

	if o.PaymentMethod != nil {
		if err := o.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody" + "." + "paymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody) UnmarshalBinary(b []byte) error {
	var res CheckoutGuestPaymentInformationManagementV1SavePaymentInformationPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
