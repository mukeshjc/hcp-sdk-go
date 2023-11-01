// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointTokenTrigger The Trigger message is a kind of token type that is only used for
// authenticated trigger URL requests. It should not have any other
// authorized access to make requests in any other API endpoint.
//
// swagger:model hashicorp.waypoint.Token.Trigger
type HashicorpWaypointTokenTrigger struct {

	// The user that initiated the trigger token generation
	FromUserID string `json:"from_user_id,omitempty"`
}

// Validate validates this hashicorp waypoint token trigger
func (m *HashicorpWaypointTokenTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint token trigger based on context it is used
func (m *HashicorpWaypointTokenTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointTokenTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointTokenTrigger) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointTokenTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
