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

// SalesDataCreditmemoItemInterface Credit memo item interface. After a customer places and pays for an order and an invoice has been issued, the merchant can create a credit memo to refund all or part of the amount paid for any returned or undelivered items. The memo restores funds to the customer account so that the customer can make future purchases. A credit memo item is an invoiced item for which a merchant creates a credit memo.
// swagger:model sales-data-creditmemo-item-interface
type SalesDataCreditmemoItemInterface struct {

	// Additional data.
	AdditionalData string `json:"additional_data,omitempty"`

	// The base cost for a credit memo item.
	// Required: true
	BaseCost *float64 `json:"base_cost"`

	// The base discount amount for a credit memo item.
	BaseDiscountAmount float64 `json:"base_discount_amount,omitempty"`

	// The base discount tax compensation amount for a credit memo item.
	BaseDiscountTaxCompensationAmount float64 `json:"base_discount_tax_compensation_amount,omitempty"`

	// The base price for a credit memo item.
	// Required: true
	BasePrice *float64 `json:"base_price"`

	// Base price including tax.
	BasePriceInclTax float64 `json:"base_price_incl_tax,omitempty"`

	// Base row total.
	BaseRowTotal float64 `json:"base_row_total,omitempty"`

	// Base row total including tax.
	BaseRowTotalInclTax float64 `json:"base_row_total_incl_tax,omitempty"`

	// Base tax amount.
	BaseTaxAmount float64 `json:"base_tax_amount,omitempty"`

	// Base WEEE tax applied amount.
	BaseWeeeTaxAppliedAmount float64 `json:"base_weee_tax_applied_amount,omitempty"`

	// Base WEEE tax applied row amount.
	BaseWeeeTaxAppliedRowAmnt float64 `json:"base_weee_tax_applied_row_amnt,omitempty"`

	// Base WEEE tax disposition.
	BaseWeeeTaxDisposition float64 `json:"base_weee_tax_disposition,omitempty"`

	// Base WEEE tax row disposition.
	BaseWeeeTaxRowDisposition float64 `json:"base_weee_tax_row_disposition,omitempty"`

	// Description.
	Description string `json:"description,omitempty"`

	// Discount amount.
	DiscountAmount float64 `json:"discount_amount,omitempty"`

	// Discount tax compensation amount.
	DiscountTaxCompensationAmount float64 `json:"discount_tax_compensation_amount,omitempty"`

	// Credit memo item ID.
	// Required: true
	EntityID *int64 `json:"entity_id"`

	// extension attributes
	ExtensionAttributes SalesDataCreditmemoItemExtensionInterface `json:"extension_attributes,omitempty"`

	// Name.
	Name string `json:"name,omitempty"`

	// Order item ID.
	// Required: true
	OrderItemID *int64 `json:"order_item_id"`

	// Parent ID.
	ParentID int64 `json:"parent_id,omitempty"`

	// Price.
	Price float64 `json:"price,omitempty"`

	// Price including tax.
	PriceInclTax float64 `json:"price_incl_tax,omitempty"`

	// Product ID.
	ProductID int64 `json:"product_id,omitempty"`

	// Quantity.
	// Required: true
	Qty *float64 `json:"qty"`

	// Row total.
	RowTotal float64 `json:"row_total,omitempty"`

	// Row total including tax.
	RowTotalInclTax float64 `json:"row_total_incl_tax,omitempty"`

	// SKU.
	Sku string `json:"sku,omitempty"`

	// Tax amount.
	TaxAmount float64 `json:"tax_amount,omitempty"`

	// WEEE tax applied.
	WeeeTaxApplied string `json:"weee_tax_applied,omitempty"`

	// WEEE tax applied amount.
	WeeeTaxAppliedAmount float64 `json:"weee_tax_applied_amount,omitempty"`

	// WEEE tax applied row amount.
	WeeeTaxAppliedRowAmount float64 `json:"weee_tax_applied_row_amount,omitempty"`

	// WEEE tax disposition.
	WeeeTaxDisposition float64 `json:"weee_tax_disposition,omitempty"`

	// WEEE tax row disposition.
	WeeeTaxRowDisposition float64 `json:"weee_tax_row_disposition,omitempty"`
}

// Validate validates this sales data creditmemo item interface
func (m *SalesDataCreditmemoItemInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

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

func (m *SalesDataCreditmemoItemInterface) validateBaseCost(formats strfmt.Registry) error {

	if err := validate.Required("base_cost", "body", m.BaseCost); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataCreditmemoItemInterface) validateBasePrice(formats strfmt.Registry) error {

	if err := validate.Required("base_price", "body", m.BasePrice); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataCreditmemoItemInterface) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entity_id", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataCreditmemoItemInterface) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("order_item_id", "body", m.OrderItemID); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataCreditmemoItemInterface) validateQty(formats strfmt.Registry) error {

	if err := validate.Required("qty", "body", m.Qty); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataCreditmemoItemInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataCreditmemoItemInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataCreditmemoItemInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
