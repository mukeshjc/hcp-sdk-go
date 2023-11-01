// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamAcceptOrganizationInvitationRequest AcceptOrganizationInvitationRequest is a request to accept an organization invitation.
//
// swagger:model hashicorp.cloud.iam.AcceptOrganizationInvitationRequest
type HashicorpCloudIamAcceptOrganizationInvitationRequest struct {

	// invitation_token is the token to be consumed on successful invitation acceptance.
	InvitationToken string `json:"invitation_token,omitempty"`
}

// Validate validates this hashicorp cloud iam accept organization invitation request
func (m *HashicorpCloudIamAcceptOrganizationInvitationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam accept organization invitation request based on context it is used
func (m *HashicorpCloudIamAcceptOrganizationInvitationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamAcceptOrganizationInvitationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamAcceptOrganizationInvitationRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamAcceptOrganizationInvitationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
