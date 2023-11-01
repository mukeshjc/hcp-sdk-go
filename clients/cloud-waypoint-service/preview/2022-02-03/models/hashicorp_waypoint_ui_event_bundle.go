// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointUIEventBundle An event packaged alongside prefetched related messages.
//
// swagger:model hashicorp.waypoint.UI.EventBundle
type HashicorpWaypointUIEventBundle struct {

	// event data
	// Format: byte
	EventData strfmt.Base64 `json:"event_data,omitempty"`

	// event type
	EventType *HashicorpWaypointUIEventBundleEventType `json:"event_type,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this hashicorp waypoint UI event bundle
func (m *HashicorpWaypointUIEventBundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIEventBundle) validateEventType(formats strfmt.Registry) error {
	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	if m.EventType != nil {
		if err := m.EventType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_type")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIEventBundle) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint UI event bundle based on the context it is used
func (m *HashicorpWaypointUIEventBundle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIEventBundle) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

	if m.EventType != nil {

		if swag.IsZero(m.EventType) { // not required
			return nil
		}

		if err := m.EventType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUIEventBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUIEventBundle) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUIEventBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
