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

// HashicorpWaypointGetTFRunStatusResponse GetTFRunStatusResponse is the response containing the status of the latest
// run of a Terraform workspace
//
// swagger:model hashicorp.waypoint.GetTFRunStatusResponse
type HashicorpWaypointGetTFRunStatusResponse struct {

	// the state of the TF run
	State *HashicorpWaypointTerraformTFRunState `json:"state,omitempty"`

	// the URL of the TF run
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp waypoint get t f run status response
func (m *HashicorpWaypointGetTFRunStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetTFRunStatusResponse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint get t f run status response based on the context it is used
func (m *HashicorpWaypointGetTFRunStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetTFRunStatusResponse) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetTFRunStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetTFRunStatusResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetTFRunStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
