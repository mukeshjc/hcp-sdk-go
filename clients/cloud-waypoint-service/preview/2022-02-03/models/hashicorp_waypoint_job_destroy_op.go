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

// HashicorpWaypointJobDestroyOp hashicorp waypoint job destroy op
//
// swagger:model hashicorp.waypoint.Job.DestroyOp
type HashicorpWaypointJobDestroyOp struct {

	// deployment
	Deployment *HashicorpWaypointDeployment `json:"deployment,omitempty"`

	// workspace will delete the app in the workspace that the job
	// is targeting.
	Workspace interface{} `json:"workspace,omitempty"`
}

// Validate validates this hashicorp waypoint job destroy op
func (m *HashicorpWaypointJobDestroyOp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobDestroyOp) validateDeployment(formats strfmt.Registry) error {
	if swag.IsZero(m.Deployment) { // not required
		return nil
	}

	if m.Deployment != nil {
		if err := m.Deployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint job destroy op based on the context it is used
func (m *HashicorpWaypointJobDestroyOp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobDestroyOp) contextValidateDeployment(ctx context.Context, formats strfmt.Registry) error {

	if m.Deployment != nil {

		if swag.IsZero(m.Deployment) { // not required
			return nil
		}

		if err := m.Deployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobDestroyOp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobDestroyOp) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobDestroyOp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}