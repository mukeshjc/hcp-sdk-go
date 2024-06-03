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

// Secrets20231128ListGcpDynamicSecretsResponse secrets 20231128 list gcp dynamic secrets response
//
// swagger:model secrets_20231128ListGcpDynamicSecretsResponse
type Secrets20231128ListGcpDynamicSecretsResponse struct {

	// secrets
	Secrets []*Secrets20231128GcpDynamicSecret `json:"secrets"`
}

// Validate validates this secrets 20231128 list gcp dynamic secrets response
func (m *Secrets20231128ListGcpDynamicSecretsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecrets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ListGcpDynamicSecretsResponse) validateSecrets(formats strfmt.Registry) error {
	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	for i := 0; i < len(m.Secrets); i++ {
		if swag.IsZero(m.Secrets[i]) { // not required
			continue
		}

		if m.Secrets[i] != nil {
			if err := m.Secrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 list gcp dynamic secrets response based on the context it is used
func (m *Secrets20231128ListGcpDynamicSecretsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ListGcpDynamicSecretsResponse) contextValidateSecrets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Secrets); i++ {

		if m.Secrets[i] != nil {

			if swag.IsZero(m.Secrets[i]) { // not required
				return nil
			}

			if err := m.Secrets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128ListGcpDynamicSecretsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128ListGcpDynamicSecretsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128ListGcpDynamicSecretsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
