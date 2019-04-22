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

// FrameworkMetadataObjectInterface Provides metadata about an attribute.
// swagger:model framework-metadata-object-interface
type FrameworkMetadataObjectInterface struct {

	// Code of the attribute.
	// Required: true
	AttributeCode *string `json:"attribute_code"`
}

// Validate validates this framework metadata object interface
func (m *FrameworkMetadataObjectInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameworkMetadataObjectInterface) validateAttributeCode(formats strfmt.Registry) error {

	if err := validate.Required("attribute_code", "body", m.AttributeCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrameworkMetadataObjectInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrameworkMetadataObjectInterface) UnmarshalBinary(b []byte) error {
	var res FrameworkMetadataObjectInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}