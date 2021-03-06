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

// SalesDataShipmentItemInterface Shipment item interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package. A product is an item in a shipment.
// swagger:model sales-data-shipment-item-interface
type SalesDataShipmentItemInterface struct {

	// Additional data.
	AdditionalData string `json:"additional_data,omitempty"`

	// Description.
	Description string `json:"description,omitempty"`

	// Shipment item ID.
	EntityID int64 `json:"entity_id,omitempty"`

	// extension attributes
	ExtensionAttributes SalesDataShipmentItemExtensionInterface `json:"extension_attributes,omitempty"`

	// Name.
	Name string `json:"name,omitempty"`

	// Order item ID.
	// Required: true
	OrderItemID *int64 `json:"order_item_id"`

	// Parent ID.
	ParentID int64 `json:"parent_id,omitempty"`

	// Price.
	Price float64 `json:"price,omitempty"`

	// Product ID.
	ProductID int64 `json:"product_id,omitempty"`

	// Quantity.
	// Required: true
	Qty *float64 `json:"qty"`

	// Row total.
	RowTotal float64 `json:"row_total,omitempty"`

	// SKU.
	Sku string `json:"sku,omitempty"`

	// Weight.
	Weight float64 `json:"weight,omitempty"`
}

// Validate validates this sales data shipment item interface
func (m *SalesDataShipmentItemInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesDataShipmentItemInterface) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("order_item_id", "body", m.OrderItemID); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataShipmentItemInterface) validateQty(formats strfmt.Registry) error {

	if err := validate.Required("qty", "body", m.Qty); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataShipmentItemInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataShipmentItemInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataShipmentItemInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
