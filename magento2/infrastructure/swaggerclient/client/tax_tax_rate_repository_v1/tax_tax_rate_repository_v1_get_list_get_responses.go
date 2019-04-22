// Code generated by go-swagger; DO NOT EDIT.

package tax_tax_rate_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// TaxTaxRateRepositoryV1GetListGetReader is a Reader for the TaxTaxRateRepositoryV1GetListGet structure.
type TaxTaxRateRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaxTaxRateRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTaxTaxRateRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTaxTaxRateRepositoryV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTaxTaxRateRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTaxTaxRateRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaxTaxRateRepositoryV1GetListGetOK creates a TaxTaxRateRepositoryV1GetListGetOK with default headers values
func NewTaxTaxRateRepositoryV1GetListGetOK() *TaxTaxRateRepositoryV1GetListGetOK {
	return &TaxTaxRateRepositoryV1GetListGetOK{}
}

/*TaxTaxRateRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type TaxTaxRateRepositoryV1GetListGetOK struct {
	Payload *models.TaxDataTaxRateSearchResultsInterface
}

func (o *TaxTaxRateRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/taxRates/search][%d] taxTaxRateRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *TaxTaxRateRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxDataTaxRateSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1GetListGetBadRequest creates a TaxTaxRateRepositoryV1GetListGetBadRequest with default headers values
func NewTaxTaxRateRepositoryV1GetListGetBadRequest() *TaxTaxRateRepositoryV1GetListGetBadRequest {
	return &TaxTaxRateRepositoryV1GetListGetBadRequest{}
}

/*TaxTaxRateRepositoryV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type TaxTaxRateRepositoryV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/taxRates/search][%d] taxTaxRateRepositoryV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *TaxTaxRateRepositoryV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1GetListGetUnauthorized creates a TaxTaxRateRepositoryV1GetListGetUnauthorized with default headers values
func NewTaxTaxRateRepositoryV1GetListGetUnauthorized() *TaxTaxRateRepositoryV1GetListGetUnauthorized {
	return &TaxTaxRateRepositoryV1GetListGetUnauthorized{}
}

/*TaxTaxRateRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type TaxTaxRateRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *TaxTaxRateRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/taxRates/search][%d] taxTaxRateRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *TaxTaxRateRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxTaxRateRepositoryV1GetListGetDefault creates a TaxTaxRateRepositoryV1GetListGetDefault with default headers values
func NewTaxTaxRateRepositoryV1GetListGetDefault(code int) *TaxTaxRateRepositoryV1GetListGetDefault {
	return &TaxTaxRateRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*TaxTaxRateRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type TaxTaxRateRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the tax tax rate repository v1 get list get default response
func (o *TaxTaxRateRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *TaxTaxRateRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/taxRates/search][%d] taxTaxRateRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *TaxTaxRateRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}