// Code generated by go-swagger; DO NOT EDIT.

package company_company_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CompanyCompanyRepositoryV1GetGetReader is a Reader for the CompanyCompanyRepositoryV1GetGet structure.
type CompanyCompanyRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyCompanyRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyCompanyRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompanyCompanyRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCompanyCompanyRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyCompanyRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyCompanyRepositoryV1GetGetOK creates a CompanyCompanyRepositoryV1GetGetOK with default headers values
func NewCompanyCompanyRepositoryV1GetGetOK() *CompanyCompanyRepositoryV1GetGetOK {
	return &CompanyCompanyRepositoryV1GetGetOK{}
}

/*CompanyCompanyRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type CompanyCompanyRepositoryV1GetGetOK struct {
	Payload *models.CompanyDataCompanyInterface
}

func (o *CompanyCompanyRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/company/{companyId}][%d] companyCompanyRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyDataCompanyInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetGetBadRequest creates a CompanyCompanyRepositoryV1GetGetBadRequest with default headers values
func NewCompanyCompanyRepositoryV1GetGetBadRequest() *CompanyCompanyRepositoryV1GetGetBadRequest {
	return &CompanyCompanyRepositoryV1GetGetBadRequest{}
}

/*CompanyCompanyRepositoryV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CompanyCompanyRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/company/{companyId}][%d] companyCompanyRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetGetUnauthorized creates a CompanyCompanyRepositoryV1GetGetUnauthorized with default headers values
func NewCompanyCompanyRepositoryV1GetGetUnauthorized() *CompanyCompanyRepositoryV1GetGetUnauthorized {
	return &CompanyCompanyRepositoryV1GetGetUnauthorized{}
}

/*CompanyCompanyRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyCompanyRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/company/{companyId}][%d] companyCompanyRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetGetDefault creates a CompanyCompanyRepositoryV1GetGetDefault with default headers values
func NewCompanyCompanyRepositoryV1GetGetDefault(code int) *CompanyCompanyRepositoryV1GetGetDefault {
	return &CompanyCompanyRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*CompanyCompanyRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type CompanyCompanyRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company company repository v1 get get default response
func (o *CompanyCompanyRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CompanyCompanyRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/company/{companyId}][%d] companyCompanyRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
