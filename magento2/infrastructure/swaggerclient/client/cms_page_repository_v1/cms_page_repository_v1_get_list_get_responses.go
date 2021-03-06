// Code generated by go-swagger; DO NOT EDIT.

package cms_page_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CmsPageRepositoryV1GetListGetReader is a Reader for the CmsPageRepositoryV1GetListGet structure.
type CmsPageRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CmsPageRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCmsPageRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCmsPageRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCmsPageRepositoryV1GetListGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCmsPageRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCmsPageRepositoryV1GetListGetOK creates a CmsPageRepositoryV1GetListGetOK with default headers values
func NewCmsPageRepositoryV1GetListGetOK() *CmsPageRepositoryV1GetListGetOK {
	return &CmsPageRepositoryV1GetListGetOK{}
}

/*CmsPageRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CmsPageRepositoryV1GetListGetOK struct {
	Payload *models.CmsDataPageSearchResultsInterface
}

func (o *CmsPageRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/search][%d] cmsPageRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CmsPageRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CmsDataPageSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsPageRepositoryV1GetListGetUnauthorized creates a CmsPageRepositoryV1GetListGetUnauthorized with default headers values
func NewCmsPageRepositoryV1GetListGetUnauthorized() *CmsPageRepositoryV1GetListGetUnauthorized {
	return &CmsPageRepositoryV1GetListGetUnauthorized{}
}

/*CmsPageRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CmsPageRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CmsPageRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/search][%d] cmsPageRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CmsPageRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsPageRepositoryV1GetListGetInternalServerError creates a CmsPageRepositoryV1GetListGetInternalServerError with default headers values
func NewCmsPageRepositoryV1GetListGetInternalServerError() *CmsPageRepositoryV1GetListGetInternalServerError {
	return &CmsPageRepositoryV1GetListGetInternalServerError{}
}

/*CmsPageRepositoryV1GetListGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CmsPageRepositoryV1GetListGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CmsPageRepositoryV1GetListGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/search][%d] cmsPageRepositoryV1GetListGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CmsPageRepositoryV1GetListGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsPageRepositoryV1GetListGetDefault creates a CmsPageRepositoryV1GetListGetDefault with default headers values
func NewCmsPageRepositoryV1GetListGetDefault(code int) *CmsPageRepositoryV1GetListGetDefault {
	return &CmsPageRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CmsPageRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CmsPageRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cms page repository v1 get list get default response
func (o *CmsPageRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CmsPageRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/search][%d] cmsPageRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CmsPageRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
