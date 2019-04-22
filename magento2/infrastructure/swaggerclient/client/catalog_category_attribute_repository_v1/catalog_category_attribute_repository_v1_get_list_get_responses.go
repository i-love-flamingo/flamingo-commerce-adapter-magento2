// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_attribute_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogCategoryAttributeRepositoryV1GetListGetReader is a Reader for the CatalogCategoryAttributeRepositoryV1GetListGet structure.
type CatalogCategoryAttributeRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryAttributeRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogCategoryAttributeRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogCategoryAttributeRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogCategoryAttributeRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryAttributeRepositoryV1GetListGetOK creates a CatalogCategoryAttributeRepositoryV1GetListGetOK with default headers values
func NewCatalogCategoryAttributeRepositoryV1GetListGetOK() *CatalogCategoryAttributeRepositoryV1GetListGetOK {
	return &CatalogCategoryAttributeRepositoryV1GetListGetOK{}
}

/*CatalogCategoryAttributeRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CatalogCategoryAttributeRepositoryV1GetListGetOK struct {
	Payload *models.CatalogDataCategoryAttributeSearchResultsInterface
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes][%d] catalogCategoryAttributeRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataCategoryAttributeSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryAttributeRepositoryV1GetListGetUnauthorized creates a CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized with default headers values
func NewCatalogCategoryAttributeRepositoryV1GetListGetUnauthorized() *CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized {
	return &CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized{}
}

/*CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes][%d] catalogCategoryAttributeRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryAttributeRepositoryV1GetListGetDefault creates a CatalogCategoryAttributeRepositoryV1GetListGetDefault with default headers values
func NewCatalogCategoryAttributeRepositoryV1GetListGetDefault(code int) *CatalogCategoryAttributeRepositoryV1GetListGetDefault {
	return &CatalogCategoryAttributeRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CatalogCategoryAttributeRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogCategoryAttributeRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category attribute repository v1 get list get default response
func (o *CatalogCategoryAttributeRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes][%d] catalogCategoryAttributeRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogCategoryAttributeRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}