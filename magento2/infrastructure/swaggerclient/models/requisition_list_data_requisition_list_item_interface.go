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

// RequisitionListDataRequisitionListItemInterface Interface RequisitionListItemInterface
// swagger:model requisition-list-data-requisition-list-item-interface
type RequisitionListDataRequisitionListItemInterface struct {

	// Added_at value.
	// Required: true
	AddedAt *string `json:"added_at"`

	// extension attributes
	ExtensionAttributes RequisitionListDataRequisitionListItemExtensionInterface `json:"extension_attributes,omitempty"`

	// Requisition List ID.
	// Required: true
	ID *int64 `json:"id"`

	// Requisition list item options.
	// Required: true
	Options []string `json:"options"`

	// Product Qty.
	// Required: true
	Qty *float64 `json:"qty"`

	// Requisition List ID.
	// Required: true
	RequisitionListID *int64 `json:"requisition_list_id"`

	// Product SKU.
	// Required: true
	Sku *int64 `json:"sku"`

	// Store ID.
	// Required: true
	StoreID *int64 `json:"store_id"`
}

// Validate validates this requisition list data requisition list item interface
func (m *RequisitionListDataRequisitionListItemInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequisitionListID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateAddedAt(formats strfmt.Registry) error {

	if err := validate.Required("added_at", "body", m.AddedAt); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateQty(formats strfmt.Registry) error {

	if err := validate.Required("qty", "body", m.Qty); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateRequisitionListID(formats strfmt.Registry) error {

	if err := validate.Required("requisition_list_id", "body", m.RequisitionListID); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

func (m *RequisitionListDataRequisitionListItemInterface) validateStoreID(formats strfmt.Registry) error {

	if err := validate.Required("store_id", "body", m.StoreID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequisitionListDataRequisitionListItemInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequisitionListDataRequisitionListItemInterface) UnmarshalBinary(b []byte) error {
	var res RequisitionListDataRequisitionListItemInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
