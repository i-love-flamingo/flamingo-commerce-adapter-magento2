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

// QuoteDataPaymentMethodInterface Interface PaymentMethodInterface
// swagger:model quote-data-payment-method-interface
type QuoteDataPaymentMethodInterface struct {

	// Payment method code
	// Required: true
	Code *string `json:"code"`

	// Payment method title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this quote data payment method interface
func (m *QuoteDataPaymentMethodInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataPaymentMethodInterface) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataPaymentMethodInterface) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteDataPaymentMethodInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteDataPaymentMethodInterface) UnmarshalBinary(b []byte) error {
	var res QuoteDataPaymentMethodInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}