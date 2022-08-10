// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudPacker20220411CreateBucketRequest hashicorp cloud packer 20220411 create bucket request
//
// swagger:model hashicorp.cloud.packer_20220411.CreateBucketRequest
type HashicorpCloudPacker20220411CreateBucketRequest struct {

	// Human-readable name for the bucket.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// A short description of what this bucket's images are for.
	Description string `json:"description,omitempty"`

	// A key:value map for custom, user-settable metadata about your bucket.
	Labels map[string]string `json:"labels,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 create bucket request
func (m *HashicorpCloudPacker20220411CreateBucketRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411CreateBucketRequest) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411CreateBucketRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411CreateBucketRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411CreateBucketRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
