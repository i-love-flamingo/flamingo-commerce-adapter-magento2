// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_repository_v1

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

// CatalogCategoryRepositoryV1SavePutReader is a Reader for the CatalogCategoryRepositoryV1SavePut structure.
type CatalogCategoryRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogCategoryRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogCategoryRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogCategoryRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogCategoryRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryRepositoryV1SavePutOK creates a CatalogCategoryRepositoryV1SavePutOK with default headers values
func NewCatalogCategoryRepositoryV1SavePutOK() *CatalogCategoryRepositoryV1SavePutOK {
	return &CatalogCategoryRepositoryV1SavePutOK{}
}

/*CatalogCategoryRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CatalogCategoryRepositoryV1SavePutOK struct {
	Payload *models.CatalogDataCategoryInterface
}

func (o *CatalogCategoryRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/categories/{id}][%d] catalogCategoryRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CatalogCategoryRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataCategoryInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryRepositoryV1SavePutBadRequest creates a CatalogCategoryRepositoryV1SavePutBadRequest with default headers values
func NewCatalogCategoryRepositoryV1SavePutBadRequest() *CatalogCategoryRepositoryV1SavePutBadRequest {
	return &CatalogCategoryRepositoryV1SavePutBadRequest{}
}

/*CatalogCategoryRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogCategoryRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/categories/{id}][%d] catalogCategoryRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCategoryRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryRepositoryV1SavePutUnauthorized creates a CatalogCategoryRepositoryV1SavePutUnauthorized with default headers values
func NewCatalogCategoryRepositoryV1SavePutUnauthorized() *CatalogCategoryRepositoryV1SavePutUnauthorized {
	return &CatalogCategoryRepositoryV1SavePutUnauthorized{}
}

/*CatalogCategoryRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogCategoryRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/categories/{id}][%d] catalogCategoryRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCategoryRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryRepositoryV1SavePutDefault creates a CatalogCategoryRepositoryV1SavePutDefault with default headers values
func NewCatalogCategoryRepositoryV1SavePutDefault(code int) *CatalogCategoryRepositoryV1SavePutDefault {
	return &CatalogCategoryRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CatalogCategoryRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CatalogCategoryRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category repository v1 save put default response
func (o *CatalogCategoryRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/categories/{id}][%d] catalogCategoryRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogCategoryRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogCategoryRepositoryV1SavePutBody catalog category repository v1 save put body
swagger:model CatalogCategoryRepositoryV1SavePutBody
*/
type CatalogCategoryRepositoryV1SavePutBody struct {

	// category
	// Required: true
	Category *models.CatalogDataCategoryInterface `json:"category"`
}

// Validate validates this catalog category repository v1 save put body
func (o *CatalogCategoryRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogCategoryRepositoryV1SavePutBody) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("catalogCategoryRepositoryV1SavePutBody"+"."+"category", "body", o.Category); err != nil {
		return err
	}

	if o.Category != nil {
		if err := o.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogCategoryRepositoryV1SavePutBody" + "." + "category")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogCategoryRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogCategoryRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CatalogCategoryRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
