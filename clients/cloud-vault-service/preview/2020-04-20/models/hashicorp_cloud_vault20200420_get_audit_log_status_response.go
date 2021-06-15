// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420GetAuditLogStatusResponse hashicorp cloud vault 20200420 get audit log status response
//
// swagger:model hashicorp.cloud.vault_20200420.GetAuditLogStatusResponse
type HashicorpCloudVault20200420GetAuditLogStatusResponse struct {

	// log
	Log *HashicorpCloudVault20200420AuditLog `json:"log,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 get audit log status response
func (m *HashicorpCloudVault20200420GetAuditLogStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420GetAuditLogStatusResponse) validateLog(formats strfmt.Registry) error {

	if swag.IsZero(m.Log) { // not required
		return nil
	}

	if m.Log != nil {
		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetAuditLogStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetAuditLogStatusResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420GetAuditLogStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
