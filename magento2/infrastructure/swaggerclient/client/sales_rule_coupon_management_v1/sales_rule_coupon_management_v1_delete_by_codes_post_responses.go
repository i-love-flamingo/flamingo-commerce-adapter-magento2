// Code generated by go-swagger; DO NOT EDIT.

package sales_rule_coupon_management_v1

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

// SalesRuleCouponManagementV1DeleteByCodesPostReader is a Reader for the SalesRuleCouponManagementV1DeleteByCodesPost structure.
type SalesRuleCouponManagementV1DeleteByCodesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleCouponManagementV1DeleteByCodesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleCouponManagementV1DeleteByCodesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSalesRuleCouponManagementV1DeleteByCodesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSalesRuleCouponManagementV1DeleteByCodesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleCouponManagementV1DeleteByCodesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleCouponManagementV1DeleteByCodesPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleCouponManagementV1DeleteByCodesPostOK creates a SalesRuleCouponManagementV1DeleteByCodesPostOK with default headers values
func NewSalesRuleCouponManagementV1DeleteByCodesPostOK() *SalesRuleCouponManagementV1DeleteByCodesPostOK {
	return &SalesRuleCouponManagementV1DeleteByCodesPostOK{}
}

/*SalesRuleCouponManagementV1DeleteByCodesPostOK handles this case with default header values.

200 Success.
*/
type SalesRuleCouponManagementV1DeleteByCodesPostOK struct {
	Payload *models.SalesRuleDataCouponMassDeleteResultInterface
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/deleteByCodes][%d] salesRuleCouponManagementV1DeleteByCodesPostOK  %+v", 200, o.Payload)
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesRuleDataCouponMassDeleteResultInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1DeleteByCodesPostBadRequest creates a SalesRuleCouponManagementV1DeleteByCodesPostBadRequest with default headers values
func NewSalesRuleCouponManagementV1DeleteByCodesPostBadRequest() *SalesRuleCouponManagementV1DeleteByCodesPostBadRequest {
	return &SalesRuleCouponManagementV1DeleteByCodesPostBadRequest{}
}

/*SalesRuleCouponManagementV1DeleteByCodesPostBadRequest handles this case with default header values.

400 Bad Request
*/
type SalesRuleCouponManagementV1DeleteByCodesPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/deleteByCodes][%d] salesRuleCouponManagementV1DeleteByCodesPostBadRequest  %+v", 400, o.Payload)
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1DeleteByCodesPostUnauthorized creates a SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized with default headers values
func NewSalesRuleCouponManagementV1DeleteByCodesPostUnauthorized() *SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized {
	return &SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized{}
}

/*SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/deleteByCodes][%d] salesRuleCouponManagementV1DeleteByCodesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1DeleteByCodesPostInternalServerError creates a SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError with default headers values
func NewSalesRuleCouponManagementV1DeleteByCodesPostInternalServerError() *SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError {
	return &SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError{}
}

/*SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/deleteByCodes][%d] salesRuleCouponManagementV1DeleteByCodesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1DeleteByCodesPostDefault creates a SalesRuleCouponManagementV1DeleteByCodesPostDefault with default headers values
func NewSalesRuleCouponManagementV1DeleteByCodesPostDefault(code int) *SalesRuleCouponManagementV1DeleteByCodesPostDefault {
	return &SalesRuleCouponManagementV1DeleteByCodesPostDefault{
		_statusCode: code,
	}
}

/*SalesRuleCouponManagementV1DeleteByCodesPostDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleCouponManagementV1DeleteByCodesPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule coupon management v1 delete by codes post default response
func (o *SalesRuleCouponManagementV1DeleteByCodesPostDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/deleteByCodes][%d] salesRuleCouponManagementV1DeleteByCodesPost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesRuleCouponManagementV1DeleteByCodesPostBody sales rule coupon management v1 delete by codes post body
swagger:model SalesRuleCouponManagementV1DeleteByCodesPostBody
*/
type SalesRuleCouponManagementV1DeleteByCodesPostBody struct {

	// codes
	// Required: true
	Codes []string `json:"codes"`

	// ignore invalid coupons
	IgnoreInvalidCoupons bool `json:"ignoreInvalidCoupons,omitempty"`
}

// Validate validates this sales rule coupon management v1 delete by codes post body
func (o *SalesRuleCouponManagementV1DeleteByCodesPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesRuleCouponManagementV1DeleteByCodesPostBody) validateCodes(formats strfmt.Registry) error {

	if err := validate.Required("salesRuleCouponManagementV1DeleteByCodesPostBody"+"."+"codes", "body", o.Codes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesRuleCouponManagementV1DeleteByCodesPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesRuleCouponManagementV1DeleteByCodesPostBody) UnmarshalBinary(b []byte) error {
	var res SalesRuleCouponManagementV1DeleteByCodesPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
