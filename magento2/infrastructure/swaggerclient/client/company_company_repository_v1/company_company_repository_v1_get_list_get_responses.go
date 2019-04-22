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

// CompanyCompanyRepositoryV1GetListGetReader is a Reader for the CompanyCompanyRepositoryV1GetListGet structure.
type CompanyCompanyRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyCompanyRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyCompanyRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCompanyCompanyRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCompanyCompanyRepositoryV1GetListGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyCompanyRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyCompanyRepositoryV1GetListGetOK creates a CompanyCompanyRepositoryV1GetListGetOK with default headers values
func NewCompanyCompanyRepositoryV1GetListGetOK() *CompanyCompanyRepositoryV1GetListGetOK {
	return &CompanyCompanyRepositoryV1GetListGetOK{}
}

/*CompanyCompanyRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type CompanyCompanyRepositoryV1GetListGetOK struct {
	Payload *models.CompanyDataCompanySearchResultsInterface
}

func (o *CompanyCompanyRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/company/][%d] companyCompanyRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyDataCompanySearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetListGetUnauthorized creates a CompanyCompanyRepositoryV1GetListGetUnauthorized with default headers values
func NewCompanyCompanyRepositoryV1GetListGetUnauthorized() *CompanyCompanyRepositoryV1GetListGetUnauthorized {
	return &CompanyCompanyRepositoryV1GetListGetUnauthorized{}
}

/*CompanyCompanyRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyCompanyRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/company/][%d] companyCompanyRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetListGetInternalServerError creates a CompanyCompanyRepositoryV1GetListGetInternalServerError with default headers values
func NewCompanyCompanyRepositoryV1GetListGetInternalServerError() *CompanyCompanyRepositoryV1GetListGetInternalServerError {
	return &CompanyCompanyRepositoryV1GetListGetInternalServerError{}
}

/*CompanyCompanyRepositoryV1GetListGetInternalServerError handles this case with default header values.

Internal Server error
*/
type CompanyCompanyRepositoryV1GetListGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CompanyCompanyRepositoryV1GetListGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/company/][%d] companyCompanyRepositoryV1GetListGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetListGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyCompanyRepositoryV1GetListGetDefault creates a CompanyCompanyRepositoryV1GetListGetDefault with default headers values
func NewCompanyCompanyRepositoryV1GetListGetDefault(code int) *CompanyCompanyRepositoryV1GetListGetDefault {
	return &CompanyCompanyRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*CompanyCompanyRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type CompanyCompanyRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company company repository v1 get list get default response
func (o *CompanyCompanyRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CompanyCompanyRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/company/][%d] companyCompanyRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyCompanyRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}