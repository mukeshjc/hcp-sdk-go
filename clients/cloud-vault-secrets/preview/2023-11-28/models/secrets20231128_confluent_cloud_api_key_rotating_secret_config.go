// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig secrets 20231128 confluent cloud Api key rotating secret config
//
// swagger:model secrets_20231128ConfluentCloudApiKeyRotatingSecretConfig
type Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig struct {

	// app name
	AppName string `json:"app_name,omitempty"`

	// confluent cloud api key params
	ConfluentCloudAPIKeyParams *Secrets20231128ConfluentCloudAPIKeyParams `json:"confluent_cloud_api_key_params,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`
}

// Validate validates this secrets 20231128 confluent cloud Api key rotating secret config
func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfluentCloudAPIKeyParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) validateConfluentCloudAPIKeyParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfluentCloudAPIKeyParams) { // not required
		return nil
	}

	if m.ConfluentCloudAPIKeyParams != nil {
		if err := m.ConfluentCloudAPIKeyParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("confluent_cloud_api_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("confluent_cloud_api_key_params")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this secrets 20231128 confluent cloud Api key rotating secret config based on the context it is used
func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfluentCloudAPIKeyParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) contextValidateConfluentCloudAPIKeyParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfluentCloudAPIKeyParams != nil {

		if swag.IsZero(m.ConfluentCloudAPIKeyParams) { // not required
			return nil
		}

		if err := m.ConfluentCloudAPIKeyParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("confluent_cloud_api_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("confluent_cloud_api_key_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig) UnmarshalBinary(b []byte) error {
	var res Secrets20231128ConfluentCloudAPIKeyRotatingSecretConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}