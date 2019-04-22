// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CompanyDataCompanyExtensionInterface ExtensionInterface class for @see \Magento\Company\Api\Data\CompanyInterface
// swagger:model company-data-company-extension-interface
type CompanyDataCompanyExtensionInterface struct {

	// applicable payment method
	ApplicablePaymentMethod int64 `json:"applicable_payment_method,omitempty"`

	// available payment methods
	AvailablePaymentMethods string `json:"available_payment_methods,omitempty"`

	// quote config
	QuoteConfig *NegotiableQuoteDataCompanyQuoteConfigInterface `json:"quote_config,omitempty"`

	// use config settings
	UseConfigSettings int64 `json:"use_config_settings,omitempty"`
}

// Validate validates this company data company extension interface
func (m *CompanyDataCompanyExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuoteConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyDataCompanyExtensionInterface) validateQuoteConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteConfig) { // not required
		return nil
	}

	if m.QuoteConfig != nil {
		if err := m.QuoteConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quote_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompanyDataCompanyExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyDataCompanyExtensionInterface) UnmarshalBinary(b []byte) error {
	var res CompanyDataCompanyExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
