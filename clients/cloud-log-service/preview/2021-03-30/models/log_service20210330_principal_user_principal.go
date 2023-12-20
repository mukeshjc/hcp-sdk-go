// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogService20210330PrincipalUserPrincipal UserPrincipal represents a user who performed an action in HCP.
//
// swagger:model log_service_20210330PrincipalUserPrincipal
type LogService20210330PrincipalUserPrincipal struct {

	// email is the user's email address.
	Email string `json:"email,omitempty"`

	// full_name is the user's name.
	FullName string `json:"full_name,omitempty"`
}

// Validate validates this log service 20210330 principal user principal
func (m *LogService20210330PrincipalUserPrincipal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log service 20210330 principal user principal based on context it is used
func (m *LogService20210330PrincipalUserPrincipal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330PrincipalUserPrincipal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330PrincipalUserPrincipal) UnmarshalBinary(b []byte) error {
	var res LogService20210330PrincipalUserPrincipal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}