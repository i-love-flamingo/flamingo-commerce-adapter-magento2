// Code generated by go-swagger; DO NOT EDIT.

package bundle_product_option_management_v1

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

// BundleProductOptionManagementV1SavePutReader is a Reader for the BundleProductOptionManagementV1SavePut structure.
type BundleProductOptionManagementV1SavePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BundleProductOptionManagementV1SavePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBundleProductOptionManagementV1SavePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBundleProductOptionManagementV1SavePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBundleProductOptionManagementV1SavePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewBundleProductOptionManagementV1SavePutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBundleProductOptionManagementV1SavePutOK creates a BundleProductOptionManagementV1SavePutOK with default headers values
func NewBundleProductOptionManagementV1SavePutOK() *BundleProductOptionManagementV1SavePutOK {
	return &BundleProductOptionManagementV1SavePutOK{}
}

/*BundleProductOptionManagementV1SavePutOK handles this case with default header values.

200 Success.
*/
type BundleProductOptionManagementV1SavePutOK struct {
	Payload int64
}

func (o *BundleProductOptionManagementV1SavePutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/bundle-products/options/{optionId}][%d] bundleProductOptionManagementV1SavePutOK  %+v", 200, o.Payload)
}

func (o *BundleProductOptionManagementV1SavePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionManagementV1SavePutBadRequest creates a BundleProductOptionManagementV1SavePutBadRequest with default headers values
func NewBundleProductOptionManagementV1SavePutBadRequest() *BundleProductOptionManagementV1SavePutBadRequest {
	return &BundleProductOptionManagementV1SavePutBadRequest{}
}

/*BundleProductOptionManagementV1SavePutBadRequest handles this case with default header values.

400 Bad Request
*/
type BundleProductOptionManagementV1SavePutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *BundleProductOptionManagementV1SavePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /V1/bundle-products/options/{optionId}][%d] bundleProductOptionManagementV1SavePutBadRequest  %+v", 400, o.Payload)
}

func (o *BundleProductOptionManagementV1SavePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionManagementV1SavePutUnauthorized creates a BundleProductOptionManagementV1SavePutUnauthorized with default headers values
func NewBundleProductOptionManagementV1SavePutUnauthorized() *BundleProductOptionManagementV1SavePutUnauthorized {
	return &BundleProductOptionManagementV1SavePutUnauthorized{}
}

/*BundleProductOptionManagementV1SavePutUnauthorized handles this case with default header values.

401 Unauthorized
*/
type BundleProductOptionManagementV1SavePutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *BundleProductOptionManagementV1SavePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /V1/bundle-products/options/{optionId}][%d] bundleProductOptionManagementV1SavePutUnauthorized  %+v", 401, o.Payload)
}

func (o *BundleProductOptionManagementV1SavePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionManagementV1SavePutDefault creates a BundleProductOptionManagementV1SavePutDefault with default headers values
func NewBundleProductOptionManagementV1SavePutDefault(code int) *BundleProductOptionManagementV1SavePutDefault {
	return &BundleProductOptionManagementV1SavePutDefault{
		_statusCode: code,
	}
}

/*BundleProductOptionManagementV1SavePutDefault handles this case with default header values.

Unexpected error
*/
type BundleProductOptionManagementV1SavePutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bundle product option management v1 save put default response
func (o *BundleProductOptionManagementV1SavePutDefault) Code() int {
	return o._statusCode
}

func (o *BundleProductOptionManagementV1SavePutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/bundle-products/options/{optionId}][%d] bundleProductOptionManagementV1SavePut default  %+v", o._statusCode, o.Payload)
}

func (o *BundleProductOptionManagementV1SavePutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*BundleProductOptionManagementV1SavePutBody bundle product option management v1 save put body
swagger:model BundleProductOptionManagementV1SavePutBody
*/
type BundleProductOptionManagementV1SavePutBody struct {

	// option
	// Required: true
	Option *models.BundleDataOptionInterface `json:"option"`
}

// Validate validates this bundle product option management v1 save put body
func (o *BundleProductOptionManagementV1SavePutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BundleProductOptionManagementV1SavePutBody) validateOption(formats strfmt.Registry) error {

	if err := validate.Required("bundleProductOptionManagementV1SavePutBody"+"."+"option", "body", o.Option); err != nil {
		return err
	}

	if o.Option != nil {
		if err := o.Option.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundleProductOptionManagementV1SavePutBody" + "." + "option")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BundleProductOptionManagementV1SavePutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BundleProductOptionManagementV1SavePutBody) UnmarshalBinary(b []byte) error {
	var res BundleProductOptionManagementV1SavePutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
