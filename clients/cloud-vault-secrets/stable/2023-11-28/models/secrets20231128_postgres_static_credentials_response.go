// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128PostgresStaticCredentialsResponse secrets 20231128 postgres static credentials response
//
// swagger:model secrets_20231128PostgresStaticCredentialsResponse
type Secrets20231128PostgresStaticCredentialsResponse struct {

	// connection string
	ConnectionString string `json:"connection_string,omitempty"`
}

// Validate validates this secrets 20231128 postgres static credentials response
func (m *Secrets20231128PostgresStaticCredentialsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 postgres static credentials response based on context it is used
func (m *Secrets20231128PostgresStaticCredentialsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128PostgresStaticCredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128PostgresStaticCredentialsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128PostgresStaticCredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}