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

// CheckoutAgreementsDataAgreementInterface Interface AgreementInterface
// swagger:model checkout-agreements-data-agreement-interface
type CheckoutAgreementsDataAgreementInterface struct {

	// Agreement ID.
	// Required: true
	AgreementID *int64 `json:"agreement_id"`

	// Agreement checkbox text.
	// Required: true
	CheckboxText *string `json:"checkbox_text"`

	// Agreement content.
	// Required: true
	Content *string `json:"content"`

	// Agreement content height. Otherwise, null.
	ContentHeight string `json:"content_height,omitempty"`

	// extension attributes
	ExtensionAttributes CheckoutAgreementsDataAgreementExtensionInterface `json:"extension_attributes,omitempty"`

	// Agreement status.
	// Required: true
	IsActive *bool `json:"is_active"`

	// * true - HTML. * false - plain text.
	// Required: true
	IsHTML *bool `json:"is_html"`

	// The agreement applied mode.
	// Required: true
	Mode *int64 `json:"mode"`

	// Agreement name.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this checkout agreements data agreement interface
func (m *CheckoutAgreementsDataAgreementInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreementID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckboxText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsHTML(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateAgreementID(formats strfmt.Registry) error {

	if err := validate.Required("agreement_id", "body", m.AgreementID); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateCheckboxText(formats strfmt.Registry) error {

	if err := validate.Required("checkbox_text", "body", m.CheckboxText); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateIsHTML(formats strfmt.Registry) error {

	if err := validate.Required("is_html", "body", m.IsHTML); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *CheckoutAgreementsDataAgreementInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckoutAgreementsDataAgreementInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckoutAgreementsDataAgreementInterface) UnmarshalBinary(b []byte) error {
	var res CheckoutAgreementsDataAgreementInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
