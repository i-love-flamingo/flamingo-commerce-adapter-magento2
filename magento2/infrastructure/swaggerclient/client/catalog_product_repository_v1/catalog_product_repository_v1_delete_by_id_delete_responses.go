// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductRepositoryV1DeleteByIDDeleteReader is a Reader for the CatalogProductRepositoryV1DeleteByIDDelete structure.
type CatalogProductRepositoryV1DeleteByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductRepositoryV1DeleteByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductRepositoryV1DeleteByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductRepositoryV1DeleteByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductRepositoryV1DeleteByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductRepositoryV1DeleteByIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductRepositoryV1DeleteByIDDeleteOK creates a CatalogProductRepositoryV1DeleteByIDDeleteOK with default headers values
func NewCatalogProductRepositoryV1DeleteByIDDeleteOK() *CatalogProductRepositoryV1DeleteByIDDeleteOK {
	return &CatalogProductRepositoryV1DeleteByIDDeleteOK{}
}

/*CatalogProductRepositoryV1DeleteByIDDeleteOK handles this case with default header values.

200 Success.
*/
type CatalogProductRepositoryV1DeleteByIDDeleteOK struct {
	Payload bool
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}][%d] catalogProductRepositoryV1DeleteByIdDeleteOK  %+v", 200, o.Payload)
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1DeleteByIDDeleteBadRequest creates a CatalogProductRepositoryV1DeleteByIDDeleteBadRequest with default headers values
func NewCatalogProductRepositoryV1DeleteByIDDeleteBadRequest() *CatalogProductRepositoryV1DeleteByIDDeleteBadRequest {
	return &CatalogProductRepositoryV1DeleteByIDDeleteBadRequest{}
}

/*CatalogProductRepositoryV1DeleteByIDDeleteBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductRepositoryV1DeleteByIDDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}][%d] catalogProductRepositoryV1DeleteByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1DeleteByIDDeleteUnauthorized creates a CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized with default headers values
func NewCatalogProductRepositoryV1DeleteByIDDeleteUnauthorized() *CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized {
	return &CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized{}
}

/*CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}][%d] catalogProductRepositoryV1DeleteByIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1DeleteByIDDeleteDefault creates a CatalogProductRepositoryV1DeleteByIDDeleteDefault with default headers values
func NewCatalogProductRepositoryV1DeleteByIDDeleteDefault(code int) *CatalogProductRepositoryV1DeleteByIDDeleteDefault {
	return &CatalogProductRepositoryV1DeleteByIDDeleteDefault{
		_statusCode: code,
	}
}

/*CatalogProductRepositoryV1DeleteByIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductRepositoryV1DeleteByIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product repository v1 delete by Id delete default response
func (o *CatalogProductRepositoryV1DeleteByIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/products/{sku}][%d] catalogProductRepositoryV1DeleteByIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductRepositoryV1DeleteByIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
