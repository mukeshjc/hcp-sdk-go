// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907AzurePeeringTarget PeeringTarget is Azure specific details about a target for network peering
//
// swagger:model hashicorp.cloud.network_20200907.Azure.PeeringTarget
type HashicorpCloudNetwork20200907AzurePeeringTarget struct {

	// application_id is the ID of the Active Directory application
	// registered for a single peering request.
	ApplicationID string `json:"application_id,omitempty"`

	// region is the region where the VNet is located.
	Region string `json:"region,omitempty"`

	// resource_group_name is the name of the resource group where the VNet
	// is located.
	ResourceGroupName string `json:"resource_group_name,omitempty"`

	// subscription_id is the ID of the customer subscription where the VNet
	// is located.
	SubscriptionID string `json:"subscription_id,omitempty"`

	// tenant_id is the ID of the customer tenant where the VNet is located.
	TenantID string `json:"tenant_id,omitempty"`

	// vnet_name is the name of the VNet.
	VnetName string `json:"vnet_name,omitempty"`

	AllowForwardedTraffic bool `json:"allow_forwarded_traffic,omitempty"`
	UseRemoteGateways     bool `json:"use_remote_gateways,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 azure peering target
func (m *HashicorpCloudNetwork20200907AzurePeeringTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 azure peering target based on context it is used
func (m *HashicorpCloudNetwork20200907AzurePeeringTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AzurePeeringTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AzurePeeringTarget) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907AzurePeeringTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
