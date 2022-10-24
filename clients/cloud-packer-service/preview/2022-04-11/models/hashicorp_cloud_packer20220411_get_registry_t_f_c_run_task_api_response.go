// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse hashicorp cloud packer 20220411 get registry t f c run task API response
//
// swagger:model hashicorp.cloud.packer_20220411.GetRegistryTFCRunTaskAPIResponse
type HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse struct {

	// URL of the API used by Terraform Cloud to run HCP Packer Run Tasks.
	APIURL string `json:"api_url,omitempty"`

	// HMAC key used by Terraform Cloud to sign the requests to the HCP Packer run task API.
	HmacKey string `json:"hmac_key,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 get registry t f c run task API response
func (m *HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20220411 get registry t f c run task API response based on context it is used
func (m *HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411GetRegistryTFCRunTaskAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}