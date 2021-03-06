// Code generated by go-swagger; DO NOT EDIT.

package sales_shipment_track_repository_v1

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

// SalesShipmentTrackRepositoryV1SavePostReader is a Reader for the SalesShipmentTrackRepositoryV1SavePost structure.
type SalesShipmentTrackRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesShipmentTrackRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesShipmentTrackRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesShipmentTrackRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesShipmentTrackRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesShipmentTrackRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesShipmentTrackRepositoryV1SavePostOK creates a SalesShipmentTrackRepositoryV1SavePostOK with default headers values
func NewSalesShipmentTrackRepositoryV1SavePostOK() *SalesShipmentTrackRepositoryV1SavePostOK {
	return &SalesShipmentTrackRepositoryV1SavePostOK{}
}

/*SalesShipmentTrackRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type SalesShipmentTrackRepositoryV1SavePostOK struct {
	Payload *models.SalesDataShipmentTrackInterface
}

func (o *SalesShipmentTrackRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/track][%d] salesShipmentTrackRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *SalesShipmentTrackRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataShipmentTrackInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentTrackRepositoryV1SavePostBadRequest creates a SalesShipmentTrackRepositoryV1SavePostBadRequest with default headers values
func NewSalesShipmentTrackRepositoryV1SavePostBadRequest() *SalesShipmentTrackRepositoryV1SavePostBadRequest {
	return &SalesShipmentTrackRepositoryV1SavePostBadRequest{}
}

/*SalesShipmentTrackRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesShipmentTrackRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentTrackRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/track][%d] salesShipmentTrackRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *SalesShipmentTrackRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentTrackRepositoryV1SavePostUnauthorized creates a SalesShipmentTrackRepositoryV1SavePostUnauthorized with default headers values
func NewSalesShipmentTrackRepositoryV1SavePostUnauthorized() *SalesShipmentTrackRepositoryV1SavePostUnauthorized {
	return &SalesShipmentTrackRepositoryV1SavePostUnauthorized{}
}

/*SalesShipmentTrackRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesShipmentTrackRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentTrackRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/track][%d] salesShipmentTrackRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesShipmentTrackRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentTrackRepositoryV1SavePostDefault creates a SalesShipmentTrackRepositoryV1SavePostDefault with default headers values
func NewSalesShipmentTrackRepositoryV1SavePostDefault(code int) *SalesShipmentTrackRepositoryV1SavePostDefault {
	return &SalesShipmentTrackRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*SalesShipmentTrackRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type SalesShipmentTrackRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales shipment track repository v1 save post default response
func (o *SalesShipmentTrackRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesShipmentTrackRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/track][%d] salesShipmentTrackRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesShipmentTrackRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesShipmentTrackRepositoryV1SavePostBody sales shipment track repository v1 save post body
swagger:model SalesShipmentTrackRepositoryV1SavePostBody
*/
type SalesShipmentTrackRepositoryV1SavePostBody struct {

	// entity
	// Required: true
	Entity *models.SalesDataShipmentTrackInterface `json:"entity"`
}

// Validate validates this sales shipment track repository v1 save post body
func (o *SalesShipmentTrackRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesShipmentTrackRepositoryV1SavePostBody) validateEntity(formats strfmt.Registry) error {

	if err := validate.Required("salesShipmentTrackRepositoryV1SavePostBody"+"."+"entity", "body", o.Entity); err != nil {
		return err
	}

	if o.Entity != nil {
		if err := o.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesShipmentTrackRepositoryV1SavePostBody" + "." + "entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesShipmentTrackRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesShipmentTrackRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res SalesShipmentTrackRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
