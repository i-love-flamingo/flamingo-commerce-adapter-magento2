// Code generated by go-swagger; DO NOT EDIT.

package catalog_tier_price_storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogTierPriceStorageV1GetPostReader is a Reader for the CatalogTierPriceStorageV1GetPost structure.
type CatalogTierPriceStorageV1GetPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogTierPriceStorageV1GetPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogTierPriceStorageV1GetPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogTierPriceStorageV1GetPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogTierPriceStorageV1GetPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogTierPriceStorageV1GetPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogTierPriceStorageV1GetPostOK creates a CatalogTierPriceStorageV1GetPostOK with default headers values
func NewCatalogTierPriceStorageV1GetPostOK() *CatalogTierPriceStorageV1GetPostOK {
	return &CatalogTierPriceStorageV1GetPostOK{}
}

/*CatalogTierPriceStorageV1GetPostOK handles this case with default header values.

200 Success.
*/
type CatalogTierPriceStorageV1GetPostOK struct {
	Payload []*models.CatalogDataTierPriceInterface
}

func (o *CatalogTierPriceStorageV1GetPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/products/tier-prices-information][%d] catalogTierPriceStorageV1GetPostOK  %+v", 200, o.Payload)
}

func (o *CatalogTierPriceStorageV1GetPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogTierPriceStorageV1GetPostBadRequest creates a CatalogTierPriceStorageV1GetPostBadRequest with default headers values
func NewCatalogTierPriceStorageV1GetPostBadRequest() *CatalogTierPriceStorageV1GetPostBadRequest {
	return &CatalogTierPriceStorageV1GetPostBadRequest{}
}

/*CatalogTierPriceStorageV1GetPostBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogTierPriceStorageV1GetPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogTierPriceStorageV1GetPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/products/tier-prices-information][%d] catalogTierPriceStorageV1GetPostBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogTierPriceStorageV1GetPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogTierPriceStorageV1GetPostUnauthorized creates a CatalogTierPriceStorageV1GetPostUnauthorized with default headers values
func NewCatalogTierPriceStorageV1GetPostUnauthorized() *CatalogTierPriceStorageV1GetPostUnauthorized {
	return &CatalogTierPriceStorageV1GetPostUnauthorized{}
}

/*CatalogTierPriceStorageV1GetPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogTierPriceStorageV1GetPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogTierPriceStorageV1GetPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/products/tier-prices-information][%d] catalogTierPriceStorageV1GetPostUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogTierPriceStorageV1GetPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogTierPriceStorageV1GetPostDefault creates a CatalogTierPriceStorageV1GetPostDefault with default headers values
func NewCatalogTierPriceStorageV1GetPostDefault(code int) *CatalogTierPriceStorageV1GetPostDefault {
	return &CatalogTierPriceStorageV1GetPostDefault{
		_statusCode: code,
	}
}

/*CatalogTierPriceStorageV1GetPostDefault handles this case with default header values.

Unexpected error
*/
type CatalogTierPriceStorageV1GetPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog tier price storage v1 get post default response
func (o *CatalogTierPriceStorageV1GetPostDefault) Code() int {
	return o._statusCode
}

func (o *CatalogTierPriceStorageV1GetPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/products/tier-prices-information][%d] catalogTierPriceStorageV1GetPost default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogTierPriceStorageV1GetPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogTierPriceStorageV1GetPostBody catalog tier price storage v1 get post body
swagger:model CatalogTierPriceStorageV1GetPostBody
*/
type CatalogTierPriceStorageV1GetPostBody struct {

	// skus
	// Required: true
	Skus []string `json:"skus"`
}

// Validate validates this catalog tier price storage v1 get post body
func (o *CatalogTierPriceStorageV1GetPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSkus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogTierPriceStorageV1GetPostBody) validateSkus(formats strfmt.Registry) error {

	if err := validate.Required("catalogTierPriceStorageV1GetPostBody"+"."+"skus", "body", o.Skus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogTierPriceStorageV1GetPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogTierPriceStorageV1GetPostBody) UnmarshalBinary(b []byte) error {
	var res CatalogTierPriceStorageV1GetPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}