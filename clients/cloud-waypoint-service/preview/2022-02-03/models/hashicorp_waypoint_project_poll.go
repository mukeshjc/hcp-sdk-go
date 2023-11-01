// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointProjectPoll hashicorp waypoint project poll
//
// swagger:model hashicorp.waypoint.Project.Poll
type HashicorpWaypointProjectPoll struct {

	// enabled must be set to true to enable polling.
	Enabled bool `json:"enabled,omitempty"`

	// interval is a duration string of how often to poll, such as "5s".
	// The server may enforce minimum values, in which case a value lower
	// than the minimum will be ignored.
	Interval string `json:"interval,omitempty"`
}

// Validate validates this hashicorp waypoint project poll
func (m *HashicorpWaypointProjectPoll) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint project poll based on context it is used
func (m *HashicorpWaypointProjectPoll) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointProjectPoll) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointProjectPoll) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointProjectPoll
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
