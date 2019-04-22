// Code generated by go-swagger; DO NOT EDIT.

package gift_wrapping_wrapping_repository_v1

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

// GiftWrappingWrappingRepositoryV1SavePutReader is a Reader for the GiftWrappingWrappingRepositoryV1SavePut structure.
type GiftWrappingWrappingRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftWrappingWrappingRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGiftWrappingWrappingRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGiftWrappingWrappingRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGiftWrappingWrappingRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGiftWrappingWrappingRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftWrappingWrappingRepositoryV1SavePutOK creates a GiftWrappingWrappingRepositoryV1SavePutOK with default headers values
func NewGiftWrappingWrappingRepositoryV1SavePutOK() *GiftWrappingWrappingRepositoryV1SavePutOK {
	return &GiftWrappingWrappingRepositoryV1SavePutOK{}
}

/*GiftWrappingWrappingRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type GiftWrappingWrappingRepositoryV1SavePutOK struct {
	Payload *models.GiftWrappingDataWrappingInterface
}

func (o *GiftWrappingWrappingRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/gift-wrappings/{wrappingId}][%d] giftWrappingWrappingRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GiftWrappingDataWrappingInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1SavePutBadRequest creates a GiftWrappingWrappingRepositoryV1SavePutBadRequest with default headers values
func NewGiftWrappingWrappingRepositoryV1SavePutBadRequest() *GiftWrappingWrappingRepositoryV1SavePutBadRequest {
	return &GiftWrappingWrappingRepositoryV1SavePutBadRequest{}
}

/*GiftWrappingWrappingRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type GiftWrappingWrappingRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftWrappingWrappingRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/gift-wrappings/{wrappingId}][%d] giftWrappingWrappingRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1SavePutUnauthorized creates a GiftWrappingWrappingRepositoryV1SavePutUnauthorized with default headers values
func NewGiftWrappingWrappingRepositoryV1SavePutUnauthorized() *GiftWrappingWrappingRepositoryV1SavePutUnauthorized {
	return &GiftWrappingWrappingRepositoryV1SavePutUnauthorized{}
}

/*GiftWrappingWrappingRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type GiftWrappingWrappingRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GiftWrappingWrappingRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/gift-wrappings/{wrappingId}][%d] giftWrappingWrappingRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1SavePutDefault creates a GiftWrappingWrappingRepositoryV1SavePutDefault with default headers values
func NewGiftWrappingWrappingRepositoryV1SavePutDefault(code int) *GiftWrappingWrappingRepositoryV1SavePutDefault {
	return &GiftWrappingWrappingRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*GiftWrappingWrappingRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type GiftWrappingWrappingRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift wrapping wrapping repository v1 save put default response
func (o *GiftWrappingWrappingRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *GiftWrappingWrappingRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/gift-wrappings/{wrappingId}][%d] giftWrappingWrappingRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GiftWrappingWrappingRepositoryV1SavePutBody gift wrapping wrapping repository v1 save put body
swagger:model GiftWrappingWrappingRepositoryV1SavePutBody
*/
type GiftWrappingWrappingRepositoryV1SavePutBody struct {

	// data
	// Required: true
	Data *models.GiftWrappingDataWrappingInterface `json:"data"`

	// store Id
	StoreID int64 `json:"storeId,omitempty"`
}

// Validate validates this gift wrapping wrapping repository v1 save put body
func (o *GiftWrappingWrappingRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftWrappingWrappingRepositoryV1SavePutBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("giftWrappingWrappingRepositoryV1SavePutBody"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftWrappingWrappingRepositoryV1SavePutBody" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GiftWrappingWrappingRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GiftWrappingWrappingRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res GiftWrappingWrappingRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}