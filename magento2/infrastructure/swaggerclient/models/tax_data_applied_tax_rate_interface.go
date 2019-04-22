// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TaxDataAppliedTaxRateInterface Applied tax rate interface.
// swagger:model tax-data-applied-tax-rate-interface
type TaxDataAppliedTaxRateInterface struct {

	// Code
	Code string `json:"code,omitempty"`

	// extension attributes
	ExtensionAttributes TaxDataAppliedTaxRateExtensionInterface `json:"extension_attributes,omitempty"`

	// Tax Percent
	Percent float64 `json:"percent,omitempty"`

	// Title
	Title string `json:"title,omitempty"`
}

// Validate validates this tax data applied tax rate interface
func (m *TaxDataAppliedTaxRateInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxDataAppliedTaxRateInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDataAppliedTaxRateInterface) UnmarshalBinary(b []byte) error {
	var res TaxDataAppliedTaxRateInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}