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

// HashicorpWaypointGetJobStreamResponseTerminalEvent hashicorp waypoint get job stream response terminal event
//
// swagger:model hashicorp.waypoint.GetJobStreamResponse.Terminal.Event
type HashicorpWaypointGetJobStreamResponseTerminalEvent struct {

	// line
	Line *HashicorpWaypointGetJobStreamResponseTerminalEventLine `json:"line,omitempty"`

	// named values
	NamedValues *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues `json:"named_values,omitempty"`

	// raw
	Raw *HashicorpWaypointGetJobStreamResponseTerminalEventRaw `json:"raw,omitempty"`

	// status
	Status *HashicorpWaypointGetJobStreamResponseTerminalEventStatus `json:"status,omitempty"`

	// step
	Step *HashicorpWaypointGetJobStreamResponseTerminalEventStep `json:"step,omitempty"`

	// step group
	StepGroup *HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup `json:"step_group,omitempty"`

	// table
	Table *HashicorpWaypointGetJobStreamResponseTerminalEventTable `json:"table,omitempty"`

	// timestamp of the event as seen by the runner. This might be
	// skewed from the server or the client but relative to all other
	// line output, it will be accurate.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this hashicorp waypoint get job stream response terminal event
func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamedValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTable(formats); err != nil {
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

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateLine(formats strfmt.Registry) error {
	if swag.IsZero(m.Line) { // not required
		return nil
	}

	if m.Line != nil {
		if err := m.Line.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("line")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("line")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateNamedValues(formats strfmt.Registry) error {
	if swag.IsZero(m.NamedValues) { // not required
		return nil
	}

	if m.NamedValues != nil {
		if err := m.NamedValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("named_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("named_values")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.Raw) { // not required
		return nil
	}

	if m.Raw != nil {
		if err := m.Raw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateStep(formats strfmt.Registry) error {
	if swag.IsZero(m.Step) { // not required
		return nil
	}

	if m.Step != nil {
		if err := m.Step.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateStepGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.StepGroup) { // not required
		return nil
	}

	if m.StepGroup != nil {
		if err := m.StepGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step_group")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateTable(formats strfmt.Registry) error {
	if swag.IsZero(m.Table) { // not required
		return nil
	}

	if m.Table != nil {
		if err := m.Table.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("table")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("table")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint get job stream response terminal event based on the context it is used
func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamedValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStepGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateLine(ctx context.Context, formats strfmt.Registry) error {

	if m.Line != nil {

		if swag.IsZero(m.Line) { // not required
			return nil
		}

		if err := m.Line.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("line")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("line")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateNamedValues(ctx context.Context, formats strfmt.Registry) error {

	if m.NamedValues != nil {

		if swag.IsZero(m.NamedValues) { // not required
			return nil
		}

		if err := m.NamedValues.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("named_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("named_values")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.Raw != nil {

		if swag.IsZero(m.Raw) { // not required
			return nil
		}

		if err := m.Raw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	if m.Step != nil {

		if swag.IsZero(m.Step) { // not required
			return nil
		}

		if err := m.Step.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateStepGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.StepGroup != nil {

		if swag.IsZero(m.StepGroup) { // not required
			return nil
		}

		if err := m.StepGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step_group")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) contextValidateTable(ctx context.Context, formats strfmt.Registry) error {

	if m.Table != nil {

		if swag.IsZero(m.Table) { // not required
			return nil
		}

		if err := m.Table.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("table")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("table")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEvent) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetJobStreamResponseTerminalEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
