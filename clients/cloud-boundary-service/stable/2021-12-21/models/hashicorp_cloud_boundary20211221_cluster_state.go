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

// HashicorpCloudBoundary20211221ClusterState State is the state of the HCP Boundary cluster.
//
//   - STATE_INVALID_UNSPECIFIED: STATE_INVALID_UNSPECIFIED is used when state is not valid.
//   - STATE_PENDING: STATE_PENDING is used when cluster is waiting for creation.
//   - STATE_CREATING: STATE_CREATING is used when cluster is creating.
//   - STATE_RUNNING: STATE_RUNNING is used when cluster is up and running.
//   - STATE_FAILED: STATE_FAILED is used when cluster failed to be created.
//   - STATE_DELETING: STATE_DELETING is used when cluster received a deletion request.
//   - STATE_DELETED: STATE_DELETED is used when cluster finished deleting.
//   - STATE_UPDATING: STATE_UPDATING is used when cluster finished updating.
//
// swagger:model hashicorp.cloud.boundary_20211221.Cluster.State
type HashicorpCloudBoundary20211221ClusterState string

func NewHashicorpCloudBoundary20211221ClusterState(value HashicorpCloudBoundary20211221ClusterState) *HashicorpCloudBoundary20211221ClusterState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudBoundary20211221ClusterState.
func (m HashicorpCloudBoundary20211221ClusterState) Pointer() *HashicorpCloudBoundary20211221ClusterState {
	return &m
}

const (

	// HashicorpCloudBoundary20211221ClusterStateSTATEINVALIDUNSPECIFIED captures enum value "STATE_INVALID_UNSPECIFIED"
	HashicorpCloudBoundary20211221ClusterStateSTATEINVALIDUNSPECIFIED HashicorpCloudBoundary20211221ClusterState = "STATE_INVALID_UNSPECIFIED"

	// HashicorpCloudBoundary20211221ClusterStateSTATEPENDING captures enum value "STATE_PENDING"
	HashicorpCloudBoundary20211221ClusterStateSTATEPENDING HashicorpCloudBoundary20211221ClusterState = "STATE_PENDING"

	// HashicorpCloudBoundary20211221ClusterStateSTATECREATING captures enum value "STATE_CREATING"
	HashicorpCloudBoundary20211221ClusterStateSTATECREATING HashicorpCloudBoundary20211221ClusterState = "STATE_CREATING"

	// HashicorpCloudBoundary20211221ClusterStateSTATERUNNING captures enum value "STATE_RUNNING"
	HashicorpCloudBoundary20211221ClusterStateSTATERUNNING HashicorpCloudBoundary20211221ClusterState = "STATE_RUNNING"

	// HashicorpCloudBoundary20211221ClusterStateSTATEFAILED captures enum value "STATE_FAILED"
	HashicorpCloudBoundary20211221ClusterStateSTATEFAILED HashicorpCloudBoundary20211221ClusterState = "STATE_FAILED"

	// HashicorpCloudBoundary20211221ClusterStateSTATEDELETING captures enum value "STATE_DELETING"
	HashicorpCloudBoundary20211221ClusterStateSTATEDELETING HashicorpCloudBoundary20211221ClusterState = "STATE_DELETING"

	// HashicorpCloudBoundary20211221ClusterStateSTATEDELETED captures enum value "STATE_DELETED"
	HashicorpCloudBoundary20211221ClusterStateSTATEDELETED HashicorpCloudBoundary20211221ClusterState = "STATE_DELETED"

	// HashicorpCloudBoundary20211221ClusterStateSTATEUPDATING captures enum value "STATE_UPDATING"
	HashicorpCloudBoundary20211221ClusterStateSTATEUPDATING HashicorpCloudBoundary20211221ClusterState = "STATE_UPDATING"
)

// for schema
var hashicorpCloudBoundary20211221ClusterStateEnum []interface{}

func init() {
	var res []HashicorpCloudBoundary20211221ClusterState
	if err := json.Unmarshal([]byte(`["STATE_INVALID_UNSPECIFIED","STATE_PENDING","STATE_CREATING","STATE_RUNNING","STATE_FAILED","STATE_DELETING","STATE_DELETED","STATE_UPDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudBoundary20211221ClusterStateEnum = append(hashicorpCloudBoundary20211221ClusterStateEnum, v)
	}
}

func (m HashicorpCloudBoundary20211221ClusterState) validateHashicorpCloudBoundary20211221ClusterStateEnum(path, location string, value HashicorpCloudBoundary20211221ClusterState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudBoundary20211221ClusterStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud boundary 20211221 cluster state
func (m HashicorpCloudBoundary20211221ClusterState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudBoundary20211221ClusterStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud boundary 20211221 cluster state based on context it is used
func (m HashicorpCloudBoundary20211221ClusterState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
