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

// DownloadableSampleRepositoryV1DeleteDeleteReader is a Reader for the DownloadableSampleRepositoryV1DeleteDelete structure.
type DownloadableSampleRepositoryV1DeleteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadableSampleRepositoryV1DeleteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDownloadableSampleRepositoryV1DeleteDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDownloadableSampleRepositoryV1DeleteDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDownloadableSampleRepositoryV1DeleteDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownloadableSampleRepositoryV1DeleteDeleteOK creates a DownloadableSampleRepositoryV1DeleteDeleteOK with default headers values
func NewDownloadableSampleRepositoryV1DeleteDeleteOK() *DownloadableSampleRepositoryV1DeleteDeleteOK {
	return &DownloadableSampleRepositoryV1DeleteDeleteOK{}
}

/*DownloadableSampleRepositoryV1DeleteDeleteOK handles this case with default header values.

200 Success.
*/
type DownloadableSampleRepositoryV1DeleteDeleteOK struct {
	Payload bool
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/downloadable-links/samples/{id}][%d] downloadableSampleRepositoryV1DeleteDeleteOK  %+v", 200, o.Payload)
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadableSampleRepositoryV1DeleteDeleteUnauthorized creates a DownloadableSampleRepositoryV1DeleteDeleteUnauthorized with default headers values
func NewDownloadableSampleRepositoryV1DeleteDeleteUnauthorized() *DownloadableSampleRepositoryV1DeleteDeleteUnauthorized {
	return &DownloadableSampleRepositoryV1DeleteDeleteUnauthorized{}
}

/*DownloadableSampleRepositoryV1DeleteDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type DownloadableSampleRepositoryV1DeleteDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/downloadable-links/samples/{id}][%d] downloadableSampleRepositoryV1DeleteDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadableSampleRepositoryV1DeleteDeleteDefault creates a DownloadableSampleRepositoryV1DeleteDeleteDefault with default headers values
func NewDownloadableSampleRepositoryV1DeleteDeleteDefault(code int) *DownloadableSampleRepositoryV1DeleteDeleteDefault {
	return &DownloadableSampleRepositoryV1DeleteDeleteDefault{
		_statusCode: code,
	}
}

/*DownloadableSampleRepositoryV1DeleteDeleteDefault handles this case with default header values.

Unexpected error
*/
type DownloadableSampleRepositoryV1DeleteDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the downloadable sample repository v1 delete delete default response
func (o *DownloadableSampleRepositoryV1DeleteDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/downloadable-links/samples/{id}][%d] downloadableSampleRepositoryV1DeleteDelete default  %+v", o._statusCode, o.Payload)
}

func (o *DownloadableSampleRepositoryV1DeleteDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
