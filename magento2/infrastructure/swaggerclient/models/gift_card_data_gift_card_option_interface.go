// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GiftCardDataGiftCardOptionInterface Interface GiftCardOptionInterface
// swagger:model gift-card-data-gift-card-option-interface
type GiftCardDataGiftCardOptionInterface struct {

	// Gift card open amount value.
	CustomGiftcardAmount float64 `json:"custom_giftcard_amount,omitempty"`

	// extension attributes
	ExtensionAttributes GiftCardDataGiftCardOptionExtensionInterface `json:"extension_attributes,omitempty"`

	// Gift card amount.
	// Required: true
	GiftcardAmount *string `json:"giftcard_amount"`

	// Giftcard message.
	GiftcardMessage string `json:"giftcard_message,omitempty"`

	// Gift card recipient email.
	// Required: true
	GiftcardRecipientEmail *string `json:"giftcard_recipient_email"`

	// Gift card recipient name.
	// Required: true
	GiftcardRecipientName *string `json:"giftcard_recipient_name"`

	// Gift card sender email.
	// Required: true
	GiftcardSenderEmail *string `json:"giftcard_sender_email"`

	// Gift card sender name.
	// Required: true
	GiftcardSenderName *string `json:"giftcard_sender_name"`
}

// Validate validates this gift card data gift card option interface
func (m *GiftCardDataGiftCardOptionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGiftcardAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftcardRecipientEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftcardRecipientName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftcardSenderEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftcardSenderName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GiftCardDataGiftCardOptionInterface) validateGiftcardAmount(formats strfmt.Registry) error {

	if err := validate.Required("giftcard_amount", "body", m.GiftcardAmount); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftCardOptionInterface) validateGiftcardRecipientEmail(formats strfmt.Registry) error {

	if err := validate.Required("giftcard_recipient_email", "body", m.GiftcardRecipientEmail); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftCardOptionInterface) validateGiftcardRecipientName(formats strfmt.Registry) error {

	if err := validate.Required("giftcard_recipient_name", "body", m.GiftcardRecipientName); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftCardOptionInterface) validateGiftcardSenderEmail(formats strfmt.Registry) error {

	if err := validate.Required("giftcard_sender_email", "body", m.GiftcardSenderEmail); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftCardOptionInterface) validateGiftcardSenderName(formats strfmt.Registry) error {

	if err := validate.Required("giftcard_sender_name", "body", m.GiftcardSenderName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GiftCardDataGiftCardOptionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftCardDataGiftCardOptionInterface) UnmarshalBinary(b []byte) error {
	var res GiftCardDataGiftCardOptionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}