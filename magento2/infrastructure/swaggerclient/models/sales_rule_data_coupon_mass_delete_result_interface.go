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

// SalesRuleDataCouponMassDeleteResultInterface Coupon mass delete results interface.
// swagger:model sales-rule-data-coupon-mass-delete-result-interface
type SalesRuleDataCouponMassDeleteResultInterface struct {

	// List of failed items.
	// Required: true
	FailedItems []string `json:"failed_items"`

	// List of missing items.
	// Required: true
	MissingItems []string `json:"missing_items"`
}

// Validate validates this sales rule data coupon mass delete result interface
func (m *SalesRuleDataCouponMassDeleteResultInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissingItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesRuleDataCouponMassDeleteResultInterface) validateFailedItems(formats strfmt.Registry) error {

	if err := validate.Required("failed_items", "body", m.FailedItems); err != nil {
		return err
	}

	return nil
}

func (m *SalesRuleDataCouponMassDeleteResultInterface) validateMissingItems(formats strfmt.Registry) error {

	if err := validate.Required("missing_items", "body", m.MissingItems); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesRuleDataCouponMassDeleteResultInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesRuleDataCouponMassDeleteResultInterface) UnmarshalBinary(b []byte) error {
	var res SalesRuleDataCouponMassDeleteResultInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}