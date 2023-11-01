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

// HashicorpWaypointCreateProjectFromTemplateRequest CreateProjectFromTemplateRequest is a message intended for use in requests
// to render project templates.
//
// swagger:model hashicorp.waypoint.CreateProjectFromTemplateRequest
type HashicorpWaypointCreateProjectFromTemplateRequest struct {

	// project name
	ProjectName string `json:"project_name,omitempty"`

	// project template
	ProjectTemplate *HashicorpWaypointRefProjectTemplate `json:"project_template,omitempty"`
}

// Validate validates this hashicorp waypoint create project from template request
func (m *HashicorpWaypointCreateProjectFromTemplateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointCreateProjectFromTemplateRequest) validateProjectTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectTemplate) { // not required
		return nil
	}

	if m.ProjectTemplate != nil {
		if err := m.ProjectTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project_template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint create project from template request based on the context it is used
func (m *HashicorpWaypointCreateProjectFromTemplateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjectTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointCreateProjectFromTemplateRequest) contextValidateProjectTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ProjectTemplate != nil {

		if swag.IsZero(m.ProjectTemplate) { // not required
			return nil
		}

		if err := m.ProjectTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project_template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointCreateProjectFromTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointCreateProjectFromTemplateRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointCreateProjectFromTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
