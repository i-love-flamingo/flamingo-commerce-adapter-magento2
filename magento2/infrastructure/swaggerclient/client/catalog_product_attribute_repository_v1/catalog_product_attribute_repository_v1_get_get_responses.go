// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductAttributeRepositoryV1GetGetReader is a Reader for the CatalogProductAttributeRepositoryV1GetGet structure.
type CatalogProductAttributeRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductAttributeRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductAttributeRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductAttributeRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductAttributeRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductAttributeRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductAttributeRepositoryV1GetGetOK creates a CatalogProductAttributeRepositoryV1GetGetOK with default headers values
func NewCatalogProductAttributeRepositoryV1GetGetOK() *CatalogProductAttributeRepositoryV1GetGetOK {
	return &CatalogProductAttributeRepositoryV1GetGetOK{}
}

/*CatalogProductAttributeRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type CatalogProductAttributeRepositoryV1GetGetOK struct {
	Payload *models.CatalogDataProductAttributeInterface
}

func (o *CatalogProductAttributeRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes/{attributeCode}][%d] catalogProductAttributeRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductAttributeInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeRepositoryV1GetGetBadRequest creates a CatalogProductAttributeRepositoryV1GetGetBadRequest with default headers values
func NewCatalogProductAttributeRepositoryV1GetGetBadRequest() *CatalogProductAttributeRepositoryV1GetGetBadRequest {
	return &CatalogProductAttributeRepositoryV1GetGetBadRequest{}
}

/*CatalogProductAttributeRepositoryV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductAttributeRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes/{attributeCode}][%d] catalogProductAttributeRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeRepositoryV1GetGetUnauthorized creates a CatalogProductAttributeRepositoryV1GetGetUnauthorized with default headers values
func NewCatalogProductAttributeRepositoryV1GetGetUnauthorized() *CatalogProductAttributeRepositoryV1GetGetUnauthorized {
	return &CatalogProductAttributeRepositoryV1GetGetUnauthorized{}
}

/*CatalogProductAttributeRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductAttributeRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes/{attributeCode}][%d] catalogProductAttributeRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeRepositoryV1GetGetDefault creates a CatalogProductAttributeRepositoryV1GetGetDefault with default headers values
func NewCatalogProductAttributeRepositoryV1GetGetDefault(code int) *CatalogProductAttributeRepositoryV1GetGetDefault {
	return &CatalogProductAttributeRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*CatalogProductAttributeRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductAttributeRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product attribute repository v1 get get default response
func (o *CatalogProductAttributeRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductAttributeRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes/{attributeCode}][%d] catalogProductAttributeRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
