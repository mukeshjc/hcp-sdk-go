// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefRunnerLabels RunnerLabels references a runner by labels.
// The labels can be a subset match or an exact match.
//
// swagger:model hashicorp.waypoint.Ref.RunnerLabels
type HashicorpWaypointRefRunnerLabels struct {

	// labels
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this hashicorp waypoint ref runner labels
func (m *HashicorpWaypointRefRunnerLabels) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint ref runner labels based on context it is used
func (m *HashicorpWaypointRefRunnerLabels) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefRunnerLabels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefRunnerLabels) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefRunnerLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
