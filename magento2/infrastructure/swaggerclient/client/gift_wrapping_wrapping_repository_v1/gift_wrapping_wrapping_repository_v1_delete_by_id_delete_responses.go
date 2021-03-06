// Code generated by go-swagger; DO NOT EDIT.

package gift_wrapping_wrapping_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// GiftWrappingWrappingRepositoryV1DeleteByIDDeleteReader is a Reader for the GiftWrappingWrappingRepositoryV1DeleteByIDDelete structure.
type GiftWrappingWrappingRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK creates a GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK with default headers values
func NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK() *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK {
	return &GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK{}
}

/*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/gift-wrappings/{id}][%d] giftWrappingWrappingRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest creates a GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest() *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest {
	return &GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/gift-wrappings/{id}][%d] giftWrappingWrappingRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized creates a GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized() *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized {
	return &GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/gift-wrappings/{id}][%d] giftWrappingWrappingRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault creates a GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewGiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault(code int) *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault {
	return &GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift wrapping wrapping repository v1 delete by Id delete default response
func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/gift-wrappings/{id}][%d] giftWrappingWrappingRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
