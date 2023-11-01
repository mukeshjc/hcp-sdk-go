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

// HashicorpCloudWaypointUpdateApplicationTemplateRequest hashicorp cloud waypoint update application template request
//
// swagger:model hashicorp.cloud.waypoint.UpdateApplicationTemplateRequest
type HashicorpCloudWaypointUpdateApplicationTemplateRequest struct {

	// application_template resembles the desired updated state of the existing
	// application template.
	ApplicationTemplate *HashicorpCloudWaypointApplicationTemplate `json:"application_template,omitempty"`

	// existing_application_template refers to the application template being
	// updated.
	ExistingApplicationTemplate *HashicorpCloudWaypointRefApplicationTemplate `json:"existing_application_template,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointRefNamespace `json:"namespace,omitempty"`
}

// Validate validates this hashicorp cloud waypoint update application template request
func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingApplicationTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) validateApplicationTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationTemplate) { // not required
		return nil
	}

	if m.ApplicationTemplate != nil {
		if err := m.ApplicationTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) validateExistingApplicationTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingApplicationTemplate) { // not required
		return nil
	}

	if m.ExistingApplicationTemplate != nil {
		if err := m.ExistingApplicationTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existing_application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existing_application_template")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint update application template request based on the context it is used
func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExistingApplicationTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) contextValidateApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationTemplate != nil {

		if swag.IsZero(m.ApplicationTemplate) { // not required
			return nil
		}

		if err := m.ApplicationTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) contextValidateExistingApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingApplicationTemplate != nil {

		if swag.IsZero(m.ExistingApplicationTemplate) { // not required
			return nil
		}

		if err := m.ExistingApplicationTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existing_application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existing_application_template")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {

		if swag.IsZero(m.Namespace) { // not required
			return nil
		}

		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUpdateApplicationTemplateRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUpdateApplicationTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}