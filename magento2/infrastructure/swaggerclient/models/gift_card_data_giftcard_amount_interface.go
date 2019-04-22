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

// GiftCardDataGiftcardAmountInterface Interface GiftcardAmountInterface: this interface is used to serialize and deserialize EAV attribute giftcard_amounts
// swagger:model gift-card-data-giftcard-amount-interface
type GiftCardDataGiftcardAmountInterface struct {

	// attribute id
	// Required: true
	AttributeID *int64 `json:"attribute_id"`

	// extension attributes
	ExtensionAttributes GiftCardDataGiftcardAmountExtensionInterface `json:"extension_attributes,omitempty"`

	// value
	// Required: true
	Value *float64 `json:"value"`

	// website id
	// Required: true
	WebsiteID *int64 `json:"website_id"`

	// website value
	// Required: true
	WebsiteValue *float64 `json:"website_value"`
}

// Validate validates this gift card data giftcard amount interface
func (m *GiftCardDataGiftcardAmountInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsiteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsiteValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GiftCardDataGiftcardAmountInterface) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", m.AttributeID); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftcardAmountInterface) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftcardAmountInterface) validateWebsiteID(formats strfmt.Registry) error {

	if err := validate.Required("website_id", "body", m.WebsiteID); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardDataGiftcardAmountInterface) validateWebsiteValue(formats strfmt.Registry) error {

	if err := validate.Required("website_value", "body", m.WebsiteValue); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GiftCardDataGiftcardAmountInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftCardDataGiftcardAmountInterface) UnmarshalBinary(b []byte) error {
	var res GiftCardDataGiftcardAmountInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}