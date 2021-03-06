// Code generated by go-swagger; DO NOT EDIT.

package negotiable_quote_gift_card_account_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
)

// NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteReader is a Reader for the NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDelete structure.
type NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK creates a NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK with default headers values
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK() *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK {
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK{}
}

/*NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK handles this case with default header values.

200 Success.
*/
type NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK struct {
	Payload bool
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /V1/negotiable-carts/{cartId}/giftCards/{giftCardCode}][%d] negotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIdDeleteOK  %+v", 200, o.Payload)
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized creates a NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized with default headers values
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized() *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized {
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized{}
}

/*NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized handles this case with default header values.

401 Unauthorized
*/
type NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /V1/negotiable-carts/{cartId}/giftCards/{giftCardCode}][%d] negotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault creates a NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault with default headers values
func NewNegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault(code int) *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault {
	return &NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault{
		_statusCode: code,
	}
}

/*NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault handles this case with default header values.

Unexpected error
*/
type NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the negotiable quote gift card account management v1 delete by quote Id delete default response
func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/negotiable-carts/{cartId}/giftCards/{giftCardCode}][%d] negotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *NegotiableQuoteGiftCardAccountManagementV1DeleteByQuoteIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
