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

// HashicorpCloudWaypointWorkflowProgram WorkflowProgram represents a program stored in the library that can
// be invoked.
//
// swagger:model hashicorp.cloud.waypoint.WorkflowProgram
type HashicorpCloudWaypointWorkflowProgram struct {

	// id
	ID string `json:"id,omitempty"`

	// labels
	Labels *HashicorpCloudWaypointLabelSet `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// program
	Program string `json:"program,omitempty"`
}

// Validate validates this hashicorp cloud waypoint workflow program
func (m *HashicorpCloudWaypointWorkflowProgram) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWorkflowProgram) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint workflow program based on the context it is used
func (m *HashicorpCloudWaypointWorkflowProgram) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWorkflowProgram) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if m.Labels != nil {

		if swag.IsZero(m.Labels) { // not required
			return nil
		}

		if err := m.Labels.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWorkflowProgram) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWorkflowProgram) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWorkflowProgram
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
