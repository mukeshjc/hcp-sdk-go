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

// SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody secret service update gcp service account key rotating secret body
//
// swagger:model SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody
type SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody struct {

	// gcp service account key params
	GcpServiceAccountKeyParams *Secrets20231128GcpServiceAccountKeyParams `json:"gcp_service_account_key_params,omitempty"`

	// rotate on update
	RotateOnUpdate bool `json:"rotate_on_update,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`
}

// Validate validates this secret service update gcp service account key rotating secret body
func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcpServiceAccountKeyParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) validateGcpServiceAccountKeyParams(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpServiceAccountKeyParams) { // not required
		return nil
	}

	if m.GcpServiceAccountKeyParams != nil {
		if err := m.GcpServiceAccountKeyParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_service_account_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_service_account_key_params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service update gcp service account key rotating secret body based on the context it is used
func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGcpServiceAccountKeyParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) contextValidateGcpServiceAccountKeyParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpServiceAccountKeyParams != nil {

		if swag.IsZero(m.GcpServiceAccountKeyParams) { // not required
			return nil
		}

		if err := m.GcpServiceAccountKeyParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_service_account_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_service_account_key_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
