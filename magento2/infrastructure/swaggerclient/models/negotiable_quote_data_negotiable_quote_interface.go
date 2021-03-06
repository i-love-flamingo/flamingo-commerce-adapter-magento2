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

// NegotiableQuoteDataNegotiableQuoteInterface Interface NegotiableQuoteInterface
// swagger:model negotiable-quote-data-negotiable-quote-interface
type NegotiableQuoteDataNegotiableQuoteInterface struct {

	// Quote rules.
	// Required: true
	AppliedRuleIds *string `json:"applied_rule_ids"`

	// Quote negotiated total price in base currency.
	BaseNegotiatedTotalPrice float64 `json:"base_negotiated_total_price,omitempty"`

	// Quote original total price in base currency.
	BaseOriginalTotalPrice float64 `json:"base_original_total_price,omitempty"`

	// Quote creator id.
	// Required: true
	CreatorID *int64 `json:"creator_id"`

	// Quote creator type.
	// Required: true
	CreatorType *int64 `json:"creator_type"`

	// Deleted products sku.
	// Required: true
	DeletedSku *string `json:"deleted_sku"`

	// Email notification status.
	// Required: true
	EmailNotificationStatus *int64 `json:"email_notification_status"`

	// Expiration period.
	// Required: true
	ExpirationPeriod *string `json:"expiration_period"`

	// extension attributes
	ExtensionAttributes NegotiableQuoteDataNegotiableQuoteExtensionInterface `json:"extension_attributes,omitempty"`

	// Has unconfirmed changes.
	// Required: true
	HasUnconfirmedChanges *bool `json:"has_unconfirmed_changes"`

	// Is address draft.
	// Required: true
	IsAddressDraft *bool `json:"is_address_draft"`

	// Customer price changes.
	// Required: true
	IsCustomerPriceChanged *bool `json:"is_customer_price_changed"`

	// Is regular quote.
	// Required: true
	IsRegularQuote *bool `json:"is_regular_quote"`

	// Shipping tax changes.
	// Required: true
	IsShippingTaxChanged *bool `json:"is_shipping_tax_changed"`

	// Negotiated price type.
	// Required: true
	NegotiatedPriceType *int64 `json:"negotiated_price_type"`

	// Negotiated price value.
	// Required: true
	NegotiatedPriceValue *float64 `json:"negotiated_price_value"`

	// Quote negotiated total price.
	NegotiatedTotalPrice float64 `json:"negotiated_total_price,omitempty"`

	// Quote notifications.
	// Required: true
	Notifications *int64 `json:"notifications"`

	// Quote original total price.
	OriginalTotalPrice float64 `json:"original_total_price,omitempty"`

	// Negotiable quote ID.
	// Required: true
	QuoteID *int64 `json:"quote_id"`

	// Negotiable quote name.
	// Required: true
	QuoteName *string `json:"quote_name"`

	// Proposed shipping price.
	// Required: true
	ShippingPrice *float64 `json:"shipping_price"`

	// Negotiable quote status.
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this negotiable quote data negotiable quote interface
func (m *NegotiableQuoteDataNegotiableQuoteInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppliedRuleIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailNotificationStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasUnconfirmedChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAddressDraft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsCustomerPriceChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRegularQuote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsShippingTaxChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegotiatedPriceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegotiatedPriceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateAppliedRuleIds(formats strfmt.Registry) error {

	if err := validate.Required("applied_rule_ids", "body", m.AppliedRuleIds); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("creator_id", "body", m.CreatorID); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateCreatorType(formats strfmt.Registry) error {

	if err := validate.Required("creator_type", "body", m.CreatorType); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateDeletedSku(formats strfmt.Registry) error {

	if err := validate.Required("deleted_sku", "body", m.DeletedSku); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateEmailNotificationStatus(formats strfmt.Registry) error {

	if err := validate.Required("email_notification_status", "body", m.EmailNotificationStatus); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateExpirationPeriod(formats strfmt.Registry) error {

	if err := validate.Required("expiration_period", "body", m.ExpirationPeriod); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateHasUnconfirmedChanges(formats strfmt.Registry) error {

	if err := validate.Required("has_unconfirmed_changes", "body", m.HasUnconfirmedChanges); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateIsAddressDraft(formats strfmt.Registry) error {

	if err := validate.Required("is_address_draft", "body", m.IsAddressDraft); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateIsCustomerPriceChanged(formats strfmt.Registry) error {

	if err := validate.Required("is_customer_price_changed", "body", m.IsCustomerPriceChanged); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateIsRegularQuote(formats strfmt.Registry) error {

	if err := validate.Required("is_regular_quote", "body", m.IsRegularQuote); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateIsShippingTaxChanged(formats strfmt.Registry) error {

	if err := validate.Required("is_shipping_tax_changed", "body", m.IsShippingTaxChanged); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateNegotiatedPriceType(formats strfmt.Registry) error {

	if err := validate.Required("negotiated_price_type", "body", m.NegotiatedPriceType); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateNegotiatedPriceValue(formats strfmt.Registry) error {

	if err := validate.Required("negotiated_price_value", "body", m.NegotiatedPriceValue); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateNotifications(formats strfmt.Registry) error {

	if err := validate.Required("notifications", "body", m.Notifications); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateQuoteID(formats strfmt.Registry) error {

	if err := validate.Required("quote_id", "body", m.QuoteID); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateQuoteName(formats strfmt.Registry) error {

	if err := validate.Required("quote_name", "body", m.QuoteName); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateShippingPrice(formats strfmt.Registry) error {

	if err := validate.Required("shipping_price", "body", m.ShippingPrice); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataNegotiableQuoteInterface) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NegotiableQuoteDataNegotiableQuoteInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NegotiableQuoteDataNegotiableQuoteInterface) UnmarshalBinary(b []byte) error {
	var res NegotiableQuoteDataNegotiableQuoteInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
