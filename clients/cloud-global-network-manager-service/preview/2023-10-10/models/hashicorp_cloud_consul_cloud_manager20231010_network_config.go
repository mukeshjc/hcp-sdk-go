// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulCloudManager20231010NetworkConfig NetworkConfig configures the network of the Consul cluster.
//
// swagger:model hashicorp.cloud.consul_cloud_manager_20231010.NetworkConfig
type HashicorpCloudConsulCloudManager20231010NetworkConfig struct {

	// A list of IP addresses used to restrict access to a cluster.
	IPAllowlist []*HashicorpCloudConsulCloudManager20231010CidrRange `json:"ip_allowlist"`

	// `private` indicates if this cluster's instances should be configured in a
	// non-externally accessible way.
	Private bool `json:"private,omitempty"`

	// This parameter is the network resource's `id`.
	ResourceID string `json:"resource_id,omitempty"`

	// This parameter is the network resource's `name`.
	ResourceName string `json:"resource_name,omitempty"`
}

// Validate validates this hashicorp cloud consul cloud manager 20231010 network config
func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllowlist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) validateIPAllowlist(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAllowlist) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAllowlist); i++ {
		if swag.IsZero(m.IPAllowlist[i]) { // not required
			continue
		}

		if m.IPAllowlist[i] != nil {
			if err := m.IPAllowlist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul cloud manager 20231010 network config based on the context it is used
func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAllowlist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) contextValidateIPAllowlist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAllowlist); i++ {

		if m.IPAllowlist[i] != nil {
			if err := m.IPAllowlist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010NetworkConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulCloudManager20231010NetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}