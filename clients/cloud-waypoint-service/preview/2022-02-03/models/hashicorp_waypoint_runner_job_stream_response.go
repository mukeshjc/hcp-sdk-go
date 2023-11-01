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

// HashicorpWaypointRunnerJobStreamResponse hashicorp waypoint runner job stream response
//
// swagger:model hashicorp.waypoint.RunnerJobStreamResponse
type HashicorpWaypointRunnerJobStreamResponse struct {

	// assignment is when a job is assigned to this job stream. This
	// will happen ONLY in response to a "Request" message from the client.
	//
	// This is sent down for the reattach use case (if reattach is set
	// in Request), too, and the client is expected to Ack it. This
	// verifies that both sides are ready to continue with the job.
	Assignment *HashicorpWaypointRunnerJobStreamResponseJobAssignment `json:"assignment,omitempty"`

	// cancel is sent when a cancel request is made.
	Cancel *HashicorpWaypointRunnerJobStreamResponseJobCancel `json:"cancel,omitempty"`
}

// Validate validates this hashicorp waypoint runner job stream response
func (m *HashicorpWaypointRunnerJobStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointRunnerJobStreamResponse) validateAssignment(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignment) { // not required
		return nil
	}

	if m.Assignment != nil {
		if err := m.Assignment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignment")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointRunnerJobStreamResponse) validateCancel(formats strfmt.Registry) error {
	if swag.IsZero(m.Cancel) { // not required
		return nil
	}

	if m.Cancel != nil {
		if err := m.Cancel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint runner job stream response based on the context it is used
func (m *HashicorpWaypointRunnerJobStreamResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCancel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointRunnerJobStreamResponse) contextValidateAssignment(ctx context.Context, formats strfmt.Registry) error {

	if m.Assignment != nil {

		if swag.IsZero(m.Assignment) { // not required
			return nil
		}

		if err := m.Assignment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignment")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointRunnerJobStreamResponse) contextValidateCancel(ctx context.Context, formats strfmt.Registry) error {

	if m.Cancel != nil {

		if swag.IsZero(m.Cancel) { // not required
			return nil
		}

		if err := m.Cancel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRunnerJobStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRunnerJobStreamResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRunnerJobStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
