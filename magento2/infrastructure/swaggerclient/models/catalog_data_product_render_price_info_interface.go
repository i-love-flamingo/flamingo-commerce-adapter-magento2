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

// CatalogDataProductRenderPriceInfoInterface Price interface.
// swagger:model catalog-data-product-render-price-info-interface
type CatalogDataProductRenderPriceInfoInterface struct {

	// extension attributes
	ExtensionAttributes *CatalogDataProductRenderPriceInfoExtensionInterface `json:"extension_attributes,omitempty"`

	// Final price
	// Required: true
	FinalPrice *float64 `json:"final_price"`

	// formatted prices
	// Required: true
	FormattedPrices *CatalogDataProductRenderFormattedPriceInfoInterface `json:"formatted_prices"`

	// Max price of a product
	// Required: true
	MaxPrice *float64 `json:"max_price"`

	// Max regular price
	// Required: true
	MaxRegularPrice *float64 `json:"max_regular_price"`

	// minimal price
	// Required: true
	MinimalPrice *float64 `json:"minimal_price"`

	// Minimal regular price
	// Required: true
	MinimalRegularPrice *float64 `json:"minimal_regular_price"`

	// Regular price
	// Required: true
	RegularPrice *float64 `json:"regular_price"`

	// Special price
	// Required: true
	SpecialPrice *float64 `json:"special_price"`
}

// Validate validates this catalog data product render price info interface
func (m *CatalogDataProductRenderPriceInfoInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormattedPrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxRegularPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimalPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimalRegularPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateExtensionAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtensionAttributes) { // not required
		return nil
	}

	if m.ExtensionAttributes != nil {
		if err := m.ExtensionAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension_attributes")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateFinalPrice(formats strfmt.Registry) error {

	if err := validate.Required("final_price", "body", m.FinalPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateFormattedPrices(formats strfmt.Registry) error {

	if err := validate.Required("formatted_prices", "body", m.FormattedPrices); err != nil {
		return err
	}

	if m.FormattedPrices != nil {
		if err := m.FormattedPrices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formatted_prices")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateMaxPrice(formats strfmt.Registry) error {

	if err := validate.Required("max_price", "body", m.MaxPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateMaxRegularPrice(formats strfmt.Registry) error {

	if err := validate.Required("max_regular_price", "body", m.MaxRegularPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateMinimalPrice(formats strfmt.Registry) error {

	if err := validate.Required("minimal_price", "body", m.MinimalPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateMinimalRegularPrice(formats strfmt.Registry) error {

	if err := validate.Required("minimal_regular_price", "body", m.MinimalRegularPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateRegularPrice(formats strfmt.Registry) error {

	if err := validate.Required("regular_price", "body", m.RegularPrice); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoInterface) validateSpecialPrice(formats strfmt.Registry) error {

	if err := validate.Required("special_price", "body", m.SpecialPrice); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductRenderPriceInfoInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductRenderPriceInfoInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductRenderPriceInfoInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
