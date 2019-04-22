// Code generated by go-swagger; DO NOT EDIT.

package sales_order_address_repository_v1

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

// SalesOrderAddressRepositoryV1SavePutReader is a Reader for the SalesOrderAddressRepositoryV1SavePut structure.
type SalesOrderAddressRepositoryV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesOrderAddressRepositoryV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesOrderAddressRepositoryV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesOrderAddressRepositoryV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesOrderAddressRepositoryV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesOrderAddressRepositoryV1SavePutOK creates a SalesOrderAddressRepositoryV1SavePutOK with default headers values
func NewSalesOrderAddressRepositoryV1SavePutOK() *SalesOrderAddressRepositoryV1SavePutOK {
	return &SalesOrderAddressRepositoryV1SavePutOK{}
}

/*SalesOrderAddressRepositoryV1SavePutOK handles this case with default header values.

200 Success.
*/
type SalesOrderAddressRepositoryV1SavePutOK struct {
	Payload *models.SalesDataOrderAddressInterface
}

func (o *SalesOrderAddressRepositoryV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/orders/{parent_id}][%d] salesOrderAddressRepositoryV1SavePutOK  %+v", 200, o.Payload)
}

func (o *SalesOrderAddressRepositoryV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataOrderAddressInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderAddressRepositoryV1SavePutUnauthorized creates a SalesOrderAddressRepositoryV1SavePutUnauthorized with default headers values
func NewSalesOrderAddressRepositoryV1SavePutUnauthorized() *SalesOrderAddressRepositoryV1SavePutUnauthorized {
	return &SalesOrderAddressRepositoryV1SavePutUnauthorized{}
}

/*SalesOrderAddressRepositoryV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesOrderAddressRepositoryV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesOrderAddressRepositoryV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/orders/{parent_id}][%d] salesOrderAddressRepositoryV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesOrderAddressRepositoryV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesOrderAddressRepositoryV1SavePutDefault creates a SalesOrderAddressRepositoryV1SavePutDefault with default headers values
func NewSalesOrderAddressRepositoryV1SavePutDefault(code int) *SalesOrderAddressRepositoryV1SavePutDefault {
	return &SalesOrderAddressRepositoryV1SavePutDefault{
		_statusCode: code,
	}
}

/*SalesOrderAddressRepositoryV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type SalesOrderAddressRepositoryV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales order address repository v1 save put default response
func (o *SalesOrderAddressRepositoryV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *SalesOrderAddressRepositoryV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/orders/{parent_id}][%d] salesOrderAddressRepositoryV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *SalesOrderAddressRepositoryV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesOrderAddressRepositoryV1SavePutBody sales order address repository v1 save put body
swagger:model SalesOrderAddressRepositoryV1SavePutBody
*/
type SalesOrderAddressRepositoryV1SavePutBody struct {

	// entity
	// Required: true
	Entity *models.SalesDataOrderAddressInterface `json:"entity"`
}

// Validate validates this sales order address repository v1 save put body
func (o *SalesOrderAddressRepositoryV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesOrderAddressRepositoryV1SavePutBody) validateEntity(formats strfmt.Registry) error {

	if err := validate.Required("salesOrderAddressRepositoryV1SavePutBody"+"."+"entity", "body", o.Entity); err != nil {
		return err
	}

	if o.Entity != nil {
		if err := o.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesOrderAddressRepositoryV1SavePutBody" + "." + "entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesOrderAddressRepositoryV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesOrderAddressRepositoryV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res SalesOrderAddressRepositoryV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}