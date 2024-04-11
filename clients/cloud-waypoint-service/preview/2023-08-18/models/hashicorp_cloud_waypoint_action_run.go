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
	"github.com/go-openapi/validate"
)

// HashicorpCloudWaypointActionRun hashicorp cloud waypoint action run
//
// swagger:model hashicorp.cloud.waypoint.ActionRun
type HashicorpCloudWaypointActionRun struct {

	// The full action config that was ran in the moment it was invoked
	ActionConfig *HashicorpCloudWaypointActionConfig `json:"action_config,omitempty"`

	// The status of the background job that was used to run this action
	BackgroundJob *HashicorpCloudWaypointJob `json:"background_job,omitempty"`

	// Time that the action run has finished
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	// Time that the action run was invoked
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The unique identifer for this action run
	ID string `json:"id,omitempty"`

	// The namespace the action will run in
	Namespace *HashicorpCloudWaypointRefNamespace `json:"namespace,omitempty"`

	// The exact request used from the action config that launched this run
	Request *HashicorpCloudWaypointActionConfigRequest `json:"request,omitempty"`

	// The response status returned to Waypoint by the external endpoint when an
	// Action is triggered by Waypoint
	ResponseStatus *HashicorpCloudWaypointActionRunResponseStatus `json:"response_status,omitempty"`

	// The service used to run the action
	//
	// NOTE(briancain): Apparently GoWorker/Redis will have job information
	//  so maybe we can store that here directly on the run action message
	RunBy string `json:"run_by,omitempty"`

	// The scope for which this run was invoked in.
	Scope *HashicorpCloudWaypointActionRunScope `json:"scope,omitempty"`

	// The sequence number used to determine how many times this has been run
	Sequence string `json:"sequence,omitempty"`

	// A list of 0-Many status logs.
	StatusLog []*HashicorpCloudWaypointStatusLog `json:"status_log"`
}

// Validate validates this hashicorp cloud waypoint action run
func (m *HashicorpCloudWaypointActionRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackgroundJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateActionConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionConfig) { // not required
		return nil
	}

	if m.ActionConfig != nil {
		if err := m.ActionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateBackgroundJob(formats strfmt.Registry) error {
	if swag.IsZero(m.BackgroundJob) { // not required
		return nil
	}

	if m.BackgroundJob != nil {
		if err := m.BackgroundJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("background_job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("background_job")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateCompletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointActionRun) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateResponseStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseStatus) { // not required
		return nil
	}

	if m.ResponseStatus != nil {
		if err := m.ResponseStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) validateStatusLog(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusLog) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusLog); i++ {
		if swag.IsZero(m.StatusLog[i]) { // not required
			continue
		}

		if m.StatusLog[i] != nil {
			if err := m.StatusLog[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_log" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint action run based on the context it is used
func (m *HashicorpCloudWaypointActionRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackgroundJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponseStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateActionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionConfig != nil {

		if swag.IsZero(m.ActionConfig) { // not required
			return nil
		}

		if err := m.ActionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateBackgroundJob(ctx context.Context, formats strfmt.Registry) error {

	if m.BackgroundJob != nil {

		if swag.IsZero(m.BackgroundJob) { // not required
			return nil
		}

		if err := m.BackgroundJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("background_job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("background_job")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointActionRun) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if swag.IsZero(m.Request) { // not required
			return nil
		}

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateResponseStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ResponseStatus != nil {

		if swag.IsZero(m.ResponseStatus) { // not required
			return nil
		}

		if err := m.ResponseStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {

		if swag.IsZero(m.Scope) { // not required
			return nil
		}

		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionRun) contextValidateStatusLog(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusLog); i++ {

		if m.StatusLog[i] != nil {

			if swag.IsZero(m.StatusLog[i]) { // not required
				return nil
			}

			if err := m.StatusLog[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_log" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionRun) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointActionRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
