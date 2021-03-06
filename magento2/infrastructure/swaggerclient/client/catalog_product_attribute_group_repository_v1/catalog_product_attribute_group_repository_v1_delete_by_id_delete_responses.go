// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_group_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteReader is a Reader for the CatalogProductAttributeGroupRepositoryV1DeleteByIDDelete structure.
type CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK creates a CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK with default headers values
func NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK() *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK {
	return &CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK{}
}

/*CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/attribute-sets/groups/{groupId}][%d] catalogProductAttributeGroupRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized creates a CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized() *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized {
	return &CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/attribute-sets/groups/{groupId}][%d] catalogProductAttributeGroupRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault creates a CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewCatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault(code int) *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault {
	return &CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product attribute group repository v1 delete by Id delete default response
func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/attribute-sets/groups/{groupId}][%d] catalogProductAttributeGroupRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductAttributeGroupRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
