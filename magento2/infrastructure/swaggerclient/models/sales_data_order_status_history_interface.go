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

// SalesDataOrderStatusHistoryInterface Order status history interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.
// swagger:model sales-data-order-status-history-interface
type SalesDataOrderStatusHistoryInterface struct {

	// Comment.
	// Required: true
	Comment *string `json:"comment"`

	// Created-at timestamp.
	CreatedAt string `json:"created_at,omitempty"`

	// Order status history ID.
	EntityID int64 `json:"entity_id,omitempty"`

	// Entity name.
	EntityName string `json:"entity_name,omitempty"`

	// extension attributes
	ExtensionAttributes SalesDataOrderStatusHistoryExtensionInterface `json:"extension_attributes,omitempty"`

	// Is-customer-notified flag value.
	// Required: true
	IsCustomerNotified *int64 `json:"is_customer_notified"`

	// Is-visible-on-storefront flag value.
	// Required: true
	IsVisibleOnFront *int64 `json:"is_visible_on_front"`

	// Parent ID.
	// Required: true
	ParentID *int64 `json:"parent_id"`

	// Status.
	Status string `json:"status,omitempty"`
}

// Validate validates this sales data order status history interface
func (m *SalesDataOrderStatusHistoryInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsCustomerNotified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsVisibleOnFront(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesDataOrderStatusHistoryInterface) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataOrderStatusHistoryInterface) validateIsCustomerNotified(formats strfmt.Registry) error {

	if err := validate.Required("is_customer_notified", "body", m.IsCustomerNotified); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataOrderStatusHistoryInterface) validateIsVisibleOnFront(formats strfmt.Registry) error {

	if err := validate.Required("is_visible_on_front", "body", m.IsVisibleOnFront); err != nil {
		return err
	}

	return nil
}

func (m *SalesDataOrderStatusHistoryInterface) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesDataOrderStatusHistoryInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesDataOrderStatusHistoryInterface) UnmarshalBinary(b []byte) error {
	var res SalesDataOrderStatusHistoryInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
