// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907AzureRoute hashicorp cloud network 20200907 azure route
//
// swagger:model hashicorp.cloud.network_20200907.AzureRoute
type HashicorpCloudNetwork20200907AzureRoute struct {

	// Only applicable for VIRTUAL_APPLIANCE. Specifies the IP address of the appliance to forward the traffic to.
	NextHopIPAddress string `json:"next_hop_ip_address,omitempty"`

	// The type of Azure hop the packet should be sent to.
	NextHopType *HashicorpCloudNetwork20200907AzureHopType `json:"next_hop_type,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 azure route
func (m *HashicorpCloudNetwork20200907AzureRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextHopType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907AzureRoute) validateNextHopType(formats strfmt.Registry) error {
	if swag.IsZero(m.NextHopType) { // not required
		return nil
	}

	if m.NextHopType != nil {
		if err := m.NextHopType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_hop_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_hop_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 azure route based on the context it is used
func (m *HashicorpCloudNetwork20200907AzureRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextHopType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907AzureRoute) contextValidateNextHopType(ctx context.Context, formats strfmt.Registry) error {

	if m.NextHopType != nil {

		if swag.IsZero(m.NextHopType) { // not required
			return nil
		}

		if err := m.NextHopType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_hop_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_hop_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AzureRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AzureRoute) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907AzureRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
