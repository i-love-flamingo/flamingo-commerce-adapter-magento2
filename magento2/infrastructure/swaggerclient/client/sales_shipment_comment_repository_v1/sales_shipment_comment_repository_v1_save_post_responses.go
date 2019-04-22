// Code generated by go-swagger; DO NOT EDIT.

package sales_shipment_comment_repository_v1

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

// SalesShipmentCommentRepositoryV1SavePostReader is a Reader for the SalesShipmentCommentRepositoryV1SavePost structure.
type SalesShipmentCommentRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesShipmentCommentRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesShipmentCommentRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesShipmentCommentRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesShipmentCommentRepositoryV1SavePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesShipmentCommentRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesShipmentCommentRepositoryV1SavePostOK creates a SalesShipmentCommentRepositoryV1SavePostOK with default headers values
func NewSalesShipmentCommentRepositoryV1SavePostOK() *SalesShipmentCommentRepositoryV1SavePostOK {
	return &SalesShipmentCommentRepositoryV1SavePostOK{}
}

/*SalesShipmentCommentRepositoryV1SavePostOK handles this case with default header values.

200 Success.
*/
type SalesShipmentCommentRepositoryV1SavePostOK struct {
	Payload *models.SalesDataShipmentCommentInterface
}

func (o *SalesShipmentCommentRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/{id}/comments][%d] salesShipmentCommentRepositoryV1SavePostOK  %+v", 200, o.Payload)
}

func (o *SalesShipmentCommentRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataShipmentCommentInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentCommentRepositoryV1SavePostBadRequest creates a SalesShipmentCommentRepositoryV1SavePostBadRequest with default headers values
func NewSalesShipmentCommentRepositoryV1SavePostBadRequest() *SalesShipmentCommentRepositoryV1SavePostBadRequest {
	return &SalesShipmentCommentRepositoryV1SavePostBadRequest{}
}

/*SalesShipmentCommentRepositoryV1SavePostBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesShipmentCommentRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentCommentRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/{id}/comments][%d] salesShipmentCommentRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}

func (o *SalesShipmentCommentRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentCommentRepositoryV1SavePostUnauthorized creates a SalesShipmentCommentRepositoryV1SavePostUnauthorized with default headers values
func NewSalesShipmentCommentRepositoryV1SavePostUnauthorized() *SalesShipmentCommentRepositoryV1SavePostUnauthorized {
	return &SalesShipmentCommentRepositoryV1SavePostUnauthorized{}
}

/*SalesShipmentCommentRepositoryV1SavePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesShipmentCommentRepositoryV1SavePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipmentCommentRepositoryV1SavePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/{id}/comments][%d] salesShipmentCommentRepositoryV1SavePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesShipmentCommentRepositoryV1SavePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipmentCommentRepositoryV1SavePostDefault creates a SalesShipmentCommentRepositoryV1SavePostDefault with default headers values
func NewSalesShipmentCommentRepositoryV1SavePostDefault(code int) *SalesShipmentCommentRepositoryV1SavePostDefault {
	return &SalesShipmentCommentRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/*SalesShipmentCommentRepositoryV1SavePostDefault handles this case with default header values.

Unexpected error
*/
type SalesShipmentCommentRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales shipment comment repository v1 save post default response
func (o *SalesShipmentCommentRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesShipmentCommentRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/shipment/{id}/comments][%d] salesShipmentCommentRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesShipmentCommentRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesShipmentCommentRepositoryV1SavePostBody sales shipment comment repository v1 save post body
swagger:model SalesShipmentCommentRepositoryV1SavePostBody
*/
type SalesShipmentCommentRepositoryV1SavePostBody struct {

	// entity
	// Required: true
	Entity *models.SalesDataShipmentCommentInterface `json:"entity"`
}

// Validate validates this sales shipment comment repository v1 save post body
func (o *SalesShipmentCommentRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesShipmentCommentRepositoryV1SavePostBody) validateEntity(formats strfmt.Registry) error {

	if err := validate.Required("salesShipmentCommentRepositoryV1SavePostBody"+"."+"entity", "body", o.Entity); err != nil {
		return err
	}

	if o.Entity != nil {
		if err := o.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesShipmentCommentRepositoryV1SavePostBody" + "." + "entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesShipmentCommentRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesShipmentCommentRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res SalesShipmentCommentRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
