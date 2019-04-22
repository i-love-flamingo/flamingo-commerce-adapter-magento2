// Code generated by go-swagger; DO NOT EDIT.

package shared_catalog_shared_catalog_repository_v1

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

// SharedCatalogSharedCatalogRepositoryV1SavePostReader is a Reader for the SharedCatalogSharedCatalogRepositoryV1SavePost structure.
type SharedCatalogSharedCatalogRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SharedCatalogSharedCatalogRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSharedCatalogSharedCatalogRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSharedCatalogSharedCatalogRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSharedCatalogSharedCatalogRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSharedCatalogSharedCatalogRepositoryV1SavePostOK creates a SharedCatalogSharedCatalogRepositoryV1SavePostOK with default headers values
func NewSharedCatalogSharedCatalogRepositoryV1SavePostOK() *SharedCatalogSharedCatalogRepositoryV1SavePostOK {
	return &SharedCatalogSharedCatalogRepositoryV1SavePostOK{}
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostOK struct {
	Payload int64
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/sharedCatalog][%d] sharedCatalogSharedCatalogRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSharedCatalogSharedCatalogRepositoryV1SavePostBadRequest creates a SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest with default headers values
func NewSharedCatalogSharedCatalogRepositoryV1SavePostBadRequest() *SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest {
	return &SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest{}
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/sharedCatalog][%d] sharedCatalogSharedCatalogRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized creates a SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized with default headers values
func NewSharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized() *SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized {
	return &SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized{}
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/sharedCatalog][%d] sharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError creates a SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError with default headers values
func NewSharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError() *SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError {
	return &SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError{}
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError handles this case with default header values.

Internal Server error
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/sharedCatalog][%d] sharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError  %+v", 500, o.Payload)
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSharedCatalogSharedCatalogRepositoryV1SavePostDefault creates a SharedCatalogSharedCatalogRepositoryV1SavePostDefault with default headers values
func NewSharedCatalogSharedCatalogRepositoryV1SavePostDefault(code int) *SharedCatalogSharedCatalogRepositoryV1SavePostDefault {
	return &SharedCatalogSharedCatalogRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the shared catalog shared catalog repository v1 save post default response
func (o *SharedCatalogSharedCatalogRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/sharedCatalog][%d] sharedCatalogSharedCatalogRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SharedCatalogSharedCatalogRepositoryV1SavePostBody shared catalog shared catalog repository v1 save post body
swagger:model SharedCatalogSharedCatalogRepositoryV1SavePostBody
*/
type SharedCatalogSharedCatalogRepositoryV1SavePostBody struct {

	// shared catalog
	// Required: true
	SharedCatalog *models.SharedCatalogDataSharedCatalogInterface `json:"sharedCatalog"`
}

// Validate validates this shared catalog shared catalog repository v1 save post body
func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSharedCatalog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBody) validateSharedCatalog(formats strfmt.Registry) error {

	if err := validate.Required("sharedCatalogSharedCatalogRepositoryV1SavePostBody"+"."+"sharedCatalog", "body", o.SharedCatalog); err != nil {
		return err
	}

	if o.SharedCatalog != nil {
		if err := o.SharedCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedCatalogSharedCatalogRepositoryV1SavePostBody" + "." + "sharedCatalog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SharedCatalogSharedCatalogRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res SharedCatalogSharedCatalogRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
