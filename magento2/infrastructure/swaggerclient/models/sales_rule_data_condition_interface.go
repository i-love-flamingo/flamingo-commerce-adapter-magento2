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

// SalesRuleDataConditionInterface Interface ConditionInterface
// swagger:model sales-rule-data-condition-interface
type SalesRuleDataConditionInterface struct {

	// The aggregator type
	AggregatorType string `json:"aggregator_type,omitempty"`

	// The attribute name of the condition
	AttributeName string `json:"attribute_name,omitempty"`

	// Condition type
	// Required: true
	ConditionType *string `json:"condition_type"`

	// List of conditions
	Conditions []*SalesRuleDataConditionInterface `json:"conditions"`

	// extension attributes
	ExtensionAttributes SalesRuleDataConditionExtensionInterface `json:"extension_attributes,omitempty"`

	// The operator of the condition
	// Required: true
	Operator *string `json:"operator"`

	// The value of the condition
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this sales rule data condition interface
func (m *SalesRuleDataConditionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesRuleDataConditionInterface) validateConditionType(formats strfmt.Registry) error {

	if err := validate.Required("condition_type", "body", m.ConditionType); err != nil {
		return err
	}

	return nil
}

func (m *SalesRuleDataConditionInterface) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SalesRuleDataConditionInterface) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *SalesRuleDataConditionInterface) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesRuleDataConditionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesRuleDataConditionInterface) UnmarshalBinary(b []byte) error {
	var res SalesRuleDataConditionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}