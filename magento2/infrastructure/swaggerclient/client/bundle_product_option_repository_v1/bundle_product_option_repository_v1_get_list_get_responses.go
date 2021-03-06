// Code generated by go-swagger; DO NOT EDIT.

package bundle_product_option_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// BundleProductOptionRepositoryV1GetListGetReader is a Reader for the BundleProductOptionRepositoryV1GetListGet structure.
type BundleProductOptionRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BundleProductOptionRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBundleProductOptionRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBundleProductOptionRepositoryV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBundleProductOptionRepositoryV1GetListGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewBundleProductOptionRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBundleProductOptionRepositoryV1GetListGetOK creates a BundleProductOptionRepositoryV1GetListGetOK with default headers values
func NewBundleProductOptionRepositoryV1GetListGetOK() *BundleProductOptionRepositoryV1GetListGetOK {
	return &BundleProductOptionRepositoryV1GetListGetOK{}
}

/*BundleProductOptionRepositoryV1GetListGetOK handles this case with default header values.

200 Success.
*/
type BundleProductOptionRepositoryV1GetListGetOK struct {
	Payload []*models.BundleDataOptionInterface
}

func (o *BundleProductOptionRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/bundle-products/{sku}/options/all][%d] bundleProductOptionRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}

func (o *BundleProductOptionRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionRepositoryV1GetListGetBadRequest creates a BundleProductOptionRepositoryV1GetListGetBadRequest with default headers values
func NewBundleProductOptionRepositoryV1GetListGetBadRequest() *BundleProductOptionRepositoryV1GetListGetBadRequest {
	return &BundleProductOptionRepositoryV1GetListGetBadRequest{}
}

/*BundleProductOptionRepositoryV1GetListGetBadRequest handles this case with default header values.

400 Bad Request
*/
type BundleProductOptionRepositoryV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *BundleProductOptionRepositoryV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/bundle-products/{sku}/options/all][%d] bundleProductOptionRepositoryV1GetListGetBadRequest  %+v", 400, o.Payload)
}

func (o *BundleProductOptionRepositoryV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionRepositoryV1GetListGetUnauthorized creates a BundleProductOptionRepositoryV1GetListGetUnauthorized with default headers values
func NewBundleProductOptionRepositoryV1GetListGetUnauthorized() *BundleProductOptionRepositoryV1GetListGetUnauthorized {
	return &BundleProductOptionRepositoryV1GetListGetUnauthorized{}
}

/*BundleProductOptionRepositoryV1GetListGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type BundleProductOptionRepositoryV1GetListGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *BundleProductOptionRepositoryV1GetListGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/bundle-products/{sku}/options/all][%d] bundleProductOptionRepositoryV1GetListGetUnauthorized  %+v", 401, o.Payload)
}

func (o *BundleProductOptionRepositoryV1GetListGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBundleProductOptionRepositoryV1GetListGetDefault creates a BundleProductOptionRepositoryV1GetListGetDefault with default headers values
func NewBundleProductOptionRepositoryV1GetListGetDefault(code int) *BundleProductOptionRepositoryV1GetListGetDefault {
	return &BundleProductOptionRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/*BundleProductOptionRepositoryV1GetListGetDefault handles this case with default header values.

Unexpected error
*/
type BundleProductOptionRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bundle product option repository v1 get list get default response
func (o *BundleProductOptionRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *BundleProductOptionRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/bundle-products/{sku}/options/all][%d] bundleProductOptionRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}

func (o *BundleProductOptionRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
