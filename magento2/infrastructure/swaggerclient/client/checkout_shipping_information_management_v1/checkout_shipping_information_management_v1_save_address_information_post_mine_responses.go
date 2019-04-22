// Code generated by go-swagger; DO NOT EDIT.

package checkout_shipping_information_management_v1

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

// CheckoutShippingInformationManagementV1SaveAddressInformationPostMineReader is a Reader for the CheckoutShippingInformationManagementV1SaveAddressInformationPostMine structure.
type CheckoutShippingInformationManagementV1SaveAddressInformationPostMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK creates a CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK with default headers values
func NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK() *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK {
	return &CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK{}
}

/*CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK handles this case with default header values.

200 Success.
*/
type CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK struct {
	Payload *models.CheckoutDataPaymentDetailsInterface
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/shipping-information][%d] checkoutShippingInformationManagementV1SaveAddressInformationPostMineOK  %+v", 200, o.Payload)
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CheckoutDataPaymentDetailsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized creates a CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized with default headers values
func NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized() *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized {
	return &CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized{}
}

/*CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/shipping-information][%d] checkoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault creates a CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault with default headers values
func NewCheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault(code int) *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault {
	return &CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault{
		_statusCode: code,
	}
}

/*CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault handles this case with default header values.

Unexpected error
*/
type CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the checkout shipping information management v1 save address information post mine default response
func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/shipping-information][%d] checkoutShippingInformationManagementV1SaveAddressInformationPostMine default  %+v", o._statusCode, o.Payload)
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody checkout shipping information management v1 save address information post mine body
swagger:model CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody
*/
type CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody struct {

	// address information
	// Required: true
	AddressInformation *models.CheckoutDataShippingInformationInterface `json:"addressInformation"`
}

// Validate validates this checkout shipping information management v1 save address information post mine body
func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddressInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody) validateAddressInformation(formats strfmt.Registry) error {

	if err := validate.Required("checkoutShippingInformationManagementV1SaveAddressInformationPostBody"+"."+"addressInformation", "body", o.AddressInformation); err != nil {
		return err
	}

	if o.AddressInformation != nil {
		if err := o.AddressInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutShippingInformationManagementV1SaveAddressInformationPostBody" + "." + "addressInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody) UnmarshalBinary(b []byte) error {
	var res CheckoutShippingInformationManagementV1SaveAddressInformationPostMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
