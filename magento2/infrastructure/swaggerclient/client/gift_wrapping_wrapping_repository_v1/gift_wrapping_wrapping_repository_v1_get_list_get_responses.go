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

// GiftWrappingWrappingRepositoryV1GetListGetReader is a Reader for the GiftWrappingWrappingRepositoryV1GetListGet structure.
type GiftWrappingWrappingRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftWrappingWrappingRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGiftWrappingWrappingRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGiftWrappingWrappingRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGiftWrappingWrappingRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftWrappingWrappingRepositoryV1GetListGetOK creates a GiftWrappingWrappingRepositoryV1GetListGetOK with default headers values
func NewGiftWrappingWrappingRepositoryV1GetListGetOK() *GiftWrappingWrappingRepositoryV1GetListGetOK {
	return &GiftWrappingWrappingRepositoryV1GetListGetOK{}
}

/*GiftWrappingWrappingRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type GiftWrappingWrappingRepositoryV1GetListGetOK struct {
	Payload *models.GiftWrappingDataWrappingSearchResultsInterface
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/gift-wrappings][%d] giftWrappingWrappingRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GiftWrappingDataWrappingSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1GetListGetUnauthorized creates a GiftWrappingWrappingRepositoryV1GetListGetUnauthorized with default headers values
func NewGiftWrappingWrappingRepositoryV1GetListGetUnauthorized() *GiftWrappingWrappingRepositoryV1GetListGetUnauthorized {
	return &GiftWrappingWrappingRepositoryV1GetListGetUnauthorized{}
}

/*GiftWrappingWrappingRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type GiftWrappingWrappingRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/gift-wrappings][%d] giftWrappingWrappingRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftWrappingWrappingRepositoryV1GetListGetDefault creates a GiftWrappingWrappingRepositoryV1GetListGetDefault with default headers values
func NewGiftWrappingWrappingRepositoryV1GetListGetDefault(code int) *GiftWrappingWrappingRepositoryV1GetListGetDefault {
	return &GiftWrappingWrappingRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*GiftWrappingWrappingRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type GiftWrappingWrappingRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift wrapping wrapping repository v1 get list get default response
func (o *GiftWrappingWrappingRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/gift-wrappings][%d] giftWrappingWrappingRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *GiftWrappingWrappingRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
