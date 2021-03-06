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

// SalesDataShipmentTrackCreationInterface Shipment Track Creation interface.
// swagger:model sales-data-shipment-track-creation-interface
type SalesDataShipmentTrackCreationInterface struct {

	// Carrier code.
	// Required: true
	CarrierCode *string `json:"carrier_code"`

	// extension attributes
	ExtensionAttributes SalesDataShipmentTrackCreationExtensionInterface `json:"extension_attributes,omitempty"`

	// Title.
	// Required: true
	Title *string `json:"title"`

	// Track number.
	// Required: true
	TrackNumber *string `json:"track_number"`
}

// Validate validates this sales data shipment track creation interface
func (m *SalesDataShipmentTrackCreationInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesDataShipmentTrackCreationInterface) validateCarrierCode(formats strfmt.Registry) error {

	if err := validate.Required("carrier_code", "body", m.CarrierCode); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataShipmentTrackCreationInterface) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataShipmentTrackCreationInterface) validateTrackNumber(formats strfmt.Registry) error {

	if err := validate.Required("track_number", "body", m.TrackNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataShipmentTrackCreationInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataShipmentTrackCreationInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataShipmentTrackCreationInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
