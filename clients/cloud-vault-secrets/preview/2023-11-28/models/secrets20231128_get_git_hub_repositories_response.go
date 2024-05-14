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

// Secrets20231128GetGitHubRepositoriesResponse secrets 20231128 get git hub repositories response
//
// swagger:model secrets_20231128GetGitHubRepositoriesResponse
type Secrets20231128GetGitHubRepositoriesResponse struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// repositories
	Repositories []*Secrets20231128GitHubRepository `json:"repositories"`
}

// Validate validates this secrets 20231128 get git hub repositories response
func (m *Secrets20231128GetGitHubRepositoriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepositories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetGitHubRepositoriesResponse) validateRepositories(formats strfmt.Registry) error {
	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {
		if swag.IsZero(m.Repositories[i]) { // not required
			continue
		}

		if m.Repositories[i] != nil {
			if err := m.Repositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 get git hub repositories response based on the context it is used
func (m *Secrets20231128GetGitHubRepositoriesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepositories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetGitHubRepositoriesResponse) contextValidateRepositories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Repositories); i++ {

		if m.Repositories[i] != nil {

			if swag.IsZero(m.Repositories[i]) { // not required
				return nil
			}

			if err := m.Repositories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GetGitHubRepositoriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GetGitHubRepositoriesResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GetGitHubRepositoriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
