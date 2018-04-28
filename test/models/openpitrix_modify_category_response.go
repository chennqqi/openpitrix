// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyCategoryResponse openpitrix modify category response
// swagger:model openpitrixModifyCategoryResponse
type OpenpitrixModifyCategoryResponse struct {

	// category
	Category *OpenpitrixCategory `json:"category,omitempty"`
}

// Validate validates this openpitrix modify category response
func (m *OpenpitrixModifyCategoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixModifyCategoryResponse) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {

		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyCategoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyCategoryResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyCategoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
