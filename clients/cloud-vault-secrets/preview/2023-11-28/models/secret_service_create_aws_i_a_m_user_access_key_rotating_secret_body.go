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

// SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody secret service create aws i a m user access key rotating secret body
//
// swagger:model SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody
type SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody struct {

	// aws iam user access key params
	AwsIamUserAccessKeyParams *Secrets20231128AwsIAMUserAccessKeyParams `json:"aws_iam_user_access_key_params,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`
}

// Validate validates this secret service create aws i a m user access key rotating secret body
func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsIamUserAccessKeyParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) validateAwsIamUserAccessKeyParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsIamUserAccessKeyParams) { // not required
		return nil
	}

	if m.AwsIamUserAccessKeyParams != nil {
		if err := m.AwsIamUserAccessKeyParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_iam_user_access_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_iam_user_access_key_params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create aws i a m user access key rotating secret body based on the context it is used
func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsIamUserAccessKeyParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) contextValidateAwsIamUserAccessKeyParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsIamUserAccessKeyParams != nil {

		if swag.IsZero(m.AwsIamUserAccessKeyParams) { // not required
			return nil
		}

		if err := m.AwsIamUserAccessKeyParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_iam_user_access_key_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_iam_user_access_key_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateAwsIAMUserAccessKeyRotatingSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
