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

// CustomerDataValidationResultsInterface Validation results interface.
// swagger:model customer-data-validation-results-interface
type CustomerDataValidationResultsInterface struct {

	// Error messages as array in case of validation failure, else return empty array.
	// Required: true
	Messages []string `json:"messages"`

	// If the provided data is valid.
	// Required: true
	Valid *bool `json:"valid"`
}

// Validate validates this customer data validation results interface
func (m *CustomerDataValidationResultsInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerDataValidationResultsInterface) validateMessages(formats strfmt.Registry) error {

	if err := validate.Required("messages", "body", m.Messages); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDataValidationResultsInterface) validateValid(formats strfmt.Registry) error {

	if err := validate.Required("valid", "body", m.Valid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerDataValidationResultsInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerDataValidationResultsInterface) UnmarshalBinary(b []byte) error {
	var res CustomerDataValidationResultsInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}