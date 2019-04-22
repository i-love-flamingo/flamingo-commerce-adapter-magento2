// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_negotiable_cart_repository_v1

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

// NegotiableQuoteNegotiableCartRepositoryV1SavePutReader is a Reader for the NegotiableQuoteNegotiableCartRepositoryV1SavePut structure.
type NegotiableQuoteNegotiableCartRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 401:
		result := NewNegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNegotiableQuoteNegotiableCartRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized creates a NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized with default headers values
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized() *NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized {
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized{}
}

/*NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/negotiableQuote/{quoteId}][%d] negotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteNegotiableCartRepositoryV1SavePutDefault creates a NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault with default headers values
func NewNegotiableQuoteNegotiableCartRepositoryV1SavePutDefault(code int) *NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault {
	return &NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the negotiable quote negotiable cart repository v1 save put default response
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/negotiableQuote/{quoteId}][%d] negotiableQuoteNegotiableCartRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NegotiableQuoteNegotiableCartRepositoryV1SavePutBody negotiable quote negotiable cart repository v1 save put body
swagger:model NegotiableQuoteNegotiableCartRepositoryV1SavePutBody
*/
type NegotiableQuoteNegotiableCartRepositoryV1SavePutBody struct {

	// quote
	// Required: true
	Quote *models.QuoteDataCartInterface `json:"quote"`
}

// Validate validates this negotiable quote negotiable cart repository v1 save put body
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQuote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) validateQuote(formats strfmt.Registry) error {

	if err := validate.Required("negotiableQuoteNegotiableCartRepositoryV1SavePutBody"+"."+"quote", "body", o.Quote); err != nil {
		return err
	}

	if o.Quote != nil {
		if err := o.Quote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negotiableQuoteNegotiableCartRepositoryV1SavePutBody" + "." + "quote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NegotiableQuoteNegotiableCartRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res NegotiableQuoteNegotiableCartRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
