// Code generated by go-swagger; DO NOT EDIT.

package downloadable_sample_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// DownloadableSampleRepositoryV1GetListGetReader is a Reader for the DownloadableSampleRepositoryV1GetListGet structure.
type DownloadableSampleRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadableSampleRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDownloadableSampleRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDownloadableSampleRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDownloadableSampleRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownloadableSampleRepositoryV1GetListGetOK creates a DownloadableSampleRepositoryV1GetListGetOK with default headers values
func NewDownloadableSampleRepositoryV1GetListGetOK() *DownloadableSampleRepositoryV1GetListGetOK {
	return &DownloadableSampleRepositoryV1GetListGetOK{}
}

/*DownloadableSampleRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type DownloadableSampleRepositoryV1GetListGetOK struct {
	Payload []*models.DownloadableDataSampleInterface
}

func (o *DownloadableSampleRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/downloadable-links/samples][%d] downloadableSampleRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *DownloadableSampleRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadableSampleRepositoryV1GetListGetUnauthorized creates a DownloadableSampleRepositoryV1GetListGetUnauthorized with default headers values
func NewDownloadableSampleRepositoryV1GetListGetUnauthorized() *DownloadableSampleRepositoryV1GetListGetUnauthorized {
	return &DownloadableSampleRepositoryV1GetListGetUnauthorized{}
}

/*DownloadableSampleRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type DownloadableSampleRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DownloadableSampleRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/downloadable-links/samples][%d] downloadableSampleRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadableSampleRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadableSampleRepositoryV1GetListGetDefault creates a DownloadableSampleRepositoryV1GetListGetDefault with default headers values
func NewDownloadableSampleRepositoryV1GetListGetDefault(code int) *DownloadableSampleRepositoryV1GetListGetDefault {
	return &DownloadableSampleRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*DownloadableSampleRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type DownloadableSampleRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the downloadable sample repository v1 get list get default response
func (o *DownloadableSampleRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *DownloadableSampleRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/downloadable-links/samples][%d] downloadableSampleRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *DownloadableSampleRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
