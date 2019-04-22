// Code generated by go-swagger; DO NOT EDIT.

package sales_refund_order_v1

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

// SalesRefundOrderV1ExecutePostReader is a Reader for the SalesRefundOrderV1ExecutePost structure.
type SalesRefundOrderV1ExecutePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRefundOrderV1ExecutePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRefundOrderV1ExecutePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesRefundOrderV1ExecutePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRefundOrderV1ExecutePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRefundOrderV1ExecutePostOK creates a SalesRefundOrderV1ExecutePostOK with default headers values
func NewSalesRefundOrderV1ExecutePostOK() *SalesRefundOrderV1ExecutePostOK {
	return &SalesRefundOrderV1ExecutePostOK{}
}

/*SalesRefundOrderV1ExecutePostOK handles this case with default header values.

200 Success.
*/
type SalesRefundOrderV1ExecutePostOK struct {
	Payload int64
}

func (o *SalesRefundOrderV1ExecutePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/refund][%d] salesRefundOrderV1ExecutePostOK  %+v", 200, o.Payload)
}

func (o *SalesRefundOrderV1ExecutePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRefundOrderV1ExecutePostUnauthorized creates a SalesRefundOrderV1ExecutePostUnauthorized with default headers values
func NewSalesRefundOrderV1ExecutePostUnauthorized() *SalesRefundOrderV1ExecutePostUnauthorized {
	return &SalesRefundOrderV1ExecutePostUnauthorized{}
}

/*SalesRefundOrderV1ExecutePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRefundOrderV1ExecutePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRefundOrderV1ExecutePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/refund][%d] salesRefundOrderV1ExecutePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRefundOrderV1ExecutePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRefundOrderV1ExecutePostDefault creates a SalesRefundOrderV1ExecutePostDefault with default headers values
func NewSalesRefundOrderV1ExecutePostDefault(code int) *SalesRefundOrderV1ExecutePostDefault {
	return &SalesRefundOrderV1ExecutePostDefault{
		_statusCode: code,
	}
}

/*SalesRefundOrderV1ExecutePostDefault handles this case with default header values.

Unexpected error
*/
type SalesRefundOrderV1ExecutePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales refund order v1 execute post default response
func (o *SalesRefundOrderV1ExecutePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesRefundOrderV1ExecutePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/order/{orderId}/refund][%d] salesRefundOrderV1ExecutePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRefundOrderV1ExecutePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesRefundOrderV1ExecutePostBody sales refund order v1 execute post body
swagger:model SalesRefundOrderV1ExecutePostBody
*/
type SalesRefundOrderV1ExecutePostBody struct {

	// append comment
	AppendComment bool `json:"appendComment,omitempty"`

	// arguments
	Arguments *models.SalesDataCreditmemoCreationArgumentsInterface `json:"arguments,omitempty"`

	// comment
	Comment *models.SalesDataCreditmemoCommentCreationInterface `json:"comment,omitempty"`

	// items
	Items []*models.SalesDataCreditmemoItemCreationInterface `json:"items"`

	// notify
	Notify bool `json:"notify,omitempty"`
}

// Validate validates this sales refund order v1 execute post body
func (o *SalesRefundOrderV1ExecutePostBody) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesRefundOrderV1ExecutePostBody) validateArguments(formats strfmt.Registry) error {

	if swag.IsZero(o.Arguments) { // not required
		return nil
	}

	if o.Arguments != nil {
		if err := o.Arguments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesRefundOrderV1ExecutePostBody" + "." + "arguments")
			}
			return err
		}
	}

	return nil
}

func (o *SalesRefundOrderV1ExecutePostBody) validateComment(formats strfmt.Registry) error {

	if swag.IsZero(o.Comment) { // not required
		return nil
	}

	if o.Comment != nil {
		if err := o.Comment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesRefundOrderV1ExecutePostBody" + "." + "comment")
			}
			return err
		}
	}

	return nil
}

func (o *SalesRefundOrderV1ExecutePostBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("salesRefundOrderV1ExecutePostBody" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesRefundOrderV1ExecutePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesRefundOrderV1ExecutePostBody) UnmarshalBinary(b []byte) error {
	var res SalesRefundOrderV1ExecutePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
