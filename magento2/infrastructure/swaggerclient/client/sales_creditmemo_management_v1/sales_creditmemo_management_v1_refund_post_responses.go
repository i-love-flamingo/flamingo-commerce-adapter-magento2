// Code generated by go-swagger; DO NOT EDIT.

package sales_creditmemo_management_v1

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

// SalesCreditmemoManagementV1RefundPostReader is a Reader for the SalesCreditmemoManagementV1RefundPost structure.
type SalesCreditmemoManagementV1RefundPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesCreditmemoManagementV1RefundPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesCreditmemoManagementV1RefundPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesCreditmemoManagementV1RefundPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesCreditmemoManagementV1RefundPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesCreditmemoManagementV1RefundPostOK creates a SalesCreditmemoManagementV1RefundPostOK with default headers values
func NewSalesCreditmemoManagementV1RefundPostOK() *SalesCreditmemoManagementV1RefundPostOK {
	return &SalesCreditmemoManagementV1RefundPostOK{}
}

/*SalesCreditmemoManagementV1RefundPostOK handles this case with default header values.

200 Success.
*/
type SalesCreditmemoManagementV1RefundPostOK struct {
	Payload *models.SalesDataCreditmemoInterface
}

func (o *SalesCreditmemoManagementV1RefundPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/refund][%d] salesCreditmemoManagementV1RefundPostOK  %+v", 200, o.Payload)
}

func (o *SalesCreditmemoManagementV1RefundPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesDataCreditmemoInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoManagementV1RefundPostUnauthorized creates a SalesCreditmemoManagementV1RefundPostUnauthorized with default headers values
func NewSalesCreditmemoManagementV1RefundPostUnauthorized() *SalesCreditmemoManagementV1RefundPostUnauthorized {
	return &SalesCreditmemoManagementV1RefundPostUnauthorized{}
}

/*SalesCreditmemoManagementV1RefundPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesCreditmemoManagementV1RefundPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesCreditmemoManagementV1RefundPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/refund][%d] salesCreditmemoManagementV1RefundPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesCreditmemoManagementV1RefundPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesCreditmemoManagementV1RefundPostDefault creates a SalesCreditmemoManagementV1RefundPostDefault with default headers values
func NewSalesCreditmemoManagementV1RefundPostDefault(code int) *SalesCreditmemoManagementV1RefundPostDefault {
	return &SalesCreditmemoManagementV1RefundPostDefault{
		_statusCode: code,
	}
}

/*SalesCreditmemoManagementV1RefundPostDefault handles this case with default header values.

Unexpected error
*/
type SalesCreditmemoManagementV1RefundPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales creditmemo management v1 refund post default response
func (o *SalesCreditmemoManagementV1RefundPostDefault) Code() int {
	return o._statusCode
}

func (o *SalesCreditmemoManagementV1RefundPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/creditmemo/refund][%d] salesCreditmemoManagementV1RefundPost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesCreditmemoManagementV1RefundPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesCreditmemoManagementV1RefundPostBody sales creditmemo management v1 refund post body
swagger:model SalesCreditmemoManagementV1RefundPostBody
*/
type SalesCreditmemoManagementV1RefundPostBody struct {

	// creditmemo
	// Required: true
	Creditmemo *models.SalesDataCreditmemoInterface `json:"creditmemo"`

	// offline requested
	OfflineRequested bool `json:"offlineRequested,omitempty"`
}

// Validate validates this sales creditmemo management v1 refund post body
func (o *SalesCreditmemoManagementV1RefundPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreditmemo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesCreditmemoManagementV1RefundPostBody) validateCreditmemo(formats strfmt.Registry) error {

	if err := validate.Required("salesCreditmemoManagementV1RefundPostBody"+"."+"creditmemo", "body", o.Creditmemo); err != nil {
		return err
	}

	if o.Creditmemo != nil {
		if err := o.Creditmemo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesCreditmemoManagementV1RefundPostBody" + "." + "creditmemo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesCreditmemoManagementV1RefundPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesCreditmemoManagementV1RefundPostBody) UnmarshalBinary(b []byte) error {
	var res SalesCreditmemoManagementV1RefundPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}