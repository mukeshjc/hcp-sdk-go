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

// HashicorpWaypointTokenTransport The outer structure of the token that is directly Marshaled and
// ASCII armored.
//
// swagger:model hashicorp.waypoint.TokenTransport
type HashicorpWaypointTokenTransport struct {

	// A Marshaled token, stored as bytes because we need to to validate
	// it with the given signature.
	// Format: byte
	Body strfmt.Base64 `json:"body,omitempty"`

	// The key used to generate the signature.
	KeyID string `json:"key_id,omitempty"`

	// Any configuration style metadata that can be passed along with the token
	// without invalidating the token body itself.
	Metadata map[string]string `json:"metadata,omitempty"`

	// A client that sees this populated will use the details to fetch a token
	// via oauth instead of submitting this token directly.
	OauthCreds *HashicorpWaypointTokenTransportOAuthCredentials `json:"oauth_creds,omitempty"`

	// The signature of body for validation.
	// Format: byte
	Signature strfmt.Base64 `json:"signature,omitempty"`
}

// Validate validates this hashicorp waypoint token transport
func (m *HashicorpWaypointTokenTransport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOauthCreds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointTokenTransport) validateOauthCreds(formats strfmt.Registry) error {
	if swag.IsZero(m.OauthCreds) { // not required
		return nil
	}

	if m.OauthCreds != nil {
		if err := m.OauthCreds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth_creds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth_creds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint token transport based on the context it is used
func (m *HashicorpWaypointTokenTransport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOauthCreds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointTokenTransport) contextValidateOauthCreds(ctx context.Context, formats strfmt.Registry) error {

	if m.OauthCreds != nil {

		if swag.IsZero(m.OauthCreds) { // not required
			return nil
		}

		if err := m.OauthCreds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth_creds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth_creds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointTokenTransport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointTokenTransport) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointTokenTransport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
