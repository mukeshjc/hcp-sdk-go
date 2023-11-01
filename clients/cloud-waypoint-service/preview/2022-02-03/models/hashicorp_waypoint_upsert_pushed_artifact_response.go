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

// HashicorpWaypointUpsertPushedArtifactResponse hashicorp waypoint upsert pushed artifact response
//
// swagger:model hashicorp.waypoint.UpsertPushedArtifactResponse
type HashicorpWaypointUpsertPushedArtifactResponse struct {

	// resulting push object, you should replace this with what was sent
	// since the update operation may touch up the input data (i.e. update
	// timestamps)
	Artifact *HashicorpWaypointPushedArtifact `json:"artifact,omitempty"`
}

// Validate validates this hashicorp waypoint upsert pushed artifact response
func (m *HashicorpWaypointUpsertPushedArtifactResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUpsertPushedArtifactResponse) validateArtifact(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifact) { // not required
		return nil
	}

	if m.Artifact != nil {
		if err := m.Artifact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint upsert pushed artifact response based on the context it is used
func (m *HashicorpWaypointUpsertPushedArtifactResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUpsertPushedArtifactResponse) contextValidateArtifact(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifact != nil {

		if swag.IsZero(m.Artifact) { // not required
			return nil
		}

		if err := m.Artifact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUpsertPushedArtifactResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUpsertPushedArtifactResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUpsertPushedArtifactResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
