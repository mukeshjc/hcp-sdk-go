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

// HashicorpWaypointBuildPreload hashicorp waypoint build preload
//
// swagger:model hashicorp.waypoint.Build.Preload
type HashicorpWaypointBuildPreload struct {

	// The ref that was used in the job to run this operation. This is
	// also accessible by querying the job via the job_id and should always
	// match.
	//
	// This may be null under multiple circumstances: (1) the job was
	// manually triggered with local data (no datasource) or (2) the job
	// was run in earlier versions of Waypoint before we tracked this or
	// (3) the job hasn't yet loaded the data.
	//
	// This is always pre-populated if it is exists.
	JobDataSourceRef *HashicorpWaypointJobDataSourceRef `json:"job_data_source_ref,omitempty"`
}

// Validate validates this hashicorp waypoint build preload
func (m *HashicorpWaypointBuildPreload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobDataSourceRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointBuildPreload) validateJobDataSourceRef(formats strfmt.Registry) error {
	if swag.IsZero(m.JobDataSourceRef) { // not required
		return nil
	}

	if m.JobDataSourceRef != nil {
		if err := m.JobDataSourceRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job_data_source_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job_data_source_ref")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint build preload based on the context it is used
func (m *HashicorpWaypointBuildPreload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobDataSourceRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointBuildPreload) contextValidateJobDataSourceRef(ctx context.Context, formats strfmt.Registry) error {

	if m.JobDataSourceRef != nil {

		if swag.IsZero(m.JobDataSourceRef) { // not required
			return nil
		}

		if err := m.JobDataSourceRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job_data_source_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job_data_source_ref")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointBuildPreload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointBuildPreload) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointBuildPreload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}