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

// CatalogDataProductCustomOptionTypeInterface catalog data product custom option type interface
// swagger:model catalog-data-product-custom-option-type-interface
type CatalogDataProductCustomOptionTypeInterface struct {

	// Option type code
	// Required: true
	Code *string `json:"code"`

	// extension attributes
	ExtensionAttributes CatalogDataProductCustomOptionTypeExtensionInterface `json:"extension_attributes,omitempty"`

	// Option type group
	// Required: true
	Group *string `json:"group"`

	// Option type label
	// Required: true
	Label *string `json:"label"`
}

// Validate validates this catalog data product custom option type interface
func (m *CatalogDataProductCustomOptionTypeInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductCustomOptionTypeInterface) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionTypeInterface) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionTypeInterface) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductCustomOptionTypeInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductCustomOptionTypeInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductCustomOptionTypeInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
