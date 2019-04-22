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

// GiftMessageDataMessageInterface Interface MessageInterface
// swagger:model gift-message-data-message-interface
type GiftMessageDataMessageInterface struct {

	// Customer ID. Otherwise, null.
	CustomerID int64 `json:"customer_id,omitempty"`

	// extension attributes
	ExtensionAttributes *GiftMessageDataMessageExtensionInterface `json:"extension_attributes,omitempty"`

	// Gift message ID. Otherwise, null.
	GiftMessageID int64 `json:"gift_message_id,omitempty"`

	// Message text.
	// Required: true
	Message *string `json:"message"`

	// Recipient name.
	// Required: true
	Recipient *string `json:"recipient"`

	// Sender name.
	// Required: true
	Sender *string `json:"sender"`
}

// Validate validates this gift message data message interface
func (m *GiftMessageDataMessageInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GiftMessageDataMessageInterface) validateExtensionAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtensionAttributes) { // not required
		return nil
	}

	if m.ExtensionAttributes != nil {
		if err := m.ExtensionAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension_attributes")
			}
			return err
		}
	}

	return nil
}

func (m *GiftMessageDataMessageInterface) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *GiftMessageDataMessageInterface) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	return nil
}

func (m *GiftMessageDataMessageInterface) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GiftMessageDataMessageInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftMessageDataMessageInterface) UnmarshalBinary(b []byte) error {
	var res GiftMessageDataMessageInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}