// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SalesDataOrderItemExtensionInterface ExtensionInterface class for @see \Magento\Sales\Api\Data\OrderItemInterface
// swagger:model sales-data-order-item-extension-interface
type SalesDataOrderItemExtensionInterface struct {

	// gift message
	GiftMessage *GiftMessageDataMessageInterface `json:"gift_message,omitempty"`

	// gw base price
	GwBasePrice string `json:"gw_base_price,omitempty"`

	// gw base price invoiced
	GwBasePriceInvoiced string `json:"gw_base_price_invoiced,omitempty"`

	// gw base price refunded
	GwBasePriceRefunded string `json:"gw_base_price_refunded,omitempty"`

	// gw base tax amount
	GwBaseTaxAmount string `json:"gw_base_tax_amount,omitempty"`

	// gw base tax amount invoiced
	GwBaseTaxAmountInvoiced string `json:"gw_base_tax_amount_invoiced,omitempty"`

	// gw base tax amount refunded
	GwBaseTaxAmountRefunded string `json:"gw_base_tax_amount_refunded,omitempty"`

	// gw id
	GwID string `json:"gw_id,omitempty"`

	// gw price
	GwPrice string `json:"gw_price,omitempty"`

	// gw price invoiced
	GwPriceInvoiced string `json:"gw_price_invoiced,omitempty"`

	// gw price refunded
	GwPriceRefunded string `json:"gw_price_refunded,omitempty"`

	// gw tax amount
	GwTaxAmount string `json:"gw_tax_amount,omitempty"`

	// gw tax amount invoiced
	GwTaxAmountInvoiced string `json:"gw_tax_amount_invoiced,omitempty"`

	// gw tax amount refunded
	GwTaxAmountRefunded string `json:"gw_tax_amount_refunded,omitempty"`
}

// Validate validates this sales data order item extension interface
func (m *SalesDataOrderItemExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesDataOrderItemExtensionInterface) validateGiftMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.GiftMessage) { // not required
		return nil
	}

	if m.GiftMessage != nil {
		if err := m.GiftMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gift_message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataOrderItemExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataOrderItemExtensionInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataOrderItemExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
