// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SalesDataShipmentPackageInterface Shipment package interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package.
// swagger:model sales-data-shipment-package-interface
type SalesDataShipmentPackageInterface struct {

	// extension attributes
	ExtensionAttributes SalesDataShipmentPackageExtensionInterface `json:"extension_attributes,omitempty"`
}

// Validate validates this sales data shipment package interface
func (m *SalesDataShipmentPackageInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataShipmentPackageInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataShipmentPackageInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataShipmentPackageInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}