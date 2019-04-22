// Code generated by go-swagger; DO NOT EDIT.

package eav_attribute_set_management_v1

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

// EavAttributeSetManagementV1CreatePostReader is a Reader for the EavAttributeSetManagementV1CreatePost structure.
type EavAttributeSetManagementV1CreatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EavAttributeSetManagementV1CreatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEavAttributeSetManagementV1CreatePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEavAttributeSetManagementV1CreatePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEavAttributeSetManagementV1CreatePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEavAttributeSetManagementV1CreatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEavAttributeSetManagementV1CreatePostOK creates a EavAttributeSetManagementV1CreatePostOK with default headers values
func NewEavAttributeSetManagementV1CreatePostOK() *EavAttributeSetManagementV1CreatePostOK {
	return &EavAttributeSetManagementV1CreatePostOK{}
}

/*EavAttributeSetManagementV1CreatePostOK handles this case with default header values.

200 Success.
*/
type EavAttributeSetManagementV1CreatePostOK struct {
	Payload *models.EavDataAttributeSetInterface
}

func (o *EavAttributeSetManagementV1CreatePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/eav/attribute-sets][%d] eavAttributeSetManagementV1CreatePostOK  %+v", 200, o.Payload)
}

func (o *EavAttributeSetManagementV1CreatePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EavDataAttributeSetInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEavAttributeSetManagementV1CreatePostBadRequest creates a EavAttributeSetManagementV1CreatePostBadRequest with default headers values
func NewEavAttributeSetManagementV1CreatePostBadRequest() *EavAttributeSetManagementV1CreatePostBadRequest {
	return &EavAttributeSetManagementV1CreatePostBadRequest{}
}

/*EavAttributeSetManagementV1CreatePostBadRequest handles this case with default header values.

400 Bad Request
*/
type EavAttributeSetManagementV1CreatePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *EavAttributeSetManagementV1CreatePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/eav/attribute-sets][%d] eavAttributeSetManagementV1CreatePostBadRequest  %+v", 400, o.Payload)
}

func (o *EavAttributeSetManagementV1CreatePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEavAttributeSetManagementV1CreatePostUnauthorized creates a EavAttributeSetManagementV1CreatePostUnauthorized with default headers values
func NewEavAttributeSetManagementV1CreatePostUnauthorized() *EavAttributeSetManagementV1CreatePostUnauthorized {
	return &EavAttributeSetManagementV1CreatePostUnauthorized{}
}

/*EavAttributeSetManagementV1CreatePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type EavAttributeSetManagementV1CreatePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *EavAttributeSetManagementV1CreatePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/eav/attribute-sets][%d] eavAttributeSetManagementV1CreatePostUnauthorized  %+v", 401, o.Payload)
}

func (o *EavAttributeSetManagementV1CreatePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEavAttributeSetManagementV1CreatePostDefault creates a EavAttributeSetManagementV1CreatePostDefault with default headers values
func NewEavAttributeSetManagementV1CreatePostDefault(code int) *EavAttributeSetManagementV1CreatePostDefault {
	return &EavAttributeSetManagementV1CreatePostDefault{
		_statusCode: code,
	}
}

/*EavAttributeSetManagementV1CreatePostDefault handles this case with default header values.

Unexpected error
*/
type EavAttributeSetManagementV1CreatePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the eav attribute set management v1 create post default response
func (o *EavAttributeSetManagementV1CreatePostDefault) Code() int {
	return o._statusCode
}

func (o *EavAttributeSetManagementV1CreatePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/eav/attribute-sets][%d] eavAttributeSetManagementV1CreatePost default  %+v", o._statusCode, o.Payload)
}

func (o *EavAttributeSetManagementV1CreatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EavAttributeSetManagementV1CreatePostBody eav attribute set management v1 create post body
swagger:model EavAttributeSetManagementV1CreatePostBody
*/
type EavAttributeSetManagementV1CreatePostBody struct {

	// attribute set
	// Required: true
	AttributeSet *models.EavDataAttributeSetInterface `json:"attributeSet"`

	// entity type code
	// Required: true
	EntityTypeCode *string `json:"entityTypeCode"`

	// skeleton Id
	// Required: true
	SkeletonID *int64 `json:"skeletonId"`
}

// Validate validates this eav attribute set management v1 create post body
func (o *EavAttributeSetManagementV1CreatePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributeSet(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEntityTypeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSkeletonID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EavAttributeSetManagementV1CreatePostBody) validateAttributeSet(formats strfmt.Registry) error {

	if err := validate.Required("eavAttributeSetManagementV1CreatePostBody"+"."+"attributeSet", "body", o.AttributeSet); err != nil {
		return err
	}

	if o.AttributeSet != nil {
		if err := o.AttributeSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eavAttributeSetManagementV1CreatePostBody" + "." + "attributeSet")
			}
			return err
		}
	}

	return nil
}

func (o *EavAttributeSetManagementV1CreatePostBody) validateEntityTypeCode(formats strfmt.Registry) error {

	if err := validate.Required("eavAttributeSetManagementV1CreatePostBody"+"."+"entityTypeCode", "body", o.EntityTypeCode); err != nil {
		return err
	}

	return nil
}

func (o *EavAttributeSetManagementV1CreatePostBody) validateSkeletonID(formats strfmt.Registry) error {

	if err := validate.Required("eavAttributeSetManagementV1CreatePostBody"+"."+"skeletonId", "body", o.SkeletonID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EavAttributeSetManagementV1CreatePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EavAttributeSetManagementV1CreatePostBody) UnmarshalBinary(b []byte) error {
	var res EavAttributeSetManagementV1CreatePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
