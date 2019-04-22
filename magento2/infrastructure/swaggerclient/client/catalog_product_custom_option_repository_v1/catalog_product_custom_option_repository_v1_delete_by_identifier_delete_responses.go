// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_custom_option_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteReader is a Reader for the CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDelete structure.
type CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK creates a CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK with default headers values
func NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK() *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK {
	return &CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK{}
}

/*CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK handles this case with default header values.

200 Success.
*/
type CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK struct {
	Payload bool
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}/options/{optionId}][%d] catalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK  %+v", 200, o.Payload)
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized creates a CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized with default headers values
func NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized() *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized {
	return &CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized{}
}

/*CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}/options/{optionId}][%d] catalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault creates a CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault with default headers values
func NewCatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault(code int) *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault {
	return &CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault{
		_statusCode: code,
	}
}

/*CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product custom option repository v1 delete by identifier delete default response
func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}/options/{optionId}][%d] catalogProductCustomOptionRepositoryV1DeleteByIdentifierDelete default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductCustomOptionRepositoryV1DeleteByIdentifierDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
