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

// CompanyDataRoleInterface Role data transfer object interface.
// swagger:model company-data-role-interface
type CompanyDataRoleInterface struct {

	// Company id.
	CompanyID int64 `json:"company_id,omitempty"`

	// extension attributes
	ExtensionAttributes CompanyDataRoleExtensionInterface `json:"extension_attributes,omitempty"`

	// Role id.
	ID int64 `json:"id,omitempty"`

	// Permissions.
	// Required: true
	Permissions []*CompanyDataPermissionInterface `json:"permissions"`

	// Role name.
	RoleName string `json:"role_name,omitempty"`
}

// Validate validates this company data role interface
func (m *CompanyDataRoleInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyDataRoleInterface) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompanyDataRoleInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyDataRoleInterface) UnmarshalBinary(b []byte) error {
	var res CompanyDataRoleInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
