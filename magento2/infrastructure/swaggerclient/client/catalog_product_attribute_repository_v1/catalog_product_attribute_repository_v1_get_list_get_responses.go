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

// CatalogProductAttributeRepositoryV1GetListGetReader is a Reader for the CatalogProductAttributeRepositoryV1GetListGet structure.
type CatalogProductAttributeRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductAttributeRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductAttributeRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogProductAttributeRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductAttributeRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductAttributeRepositoryV1GetListGetOK creates a CatalogProductAttributeRepositoryV1GetListGetOK with default headers values
func NewCatalogProductAttributeRepositoryV1GetListGetOK() *CatalogProductAttributeRepositoryV1GetListGetOK {
	return &CatalogProductAttributeRepositoryV1GetListGetOK{}
}

/*CatalogProductAttributeRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CatalogProductAttributeRepositoryV1GetListGetOK struct {
	Payload *models.CatalogDataProductAttributeSearchResultsInterface
}

func (o *CatalogProductAttributeRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes][%d] catalogProductAttributeRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductAttributeSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeRepositoryV1GetListGetUnauthorized creates a CatalogProductAttributeRepositoryV1GetListGetUnauthorized with default headers values
func NewCatalogProductAttributeRepositoryV1GetListGetUnauthorized() *CatalogProductAttributeRepositoryV1GetListGetUnauthorized {
	return &CatalogProductAttributeRepositoryV1GetListGetUnauthorized{}
}

/*CatalogProductAttributeRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductAttributeRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes][%d] catalogProductAttributeRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeRepositoryV1GetListGetDefault creates a CatalogProductAttributeRepositoryV1GetListGetDefault with default headers values
func NewCatalogProductAttributeRepositoryV1GetListGetDefault(code int) *CatalogProductAttributeRepositoryV1GetListGetDefault {
	return &CatalogProductAttributeRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CatalogProductAttributeRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductAttributeRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product attribute repository v1 get list get default response
func (o *CatalogProductAttributeRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductAttributeRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/attributes][%d] catalogProductAttributeRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductAttributeRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
