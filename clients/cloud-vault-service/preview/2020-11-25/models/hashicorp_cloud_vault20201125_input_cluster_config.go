// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125InputClusterConfig hashicorp cloud vault 20201125 input cluster config
//
// swagger:model hashicorp.cloud.vault_20201125.InputClusterConfig
type HashicorpCloudVault20201125InputClusterConfig struct {

	// network_config is the network configuration for the cluster
	NetworkConfig *HashicorpCloudVault20201125InputNetworkConfig `json:"network_config,omitempty"`

	// Tier is the type of Vault cluster that should be provisioned
	Tier HashicorpCloudVault20201125Tier `json:"tier,omitempty"`

	// vault_config is the Vault specific configuration
	VaultConfig *HashicorpCloudVault20201125VaultConfig `json:"vault_config,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 input cluster config
func (m *HashicorpCloudVault20201125InputClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateTier(formats strfmt.Registry) error {

	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateVaultConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.VaultConfig) { // not required
		return nil
	}

	if m.VaultConfig != nil {
		if err := m.VaultConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125InputClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
