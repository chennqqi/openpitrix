// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteAppVersionsResponse openpitrix delete app versions response
// swagger:model openpitrixDeleteAppVersionsResponse
type OpenpitrixDeleteAppVersionsResponse struct {

	// version id
	VersionID []string `json:"version_id"`
}

// Validate validates this openpitrix delete app versions response
func (m *OpenpitrixDeleteAppVersionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteAppVersionsResponse) validateVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteAppVersionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteAppVersionsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteAppVersionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
