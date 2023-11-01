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

// HashicorpWaypointDeployment hashicorp waypoint deployment
//
// swagger:model hashicorp.waypoint.Deployment
type HashicorpWaypointDeployment struct {

	// application that this deployment belongs to
	Application *HashicorpWaypointRefApplication `json:"application,omitempty"`

	// ID of the PushedArtifact that was deployed.
	ArtifactID string `json:"artifact_id,omitempty"`

	// component that initiated this deployment
	Component *HashicorpWaypointComponent `json:"component,omitempty"`

	// Resources that this deployment has created or manages.
	DeclaredResources []*HashicorpWaypointDeclaredResource `json:"declared_resources"`

	// deployment is the full raw deployment object encoded directly from
	// the plugin. The client must have all the plugins setup to properly
	// decode this.
	Deployment *OpaqueanyAny `json:"deployment,omitempty"`

	// This is the JSON-encoded protobuf structure of the Any field above.
	// This is generated by the plugin and Waypoint core does not modify this
	// value or have any enforced structure. This will be different per-plugin.
	DeploymentJSON string `json:"deployment_json,omitempty"`

	// Resources that a destroy operation has destroyed
	DestroyedResources []*HashicorpWaypointDestroyedResource `json:"destroyed_resources"`

	// See the docs for Generation.
	Generation *HashicorpWaypointGeneration `json:"generation,omitempty"`

	// True if this deployment had the environment variables available
	// for the entrypoint to talk to. If this is false, this deployment
	// should not be able to communicate back to the server even if it
	// has the entrypoint available. This means this deployment will not
	// support logs, exec, etc.
	HasEntrypointConfig bool `json:"has_entrypoint_config,omitempty"`

	// True if the deployment was done by a plugin that defined an exec plugin
	HasExecPlugin bool `json:"has_exec_plugin,omitempty"`

	// True if the deployment was done by a plugin that defined an logs plugin
	HasLogsPlugin bool `json:"has_logs_plugin,omitempty"`

	// id is the unique ID for this deployment
	ID string `json:"id,omitempty"`

	// ID of the job that created this. This may be empty.
	JobID string `json:"job_id,omitempty"`

	// labels are the set of labels that are present on this build.
	Labels map[string]string `json:"labels,omitempty"`

	// This is the populated preload data. Most of this data can be retrieved
	// through additional API calls or manually computed, but certain API
	// calls will pre-populate some of these fields for convenience. The exact
	// pre-populated fields depend on the API.
	Preload *HashicorpWaypointDeploymentPreload `json:"preload,omitempty"`

	// The sequence number for this build.
	Sequence string `json:"sequence,omitempty"`

	// state is the state of this deployment.
	State *HashicorpWaypointOperationPhysicalState `json:"state,omitempty"`

	// status tracks the status of the most recent operation (creation,
	// destroy, etc. NOTE(mitchellh): I want to separate these out so that
	// you can keep history of the status of multiple operations.
	Status *HashicorpWaypointStatus `json:"status,omitempty"`

	// template data for HCL variables and template functions, json-encoded
	// Format: byte
	TemplateData strfmt.Base64 `json:"template_data,omitempty"`

	// url is the URL to the Deployment
	// this URL might be empty, indicating that the deployment doesn't have
	// the possibility to be contacted directly (e.g: Kubernetes pod)
	// and thus the URL Service (Hashicorp Horizon) will be used instead, if enabled.
	URL string `json:"url,omitempty"`

	// The workspace that this exists in
	Workspace *HashicorpWaypointRefWorkspace `json:"workspace,omitempty"`
}

// Validate validates this hashicorp waypoint deployment
func (m *HashicorpWaypointDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeclaredResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestroyedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointDeployment) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateDeclaredResources(formats strfmt.Registry) error {
	if swag.IsZero(m.DeclaredResources) { // not required
		return nil
	}

	for i := 0; i < len(m.DeclaredResources); i++ {
		if swag.IsZero(m.DeclaredResources[i]) { // not required
			continue
		}

		if m.DeclaredResources[i] != nil {
			if err := m.DeclaredResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("declared_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("declared_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateDeployment(formats strfmt.Registry) error {
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

func (m *HashicorpWaypointDeployment) validateDestroyedResources(formats strfmt.Registry) error {
	if swag.IsZero(m.DestroyedResources) { // not required
		return nil
	}

	for i := 0; i < len(m.DestroyedResources); i++ {
		if swag.IsZero(m.DestroyedResources[i]) { // not required
			continue
		}

		if m.DestroyedResources[i] != nil {
			if err := m.DestroyedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destroyed_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destroyed_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateGeneration(formats strfmt.Registry) error {
	if swag.IsZero(m.Generation) { // not required
		return nil
	}

	if m.Generation != nil {
		if err := m.Generation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validatePreload(formats strfmt.Registry) error {
	if swag.IsZero(m.Preload) { // not required
		return nil
	}

	if m.Preload != nil {
		if err := m.Preload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preload")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint deployment based on the context it is used
func (m *HashicorpWaypointDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeclaredResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestroyedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeneration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {

		if swag.IsZero(m.Component) { // not required
			return nil
		}

		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateDeclaredResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeclaredResources); i++ {

		if m.DeclaredResources[i] != nil {

			if swag.IsZero(m.DeclaredResources[i]) { // not required
				return nil
			}

			if err := m.DeclaredResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("declared_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("declared_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateDeployment(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpWaypointDeployment) contextValidateDestroyedResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DestroyedResources); i++ {

		if m.DestroyedResources[i] != nil {

			if swag.IsZero(m.DestroyedResources[i]) { // not required
				return nil
			}

			if err := m.DestroyedResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destroyed_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destroyed_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateGeneration(ctx context.Context, formats strfmt.Registry) error {

	if m.Generation != nil {

		if swag.IsZero(m.Generation) { // not required
			return nil
		}

		if err := m.Generation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidatePreload(ctx context.Context, formats strfmt.Registry) error {

	if m.Preload != nil {

		if swag.IsZero(m.Preload) { // not required
			return nil
		}

		if err := m.Preload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preload")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDeployment) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace != nil {

		if swag.IsZero(m.Workspace) { // not required
			return nil
		}

		if err := m.Workspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointDeployment) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
