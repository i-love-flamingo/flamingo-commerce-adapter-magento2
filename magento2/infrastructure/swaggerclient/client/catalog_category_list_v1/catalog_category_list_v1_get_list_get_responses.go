// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_list_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogCategoryListV1GetListGetReader is a Reader for the CatalogCategoryListV1GetListGet structure.
type CatalogCategoryListV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryListV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogCategoryListV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogCategoryListV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogCategoryListV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryListV1GetListGetOK creates a CatalogCategoryListV1GetListGetOK with default headers values
func NewCatalogCategoryListV1GetListGetOK() *CatalogCategoryListV1GetListGetOK {
	return &CatalogCategoryListV1GetListGetOK{}
}

/*CatalogCategoryListV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CatalogCategoryListV1GetListGetOK struct {
	Payload *models.CatalogDataCategorySearchResultsInterface
}

func (o *CatalogCategoryListV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories/list][%d] catalogCategoryListV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CatalogCategoryListV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataCategorySearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryListV1GetListGetUnauthorized creates a CatalogCategoryListV1GetListGetUnauthorized with default headers values
func NewCatalogCategoryListV1GetListGetUnauthorized() *CatalogCategoryListV1GetListGetUnauthorized {
	return &CatalogCategoryListV1GetListGetUnauthorized{}
}

/*CatalogCategoryListV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogCategoryListV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryListV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/categories/list][%d] catalogCategoryListV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCategoryListV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryListV1GetListGetDefault creates a CatalogCategoryListV1GetListGetDefault with default headers values
func NewCatalogCategoryListV1GetListGetDefault(code int) *CatalogCategoryListV1GetListGetDefault {
	return &CatalogCategoryListV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CatalogCategoryListV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogCategoryListV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category list v1 get list get default response
func (o *CatalogCategoryListV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryListV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories/list][%d] catalogCategoryListV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogCategoryListV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
