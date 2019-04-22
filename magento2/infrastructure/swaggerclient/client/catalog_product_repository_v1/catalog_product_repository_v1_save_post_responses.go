// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

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

// CatalogProductRepositoryV1SavePostReader is a Reader for the CatalogProductRepositoryV1SavePost structure.
type CatalogProductRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductRepositoryV1SavePostOK creates a CatalogProductRepositoryV1SavePostOK with default headers values
func NewCatalogProductRepositoryV1SavePostOK() *CatalogProductRepositoryV1SavePostOK {
	return &CatalogProductRepositoryV1SavePostOK{}
}

/*CatalogProductRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type CatalogProductRepositoryV1SavePostOK struct {
	Payload *models.CatalogDataProductInterface
}

func (o *CatalogProductRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/products][%d] catalogProductRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *CatalogProductRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1SavePostBadRequest creates a CatalogProductRepositoryV1SavePostBadRequest with default headers values
func NewCatalogProductRepositoryV1SavePostBadRequest() *CatalogProductRepositoryV1SavePostBadRequest {
	return &CatalogProductRepositoryV1SavePostBadRequest{}
}

/*CatalogProductRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/products][%d] catalogProductRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1SavePostUnauthorized creates a CatalogProductRepositoryV1SavePostUnauthorized with default headers values
func NewCatalogProductRepositoryV1SavePostUnauthorized() *CatalogProductRepositoryV1SavePostUnauthorized {
	return &CatalogProductRepositoryV1SavePostUnauthorized{}
}

/*CatalogProductRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/products][%d] catalogProductRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1SavePostDefault creates a CatalogProductRepositoryV1SavePostDefault with default headers values
func NewCatalogProductRepositoryV1SavePostDefault(code int) *CatalogProductRepositoryV1SavePostDefault {
	return &CatalogProductRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*CatalogProductRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product repository v1 save post default response
func (o *CatalogProductRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/products][%d] catalogProductRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogProductRepositoryV1SavePostBody catalog product repository v1 save post body
swagger:model CatalogProductRepositoryV1SavePostBody
*/
type CatalogProductRepositoryV1SavePostBody struct {

	// product
	// Required: true
	Product *models.CatalogDataProductInterface `json:"product"`

	// save options
	SaveOptions bool `json:"saveOptions,omitempty"`
}

// Validate validates this catalog product repository v1 save post body
func (o *CatalogProductRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogProductRepositoryV1SavePostBody) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("catalogProductRepositoryV1SavePostBody"+"."+"product", "body", o.Product); err != nil {
		return err
	}

	if o.Product != nil {
		if err := o.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogProductRepositoryV1SavePostBody" + "." + "product")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogProductRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogProductRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res CatalogProductRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}