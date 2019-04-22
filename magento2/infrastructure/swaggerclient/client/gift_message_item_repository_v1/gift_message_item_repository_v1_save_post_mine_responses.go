// Code generated by go-swagger; DO NOT EDIT.

package gift_message_item_repository_v1

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

// GiftMessageItemRepositoryV1SavePostMineReader is a Reader for the GiftMessageItemRepositoryV1SavePostMine structure.
type GiftMessageItemRepositoryV1SavePostMineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftMessageItemRepositoryV1SavePostMineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGiftMessageItemRepositoryV1SavePostMineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGiftMessageItemRepositoryV1SavePostMineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGiftMessageItemRepositoryV1SavePostMineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGiftMessageItemRepositoryV1SavePostMineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftMessageItemRepositoryV1SavePostMineOK creates a GiftMessageItemRepositoryV1SavePostMineOK with default headers values
func NewGiftMessageItemRepositoryV1SavePostMineOK() *GiftMessageItemRepositoryV1SavePostMineOK {
	return &GiftMessageItemRepositoryV1SavePostMineOK{}
}

/*GiftMessageItemRepositoryV1SavePostMineOK handles this case with default header values.

200 Success.
*/
type GiftMessageItemRepositoryV1SavePostMineOK struct {
	Payload bool
}

func (o *GiftMessageItemRepositoryV1SavePostMineOK) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message/{itemId}][%d] giftMessageItemRepositoryV1SavePostMineOK  %+v", 200, o.Payload)
}

func (o *GiftMessageItemRepositoryV1SavePostMineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageItemRepositoryV1SavePostMineBadRequest creates a GiftMessageItemRepositoryV1SavePostMineBadRequest with default headers values
func NewGiftMessageItemRepositoryV1SavePostMineBadRequest() *GiftMessageItemRepositoryV1SavePostMineBadRequest {
	return &GiftMessageItemRepositoryV1SavePostMineBadRequest{}
}

/*GiftMessageItemRepositoryV1SavePostMineBadRequest handles this case with default header values.

400 Bad Request
*/
type GiftMessageItemRepositoryV1SavePostMineBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageItemRepositoryV1SavePostMineBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message/{itemId}][%d] giftMessageItemRepositoryV1SavePostMineBadRequest  %+v", 400, o.Payload)
}

func (o *GiftMessageItemRepositoryV1SavePostMineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageItemRepositoryV1SavePostMineUnauthorized creates a GiftMessageItemRepositoryV1SavePostMineUnauthorized with default headers values
func NewGiftMessageItemRepositoryV1SavePostMineUnauthorized() *GiftMessageItemRepositoryV1SavePostMineUnauthorized {
	return &GiftMessageItemRepositoryV1SavePostMineUnauthorized{}
}

/*GiftMessageItemRepositoryV1SavePostMineUnauthorized handles this case with default header values.

401 Unauthorized
*/
type GiftMessageItemRepositoryV1SavePostMineUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageItemRepositoryV1SavePostMineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message/{itemId}][%d] giftMessageItemRepositoryV1SavePostMineUnauthorized  %+v", 401, o.Payload)
}

func (o *GiftMessageItemRepositoryV1SavePostMineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageItemRepositoryV1SavePostMineDefault creates a GiftMessageItemRepositoryV1SavePostMineDefault with default headers values
func NewGiftMessageItemRepositoryV1SavePostMineDefault(code int) *GiftMessageItemRepositoryV1SavePostMineDefault {
	return &GiftMessageItemRepositoryV1SavePostMineDefault{
		_statusCode: code,
	}
}

/*GiftMessageItemRepositoryV1SavePostMineDefault handles this case with default header values.

Unexpected error
*/
type GiftMessageItemRepositoryV1SavePostMineDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift message item repository v1 save post mine default response
func (o *GiftMessageItemRepositoryV1SavePostMineDefault) Code() int {
	return o._statusCode
}

func (o *GiftMessageItemRepositoryV1SavePostMineDefault) Error() string {
	return fmt.Sprintf("[POST /V1/carts/mine/gift-message/{itemId}][%d] giftMessageItemRepositoryV1SavePostMine default  %+v", o._statusCode, o.Payload)
}

func (o *GiftMessageItemRepositoryV1SavePostMineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GiftMessageItemRepositoryV1SavePostMineBody gift message item repository v1 save post mine body
swagger:model GiftMessageItemRepositoryV1SavePostMineBody
*/
type GiftMessageItemRepositoryV1SavePostMineBody struct {

	// gift message
	// Required: true
	GiftMessage *models.GiftMessageDataMessageInterface `json:"giftMessage"`
}

// Validate validates this gift message item repository v1 save post mine body
func (o *GiftMessageItemRepositoryV1SavePostMineBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageItemRepositoryV1SavePostMineBody) validateGiftMessage(formats strfmt.Registry) error {

	if err := validate.Required("giftMessageItemRepositoryV1SavePostBody"+"."+"giftMessage", "body", o.GiftMessage); err != nil {
		return err
	}

	if o.GiftMessage != nil {
		if err := o.GiftMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageItemRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GiftMessageItemRepositoryV1SavePostMineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GiftMessageItemRepositoryV1SavePostMineBody) UnmarshalBinary(b []byte) error {
	var res GiftMessageItemRepositoryV1SavePostMineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}