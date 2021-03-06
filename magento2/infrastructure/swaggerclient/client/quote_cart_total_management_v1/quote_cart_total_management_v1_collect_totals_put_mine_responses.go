// Code generated by go-swagger; DO NOT EDIT.

package quote_cart_total_management_v1

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

// QuoteCartTotalManagementV1CollectTotalsPutMineReader is a Reader for the QuoteCartTotalManagementV1CollectTotalsPutMine structure.
type QuoteCartTotalManagementV1CollectTotalsPutMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCartTotalManagementV1CollectTotalsPutMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteCartTotalManagementV1CollectTotalsPutMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewQuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteCartTotalManagementV1CollectTotalsPutMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteCartTotalManagementV1CollectTotalsPutMineOK creates a QuoteCartTotalManagementV1CollectTotalsPutMineOK with default headers values
func NewQuoteCartTotalManagementV1CollectTotalsPutMineOK() *QuoteCartTotalManagementV1CollectTotalsPutMineOK {
	return &QuoteCartTotalManagementV1CollectTotalsPutMineOK{}
}

/*QuoteCartTotalManagementV1CollectTotalsPutMineOK handles this case with default header values.

200 Success.
*/
type QuoteCartTotalManagementV1CollectTotalsPutMineOK struct {
	Payload *models.QuoteDataTotalsInterface
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineOK) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/collect-totals][%d] quoteCartTotalManagementV1CollectTotalsPutMineOK  %+v", 200, o.Payload)
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataTotalsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized creates a QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized with default headers values
func NewQuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized() *QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized {
	return &QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized{}
}

/*QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/collect-totals][%d] quoteCartTotalManagementV1CollectTotalsPutMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCartTotalManagementV1CollectTotalsPutMineDefault creates a QuoteCartTotalManagementV1CollectTotalsPutMineDefault with default headers values
func NewQuoteCartTotalManagementV1CollectTotalsPutMineDefault(code int) *QuoteCartTotalManagementV1CollectTotalsPutMineDefault {
	return &QuoteCartTotalManagementV1CollectTotalsPutMineDefault{
		_statusCode: code,
	}
}

/*QuoteCartTotalManagementV1CollectTotalsPutMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteCartTotalManagementV1CollectTotalsPutMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote cart total management v1 collect totals put mine default response
func (o *QuoteCartTotalManagementV1CollectTotalsPutMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/carts/mine/collect-totals][%d] quoteCartTotalManagementV1CollectTotalsPutMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteCartTotalManagementV1CollectTotalsPutMineBody quote cart total management v1 collect totals put mine body
swagger:model QuoteCartTotalManagementV1CollectTotalsPutMineBody
*/
type QuoteCartTotalManagementV1CollectTotalsPutMineBody struct {

	// additional data
	AdditionalData *models.QuoteDataTotalsAdditionalDataInterface `json:"additionalData,omitempty"`

	// payment method
	// Required: true
	PaymentMethod *models.QuoteDataPaymentInterface `json:"paymentMethod"`

	// The carrier code.
	ShippingCarrierCode string `json:"shippingCarrierCode,omitempty"`

	// The shipping method code.
	ShippingMethodCode string `json:"shippingMethodCode,omitempty"`
}

// Validate validates this quote cart total management v1 collect totals put mine body
func (o *QuoteCartTotalManagementV1CollectTotalsPutMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdditionalData(formats); err != nil {
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

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineBody) validateAdditionalData(formats strfmt.Registry) error {

	if swag.IsZero(o.AdditionalData) { // not required
		return nil
	}

	if o.AdditionalData != nil {
		if err := o.AdditionalData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteCartTotalManagementV1CollectTotalsPutBody" + "." + "additionalData")
			}
			return err
		}
	}

	return nil
}

func (o *QuoteCartTotalManagementV1CollectTotalsPutMineBody) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("quoteCartTotalManagementV1CollectTotalsPutBody"+"."+"paymentMethod", "body", o.PaymentMethod); err != nil {
		return err
	}

	if o.PaymentMethod != nil {
		if err := o.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteCartTotalManagementV1CollectTotalsPutBody" + "." + "paymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteCartTotalManagementV1CollectTotalsPutMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteCartTotalManagementV1CollectTotalsPutMineBody) UnmarshalBinary(b []byte) error {
	var res QuoteCartTotalManagementV1CollectTotalsPutMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
