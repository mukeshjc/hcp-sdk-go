// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer hashicorp cloud global network manager 20220215 cluster partition peer
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ClusterPartitionPeer
type HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer struct {

	// display_error is an error to give to the front end to determine on-going issues for customers
	DisplayError *cloud.GoogleRPCStatus `json:"display_error,omitempty"`

	// exported_services_count is the total number of services exported to the peer cluster/partition
	ExportedServicesCount string `json:"exported_services_count,omitempty"`

	// imported_services_count is the total number of services imported from the peer cluster/partition
	ImportedServicesCount string `json:"imported_services_count,omitempty"`

	// last_heartbeat_timestamp is the timestamp the last heartbeat was received from this peer cluster/partition.
	// Format: date-time
	LastHeartbeatTimestamp strfmt.DateTime `json:"last_heartbeat_timestamp,omitempty"`

	// peer_cluster_id is the user settable GNM cluster name of the peer.
	PeerClusterID string `json:"peer_cluster_id,omitempty"`

	// peer_cluster_location is the location of the peer GNM cluster
	PeerClusterLocation *cloud.HashicorpCloudLocationLocation `json:"peer_cluster_location,omitempty"`

	// peer_licensing is the Consul licensing information of the peer GNM cluster
	PeerLicensing *HashicorpCloudGlobalNetworkManager20220215Licensing `json:"peer_licensing,omitempty"`

	// peer_name is the name of the peer
	PeerName string `json:"peer_name,omitempty"`

	// peer_partition_name is the name of the admin partition on the peer cluster
	PeerPartitionName string `json:"peer_partition_name,omitempty"`

	// peering_connection_id is the slug ID of the GNM peering connection
	PeeringConnectionID string `json:"peering_connection_id,omitempty"`

	// peering_connection_location is the location of the GNM peering connections
	PeeringConnectionLocation *cloud.HashicorpCloudLocationLocation `json:"peering_connection_location,omitempty"`

	// peering_connection_status is the overall status of the peering connection
	PeeringConnectionStatus *HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus `json:"peering_connection_status,omitempty"`

	// peering_status is the state of this individual peer
	PeeringStatus *HashicorpCloudGlobalNetworkManager20220215PeerState `json:"peering_status,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 cluster partition peer
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastHeartbeatTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerClusterLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerLicensing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeeringConnectionLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeeringConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeeringStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validateDisplayError(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayError) { // not required
		return nil
	}

	if m.DisplayError != nil {
		if err := m.DisplayError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validateLastHeartbeatTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastHeartbeatTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("last_heartbeat_timestamp", "body", "date-time", m.LastHeartbeatTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validatePeerClusterLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerClusterLocation) { // not required
		return nil
	}

	if m.PeerClusterLocation != nil {
		if err := m.PeerClusterLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster_location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validatePeerLicensing(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerLicensing) { // not required
		return nil
	}

	if m.PeerLicensing != nil {
		if err := m.PeerLicensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validatePeeringConnectionLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.PeeringConnectionLocation) { // not required
		return nil
	}

	if m.PeeringConnectionLocation != nil {
		if err := m.PeeringConnectionLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_connection_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_connection_location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validatePeeringConnectionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PeeringConnectionStatus) { // not required
		return nil
	}

	if m.PeeringConnectionStatus != nil {
		if err := m.PeeringConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_connection_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) validatePeeringStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PeeringStatus) { // not required
		return nil
	}

	if m.PeeringStatus != nil {
		if err := m.PeeringStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 cluster partition peer based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerClusterLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerLicensing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeeringConnectionLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeeringConnectionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeeringStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidateDisplayError(ctx context.Context, formats strfmt.Registry) error {

	if m.DisplayError != nil {

		if swag.IsZero(m.DisplayError) { // not required
			return nil
		}

		if err := m.DisplayError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidatePeerClusterLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.PeerClusterLocation != nil {

		if swag.IsZero(m.PeerClusterLocation) { // not required
			return nil
		}

		if err := m.PeerClusterLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster_location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidatePeerLicensing(ctx context.Context, formats strfmt.Registry) error {

	if m.PeerLicensing != nil {

		if swag.IsZero(m.PeerLicensing) { // not required
			return nil
		}

		if err := m.PeerLicensing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidatePeeringConnectionLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.PeeringConnectionLocation != nil {

		if swag.IsZero(m.PeeringConnectionLocation) { // not required
			return nil
		}

		if err := m.PeeringConnectionLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_connection_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_connection_location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidatePeeringConnectionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PeeringConnectionStatus != nil {

		if swag.IsZero(m.PeeringConnectionStatus) { // not required
			return nil
		}

		if err := m.PeeringConnectionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_connection_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) contextValidatePeeringStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PeeringStatus != nil {

		if swag.IsZero(m.PeeringStatus) { // not required
			return nil
		}

		if err := m.PeeringStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
