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

// HashicorpWaypointUIListEventSchemasResponse NOTE(Cassie): This should be implemented once pipeline_runs are app scoped
// message EventPipelineRun {
// string pipelineRun_id = 1;
// uint64 sequence = 2;
// //    Component component = 3; //aka platform
// //    Ref.Workspace workspace = 4;
// }
//
// swagger:model hashicorp.waypoint.UI.ListEventSchemasResponse
type HashicorpWaypointUIListEventSchemasResponse struct {

	// event add on schema
	EventAddOnSchema *HashicorpWaypointUIEventAddOn `json:"event_add_on_schema,omitempty"`

	// event build schema
	EventBuildSchema *HashicorpWaypointUIEventBuild `json:"event_build_schema,omitempty"`

	// event deployment schema
	EventDeploymentSchema *HashicorpWaypointUIEventDeployment `json:"event_deployment_schema,omitempty"`

	// event release schema
	EventReleaseSchema *HashicorpWaypointUIEventRelease `json:"event_release_schema,omitempty"`
}

// Validate validates this hashicorp waypoint UI list event schemas response
func (m *HashicorpWaypointUIListEventSchemasResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventAddOnSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventBuildSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventDeploymentSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventReleaseSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) validateEventAddOnSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.EventAddOnSchema) { // not required
		return nil
	}

	if m.EventAddOnSchema != nil {
		if err := m.EventAddOnSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_add_on_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_add_on_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) validateEventBuildSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.EventBuildSchema) { // not required
		return nil
	}

	if m.EventBuildSchema != nil {
		if err := m.EventBuildSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_build_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_build_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) validateEventDeploymentSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.EventDeploymentSchema) { // not required
		return nil
	}

	if m.EventDeploymentSchema != nil {
		if err := m.EventDeploymentSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_deployment_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_deployment_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) validateEventReleaseSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.EventReleaseSchema) { // not required
		return nil
	}

	if m.EventReleaseSchema != nil {
		if err := m.EventReleaseSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_release_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_release_schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint UI list event schemas response based on the context it is used
func (m *HashicorpWaypointUIListEventSchemasResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventAddOnSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventBuildSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventDeploymentSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventReleaseSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) contextValidateEventAddOnSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.EventAddOnSchema != nil {

		if swag.IsZero(m.EventAddOnSchema) { // not required
			return nil
		}

		if err := m.EventAddOnSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_add_on_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_add_on_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) contextValidateEventBuildSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.EventBuildSchema != nil {

		if swag.IsZero(m.EventBuildSchema) { // not required
			return nil
		}

		if err := m.EventBuildSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_build_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_build_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) contextValidateEventDeploymentSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.EventDeploymentSchema != nil {

		if swag.IsZero(m.EventDeploymentSchema) { // not required
			return nil
		}

		if err := m.EventDeploymentSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_deployment_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_deployment_schema")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointUIListEventSchemasResponse) contextValidateEventReleaseSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.EventReleaseSchema != nil {

		if swag.IsZero(m.EventReleaseSchema) { // not required
			return nil
		}

		if err := m.EventReleaseSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_release_schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_release_schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUIListEventSchemasResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUIListEventSchemasResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUIListEventSchemasResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
