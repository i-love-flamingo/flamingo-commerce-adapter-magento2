// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SalesDataShipmentCreationArgumentsInterface Interface for creation arguments for Shipment.
// swagger:model sales-data-shipment-creation-arguments-interface
type SalesDataShipmentCreationArgumentsInterface struct {

	// extension attributes
	ExtensionAttributes SalesDataShipmentCreationArgumentsExtensionInterface `json:"extension_attributes,omitempty"`
}

// Validate validates this sales data shipment creation arguments interface
func (m *SalesDataShipmentCreationArgumentsInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataShipmentCreationArgumentsInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataShipmentCreationArgumentsInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataShipmentCreationArgumentsInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}