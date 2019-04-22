// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_negotiable_quote_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// NegotiableQuoteNegotiableQuoteManagementV1CreatePostReader is a Reader for the NegotiableQuoteNegotiableQuoteManagementV1CreatePost structure.
type NegotiableQuoteNegotiableQuoteManagementV1CreatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostOK creates a NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK with default headers values
func NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostOK() *NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK {
	return &NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK{}
}

/*NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK handles this case with default header values.

200 Success.
*/
type NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK struct {
	Payload bool
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/negotiableQuote/request][%d] negotiableQuoteNegotiableQuoteManagementV1CreatePostOK  %+v", 200, o.Payload)
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized creates a NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized with default headers values
func NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized() *NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized {
	return &NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized{}
}

/*NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized handles this case with default header values.

401 Unauthorized
*/
type NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /V1/negotiableQuote/request][%d] negotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized  %+v", 401, o.Payload)
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault creates a NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault with default headers values
func NewNegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault(code int) *NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault {
	return &NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault{
		_statusCode: code,
	}
}

/*NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault handles this case with default header values.

Unexpected error
*/
type NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the negotiable quote negotiable quote management v1 create post default response
func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault) Code() int {
	return o._statusCode
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/negotiableQuote/request][%d] negotiableQuoteNegotiableQuoteManagementV1CreatePost default  %+v", o._statusCode, o.Payload)
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody negotiable quote negotiable quote management v1 create post body
swagger:model NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody
*/
type NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// files
	Files []*models.NegotiableQuoteDataAttachmentContentInterface `json:"files"`

	// quote Id
	// Required: true
	QuoteID *int64 `json:"quoteId"`

	// quote name
	// Required: true
	QuoteName *string `json:"quoteName"`
}

// Validate validates this negotiable quote negotiable quote management v1 create post body
func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuoteID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuoteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(o.Files) { // not required
		return nil
	}

	for i := 0; i < len(o.Files); i++ {
		if swag.IsZero(o.Files[i]) { // not required
			continue
		}

		if o.Files[i] != nil {
			if err := o.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("negotiableQuoteNegotiableQuoteManagementV1CreatePostBody" + "." + "files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) validateQuoteID(formats strfmt.Registry) error {

	if err := validate.Required("negotiableQuoteNegotiableQuoteManagementV1CreatePostBody"+"."+"quoteId", "body", o.QuoteID); err != nil {
		return err
	}

	return nil
}

func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) validateQuoteName(formats strfmt.Registry) error {

	if err := validate.Required("negotiableQuoteNegotiableQuoteManagementV1CreatePostBody"+"."+"quoteName", "body", o.QuoteName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody) UnmarshalBinary(b []byte) error {
	var res NegotiableQuoteNegotiableQuoteManagementV1CreatePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
