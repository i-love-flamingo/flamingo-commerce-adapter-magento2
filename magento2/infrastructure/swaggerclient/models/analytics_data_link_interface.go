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

// AnalyticsDataLinkInterface Interface LinkInterface Represents link with collected data and initialized vector for decryption.
// swagger:model analytics-data-link-interface
type AnalyticsDataLinkInterface struct {

	// initialization vector
	// Required: true
	InitializationVector *string `json:"initialization_vector"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this analytics data link interface
func (m *AnalyticsDataLinkInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInitializationVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsDataLinkInterface) validateInitializationVector(formats strfmt.Registry) error {

	if err := validate.Required("initialization_vector", "body", m.InitializationVector); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsDataLinkInterface) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsDataLinkInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsDataLinkInterface) UnmarshalBinary(b []byte) error {
	var res AnalyticsDataLinkInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}