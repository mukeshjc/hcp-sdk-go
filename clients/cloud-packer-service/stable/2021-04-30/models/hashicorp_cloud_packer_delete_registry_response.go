// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudPackerDeleteRegistryResponse hashicorp cloud packer delete registry response
//
// swagger:model hashicorp.cloud.packer.DeleteRegistryResponse
type HashicorpCloudPackerDeleteRegistryResponse struct {

	// operation
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`

	// registry
	Registry *HashicorpCloudPackerRegistry `json:"registry,omitempty"`
}

// Validate validates this hashicorp cloud packer delete registry response
func (m *HashicorpCloudPackerDeleteRegistryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerDeleteRegistryResponse) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerDeleteRegistryResponse) validateRegistry(formats strfmt.Registry) error {

	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerDeleteRegistryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerDeleteRegistryResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerDeleteRegistryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}