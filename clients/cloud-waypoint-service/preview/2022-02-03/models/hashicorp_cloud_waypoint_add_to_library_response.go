// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointAddToLibraryResponse hashicorp cloud waypoint add to library response
//
// swagger:model hashicorp.cloud.waypoint.AddToLibraryResponse
type HashicorpCloudWaypointAddToLibraryResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this hashicorp cloud waypoint add to library response
func (m *HashicorpCloudWaypointAddToLibraryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint add to library response based on context it is used
func (m *HashicorpCloudWaypointAddToLibraryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddToLibraryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddToLibraryResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointAddToLibraryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}