// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerGetAncestorImagesResponse hashicorp cloud packer get ancestor images response
//
// swagger:model hashicorp.cloud.packer.GetAncestorImagesResponse
type HashicorpCloudPackerGetAncestorImagesResponse struct {

	// ancestors
	Ancestors []*HashicorpCloudPackerIteration `json:"ancestors"`
}

// Validate validates this hashicorp cloud packer get ancestor images response
func (m *HashicorpCloudPackerGetAncestorImagesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAncestors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerGetAncestorImagesResponse) validateAncestors(formats strfmt.Registry) error {

	if swag.IsZero(m.Ancestors) { // not required
		return nil
	}

	for i := 0; i < len(m.Ancestors); i++ {
		if swag.IsZero(m.Ancestors[i]) { // not required
			continue
		}

		if m.Ancestors[i] != nil {
			if err := m.Ancestors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ancestors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerGetAncestorImagesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerGetAncestorImagesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerGetAncestorImagesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
