// Code generated by go-swagger; DO NOT EDIT.

package quote_shipment_estimation_v1

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

// QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineReader is a Reader for the QuoteShipmentEstimationV1EstimateByExtendedAddressPostMine structure.
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK creates a QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK with default headers values
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK() *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK {
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK{}
}

/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK handles this case with default header values.

200 Success.
*/
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK struct {
	Payload []*models.QuoteDataShippingMethodInterface
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/estimate-shipping-methods][%d] quoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK  %+v", 200, o.Payload)
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized creates a QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized with default headers values
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized() *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized {
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized{}
}

/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/estimate-shipping-methods][%d] quoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault creates a QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault with default headers values
func NewQuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault(code int) *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault {
	return &QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault{
		_statusCode: code,
	}
}

/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault handles this case with default header values.

Unexpected error
*/
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote shipment estimation v1 estimate by extended address post mine default response
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault) Code() int {
	return o._statusCode
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/estimate-shipping-methods][%d] quoteShipmentEstimationV1EstimateByExtendedAddressPostMine default  %+v", o._statusCode, o.Payload)
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody quote shipment estimation v1 estimate by extended address post mine body
swagger:model QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody
*/
type QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody struct {

	// address
	// Required: true
	Address *models.QuoteDataAddressInterface `json:"address"`
}

// Validate validates this quote shipment estimation v1 estimate by extended address post mine body
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("quoteShipmentEstimationV1EstimateByExtendedAddressPostBody"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	if o.Address != nil {
		if err := o.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteShipmentEstimationV1EstimateByExtendedAddressPostBody" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody) UnmarshalBinary(b []byte) error {
	var res QuoteShipmentEstimationV1EstimateByExtendedAddressPostMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
