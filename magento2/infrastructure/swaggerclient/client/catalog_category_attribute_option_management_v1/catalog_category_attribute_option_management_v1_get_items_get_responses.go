// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_attribute_option_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogCategoryAttributeOptionManagementV1GetItemsGetReader is a Reader for the CatalogCategoryAttributeOptionManagementV1GetItemsGet structure.
type CatalogCategoryAttributeOptionManagementV1GetItemsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogCategoryAttributeOptionManagementV1GetItemsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogCategoryAttributeOptionManagementV1GetItemsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryAttributeOptionManagementV1GetItemsGetOK creates a CatalogCategoryAttributeOptionManagementV1GetItemsGetOK with default headers values
func NewCatalogCategoryAttributeOptionManagementV1GetItemsGetOK() *CatalogCategoryAttributeOptionManagementV1GetItemsGetOK {
	return &CatalogCategoryAttributeOptionManagementV1GetItemsGetOK{}
}

/*CatalogCategoryAttributeOptionManagementV1GetItemsGetOK handles this case with default header values.

200 Success.
*/
type CatalogCategoryAttributeOptionManagementV1GetItemsGetOK struct {
	Payload []*models.EavDataAttributeOptionInterface
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes/{attributeCode}/options][%d] catalogCategoryAttributeOptionManagementV1GetItemsGetOK  %+v", 200, o.Payload)
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest creates a CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest with default headers values
func NewCatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest() *CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest {
	return &CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest{}
}

/*CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes/{attributeCode}/options][%d] catalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized creates a CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized with default headers values
func NewCatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized() *CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized {
	return &CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized{}
}

/*CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes/{attributeCode}/options][%d] catalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryAttributeOptionManagementV1GetItemsGetDefault creates a CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault with default headers values
func NewCatalogCategoryAttributeOptionManagementV1GetItemsGetDefault(code int) *CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault {
	return &CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault{
		_statusCode: code,
	}
}

/*CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category attribute option management v1 get items get default response
func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories/attributes/{attributeCode}/options][%d] catalogCategoryAttributeOptionManagementV1GetItemsGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogCategoryAttributeOptionManagementV1GetItemsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
