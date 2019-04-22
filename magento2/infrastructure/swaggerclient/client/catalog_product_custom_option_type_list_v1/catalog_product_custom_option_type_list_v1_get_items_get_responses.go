// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_custom_option_type_list_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductCustomOptionTypeListV1GetItemsGetReader is a Reader for the CatalogProductCustomOptionTypeListV1GetItemsGet structure.
type CatalogProductCustomOptionTypeListV1GetItemsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductCustomOptionTypeListV1GetItemsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductCustomOptionTypeListV1GetItemsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductCustomOptionTypeListV1GetItemsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductCustomOptionTypeListV1GetItemsGetOK creates a CatalogProductCustomOptionTypeListV1GetItemsGetOK with default headers values
func NewCatalogProductCustomOptionTypeListV1GetItemsGetOK() *CatalogProductCustomOptionTypeListV1GetItemsGetOK {
	return &CatalogProductCustomOptionTypeListV1GetItemsGetOK{}
}

/*CatalogProductCustomOptionTypeListV1GetItemsGetOK handles this case with default header values.

200 Success.
*/
type CatalogProductCustomOptionTypeListV1GetItemsGetOK struct {
	Payload []*models.CatalogDataProductCustomOptionTypeInterface
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/options/types][%d] catalogProductCustomOptionTypeListV1GetItemsGetOK  %+v", 200, o.Payload)
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized creates a CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized with default headers values
func NewCatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized() *CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized {
	return &CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized{}
}

/*CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/options/types][%d] catalogProductCustomOptionTypeListV1GetItemsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductCustomOptionTypeListV1GetItemsGetDefault creates a CatalogProductCustomOptionTypeListV1GetItemsGetDefault with default headers values
func NewCatalogProductCustomOptionTypeListV1GetItemsGetDefault(code int) *CatalogProductCustomOptionTypeListV1GetItemsGetDefault {
	return &CatalogProductCustomOptionTypeListV1GetItemsGetDefault{
		_statusCode: code,
	}
}

/*CatalogProductCustomOptionTypeListV1GetItemsGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductCustomOptionTypeListV1GetItemsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product custom option type list v1 get items get default response
func (o *CatalogProductCustomOptionTypeListV1GetItemsGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/options/types][%d] catalogProductCustomOptionTypeListV1GetItemsGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductCustomOptionTypeListV1GetItemsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
