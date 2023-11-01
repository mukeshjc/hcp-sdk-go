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

// HashicorpWaypointVariableFinalValue hashicorp waypoint variable final value
//
// swagger:model hashicorp.waypoint.Variable.FinalValue
type HashicorpWaypointVariableFinalValue struct {

	// bool
	Bool bool `json:"bool,omitempty"`

	// hcl stores values of any complex type in a raw string format
	Hcl string `json:"hcl,omitempty"`

	// num
	Num string `json:"num,omitempty"`

	// 'sensitive' values are hashed as SHA256 values for
	// the purposes of output and logging
	Sensitive string `json:"sensitive,omitempty"`

	// source
	Source *HashicorpWaypointVariableFinalValueSource `json:"source,omitempty"`

	// str
	Str string `json:"str,omitempty"`
}

// Validate validates this hashicorp waypoint variable final value
func (m *HashicorpWaypointVariableFinalValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointVariableFinalValue) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint variable final value based on the context it is used
func (m *HashicorpWaypointVariableFinalValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointVariableFinalValue) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointVariableFinalValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointVariableFinalValue) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointVariableFinalValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
