// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411GetBuildResponse hashicorp cloud packer 20220411 get build response
//
// swagger:model hashicorp.cloud.packer_20220411.GetBuildResponse
type HashicorpCloudPacker20220411GetBuildResponse struct {

	// The requested build.
	Build *HashicorpCloudPacker20220411Build `json:"build,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 get build response
func (m *HashicorpCloudPacker20220411GetBuildResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411GetBuildResponse) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411GetBuildResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411GetBuildResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411GetBuildResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
