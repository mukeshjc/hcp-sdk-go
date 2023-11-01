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

// HashicorpWaypointUIListProjectsResponse hashicorp waypoint UI list projects response
//
// swagger:model hashicorp.waypoint.UI.ListProjectsResponse
type HashicorpWaypointUIListProjectsResponse struct {

	// pagination
	Pagination *HashicorpWaypointPaginationResponse `json:"pagination,omitempty"`

	// project bundles
	ProjectBundles []*HashicorpWaypointUIProjectBundle `json:"project_bundles"`

	// total count
	TotalCount string `json:"total_count,omitempty"`
}

// Validate validates this hashicorp waypoint UI list projects response
func (m *HashicorpWaypointUIListProjectsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectBundles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListProjectsResponse) validatePagination(formats strfmt.Registry) error {
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

func (m *HashicorpWaypointUIListProjectsResponse) validateProjectBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectBundles); i++ {
		if swag.IsZero(m.ProjectBundles[i]) { // not required
			continue
		}

		if m.ProjectBundles[i] != nil {
			if err := m.ProjectBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint UI list projects response based on the context it is used
func (m *HashicorpWaypointUIListProjectsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListProjectsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpWaypointUIListProjectsResponse) contextValidateProjectBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectBundles); i++ {

		if m.ProjectBundles[i] != nil {

			if swag.IsZero(m.ProjectBundles[i]) { // not required
				return nil
			}

			if err := m.ProjectBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUIListProjectsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUIListProjectsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUIListProjectsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
