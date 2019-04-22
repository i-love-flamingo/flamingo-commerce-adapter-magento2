// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_link_type_list_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductLinkTypeListV1GetItemsGetReader is a Reader for the CatalogProductLinkTypeListV1GetItemsGet structure.
type CatalogProductLinkTypeListV1GetItemsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductLinkTypeListV1GetItemsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductLinkTypeListV1GetItemsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogProductLinkTypeListV1GetItemsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductLinkTypeListV1GetItemsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductLinkTypeListV1GetItemsGetOK creates a CatalogProductLinkTypeListV1GetItemsGetOK with default headers values
func NewCatalogProductLinkTypeListV1GetItemsGetOK() *CatalogProductLinkTypeListV1GetItemsGetOK {
	return &CatalogProductLinkTypeListV1GetItemsGetOK{}
}

/*CatalogProductLinkTypeListV1GetItemsGetOK handles this case with default header values.

200 Success.
*/
type CatalogProductLinkTypeListV1GetItemsGetOK struct {
	Payload []*models.CatalogDataProductLinkTypeInterface
}

func (o *CatalogProductLinkTypeListV1GetItemsGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/links/types][%d] catalogProductLinkTypeListV1GetItemsGetOK  %+v", 200, o.Payload)
}

func (o *CatalogProductLinkTypeListV1GetItemsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductLinkTypeListV1GetItemsGetUnauthorized creates a CatalogProductLinkTypeListV1GetItemsGetUnauthorized with default headers values
func NewCatalogProductLinkTypeListV1GetItemsGetUnauthorized() *CatalogProductLinkTypeListV1GetItemsGetUnauthorized {
	return &CatalogProductLinkTypeListV1GetItemsGetUnauthorized{}
}

/*CatalogProductLinkTypeListV1GetItemsGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductLinkTypeListV1GetItemsGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductLinkTypeListV1GetItemsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/links/types][%d] catalogProductLinkTypeListV1GetItemsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductLinkTypeListV1GetItemsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductLinkTypeListV1GetItemsGetDefault creates a CatalogProductLinkTypeListV1GetItemsGetDefault with default headers values
func NewCatalogProductLinkTypeListV1GetItemsGetDefault(code int) *CatalogProductLinkTypeListV1GetItemsGetDefault {
	return &CatalogProductLinkTypeListV1GetItemsGetDefault{
		_statusCode: code,
	}
}

/*CatalogProductLinkTypeListV1GetItemsGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductLinkTypeListV1GetItemsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product link type list v1 get items get default response
func (o *CatalogProductLinkTypeListV1GetItemsGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductLinkTypeListV1GetItemsGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/links/types][%d] catalogProductLinkTypeListV1GetItemsGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductLinkTypeListV1GetItemsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
