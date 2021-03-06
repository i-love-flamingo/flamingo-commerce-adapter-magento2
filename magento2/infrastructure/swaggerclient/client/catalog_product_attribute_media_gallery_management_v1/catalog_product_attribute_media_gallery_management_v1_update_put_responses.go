// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_media_gallery_management_v1

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

// CatalogProductAttributeMediaGalleryManagementV1UpdatePutReader is a Reader for the CatalogProductAttributeMediaGalleryManagementV1UpdatePut structure.
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutOK creates a CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK with default headers values
func NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutOK() *CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK {
	return &CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK{}
}

/*CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK handles this case with default header values.

200 Success.
*/
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK struct {
	Payload bool
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/products/{sku}/media/{entryId}][%d] catalogProductAttributeMediaGalleryManagementV1UpdatePutOK  %+v", 200, o.Payload)
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest creates a CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest with default headers values
func NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest() *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest {
	return &CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest{}
}

/*CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/products/{sku}/media/{entryId}][%d] catalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized creates a CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized with default headers values
func NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized() *CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized {
	return &CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized{}
}

/*CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/products/{sku}/media/{entryId}][%d] catalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault creates a CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault with default headers values
func NewCatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault(code int) *CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault {
	return &CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault{
		_statusCode: code,
	}
}

/*CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault handles this case with default header values.

Unexpected error
*/
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product attribute media gallery management v1 update put default response
func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/products/{sku}/media/{entryId}][%d] catalogProductAttributeMediaGalleryManagementV1UpdatePut default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody catalog product attribute media gallery management v1 update put body
swagger:model CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody
*/
type CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody struct {

	// entry
	// Required: true
	Entry *models.CatalogDataProductAttributeMediaGalleryEntryInterface `json:"entry"`
}

// Validate validates this catalog product attribute media gallery management v1 update put body
func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody) validateEntry(formats strfmt.Registry) error {

	if err := validate.Required("catalogProductAttributeMediaGalleryManagementV1UpdatePutBody"+"."+"entry", "body", o.Entry); err != nil {
		return err
	}

	if o.Entry != nil {
		if err := o.Entry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogProductAttributeMediaGalleryManagementV1UpdatePutBody" + "." + "entry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody) UnmarshalBinary(b []byte) error {
	var res CatalogProductAttributeMediaGalleryManagementV1UpdatePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
