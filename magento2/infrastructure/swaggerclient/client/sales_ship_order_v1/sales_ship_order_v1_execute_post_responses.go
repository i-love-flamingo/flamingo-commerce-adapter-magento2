// Code generated by go-swagger; DO NOT EDIT.

package sales_ship_order_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// SalesShipOrderV1ExecutePostReader is a Reader for the SalesShipOrderV1ExecutePost structure.
type SalesShipOrderV1ExecutePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesShipOrderV1ExecutePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesShipOrderV1ExecutePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesShipOrderV1ExecutePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesShipOrderV1ExecutePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesShipOrderV1ExecutePostOK creates a SalesShipOrderV1ExecutePostOK with default headers values
func NewSalesShipOrderV1ExecutePostOK() *SalesShipOrderV1ExecutePostOK {
	return &SalesShipOrderV1ExecutePostOK{}
}

/*SalesShipOrderV1ExecutePostOK handles this case with default header values.

200 Success.
*/
type SalesShipOrderV1ExecutePostOK struct {
	Payload int64
}

func (o *SalesShipOrderV1ExecutePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/ship][%d] salesShipOrderV1ExecutePostOK  %+v", 200, o.Payload)
}

func (o *SalesShipOrderV1ExecutePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipOrderV1ExecutePostUnauthorized creates a SalesShipOrderV1ExecutePostUnauthorized with default headers values
func NewSalesShipOrderV1ExecutePostUnauthorized() *SalesShipOrderV1ExecutePostUnauthorized {
	return &SalesShipOrderV1ExecutePostUnauthorized{}
}

/*SalesShipOrderV1ExecutePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesShipOrderV1ExecutePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesShipOrderV1ExecutePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/ship][%d] salesShipOrderV1ExecutePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesShipOrderV1ExecutePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesShipOrderV1ExecutePostDefault creates a SalesShipOrderV1ExecutePostDefault with default headers values
func NewSalesShipOrderV1ExecutePostDefault(code int) *SalesShipOrderV1ExecutePostDefault {
	return &SalesShipOrderV1ExecutePostDefault{
		_statusCode: code,
	}
}

/*SalesShipOrderV1ExecutePostDefault handles this case with default header values.

Unexpected error
*/
type SalesShipOrderV1ExecutePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales ship order v1 execute post default response
func (o *SalesShipOrderV1ExecutePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesShipOrderV1ExecutePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/ship][%d] salesShipOrderV1ExecutePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesShipOrderV1ExecutePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesShipOrderV1ExecutePostBody sales ship order v1 execute post body
swagger:model SalesShipOrderV1ExecutePostBody
*/
type SalesShipOrderV1ExecutePostBody struct {

	// append comment
	AppendComment bool `json:"appendComment,omitempty"`

	// arguments
	Arguments *models.SalesDataShipmentCreationArgumentsInterface `json:"arguments,omitempty"`

	// comment
	Comment *models.SalesDataShipmentCommentCreationInterface `json:"comment,omitempty"`

	// items
	Items []*models.SalesDataShipmentItemCreationInterface `json:"items"`

	// notify
	Notify bool `json:"notify,omitempty"`

	// packages
	Packages []*models.SalesDataShipmentPackageCreationInterface `json:"packages"`

	// tracks
	Tracks []*models.SalesDataShipmentTrackCreationInterface `json:"tracks"`
}

// Validate validates this sales ship order v1 execute post body
func (o *SalesShipOrderV1ExecutePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesShipOrderV1ExecutePostBody) validateArguments(formats strfmt.Registry) error {

	if swag.IsZero(o.Arguments) { // not required
		return nil
	}

	if o.Arguments != nil {
		if err := o.Arguments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesShipOrderV1ExecutePostBody" + "." + "arguments")
			}
			return err
		}
	}

	return nil
}

func (o *SalesShipOrderV1ExecutePostBody) validateComment(formats strfmt.Registry) error {

	if swag.IsZero(o.Comment) { // not required
		return nil
	}

	if o.Comment != nil {
		if err := o.Comment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesShipOrderV1ExecutePostBody" + "." + "comment")
			}
			return err
		}
	}

	return nil
}

func (o *SalesShipOrderV1ExecutePostBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("salesShipOrderV1ExecutePostBody" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SalesShipOrderV1ExecutePostBody) validatePackages(formats strfmt.Registry) error {

	if swag.IsZero(o.Packages) { // not required
		return nil
	}

	for i := 0; i < len(o.Packages); i++ {
		if swag.IsZero(o.Packages[i]) { // not required
			continue
		}

		if o.Packages[i] != nil {
			if err := o.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("salesShipOrderV1ExecutePostBody" + "." + "packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SalesShipOrderV1ExecutePostBody) validateTracks(formats strfmt.Registry) error {

	if swag.IsZero(o.Tracks) { // not required
		return nil
	}

	for i := 0; i < len(o.Tracks); i++ {
		if swag.IsZero(o.Tracks[i]) { // not required
			continue
		}

		if o.Tracks[i] != nil {
			if err := o.Tracks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("salesShipOrderV1ExecutePostBody" + "." + "tracks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesShipOrderV1ExecutePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesShipOrderV1ExecutePostBody) UnmarshalBinary(b []byte) error {
	var res SalesShipOrderV1ExecutePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
