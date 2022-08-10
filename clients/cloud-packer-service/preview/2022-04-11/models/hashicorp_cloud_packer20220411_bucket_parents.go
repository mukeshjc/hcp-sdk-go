// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411BucketParents hashicorp cloud packer 20220411 bucket parents
//
// swagger:model hashicorp.cloud.packer_20220411.BucketParents
type HashicorpCloudPacker20220411BucketParents struct {

	// The URL for the Bucket's ancestors endpoint that takes to the previous level of the ancestry tree listing all
	// parent buckets.
	Href string `json:"href,omitempty"`

	// The parents' overall status. If at least one parent is out of date, the overall status will be 'OUT_OF_DATE'.
	Status HashicorpCloudPacker20220411BucketAncestryStatus `json:"status,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 bucket parents
func (m *HashicorpCloudPacker20220411BucketParents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411BucketParents) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411BucketParents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411BucketParents) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411BucketParents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
