// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_shipping_information_management_v1

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

// CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostReader is a Reader for the CheckoutGuestShippingInformationManagementV1SaveAddressInformationPost structure.
type CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK creates a CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK with default headers values
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK() *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK{}
}

/*CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK handles this case with default header values.

200 Success.
*/
type CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK struct {
	Payload *models.CheckoutDataPaymentDetailsInterface
}

func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/shipping-information][%d] checkoutGuestShippingInformationManagementV1SaveAddressInformationPostOK  %+v", 200, o.Payload)
}

func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CheckoutDataPaymentDetailsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault creates a CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault with default headers values
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault(code int) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault{
		_statusCode: code,
	}
}

/*CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault handles this case with default header values.

Unexpected error
*/
type CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the checkout guest shipping information management v1 save address information post default response
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/shipping-information][%d] checkoutGuestShippingInformationManagementV1SaveAddressInformationPost default  %+v", o._statusCode, o.Payload)
}

func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody checkout guest shipping information management v1 save address information post body
swagger:model CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody
*/
type CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody struct {

	// address information
	// Required: true
	AddressInformation *models.CheckoutDataShippingInformationInterface `json:"addressInformation"`
}

// Validate validates this checkout guest shipping information management v1 save address information post body
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddressInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) validateAddressInformation(formats strfmt.Registry) error {

	if err := validate.Required("checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody"+"."+"addressInformation", "body", o.AddressInformation); err != nil {
		return err
	}

	if o.AddressInformation != nil {
		if err := o.AddressInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody" + "." + "addressInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) UnmarshalBinary(b []byte) error {
	var res CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
