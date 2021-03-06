// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_website_link_repository_v1

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

// CatalogProductWebsiteLinkRepositoryV1SavePostReader is a Reader for the CatalogProductWebsiteLinkRepositoryV1SavePost structure.
type CatalogProductWebsiteLinkRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductWebsiteLinkRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductWebsiteLinkRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductWebsiteLinkRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductWebsiteLinkRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductWebsiteLinkRepositoryV1SavePostOK creates a CatalogProductWebsiteLinkRepositoryV1SavePostOK with default headers values
func NewCatalogProductWebsiteLinkRepositoryV1SavePostOK() *CatalogProductWebsiteLinkRepositoryV1SavePostOK {
	return &CatalogProductWebsiteLinkRepositoryV1SavePostOK{}
}

/*CatalogProductWebsiteLinkRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type CatalogProductWebsiteLinkRepositoryV1SavePostOK struct {
	Payload bool
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/products/{sku}/websites][%d] catalogProductWebsiteLinkRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductWebsiteLinkRepositoryV1SavePostBadRequest creates a CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest with default headers values
func NewCatalogProductWebsiteLinkRepositoryV1SavePostBadRequest() *CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest {
	return &CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest{}
}

/*CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/products/{sku}/websites][%d] catalogProductWebsiteLinkRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized creates a CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized with default headers values
func NewCatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized() *CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized {
	return &CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized{}
}

/*CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/products/{sku}/websites][%d] catalogProductWebsiteLinkRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductWebsiteLinkRepositoryV1SavePostDefault creates a CatalogProductWebsiteLinkRepositoryV1SavePostDefault with default headers values
func NewCatalogProductWebsiteLinkRepositoryV1SavePostDefault(code int) *CatalogProductWebsiteLinkRepositoryV1SavePostDefault {
	return &CatalogProductWebsiteLinkRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*CatalogProductWebsiteLinkRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductWebsiteLinkRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product website link repository v1 save post default response
func (o *CatalogProductWebsiteLinkRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/products/{sku}/websites][%d] catalogProductWebsiteLinkRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogProductWebsiteLinkRepositoryV1SavePostBody catalog product website link repository v1 save post body
swagger:model CatalogProductWebsiteLinkRepositoryV1SavePostBody
*/
type CatalogProductWebsiteLinkRepositoryV1SavePostBody struct {

	// product website link
	// Required: true
	ProductWebsiteLink *models.CatalogDataProductWebsiteLinkInterface `json:"productWebsiteLink"`
}

// Validate validates this catalog product website link repository v1 save post body
func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProductWebsiteLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBody) validateProductWebsiteLink(formats strfmt.Registry) error {

	if err := validate.Required("catalogProductWebsiteLinkRepositoryV1SavePostBody"+"."+"productWebsiteLink", "body", o.ProductWebsiteLink); err != nil {
		return err
	}

	if o.ProductWebsiteLink != nil {
		if err := o.ProductWebsiteLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogProductWebsiteLinkRepositoryV1SavePostBody" + "." + "productWebsiteLink")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogProductWebsiteLinkRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res CatalogProductWebsiteLinkRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
