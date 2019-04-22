// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SalesDataShippingInterface Interface ShippingInterface
// swagger:model sales-data-shipping-interface
type SalesDataShippingInterface struct {

	// address
	Address *SalesDataOrderAddressInterface `json:"address,omitempty"`

	// extension attributes
	ExtensionAttributes SalesDataShippingExtensionInterface `json:"extension_attributes,omitempty"`

	// Shipping method
	Method string `json:"method,omitempty"`

	// total
	Total *SalesDataTotalInterface `json:"total,omitempty"`
}

// Validate validates this sales data shipping interface
func (m *SalesDataShippingInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesDataShippingInterface) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *SalesDataShippingInterface) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataShippingInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataShippingInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataShippingInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}