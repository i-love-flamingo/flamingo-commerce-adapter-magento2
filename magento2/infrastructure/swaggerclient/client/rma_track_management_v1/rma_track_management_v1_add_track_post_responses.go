// Code generated by go-swagger; DO NOT EDIT.

package rma_track_management_v1

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

// RmaTrackManagementV1AddTrackPostReader is a Reader for the RmaTrackManagementV1AddTrackPost structure.
type RmaTrackManagementV1AddTrackPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RmaTrackManagementV1AddTrackPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRmaTrackManagementV1AddTrackPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRmaTrackManagementV1AddTrackPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRmaTrackManagementV1AddTrackPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRmaTrackManagementV1AddTrackPostOK creates a RmaTrackManagementV1AddTrackPostOK with default headers values
func NewRmaTrackManagementV1AddTrackPostOK() *RmaTrackManagementV1AddTrackPostOK {
	return &RmaTrackManagementV1AddTrackPostOK{}
}

/*RmaTrackManagementV1AddTrackPostOK handles this case with default header values.

200 Success.
*/
type RmaTrackManagementV1AddTrackPostOK struct {
	Payload bool
}

func (o *RmaTrackManagementV1AddTrackPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/returns/{id}/tracking-numbers][%d] rmaTrackManagementV1AddTrackPostOK  %+v", 200, o.Payload)
}

func (o *RmaTrackManagementV1AddTrackPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaTrackManagementV1AddTrackPostUnauthorized creates a RmaTrackManagementV1AddTrackPostUnauthorized with default headers values
func NewRmaTrackManagementV1AddTrackPostUnauthorized() *RmaTrackManagementV1AddTrackPostUnauthorized {
	return &RmaTrackManagementV1AddTrackPostUnauthorized{}
}

/*RmaTrackManagementV1AddTrackPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type RmaTrackManagementV1AddTrackPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RmaTrackManagementV1AddTrackPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/returns/{id}/tracking-numbers][%d] rmaTrackManagementV1AddTrackPostUnauthorized  %+v", 401, o.Payload)
}

func (o *RmaTrackManagementV1AddTrackPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRmaTrackManagementV1AddTrackPostDefault creates a RmaTrackManagementV1AddTrackPostDefault with default headers values
func NewRmaTrackManagementV1AddTrackPostDefault(code int) *RmaTrackManagementV1AddTrackPostDefault {
	return &RmaTrackManagementV1AddTrackPostDefault{
		_statusCode: code,
	}
}

/*RmaTrackManagementV1AddTrackPostDefault handles this case with default header values.

Unexpected error
*/
type RmaTrackManagementV1AddTrackPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the rma track management v1 add track post default response
func (o *RmaTrackManagementV1AddTrackPostDefault) Code() int {
	return o._statusCode
}

func (o *RmaTrackManagementV1AddTrackPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/returns/{id}/tracking-numbers][%d] rmaTrackManagementV1AddTrackPost default  %+v", o._statusCode, o.Payload)
}

func (o *RmaTrackManagementV1AddTrackPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RmaTrackManagementV1AddTrackPostBody rma track management v1 add track post body
swagger:model RmaTrackManagementV1AddTrackPostBody
*/
type RmaTrackManagementV1AddTrackPostBody struct {

	// track
	// Required: true
	Track *models.RmaDataTrackInterface `json:"track"`
}

// Validate validates this rma track management v1 add track post body
func (o *RmaTrackManagementV1AddTrackPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTrack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RmaTrackManagementV1AddTrackPostBody) validateTrack(formats strfmt.Registry) error {

	if err := validate.Required("rmaTrackManagementV1AddTrackPostBody"+"."+"track", "body", o.Track); err != nil {
		return err
	}

	if o.Track != nil {
		if err := o.Track.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rmaTrackManagementV1AddTrackPostBody" + "." + "track")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RmaTrackManagementV1AddTrackPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RmaTrackManagementV1AddTrackPostBody) UnmarshalBinary(b []byte) error {
	var res RmaTrackManagementV1AddTrackPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
