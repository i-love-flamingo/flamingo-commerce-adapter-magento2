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

// RequisitionListDataRequisitionListInterface Interface RequisitionListInterface
// swagger:model requisition-list-data-requisition-list-interface
type RequisitionListDataRequisitionListInterface struct {

	// Customer ID
	// Required: true
	CustomerID *int64 `json:"customer_id"`

	// Requisition List Description
	// Required: true
	Description *string `json:"description"`

	// extension attributes
	ExtensionAttributes RequisitionListDataRequisitionListExtensionInterface `json:"extension_attributes,omitempty"`

	// Requisition List ID
	// Required: true
	ID *int64 `json:"id"`

	// Requisition List Items
	// Required: true
	Items []*RequisitionListDataRequisitionListItemInterface `json:"items"`

	// Requisition List Name
	// Required: true
	Name *string `json:"name"`

	// Requisition List Update Time
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this requisition list data requisition list interface
func (m *RequisitionListDataRequisitionListInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListInterface) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequisitionListDataRequisitionListInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequisitionListDataRequisitionListInterface) UnmarshalBinary(b []byte) error {
	var res RequisitionListDataRequisitionListInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
