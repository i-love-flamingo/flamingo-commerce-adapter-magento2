// Code generated by go-swagger; DO NOT EDIT.

package store_group_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// StoreGroupRepositoryV1GetListGetReader is a Reader for the StoreGroupRepositoryV1GetListGet structure.
type StoreGroupRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoreGroupRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStoreGroupRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewStoreGroupRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewStoreGroupRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoreGroupRepositoryV1GetListGetOK creates a StoreGroupRepositoryV1GetListGetOK with default headers values
func NewStoreGroupRepositoryV1GetListGetOK() *StoreGroupRepositoryV1GetListGetOK {
	return &StoreGroupRepositoryV1GetListGetOK{}
}

/*StoreGroupRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type StoreGroupRepositoryV1GetListGetOK struct {
	Payload []*models.StoreDataGroupInterface
}

func (o *StoreGroupRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/store/storeGroups][%d] storeGroupRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *StoreGroupRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoreGroupRepositoryV1GetListGetUnauthorized creates a StoreGroupRepositoryV1GetListGetUnauthorized with default headers values
func NewStoreGroupRepositoryV1GetListGetUnauthorized() *StoreGroupRepositoryV1GetListGetUnauthorized {
	return &StoreGroupRepositoryV1GetListGetUnauthorized{}
}

/*StoreGroupRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type StoreGroupRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *StoreGroupRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/store/storeGroups][%d] storeGroupRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *StoreGroupRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoreGroupRepositoryV1GetListGetDefault creates a StoreGroupRepositoryV1GetListGetDefault with default headers values
func NewStoreGroupRepositoryV1GetListGetDefault(code int) *StoreGroupRepositoryV1GetListGetDefault {
	return &StoreGroupRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*StoreGroupRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type StoreGroupRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the store group repository v1 get list get default response
func (o *StoreGroupRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *StoreGroupRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/store/storeGroups][%d] storeGroupRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *StoreGroupRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
