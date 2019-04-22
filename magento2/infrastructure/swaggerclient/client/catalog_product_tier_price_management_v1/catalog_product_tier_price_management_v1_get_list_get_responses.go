// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_tier_price_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogProductTierPriceManagementV1GetListGetReader is a Reader for the CatalogProductTierPriceManagementV1GetListGet structure.
type CatalogProductTierPriceManagementV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductTierPriceManagementV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductTierPriceManagementV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductTierPriceManagementV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductTierPriceManagementV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductTierPriceManagementV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductTierPriceManagementV1GetListGetOK creates a CatalogProductTierPriceManagementV1GetListGetOK with default headers values
func NewCatalogProductTierPriceManagementV1GetListGetOK() *CatalogProductTierPriceManagementV1GetListGetOK {
	return &CatalogProductTierPriceManagementV1GetListGetOK{}
}

/*CatalogProductTierPriceManagementV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CatalogProductTierPriceManagementV1GetListGetOK struct {
	Payload []*models.CatalogDataProductTierPriceInterface
}

func (o *CatalogProductTierPriceManagementV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/group-prices/{customerGroupId}/tiers][%d] catalogProductTierPriceManagementV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CatalogProductTierPriceManagementV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductTierPriceManagementV1GetListGetBadRequest creates a CatalogProductTierPriceManagementV1GetListGetBadRequest with default headers values
func NewCatalogProductTierPriceManagementV1GetListGetBadRequest() *CatalogProductTierPriceManagementV1GetListGetBadRequest {
	return &CatalogProductTierPriceManagementV1GetListGetBadRequest{}
}

/*CatalogProductTierPriceManagementV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductTierPriceManagementV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductTierPriceManagementV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/group-prices/{customerGroupId}/tiers][%d] catalogProductTierPriceManagementV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductTierPriceManagementV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductTierPriceManagementV1GetListGetUnauthorized creates a CatalogProductTierPriceManagementV1GetListGetUnauthorized with default headers values
func NewCatalogProductTierPriceManagementV1GetListGetUnauthorized() *CatalogProductTierPriceManagementV1GetListGetUnauthorized {
	return &CatalogProductTierPriceManagementV1GetListGetUnauthorized{}
}

/*CatalogProductTierPriceManagementV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductTierPriceManagementV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductTierPriceManagementV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/group-prices/{customerGroupId}/tiers][%d] catalogProductTierPriceManagementV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductTierPriceManagementV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductTierPriceManagementV1GetListGetDefault creates a CatalogProductTierPriceManagementV1GetListGetDefault with default headers values
func NewCatalogProductTierPriceManagementV1GetListGetDefault(code int) *CatalogProductTierPriceManagementV1GetListGetDefault {
	return &CatalogProductTierPriceManagementV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CatalogProductTierPriceManagementV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductTierPriceManagementV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product tier price management v1 get list get default response
func (o *CatalogProductTierPriceManagementV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductTierPriceManagementV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/group-prices/{customerGroupId}/tiers][%d] catalogProductTierPriceManagementV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductTierPriceManagementV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}