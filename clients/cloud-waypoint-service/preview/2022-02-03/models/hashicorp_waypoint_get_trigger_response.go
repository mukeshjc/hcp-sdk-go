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

// HashicorpWaypointGetTriggerResponse hashicorp waypoint get trigger response
//
// swagger:model hashicorp.waypoint.GetTriggerResponse
type HashicorpWaypointGetTriggerResponse struct {

	// trigger
	Trigger *HashicorpWaypointTrigger `json:"trigger,omitempty"`
}

// Validate validates this hashicorp waypoint get trigger response
func (m *HashicorpWaypointGetTriggerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetTriggerResponse) validateTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint get trigger response based on the context it is used
func (m *HashicorpWaypointGetTriggerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetTriggerResponse) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {

		if swag.IsZero(m.Trigger) { // not required
			return nil
		}

		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetTriggerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetTriggerResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetTriggerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
