// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SalesDataInvoiceExtensionInterface ExtensionInterface class for @see \Magento\Sales\Api\Data\InvoiceInterface
// swagger:model sales-data-invoice-extension-interface
type SalesDataInvoiceExtensionInterface struct {

	// base customer balance amount
	BaseCustomerBalanceAmount float64 `json:"base_customer_balance_amount,omitempty"`

	// base gift cards amount
	BaseGiftCardsAmount float64 `json:"base_gift_cards_amount,omitempty"`

	// customer balance amount
	CustomerBalanceAmount float64 `json:"customer_balance_amount,omitempty"`

	// gift cards amount
	GiftCardsAmount float64 `json:"gift_cards_amount,omitempty"`

	// gw base price
	GwBasePrice string `json:"gw_base_price,omitempty"`

	// gw base tax amount
	GwBaseTaxAmount string `json:"gw_base_tax_amount,omitempty"`

	// gw card base price
	GwCardBasePrice string `json:"gw_card_base_price,omitempty"`

	// gw card base tax amount
	GwCardBaseTaxAmount string `json:"gw_card_base_tax_amount,omitempty"`

	// gw card price
	GwCardPrice string `json:"gw_card_price,omitempty"`

	// gw card tax amount
	GwCardTaxAmount string `json:"gw_card_tax_amount,omitempty"`

	// gw items base price
	GwItemsBasePrice string `json:"gw_items_base_price,omitempty"`

	// gw items base tax amount
	GwItemsBaseTaxAmount string `json:"gw_items_base_tax_amount,omitempty"`

	// gw items price
	GwItemsPrice string `json:"gw_items_price,omitempty"`

	// gw items tax amount
	GwItemsTaxAmount string `json:"gw_items_tax_amount,omitempty"`

	// gw price
	GwPrice string `json:"gw_price,omitempty"`

	// gw tax amount
	GwTaxAmount string `json:"gw_tax_amount,omitempty"`
}

// Validate validates this sales data invoice extension interface
func (m *SalesDataInvoiceExtensionInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataInvoiceExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataInvoiceExtensionInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataInvoiceExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
