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

// HashicorpCloudVault20201125UpdateSnapshotRequest UpdateSnapshotRequest is a request to update a snapshot.
//
// swagger:model hashicorp.cloud.vault_20201125.UpdateSnapshotRequest
type HashicorpCloudVault20201125UpdateSnapshotRequest struct {

	// mask is the mask of fields to update.
	Mask *cloud.GoogleProtobufFieldMask `json:"mask,omitempty"`

	// snapshot contains the fields to update.
	//
	// Supported fields: name
	Snapshot *HashicorpCloudVault20201125Snapshot `json:"snapshot,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 update snapshot request
func (m *HashicorpCloudVault20201125UpdateSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125UpdateSnapshotRequest) validateMask(formats strfmt.Registry) error {

	if swag.IsZero(m.Mask) { // not required
		return nil
	}

	if m.Mask != nil {
		if err := m.Mask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125UpdateSnapshotRequest) validateSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125UpdateSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125UpdateSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125UpdateSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}