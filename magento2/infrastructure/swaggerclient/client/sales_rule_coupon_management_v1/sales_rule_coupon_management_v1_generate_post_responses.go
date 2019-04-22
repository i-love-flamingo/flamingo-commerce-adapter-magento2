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

// SalesRuleCouponManagementV1GeneratePostReader is a Reader for the SalesRuleCouponManagementV1GeneratePost structure.
type SalesRuleCouponManagementV1GeneratePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesRuleCouponManagementV1GeneratePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSalesRuleCouponManagementV1GeneratePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSalesRuleCouponManagementV1GeneratePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSalesRuleCouponManagementV1GeneratePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSalesRuleCouponManagementV1GeneratePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesRuleCouponManagementV1GeneratePostOK creates a SalesRuleCouponManagementV1GeneratePostOK with default headers values
func NewSalesRuleCouponManagementV1GeneratePostOK() *SalesRuleCouponManagementV1GeneratePostOK {
	return &SalesRuleCouponManagementV1GeneratePostOK{}
}

/*SalesRuleCouponManagementV1GeneratePostOK handles this case with default header values.

200 Success.
*/
type SalesRuleCouponManagementV1GeneratePostOK struct {
	Payload []string
}

func (o *SalesRuleCouponManagementV1GeneratePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/generate][%d] salesRuleCouponManagementV1GeneratePostOK  %+v", 200, o.Payload)
}

func (o *SalesRuleCouponManagementV1GeneratePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1GeneratePostUnauthorized creates a SalesRuleCouponManagementV1GeneratePostUnauthorized with default headers values
func NewSalesRuleCouponManagementV1GeneratePostUnauthorized() *SalesRuleCouponManagementV1GeneratePostUnauthorized {
	return &SalesRuleCouponManagementV1GeneratePostUnauthorized{}
}

/*SalesRuleCouponManagementV1GeneratePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type SalesRuleCouponManagementV1GeneratePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponManagementV1GeneratePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/generate][%d] salesRuleCouponManagementV1GeneratePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SalesRuleCouponManagementV1GeneratePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1GeneratePostInternalServerError creates a SalesRuleCouponManagementV1GeneratePostInternalServerError with default headers values
func NewSalesRuleCouponManagementV1GeneratePostInternalServerError() *SalesRuleCouponManagementV1GeneratePostInternalServerError {
	return &SalesRuleCouponManagementV1GeneratePostInternalServerError{}
}

/*SalesRuleCouponManagementV1GeneratePostInternalServerError handles this case with default header values.

Internal Server error
*/
type SalesRuleCouponManagementV1GeneratePostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SalesRuleCouponManagementV1GeneratePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/generate][%d] salesRuleCouponManagementV1GeneratePostInternalServerError  %+v", 500, o.Payload)
}

func (o *SalesRuleCouponManagementV1GeneratePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesRuleCouponManagementV1GeneratePostDefault creates a SalesRuleCouponManagementV1GeneratePostDefault with default headers values
func NewSalesRuleCouponManagementV1GeneratePostDefault(code int) *SalesRuleCouponManagementV1GeneratePostDefault {
	return &SalesRuleCouponManagementV1GeneratePostDefault{
		_statusCode: code,
	}
}

/*SalesRuleCouponManagementV1GeneratePostDefault handles this case with default header values.

Unexpected error
*/
type SalesRuleCouponManagementV1GeneratePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the sales rule coupon management v1 generate post default response
func (o *SalesRuleCouponManagementV1GeneratePostDefault) Code() int {
	return o._statusCode
}

func (o *SalesRuleCouponManagementV1GeneratePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/coupons/generate][%d] salesRuleCouponManagementV1GeneratePost default  %+v", o._statusCode, o.Payload)
}

func (o *SalesRuleCouponManagementV1GeneratePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SalesRuleCouponManagementV1GeneratePostBody sales rule coupon management v1 generate post body
swagger:model SalesRuleCouponManagementV1GeneratePostBody
*/
type SalesRuleCouponManagementV1GeneratePostBody struct {

	// coupon spec
	// Required: true
	CouponSpec *models.SalesRuleDataCouponGenerationSpecInterface `json:"couponSpec"`
}

// Validate validates this sales rule coupon management v1 generate post body
func (o *SalesRuleCouponManagementV1GeneratePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCouponSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SalesRuleCouponManagementV1GeneratePostBody) validateCouponSpec(formats strfmt.Registry) error {

	if err := validate.Required("salesRuleCouponManagementV1GeneratePostBody"+"."+"couponSpec", "body", o.CouponSpec); err != nil {
		return err
	}

	if o.CouponSpec != nil {
		if err := o.CouponSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("salesRuleCouponManagementV1GeneratePostBody" + "." + "couponSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SalesRuleCouponManagementV1GeneratePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SalesRuleCouponManagementV1GeneratePostBody) UnmarshalBinary(b []byte) error {
	var res SalesRuleCouponManagementV1GeneratePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
