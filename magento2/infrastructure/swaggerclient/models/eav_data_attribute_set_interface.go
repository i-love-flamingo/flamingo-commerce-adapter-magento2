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

// EavDataAttributeSetInterface Interface AttributeSetInterface
// swagger:model eav-data-attribute-set-interface
type EavDataAttributeSetInterface struct {

	// Attribute set ID
	AttributeSetID int64 `json:"attribute_set_id,omitempty"`

	// Attribute set name
	// Required: true
	AttributeSetName *string `json:"attribute_set_name"`

	// Attribute set entity type id
	EntityTypeID int64 `json:"entity_type_id,omitempty"`

	// extension attributes
	ExtensionAttributes EavDataAttributeSetExtensionInterface `json:"extension_attributes,omitempty"`

	// Attribute set sort order index
	// Required: true
	SortOrder *int64 `json:"sort_order"`
}

// Validate validates this eav data attribute set interface
func (m *EavDataAttributeSetInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeSetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EavDataAttributeSetInterface) validateAttributeSetName(formats strfmt.Registry) error {

	if err := validate.Required("attribute_set_name", "body", m.AttributeSetName); err != nil {
		return err
	}

	return nil
}

func (m *EavDataAttributeSetInterface) validateSortOrder(formats strfmt.Registry) error {

	if err := validate.Required("sort_order", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EavDataAttributeSetInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EavDataAttributeSetInterface) UnmarshalBinary(b []byte) error {
	var res EavDataAttributeSetInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}