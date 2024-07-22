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

// HashicorpCloudWaypointCreateVariableResponse hashicorp cloud waypoint create variable response
//
// swagger:model hashicorp.cloud.waypoint.CreateVariableResponse
type HashicorpCloudWaypointCreateVariableResponse struct {

	// The created variable
	Variable *HashicorpCloudWaypointVariable `json:"variable,omitempty"`
}

// Validate validates this hashicorp cloud waypoint create variable response
func (m *HashicorpCloudWaypointCreateVariableResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointCreateVariableResponse) validateVariable(formats strfmt.Registry) error {
	if swag.IsZero(m.Variable) { // not required
		return nil
	}

	if m.Variable != nil {
		if err := m.Variable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variable")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint create variable response based on the context it is used
func (m *HashicorpCloudWaypointCreateVariableResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVariable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointCreateVariableResponse) contextValidateVariable(ctx context.Context, formats strfmt.Registry) error {

	if m.Variable != nil {

		if swag.IsZero(m.Variable) { // not required
			return nil
		}

		if err := m.Variable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variable")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointCreateVariableResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointCreateVariableResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointCreateVariableResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}