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

// GiftCardAccountDataGiftCardInterface Gift Card data
// swagger:model gift-card-account-data-gift-card-interface
type GiftCardAccountDataGiftCardInterface struct {

	// Amount
	// Required: true
	Amount *float64 `json:"amount"`

	// Base Amount
	// Required: true
	BaseAmount *float64 `json:"base_amount"`

	// Code
	// Required: true
	Code *string `json:"code"`

	// Id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this gift card account data gift card interface
func (m *GiftCardAccountDataGiftCardInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GiftCardAccountDataGiftCardInterface) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardAccountDataGiftCardInterface) validateBaseAmount(formats strfmt.Registry) error {

	if err := validate.Required("base_amount", "body", m.BaseAmount); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardAccountDataGiftCardInterface) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *GiftCardAccountDataGiftCardInterface) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GiftCardAccountDataGiftCardInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftCardAccountDataGiftCardInterface) UnmarshalBinary(b []byte) error {
	var res GiftCardAccountDataGiftCardInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
