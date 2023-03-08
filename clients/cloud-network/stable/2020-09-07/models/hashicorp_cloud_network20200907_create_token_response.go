// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907CreateTokenResponse CreateTokenResponse is a response type for CreateToken endpoint.
//
// swagger:model hashicorp.cloud.network_20200907.CreateTokenResponse
type HashicorpCloudNetwork20200907CreateTokenResponse struct {

	// token is a JWT granting authentication and authorization to create a peering against an HVN.
	Token string `json:"token,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create token response
func (m *HashicorpCloudNetwork20200907CreateTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 create token response based on context it is used
func (m *HashicorpCloudNetwork20200907CreateTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTokenResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreateTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
