// Code generated by go-swagger; DO NOT EDIT.

package company_company_repository_v1

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

// CompanyCompanyRepositoryV1SavePutReader is a Reader for the CompanyCompanyRepositoryV1SavePut structure.
type CompanyCompanyRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyCompanyRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyCompanyRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompanyCompanyRepositoryV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCompanyCompanyRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyCompanyRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyCompanyRepositoryV1SavePutOK creates a CompanyCompanyRepositoryV1SavePutOK with default headers values
func NewCompanyCompanyRepositoryV1SavePutOK() *CompanyCompanyRepositoryV1SavePutOK {
	return &CompanyCompanyRepositoryV1SavePutOK{}
}

/*CompanyCompanyRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type CompanyCompanyRepositoryV1SavePutOK struct {
	Payload *models.CompanyDataCompanyInterface
}

func (o *CompanyCompanyRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/company/{companyId}][%d] companyCompanyRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *CompanyCompanyRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyDataCompanyInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1SavePutBadRequest creates a CompanyCompanyRepositoryV1SavePutBadRequest with default headers values
func NewCompanyCompanyRepositoryV1SavePutBadRequest() *CompanyCompanyRepositoryV1SavePutBadRequest {
	return &CompanyCompanyRepositoryV1SavePutBadRequest{}
}

/*CompanyCompanyRepositoryV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type CompanyCompanyRepositoryV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/company/{companyId}][%d] companyCompanyRepositoryV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *CompanyCompanyRepositoryV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1SavePutUnauthorized creates a CompanyCompanyRepositoryV1SavePutUnauthorized with default headers values
func NewCompanyCompanyRepositoryV1SavePutUnauthorized() *CompanyCompanyRepositoryV1SavePutUnauthorized {
	return &CompanyCompanyRepositoryV1SavePutUnauthorized{}
}

/*CompanyCompanyRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyCompanyRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/company/{companyId}][%d] companyCompanyRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyCompanyRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1SavePutDefault creates a CompanyCompanyRepositoryV1SavePutDefault with default headers values
func NewCompanyCompanyRepositoryV1SavePutDefault(code int) *CompanyCompanyRepositoryV1SavePutDefault {
	return &CompanyCompanyRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*CompanyCompanyRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type CompanyCompanyRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company company repository v1 save put default response
func (o *CompanyCompanyRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *CompanyCompanyRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/company/{companyId}][%d] companyCompanyRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyCompanyRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompanyCompanyRepositoryV1SavePutBody company company repository v1 save put body
swagger:model CompanyCompanyRepositoryV1SavePutBody
*/
type CompanyCompanyRepositoryV1SavePutBody struct {

	// company
	// Required: true
	Company *models.CompanyDataCompanyInterface `json:"company"`
}

// Validate validates this company company repository v1 save put body
func (o *CompanyCompanyRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CompanyCompanyRepositoryV1SavePutBody) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("companyCompanyRepositoryV1SavePutBody"+"."+"company", "body", o.Company); err != nil {
		return err
	}

	if o.Company != nil {
		if err := o.Company.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("companyCompanyRepositoryV1SavePutBody" + "." + "company")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CompanyCompanyRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompanyCompanyRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res CompanyCompanyRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}