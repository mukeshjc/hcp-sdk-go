// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudConsul20210204ClusterState State is the state of the Consul cluster. Note that this state
// represents the abstract Consul cluster itself, not necessarily whether
// Consul cluster is currently available or not.
//
//   - UNSET: UNSET is a sentinel zero value so that an uninitialized value can be
//
// detected.
//   - PENDING: PENDING is the state the cluster is in while it is waiting to be created.
//   - CREATING: CREATING is the state the cluster is in while it is being provisioned for
//
// the first time.
//   - RUNNING: RUNNING is the steady state while the cluster is running.
//   - FAILED: FAILED is a failure state in which the cluster is unavailable and may
//
// required an operator restore action to recover.
//   - UPDATING: UPDATING is the state the cluster is in while undergoing a version
//
// update.
//   - RESTORING: RESTORING is the state the cluster is in while restoring from a snapshot.
//   - DELETING: DELETING is the state the cluster is in while it is being de-provisioned.
//   - DELETED: DELETED is the state the cluster is in when it has been de-provisioned. At
//
// this point, the cluster is eligible for garbage collection.
//
// swagger:model hashicorp.cloud.consul_20210204.Cluster.State
type HashicorpCloudConsul20210204ClusterState string

func NewHashicorpCloudConsul20210204ClusterState(value HashicorpCloudConsul20210204ClusterState) *HashicorpCloudConsul20210204ClusterState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsul20210204ClusterState.
func (m HashicorpCloudConsul20210204ClusterState) Pointer() *HashicorpCloudConsul20210204ClusterState {
	return &m
}

const (

	// HashicorpCloudConsul20210204ClusterStateUNSET captures enum value "UNSET"
	HashicorpCloudConsul20210204ClusterStateUNSET HashicorpCloudConsul20210204ClusterState = "UNSET"

	// HashicorpCloudConsul20210204ClusterStatePENDING captures enum value "PENDING"
	HashicorpCloudConsul20210204ClusterStatePENDING HashicorpCloudConsul20210204ClusterState = "PENDING"

	// HashicorpCloudConsul20210204ClusterStateCREATING captures enum value "CREATING"
	HashicorpCloudConsul20210204ClusterStateCREATING HashicorpCloudConsul20210204ClusterState = "CREATING"

	// HashicorpCloudConsul20210204ClusterStateRUNNING captures enum value "RUNNING"
	HashicorpCloudConsul20210204ClusterStateRUNNING HashicorpCloudConsul20210204ClusterState = "RUNNING"

	// HashicorpCloudConsul20210204ClusterStateFAILED captures enum value "FAILED"
	HashicorpCloudConsul20210204ClusterStateFAILED HashicorpCloudConsul20210204ClusterState = "FAILED"

	// HashicorpCloudConsul20210204ClusterStateUPDATING captures enum value "UPDATING"
	HashicorpCloudConsul20210204ClusterStateUPDATING HashicorpCloudConsul20210204ClusterState = "UPDATING"

	// HashicorpCloudConsul20210204ClusterStateRESTORING captures enum value "RESTORING"
	HashicorpCloudConsul20210204ClusterStateRESTORING HashicorpCloudConsul20210204ClusterState = "RESTORING"

	// HashicorpCloudConsul20210204ClusterStateDELETING captures enum value "DELETING"
	HashicorpCloudConsul20210204ClusterStateDELETING HashicorpCloudConsul20210204ClusterState = "DELETING"

	// HashicorpCloudConsul20210204ClusterStateDELETED captures enum value "DELETED"
	HashicorpCloudConsul20210204ClusterStateDELETED HashicorpCloudConsul20210204ClusterState = "DELETED"
)

// for schema
var hashicorpCloudConsul20210204ClusterStateEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20210204ClusterState
	if err := json.Unmarshal([]byte(`["UNSET","PENDING","CREATING","RUNNING","FAILED","UPDATING","RESTORING","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20210204ClusterStateEnum = append(hashicorpCloudConsul20210204ClusterStateEnum, v)
	}
}

func (m HashicorpCloudConsul20210204ClusterState) validateHashicorpCloudConsul20210204ClusterStateEnum(path, location string, value HashicorpCloudConsul20210204ClusterState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20210204ClusterStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20210204 cluster state
func (m HashicorpCloudConsul20210204ClusterState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20210204ClusterStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20210204 cluster state based on context it is used
func (m HashicorpCloudConsul20210204ClusterState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}