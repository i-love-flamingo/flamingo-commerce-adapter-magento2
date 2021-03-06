// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_option_repository_v1

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

// ConfigurableProductOptionRepositoryV1SavePostReader is a Reader for the ConfigurableProductOptionRepositoryV1SavePost structure.
type ConfigurableProductOptionRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurableProductOptionRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConfigurableProductOptionRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConfigurableProductOptionRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConfigurableProductOptionRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConfigurableProductOptionRepositoryV1SavePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewConfigurableProductOptionRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurableProductOptionRepositoryV1SavePostOK creates a ConfigurableProductOptionRepositoryV1SavePostOK with default headers values
func NewConfigurableProductOptionRepositoryV1SavePostOK() *ConfigurableProductOptionRepositoryV1SavePostOK {
	return &ConfigurableProductOptionRepositoryV1SavePostOK{}
}

/*ConfigurableProductOptionRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type ConfigurableProductOptionRepositoryV1SavePostOK struct {
	Payload int64
}

func (o *ConfigurableProductOptionRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/configurable-products/{sku}/options][%d] configurableProductOptionRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *ConfigurableProductOptionRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1SavePostBadRequest creates a ConfigurableProductOptionRepositoryV1SavePostBadRequest with default headers values
func NewConfigurableProductOptionRepositoryV1SavePostBadRequest() *ConfigurableProductOptionRepositoryV1SavePostBadRequest {
	return &ConfigurableProductOptionRepositoryV1SavePostBadRequest{}
}

/*ConfigurableProductOptionRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type ConfigurableProductOptionRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConfigurableProductOptionRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/configurable-products/{sku}/options][%d] configurableProductOptionRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigurableProductOptionRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1SavePostUnauthorized creates a ConfigurableProductOptionRepositoryV1SavePostUnauthorized with default headers values
func NewConfigurableProductOptionRepositoryV1SavePostUnauthorized() *ConfigurableProductOptionRepositoryV1SavePostUnauthorized {
	return &ConfigurableProductOptionRepositoryV1SavePostUnauthorized{}
}

/*ConfigurableProductOptionRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type ConfigurableProductOptionRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *ConfigurableProductOptionRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/configurable-products/{sku}/options][%d] configurableProductOptionRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *ConfigurableProductOptionRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1SavePostInternalServerError creates a ConfigurableProductOptionRepositoryV1SavePostInternalServerError with default headers values
func NewConfigurableProductOptionRepositoryV1SavePostInternalServerError() *ConfigurableProductOptionRepositoryV1SavePostInternalServerError {
	return &ConfigurableProductOptionRepositoryV1SavePostInternalServerError{}
}

/*ConfigurableProductOptionRepositoryV1SavePostInternalServerError handles this case with default header values.

Internal Server error
*/
type ConfigurableProductOptionRepositoryV1SavePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConfigurableProductOptionRepositoryV1SavePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/configurable-products/{sku}/options][%d] configurableProductOptionRepositoryV1SavePostInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigurableProductOptionRepositoryV1SavePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1SavePostDefault creates a ConfigurableProductOptionRepositoryV1SavePostDefault with default headers values
func NewConfigurableProductOptionRepositoryV1SavePostDefault(code int) *ConfigurableProductOptionRepositoryV1SavePostDefault {
	return &ConfigurableProductOptionRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*ConfigurableProductOptionRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type ConfigurableProductOptionRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the configurable product option repository v1 save post default response
func (o *ConfigurableProductOptionRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurableProductOptionRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/configurable-products/{sku}/options][%d] configurableProductOptionRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *ConfigurableProductOptionRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConfigurableProductOptionRepositoryV1SavePostBody configurable product option repository v1 save post body
swagger:model ConfigurableProductOptionRepositoryV1SavePostBody
*/
type ConfigurableProductOptionRepositoryV1SavePostBody struct {

	// option
	// Required: true
	Option *models.ConfigurableProductDataOptionInterface `json:"option"`
}

// Validate validates this configurable product option repository v1 save post body
func (o *ConfigurableProductOptionRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConfigurableProductOptionRepositoryV1SavePostBody) validateOption(formats strfmt.Registry) error {

	if err := validate.Required("configurableProductOptionRepositoryV1SavePostBody"+"."+"option", "body", o.Option); err != nil {
		return err
	}

	if o.Option != nil {
		if err := o.Option.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurableProductOptionRepositoryV1SavePostBody" + "." + "option")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConfigurableProductOptionRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigurableProductOptionRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res ConfigurableProductOptionRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
