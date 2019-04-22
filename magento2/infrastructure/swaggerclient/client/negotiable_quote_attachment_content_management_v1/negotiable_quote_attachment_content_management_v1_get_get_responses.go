// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_attachment_content_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// NegotiableQuoteAttachmentContentManagementV1GetGetReader is a Reader for the NegotiableQuoteAttachmentContentManagementV1GetGet structure.
type NegotiableQuoteAttachmentContentManagementV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NegotiableQuoteAttachmentContentManagementV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNegotiableQuoteAttachmentContentManagementV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNegotiableQuoteAttachmentContentManagementV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewNegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNegotiableQuoteAttachmentContentManagementV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNegotiableQuoteAttachmentContentManagementV1GetGetOK creates a NegotiableQuoteAttachmentContentManagementV1GetGetOK with default headers values
func NewNegotiableQuoteAttachmentContentManagementV1GetGetOK() *NegotiableQuoteAttachmentContentManagementV1GetGetOK {
	return &NegotiableQuoteAttachmentContentManagementV1GetGetOK{}
}

/*NegotiableQuoteAttachmentContentManagementV1GetGetOK handles this case with default header values.

200 Success.
*/
type NegotiableQuoteAttachmentContentManagementV1GetGetOK struct {
	Payload []*models.NegotiableQuoteDataAttachmentContentInterface
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/negotiableQuote/attachmentContent][%d] negotiableQuoteAttachmentContentManagementV1GetGetOK  %+v", 200, o.Payload)
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteAttachmentContentManagementV1GetGetBadRequest creates a NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest with default headers values
func NewNegotiableQuoteAttachmentContentManagementV1GetGetBadRequest() *NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest {
	return &NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest{}
}

/*NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest handles this case with default header values.

400 Bad Request
*/
type NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/negotiableQuote/attachmentContent][%d] negotiableQuoteAttachmentContentManagementV1GetGetBadRequest  %+v", 400, o.Payload)
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized creates a NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized with default headers values
func NewNegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized() *NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized {
	return &NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized{}
}

/*NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized handles this case with default header values.

401 Unauthorized
*/
type NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /V1/negotiableQuote/attachmentContent][%d] negotiableQuoteAttachmentContentManagementV1GetGetUnauthorized  %+v", 401, o.Payload)
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteAttachmentContentManagementV1GetGetDefault creates a NegotiableQuoteAttachmentContentManagementV1GetGetDefault with default headers values
func NewNegotiableQuoteAttachmentContentManagementV1GetGetDefault(code int) *NegotiableQuoteAttachmentContentManagementV1GetGetDefault {
	return &NegotiableQuoteAttachmentContentManagementV1GetGetDefault{
		_statusCode: code,
	}
}

/*NegotiableQuoteAttachmentContentManagementV1GetGetDefault handles this case with default header values.

Unexpected error
*/
type NegotiableQuoteAttachmentContentManagementV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the negotiable quote attachment content management v1 get get default response
func (o *NegotiableQuoteAttachmentContentManagementV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/negotiableQuote/attachmentContent][%d] negotiableQuoteAttachmentContentManagementV1GetGet default  %+v", o._statusCode, o.Payload)
}

func (o *NegotiableQuoteAttachmentContentManagementV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}