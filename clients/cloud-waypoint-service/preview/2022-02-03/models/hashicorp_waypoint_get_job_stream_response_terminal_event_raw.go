// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetJobStreamResponseTerminalEventRaw hashicorp waypoint get job stream response terminal event raw
//
// swagger:model hashicorp.waypoint.GetJobStreamResponse.Terminal.Event.Raw
type HashicorpWaypointGetJobStreamResponseTerminalEventRaw struct {

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// stderr
	Stderr bool `json:"stderr,omitempty"`
}

// Validate validates this hashicorp waypoint get job stream response terminal event raw
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint get job stream response terminal event raw based on context it is used
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventRaw) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetJobStreamResponseTerminalEventRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
