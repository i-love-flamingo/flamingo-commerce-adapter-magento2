// Code generated by go-swagger; DO NOT EDIT.

package catalog_special_price_storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CatalogSpecialPriceStorageV1DeletePostReader is a Reader for the CatalogSpecialPriceStorageV1DeletePost structure.
type CatalogSpecialPriceStorageV1DeletePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogSpecialPriceStorageV1DeletePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogSpecialPriceStorageV1DeletePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogSpecialPriceStorageV1DeletePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogSpecialPriceStorageV1DeletePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogSpecialPriceStorageV1DeletePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogSpecialPriceStorageV1DeletePostOK creates a CatalogSpecialPriceStorageV1DeletePostOK with default headers values
func NewCatalogSpecialPriceStorageV1DeletePostOK() *CatalogSpecialPriceStorageV1DeletePostOK {
	return &CatalogSpecialPriceStorageV1DeletePostOK{}
}

/*CatalogSpecialPriceStorageV1DeletePostOK handles this case with default header values.

200 Success.
*/
type CatalogSpecialPriceStorageV1DeletePostOK struct {
	Payload []*models.CatalogDataPriceUpdateResultInterface
}

func (o *CatalogSpecialPriceStorageV1DeletePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/products/special-price-delete][%d] catalogSpecialPriceStorageV1DeletePostOK  %+v", 200, o.Payload)
}

func (o *CatalogSpecialPriceStorageV1DeletePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogSpecialPriceStorageV1DeletePostBadRequest creates a CatalogSpecialPriceStorageV1DeletePostBadRequest with default headers values
func NewCatalogSpecialPriceStorageV1DeletePostBadRequest() *CatalogSpecialPriceStorageV1DeletePostBadRequest {
	return &CatalogSpecialPriceStorageV1DeletePostBadRequest{}
}

/*CatalogSpecialPriceStorageV1DeletePostBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogSpecialPriceStorageV1DeletePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogSpecialPriceStorageV1DeletePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/products/special-price-delete][%d] catalogSpecialPriceStorageV1DeletePostBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogSpecialPriceStorageV1DeletePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogSpecialPriceStorageV1DeletePostUnauthorized creates a CatalogSpecialPriceStorageV1DeletePostUnauthorized with default headers values
func NewCatalogSpecialPriceStorageV1DeletePostUnauthorized() *CatalogSpecialPriceStorageV1DeletePostUnauthorized {
	return &CatalogSpecialPriceStorageV1DeletePostUnauthorized{}
}

/*CatalogSpecialPriceStorageV1DeletePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogSpecialPriceStorageV1DeletePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogSpecialPriceStorageV1DeletePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/products/special-price-delete][%d] catalogSpecialPriceStorageV1DeletePostUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogSpecialPriceStorageV1DeletePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogSpecialPriceStorageV1DeletePostDefault creates a CatalogSpecialPriceStorageV1DeletePostDefault with default headers values
func NewCatalogSpecialPriceStorageV1DeletePostDefault(code int) *CatalogSpecialPriceStorageV1DeletePostDefault {
	return &CatalogSpecialPriceStorageV1DeletePostDefault{
		_statusCode: code,
	}
}

/*CatalogSpecialPriceStorageV1DeletePostDefault handles this case with default header values.

Unexpected error
*/
type CatalogSpecialPriceStorageV1DeletePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog special price storage v1 delete post default response
func (o *CatalogSpecialPriceStorageV1DeletePostDefault) Code() int {
	return o._statusCode
}

func (o *CatalogSpecialPriceStorageV1DeletePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/products/special-price-delete][%d] catalogSpecialPriceStorageV1DeletePost default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogSpecialPriceStorageV1DeletePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogSpecialPriceStorageV1DeletePostBody catalog special price storage v1 delete post body
swagger:model CatalogSpecialPriceStorageV1DeletePostBody
*/
type CatalogSpecialPriceStorageV1DeletePostBody struct {

	// prices
	// Required: true
	Prices []*models.CatalogDataSpecialPriceInterface `json:"prices"`
}

// Validate validates this catalog special price storage v1 delete post body
func (o *CatalogSpecialPriceStorageV1DeletePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogSpecialPriceStorageV1DeletePostBody) validatePrices(formats strfmt.Registry) error {

	if err := validate.Required("catalogSpecialPriceStorageV1DeletePostBody"+"."+"prices", "body", o.Prices); err != nil {
		return err
	}

	for i := 0; i < len(o.Prices); i++ {
		if swag.IsZero(o.Prices[i]) { // not required
			continue
		}

		if o.Prices[i] != nil {
			if err := o.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogSpecialPriceStorageV1DeletePostBody" + "." + "prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogSpecialPriceStorageV1DeletePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogSpecialPriceStorageV1DeletePostBody) UnmarshalBinary(b []byte) error {
	var res CatalogSpecialPriceStorageV1DeletePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
