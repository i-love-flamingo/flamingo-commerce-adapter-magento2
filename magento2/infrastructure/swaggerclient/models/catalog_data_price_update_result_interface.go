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

// CatalogDataPriceUpdateResultInterface Interface returned in case of incorrect price passed to efficient price API.
// swagger:model catalog-data-price-update-result-interface
type CatalogDataPriceUpdateResultInterface struct {

	// extension attributes
	ExtensionAttributes CatalogDataPriceUpdateResultExtensionInterface `json:"extension_attributes,omitempty"`

	// Error message, that contains description of error occurred during price update.
	// Required: true
	Message *string `json:"message"`

	// Parameters, that could be displayed in error message placeholders.
	// Required: true
	Parameters []string `json:"parameters"`
}

// Validate validates this catalog data price update result interface
func (m *CatalogDataPriceUpdateResultInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataPriceUpdateResultInterface) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataPriceUpdateResultInterface) validateParameters(formats strfmt.Registry) error {

	if err := validate.Required("parameters", "body", m.Parameters); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataPriceUpdateResultInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataPriceUpdateResultInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataPriceUpdateResultInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
