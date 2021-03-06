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

// CompanyDataCompanySearchResultsInterface Interface for company search results
// swagger:model company-data-company-search-results-interface
type CompanyDataCompanySearchResultsInterface struct {

	// Companies list
	// Required: true
	Items []*CompanyDataCompanyInterface `json:"items"`

	// search criteria
	// Required: true
	SearchCriteria *FrameworkSearchCriteriaInterface `json:"search_criteria"`

	// Total count.
	// Required: true
	TotalCount *int64 `json:"total_count"`
}

// Validate validates this company data company search results interface
func (m *CompanyDataCompanySearchResultsInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyDataCompanySearchResultsInterface) validateItems(formats strfmt.Registry) error {

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

func (m *CompanyDataCompanySearchResultsInterface) validateSearchCriteria(formats strfmt.Registry) error {

	if err := validate.Required("search_criteria", "body", m.SearchCriteria); err != nil {
		return err
	}

	if m.SearchCriteria != nil {
		if err := m.SearchCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search_criteria")
			}
			return err
		}
	}

	return nil
}

func (m *CompanyDataCompanySearchResultsInterface) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompanyDataCompanySearchResultsInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyDataCompanySearchResultsInterface) UnmarshalBinary(b []byte) error {
	var res CompanyDataCompanySearchResultsInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
