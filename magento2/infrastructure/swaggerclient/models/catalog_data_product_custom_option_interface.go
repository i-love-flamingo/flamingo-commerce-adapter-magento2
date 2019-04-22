// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogDataProductCustomOptionInterface catalog data product custom option interface
// swagger:model catalog-data-product-custom-option-interface
type CatalogDataProductCustomOptionInterface struct {

	// extension attributes
	ExtensionAttributes CatalogDataProductCustomOptionExtensionInterface `json:"extension_attributes,omitempty"`

	// file extension
	FileExtension string `json:"file_extension,omitempty"`

	// image size x
	ImageSizeX int64 `json:"image_size_x,omitempty"`

	// image size y
	ImageSizeY int64 `json:"image_size_y,omitempty"`

	// Is require
	// Required: true
	IsRequire *bool `json:"is_require"`

	// max characters
	MaxCharacters int64 `json:"max_characters,omitempty"`

	// Option id
	OptionID int64 `json:"option_id,omitempty"`

	// Price
	Price float64 `json:"price,omitempty"`

	// Price type
	PriceType string `json:"price_type,omitempty"`

	// Product SKU
	// Required: true
	ProductSku *string `json:"product_sku"`

	// Sku
	Sku string `json:"sku,omitempty"`

	// Sort order
	// Required: true
	SortOrder *int64 `json:"sort_order"`

	// Option title
	// Required: true
	Title *string `json:"title"`

	// Option type
	// Required: true
	Type *string `json:"type"`

	// values
	Values []*CatalogDataProductCustomOptionValuesInterface `json:"values"`
}

// Validate validates this catalog data product custom option interface
func (m *CatalogDataProductCustomOptionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsRequire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateIsRequire(formats strfmt.Registry) error {

	if err := validate.Required("is_require", "body", m.IsRequire); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateProductSku(formats strfmt.Registry) error {

	if err := validate.Required("product_sku", "body", m.ProductSku); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateSortOrder(formats strfmt.Registry) error {

	if err := validate.Required("sort_order", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductCustomOptionInterface) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductCustomOptionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductCustomOptionInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductCustomOptionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
