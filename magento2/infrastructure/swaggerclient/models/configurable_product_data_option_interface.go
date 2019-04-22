// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfigurableProductDataOptionInterface Interface OptionInterface
// swagger:model configurable-product-data-option-interface
type ConfigurableProductDataOptionInterface struct {

	// attribute id
	AttributeID string `json:"attribute_id,omitempty"`

	// extension attributes
	ExtensionAttributes ConfigurableProductDataOptionExtensionInterface `json:"extension_attributes,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is use default
	IsUseDefault bool `json:"is_use_default,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// position
	Position int64 `json:"position,omitempty"`

	// product id
	ProductID int64 `json:"product_id,omitempty"`

	// values
	Values []*ConfigurableProductDataOptionValueInterface `json:"values"`
}

// Validate validates this configurable product data option interface
func (m *ConfigurableProductDataOptionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurableProductDataOptionInterface) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurableProductDataOptionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurableProductDataOptionInterface) UnmarshalBinary(b []byte) error {
	var res ConfigurableProductDataOptionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
