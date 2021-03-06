// Code generated by go-swagger; DO NOT EDIT.

package catalog_attribute_set_repository_v1

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

// CatalogAttributeSetRepositoryV1SavePutReader is a Reader for the CatalogAttributeSetRepositoryV1SavePut structure.
type CatalogAttributeSetRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogAttributeSetRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogAttributeSetRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogAttributeSetRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogAttributeSetRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCatalogAttributeSetRepositoryV1SavePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogAttributeSetRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogAttributeSetRepositoryV1SavePutOK creates a CatalogAttributeSetRepositoryV1SavePutOK with default headers values
func NewCatalogAttributeSetRepositoryV1SavePutOK() *CatalogAttributeSetRepositoryV1SavePutOK {
	return &CatalogAttributeSetRepositoryV1SavePutOK{}
}

/*CatalogAttributeSetRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CatalogAttributeSetRepositoryV1SavePutOK struct {
	Payload *models.EavDataAttributeSetInterface
}

func (o *CatalogAttributeSetRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/products/attribute-sets/{attributeSetId}][%d] catalogAttributeSetRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CatalogAttributeSetRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EavDataAttributeSetInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAttributeSetRepositoryV1SavePutBadRequest creates a CatalogAttributeSetRepositoryV1SavePutBadRequest with default headers values
func NewCatalogAttributeSetRepositoryV1SavePutBadRequest() *CatalogAttributeSetRepositoryV1SavePutBadRequest {
	return &CatalogAttributeSetRepositoryV1SavePutBadRequest{}
}

/*CatalogAttributeSetRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogAttributeSetRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogAttributeSetRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/products/attribute-sets/{attributeSetId}][%d] catalogAttributeSetRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogAttributeSetRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAttributeSetRepositoryV1SavePutUnauthorized creates a CatalogAttributeSetRepositoryV1SavePutUnauthorized with default headers values
func NewCatalogAttributeSetRepositoryV1SavePutUnauthorized() *CatalogAttributeSetRepositoryV1SavePutUnauthorized {
	return &CatalogAttributeSetRepositoryV1SavePutUnauthorized{}
}

/*CatalogAttributeSetRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogAttributeSetRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogAttributeSetRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/products/attribute-sets/{attributeSetId}][%d] catalogAttributeSetRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogAttributeSetRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAttributeSetRepositoryV1SavePutInternalServerError creates a CatalogAttributeSetRepositoryV1SavePutInternalServerError with default headers values
func NewCatalogAttributeSetRepositoryV1SavePutInternalServerError() *CatalogAttributeSetRepositoryV1SavePutInternalServerError {
	return &CatalogAttributeSetRepositoryV1SavePutInternalServerError{}
}

/*CatalogAttributeSetRepositoryV1SavePutInternalServerError handles this case with default header values.

Internal Server error
*/
type CatalogAttributeSetRepositoryV1SavePutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CatalogAttributeSetRepositoryV1SavePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/products/attribute-sets/{attributeSetId}][%d] catalogAttributeSetRepositoryV1SavePutInternalServerError  %+v", 500, o.Payload)
}

func (o *CatalogAttributeSetRepositoryV1SavePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAttributeSetRepositoryV1SavePutDefault creates a CatalogAttributeSetRepositoryV1SavePutDefault with default headers values
func NewCatalogAttributeSetRepositoryV1SavePutDefault(code int) *CatalogAttributeSetRepositoryV1SavePutDefault {
	return &CatalogAttributeSetRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CatalogAttributeSetRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CatalogAttributeSetRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog attribute set repository v1 save put default response
func (o *CatalogAttributeSetRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CatalogAttributeSetRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/products/attribute-sets/{attributeSetId}][%d] catalogAttributeSetRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogAttributeSetRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogAttributeSetRepositoryV1SavePutBody catalog attribute set repository v1 save put body
swagger:model CatalogAttributeSetRepositoryV1SavePutBody
*/
type CatalogAttributeSetRepositoryV1SavePutBody struct {

	// attribute set
	// Required: true
	AttributeSet *models.EavDataAttributeSetInterface `json:"attributeSet"`
}

// Validate validates this catalog attribute set repository v1 save put body
func (o *CatalogAttributeSetRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributeSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogAttributeSetRepositoryV1SavePutBody) validateAttributeSet(formats strfmt.Registry) error {

	if err := validate.Required("catalogAttributeSetRepositoryV1SavePutBody"+"."+"attributeSet", "body", o.AttributeSet); err != nil {
		return err
	}

	if o.AttributeSet != nil {
		if err := o.AttributeSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogAttributeSetRepositoryV1SavePutBody" + "." + "attributeSet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogAttributeSetRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogAttributeSetRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CatalogAttributeSetRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
