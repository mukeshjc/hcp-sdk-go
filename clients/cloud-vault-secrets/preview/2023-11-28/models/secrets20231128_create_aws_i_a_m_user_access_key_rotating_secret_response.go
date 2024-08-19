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

// Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse secrets 20231128 create aws i a m user access key rotating secret response
//
// swagger:model secrets_20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse
type Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse struct {

	// config
	Config *Secrets20231128AwsIAMUserAccessKeyRotatingSecretConfig `json:"config,omitempty"`
}

// Validate validates this secrets 20231128 create aws i a m user access key rotating secret response
func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 create aws i a m user access key rotating secret response based on the context it is used
func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128CreateAwsIAMUserAccessKeyRotatingSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
