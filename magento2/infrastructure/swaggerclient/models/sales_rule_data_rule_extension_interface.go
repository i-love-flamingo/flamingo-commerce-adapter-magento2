// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SalesRuleDataRuleExtensionInterface ExtensionInterface class for @see \Magento\SalesRule\Api\Data\RuleInterface
// swagger:model sales-rule-data-rule-extension-interface
type SalesRuleDataRuleExtensionInterface struct {

	// reward points delta
	RewardPointsDelta int64 `json:"reward_points_delta,omitempty"`
}

// Validate validates this sales rule data rule extension interface
func (m *SalesRuleDataRuleExtensionInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesRuleDataRuleExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesRuleDataRuleExtensionInterface) UnmarshalBinary(b []byte) error {
	var res SalesRuleDataRuleExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
