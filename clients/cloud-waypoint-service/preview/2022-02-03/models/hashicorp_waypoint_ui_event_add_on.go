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

// HashicorpWaypointUIEventAddOn hashicorp waypoint UI event add on
//
// swagger:model hashicorp.waypoint.UI.EventAddOn
type HashicorpWaypointUIEventAddOn struct {

	// add on id
	AddOnID string `json:"add_on_id,omitempty"`

	// add on operation
	AddOnOperation *HashicorpWaypointUIEventAddOnOperation `json:"add_on_operation,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp waypoint UI event add on
func (m *HashicorpWaypointUIEventAddOn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOnOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIEventAddOn) validateAddOnOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOnOperation) { // not required
		return nil
	}

	if m.AddOnOperation != nil {
		if err := m.AddOnOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on_operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint UI event add on based on the context it is used
func (m *HashicorpWaypointUIEventAddOn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOnOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIEventAddOn) contextValidateAddOnOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.AddOnOperation != nil {

		if swag.IsZero(m.AddOnOperation) { // not required
			return nil
		}

		if err := m.AddOnOperation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on_operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUIEventAddOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUIEventAddOn) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUIEventAddOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
