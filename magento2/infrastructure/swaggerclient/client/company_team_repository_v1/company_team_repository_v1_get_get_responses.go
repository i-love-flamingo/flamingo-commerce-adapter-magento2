// Code generated by go-swagger; DO NOT EDIT.

package company_team_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// CompanyTeamRepositoryV1GetGetReader is a Reader for the CompanyTeamRepositoryV1GetGet structure.
type CompanyTeamRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanyTeamRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompanyTeamRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompanyTeamRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCompanyTeamRepositoryV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCompanyTeamRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompanyTeamRepositoryV1GetGetOK creates a CompanyTeamRepositoryV1GetGetOK with default headers values
func NewCompanyTeamRepositoryV1GetGetOK() *CompanyTeamRepositoryV1GetGetOK {
	return &CompanyTeamRepositoryV1GetGetOK{}
}

/*CompanyTeamRepositoryV1GetGetOK handles this case with default header values.

200 Success.
*/
type CompanyTeamRepositoryV1GetGetOK struct {
	Payload *models.CompanyDataTeamInterface
}

func (o *CompanyTeamRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/team/{teamId}][%d] companyTeamRepositoryV1GetGetOK  %+v", 200, o.Payload)
}

func (o *CompanyTeamRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyDataTeamInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyTeamRepositoryV1GetGetBadRequest creates a CompanyTeamRepositoryV1GetGetBadRequest with default headers values
func NewCompanyTeamRepositoryV1GetGetBadRequest() *CompanyTeamRepositoryV1GetGetBadRequest {
	return &CompanyTeamRepositoryV1GetGetBadRequest{}
}

/*CompanyTeamRepositoryV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type CompanyTeamRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompanyTeamRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/team/{teamId}][%d] companyTeamRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *CompanyTeamRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyTeamRepositoryV1GetGetUnauthorized creates a CompanyTeamRepositoryV1GetGetUnauthorized with default headers values
func NewCompanyTeamRepositoryV1GetGetUnauthorized() *CompanyTeamRepositoryV1GetGetUnauthorized {
	return &CompanyTeamRepositoryV1GetGetUnauthorized{}
}

/*CompanyTeamRepositoryV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type CompanyTeamRepositoryV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompanyTeamRepositoryV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/team/{teamId}][%d] companyTeamRepositoryV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CompanyTeamRepositoryV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompanyTeamRepositoryV1GetGetDefault creates a CompanyTeamRepositoryV1GetGetDefault with default headers values
func NewCompanyTeamRepositoryV1GetGetDefault(code int) *CompanyTeamRepositoryV1GetGetDefault {
	return &CompanyTeamRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/*CompanyTeamRepositoryV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type CompanyTeamRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the company team repository v1 get get default response
func (o *CompanyTeamRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CompanyTeamRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/team/{teamId}][%d] companyTeamRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *CompanyTeamRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}