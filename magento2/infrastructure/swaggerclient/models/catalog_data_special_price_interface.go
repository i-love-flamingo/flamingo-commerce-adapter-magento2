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

// CatalogDataSpecialPriceInterface Product Special Price Interface is used to encapsulate data that can be processed by efficient price API.
// swagger:model catalog-data-special-price-interface
type CatalogDataSpecialPriceInterface struct {

	// extension attributes
	ExtensionAttributes CatalogDataSpecialPriceExtensionInterface `json:"extension_attributes,omitempty"`

	// Product special price value.
	// Required: true
	Price *float64 `json:"price"`

	// Start date for special price in Y-m-d H:i:s format.
	// Required: true
	PriceFrom *string `json:"price_from"`

	// End date for special price in Y-m-d H:i:s format.
	// Required: true
	PriceTo *string `json:"price_to"`

	// SKU of product, that contains special price value.
	// Required: true
	Sku *string `json:"sku"`

	// ID of store, that contains special price value.
	// Required: true
	StoreID *int64 `json:"store_id"`
}

// Validate validates this catalog data special price interface
func (m *CatalogDataSpecialPriceInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataSpecialPriceInterface) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataSpecialPriceInterface) validatePriceFrom(formats strfmt.Registry) error {

	if err := validate.Required("price_from", "body", m.PriceFrom); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataSpecialPriceInterface) validatePriceTo(formats strfmt.Registry) error {

	if err := validate.Required("price_to", "body", m.PriceTo); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataSpecialPriceInterface) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataSpecialPriceInterface) validateStoreID(formats strfmt.Registry) error {

	if err := validate.Required("store_id", "body", m.StoreID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataSpecialPriceInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataSpecialPriceInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataSpecialPriceInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
