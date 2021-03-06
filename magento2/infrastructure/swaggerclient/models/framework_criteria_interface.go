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

// FrameworkCriteriaInterface Interface CriteriaInterface
// swagger:model framework-criteria-interface
type FrameworkCriteriaInterface struct {

	// Criteria objects added to current Composite Criteria
	// Required: true
	CriteriaList []*FrameworkCriteriaInterface `json:"criteria_list"`

	// List of filters
	// Required: true
	Filters []string `json:"filters"`

	// Limit
	// Required: true
	Limit []string `json:"limit"`

	// Associated Mapper Interface name
	// Required: true
	MapperInterfaceName *string `json:"mapper_interface_name"`

	// Ordering criteria
	// Required: true
	Orders []string `json:"orders"`
}

// Validate validates this framework criteria interface
func (m *FrameworkCriteriaInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriteriaList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapperInterfaceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameworkCriteriaInterface) validateCriteriaList(formats strfmt.Registry) error {

	if err := validate.Required("criteria_list", "body", m.CriteriaList); err != nil {
		return err
	}

	for i := 0; i < len(m.CriteriaList); i++ {
		if swag.IsZero(m.CriteriaList[i]) { // not required
			continue
		}

		if m.CriteriaList[i] != nil {
			if err := m.CriteriaList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameworkCriteriaInterface) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	return nil
}

func (m *FrameworkCriteriaInterface) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *FrameworkCriteriaInterface) validateMapperInterfaceName(formats strfmt.Registry) error {

	if err := validate.Required("mapper_interface_name", "body", m.MapperInterfaceName); err != nil {
		return err
	}

	return nil
}

func (m *FrameworkCriteriaInterface) validateOrders(formats strfmt.Registry) error {

	if err := validate.Required("orders", "body", m.Orders); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrameworkCriteriaInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrameworkCriteriaInterface) UnmarshalBinary(b []byte) error {
	var res FrameworkCriteriaInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
