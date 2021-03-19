// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420SnapshotConfig hashicorp cloud vault 20200420 snapshot config
//
// swagger:model hashicorp.cloud.vault_20200420.SnapshotConfig
type HashicorpCloudVault20200420SnapshotConfig struct {

	// periodic_snapshots_disabled indicates if automated periodic snapshots should be disabled
	PeriodicSnapshotsDisabled bool `json:"periodic_snapshots_disabled,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 snapshot config
func (m *HashicorpCloudVault20200420SnapshotConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420SnapshotConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420SnapshotConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420SnapshotConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
