// Code generated by go-swagger; DO NOT EDIT.

package cms_block_repository_v1

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

// CmsBlockRepositoryV1SavePutReader is a Reader for the CmsBlockRepositoryV1SavePut structure.
type CmsBlockRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CmsBlockRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCmsBlockRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCmsBlockRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCmsBlockRepositoryV1SavePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCmsBlockRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCmsBlockRepositoryV1SavePutOK creates a CmsBlockRepositoryV1SavePutOK with default headers values
func NewCmsBlockRepositoryV1SavePutOK() *CmsBlockRepositoryV1SavePutOK {
	return &CmsBlockRepositoryV1SavePutOK{}
}

/*CmsBlockRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CmsBlockRepositoryV1SavePutOK struct {
	Payload *models.CmsDataBlockInterface
}

func (o *CmsBlockRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/cmsBlock/{id}][%d] cmsBlockRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CmsBlockRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CmsDataBlockInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsBlockRepositoryV1SavePutUnauthorized creates a CmsBlockRepositoryV1SavePutUnauthorized with default headers values
func NewCmsBlockRepositoryV1SavePutUnauthorized() *CmsBlockRepositoryV1SavePutUnauthorized {
	return &CmsBlockRepositoryV1SavePutUnauthorized{}
}

/*CmsBlockRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CmsBlockRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CmsBlockRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/cmsBlock/{id}][%d] cmsBlockRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CmsBlockRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsBlockRepositoryV1SavePutInternalServerError creates a CmsBlockRepositoryV1SavePutInternalServerError with default headers values
func NewCmsBlockRepositoryV1SavePutInternalServerError() *CmsBlockRepositoryV1SavePutInternalServerError {
	return &CmsBlockRepositoryV1SavePutInternalServerError{}
}

/*CmsBlockRepositoryV1SavePutInternalServerError handles this case with default header values.

Internal Server error
*/
type CmsBlockRepositoryV1SavePutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CmsBlockRepositoryV1SavePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/cmsBlock/{id}][%d] cmsBlockRepositoryV1SavePutInternalServerError  %+v", 500, o.Payload)
}

func (o *CmsBlockRepositoryV1SavePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsBlockRepositoryV1SavePutDefault creates a CmsBlockRepositoryV1SavePutDefault with default headers values
func NewCmsBlockRepositoryV1SavePutDefault(code int) *CmsBlockRepositoryV1SavePutDefault {
	return &CmsBlockRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CmsBlockRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CmsBlockRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cms block repository v1 save put default response
func (o *CmsBlockRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CmsBlockRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/cmsBlock/{id}][%d] cmsBlockRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CmsBlockRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CmsBlockRepositoryV1SavePutBody cms block repository v1 save put body
swagger:model CmsBlockRepositoryV1SavePutBody
*/
type CmsBlockRepositoryV1SavePutBody struct {

	// block
	// Required: true
	Block *models.CmsDataBlockInterface `json:"block"`
}

// Validate validates this cms block repository v1 save put body
func (o *CmsBlockRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBlock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CmsBlockRepositoryV1SavePutBody) validateBlock(formats strfmt.Registry) error {

	if err := validate.Required("cmsBlockRepositoryV1SavePutBody"+"."+"block", "body", o.Block); err != nil {
		return err
	}

	if o.Block != nil {
		if err := o.Block.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cmsBlockRepositoryV1SavePutBody" + "." + "block")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CmsBlockRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CmsBlockRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CmsBlockRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
