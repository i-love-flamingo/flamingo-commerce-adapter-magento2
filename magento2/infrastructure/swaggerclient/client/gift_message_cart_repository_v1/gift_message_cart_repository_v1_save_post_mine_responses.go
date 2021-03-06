// Code generated by go-swagger; DO NOT EDIT.

package gift_message_cart_repository_v1

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

// GiftMessageCartRepositoryV1SavePostMineReader is a Reader for the GiftMessageCartRepositoryV1SavePostMine structure.
type GiftMessageCartRepositoryV1SavePostMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftMessageCartRepositoryV1SavePostMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGiftMessageCartRepositoryV1SavePostMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGiftMessageCartRepositoryV1SavePostMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGiftMessageCartRepositoryV1SavePostMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGiftMessageCartRepositoryV1SavePostMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftMessageCartRepositoryV1SavePostMineOK creates a GiftMessageCartRepositoryV1SavePostMineOK with default headers values
func NewGiftMessageCartRepositoryV1SavePostMineOK() *GiftMessageCartRepositoryV1SavePostMineOK {
	return &GiftMessageCartRepositoryV1SavePostMineOK{}
}

/*GiftMessageCartRepositoryV1SavePostMineOK handles this case with default header values.

200 Success.
*/
type GiftMessageCartRepositoryV1SavePostMineOK struct {
	Payload bool
}

func (o *GiftMessageCartRepositoryV1SavePostMineOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message][%d] giftMessageCartRepositoryV1SavePostMineOK  %+v", 200, o.Payload)
}

func (o *GiftMessageCartRepositoryV1SavePostMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageCartRepositoryV1SavePostMineBadRequest creates a GiftMessageCartRepositoryV1SavePostMineBadRequest with default headers values
func NewGiftMessageCartRepositoryV1SavePostMineBadRequest() *GiftMessageCartRepositoryV1SavePostMineBadRequest {
	return &GiftMessageCartRepositoryV1SavePostMineBadRequest{}
}

/*GiftMessageCartRepositoryV1SavePostMineBadRequest handles this case with default header values.

400 Bad Request
*/
type GiftMessageCartRepositoryV1SavePostMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageCartRepositoryV1SavePostMineBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message][%d] giftMessageCartRepositoryV1SavePostMineBadRequest  %+v", 400, o.Payload)
}

func (o *GiftMessageCartRepositoryV1SavePostMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageCartRepositoryV1SavePostMineUnauthorized creates a GiftMessageCartRepositoryV1SavePostMineUnauthorized with default headers values
func NewGiftMessageCartRepositoryV1SavePostMineUnauthorized() *GiftMessageCartRepositoryV1SavePostMineUnauthorized {
	return &GiftMessageCartRepositoryV1SavePostMineUnauthorized{}
}

/*GiftMessageCartRepositoryV1SavePostMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type GiftMessageCartRepositoryV1SavePostMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageCartRepositoryV1SavePostMineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message][%d] giftMessageCartRepositoryV1SavePostMineUnauthorized  %+v", 401, o.Payload)
}

func (o *GiftMessageCartRepositoryV1SavePostMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageCartRepositoryV1SavePostMineDefault creates a GiftMessageCartRepositoryV1SavePostMineDefault with default headers values
func NewGiftMessageCartRepositoryV1SavePostMineDefault(code int) *GiftMessageCartRepositoryV1SavePostMineDefault {
	return &GiftMessageCartRepositoryV1SavePostMineDefault{
		_statusCode: code,
	}
}

/*GiftMessageCartRepositoryV1SavePostMineDefault handles this case with default header values.

Unexpected error
*/
type GiftMessageCartRepositoryV1SavePostMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift message cart repository v1 save post mine default response
func (o *GiftMessageCartRepositoryV1SavePostMineDefault) Code() int {
	return o._statusCode
}

func (o *GiftMessageCartRepositoryV1SavePostMineDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message][%d] giftMessageCartRepositoryV1SavePostMine default  %+v", o._statusCode, o.Payload)
}

func (o *GiftMessageCartRepositoryV1SavePostMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GiftMessageCartRepositoryV1SavePostMineBody gift message cart repository v1 save post mine body
swagger:model GiftMessageCartRepositoryV1SavePostMineBody
*/
type GiftMessageCartRepositoryV1SavePostMineBody struct {

	// gift message
	// Required: true
	GiftMessage *models.GiftMessageDataMessageInterface `json:"giftMessage"`
}

// Validate validates this gift message cart repository v1 save post mine body
func (o *GiftMessageCartRepositoryV1SavePostMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageCartRepositoryV1SavePostMineBody) validateGiftMessage(formats strfmt.Registry) error {

	if err := validate.Required("giftMessageCartRepositoryV1SavePostBody"+"."+"giftMessage", "body", o.GiftMessage); err != nil {
		return err
	}

	if o.GiftMessage != nil {
		if err := o.GiftMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageCartRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GiftMessageCartRepositoryV1SavePostMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GiftMessageCartRepositoryV1SavePostMineBody) UnmarshalBinary(b []byte) error {
	var res GiftMessageCartRepositoryV1SavePostMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
