// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointUIListPipelineRunsResponse hashicorp waypoint UI list pipeline runs response
//
// swagger:model hashicorp.waypoint.UI.ListPipelineRunsResponse
type HashicorpWaypointUIListPipelineRunsResponse struct {

	// pagination
	Pagination *HashicorpWaypointPaginationResponse `json:"pagination,omitempty"`

	// pipeline run bundles
	PipelineRunBundles []*HashicorpWaypointUIPipelineRunBundle `json:"pipeline_run_bundles"`
}

// Validate validates this hashicorp waypoint UI list pipeline runs response
func (m *HashicorpWaypointUIListPipelineRunsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelineRunBundles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListPipelineRunsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListPipelineRunsResponse) validatePipelineRunBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.PipelineRunBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.PipelineRunBundles); i++ {
		if swag.IsZero(m.PipelineRunBundles[i]) { // not required
			continue
		}

		if m.PipelineRunBundles[i] != nil {
			if err := m.PipelineRunBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipeline_run_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pipeline_run_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint UI list pipeline runs response based on the context it is used
func (m *HashicorpWaypointUIListPipelineRunsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipelineRunBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListPipelineRunsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListPipelineRunsResponse) contextValidatePipelineRunBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PipelineRunBundles); i++ {

		if m.PipelineRunBundles[i] != nil {

			if swag.IsZero(m.PipelineRunBundles[i]) { // not required
				return nil
			}

			if err := m.PipelineRunBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipeline_run_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pipeline_run_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUIListPipelineRunsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUIListPipelineRunsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUIListPipelineRunsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}