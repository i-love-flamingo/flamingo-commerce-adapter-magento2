// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_rule_repository_v1

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

// SalesRuleRuleRepositoryV1SavePostReader is a Reader for the SalesRuleRuleRepositoryV1SavePost structure.
type SalesRuleRuleRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleRuleRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleRuleRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesRuleRuleRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesRuleRuleRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleRuleRepositoryV1SavePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleRuleRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleRuleRepositoryV1SavePostOK creates a SalesRuleRuleRepositoryV1SavePostOK with default headers values
func NewSalesRuleRuleRepositoryV1SavePostOK() *SalesRuleRuleRepositoryV1SavePostOK {
	return &SalesRuleRuleRepositoryV1SavePostOK{}
}

/*SalesRuleRuleRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type SalesRuleRuleRepositoryV1SavePostOK struct {
	Payload *models.SalesRuleDataRuleInterface
}

func (o *SalesRuleRuleRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/salesRules][%d] salesRuleRuleRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesRuleDataRuleInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1SavePostBadRequest creates a SalesRuleRuleRepositoryV1SavePostBadRequest with default headers values
func NewSalesRuleRuleRepositoryV1SavePostBadRequest() *SalesRuleRuleRepositoryV1SavePostBadRequest {
	return &SalesRuleRuleRepositoryV1SavePostBadRequest{}
}

/*SalesRuleRuleRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesRuleRuleRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/salesRules][%d] salesRuleRuleRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1SavePostUnauthorized creates a SalesRuleRuleRepositoryV1SavePostUnauthorized with default headers values
func NewSalesRuleRuleRepositoryV1SavePostUnauthorized() *SalesRuleRuleRepositoryV1SavePostUnauthorized {
	return &SalesRuleRuleRepositoryV1SavePostUnauthorized{}
}

/*SalesRuleRuleRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleRuleRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/salesRules][%d] salesRuleRuleRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1SavePostInternalServerError creates a SalesRuleRuleRepositoryV1SavePostInternalServerError with default headers values
func NewSalesRuleRuleRepositoryV1SavePostInternalServerError() *SalesRuleRuleRepositoryV1SavePostInternalServerError {
	return &SalesRuleRuleRepositoryV1SavePostInternalServerError{}
}

/*SalesRuleRuleRepositoryV1SavePostInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleRuleRepositoryV1SavePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleRuleRepositoryV1SavePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/salesRules][%d] salesRuleRuleRepositoryV1SavePostInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1SavePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleRuleRepositoryV1SavePostDefault creates a SalesRuleRuleRepositoryV1SavePostDefault with default headers values
func NewSalesRuleRuleRepositoryV1SavePostDefault(code int) *SalesRuleRuleRepositoryV1SavePostDefault {
	return &SalesRuleRuleRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*SalesRuleRuleRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleRuleRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule rule repository v1 save post default response
func (o *SalesRuleRuleRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleRuleRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/salesRules][%d] salesRuleRuleRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleRuleRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesRuleRuleRepositoryV1SavePostBody sales rule rule repository v1 save post body
swagger:model SalesRuleRuleRepositoryV1SavePostBody
*/
type SalesRuleRuleRepositoryV1SavePostBody struct {

	// rule
	// Required: true
	Rule *models.SalesRuleDataRuleInterface `json:"rule"`
}

// Validate validates this sales rule rule repository v1 save post body
func (o *SalesRuleRuleRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesRuleRuleRepositoryV1SavePostBody) validateRule(formats strfmt.Registry) error {

	if err := validate.Required("salesRuleRuleRepositoryV1SavePostBody"+"."+"rule", "body", o.Rule); err != nil {
		return err
	}

	if o.Rule != nil {
		if err := o.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesRuleRuleRepositoryV1SavePostBody" + "." + "rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesRuleRuleRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesRuleRuleRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res SalesRuleRuleRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
