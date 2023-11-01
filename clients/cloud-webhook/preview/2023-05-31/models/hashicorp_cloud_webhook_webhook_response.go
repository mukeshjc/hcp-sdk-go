// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWebhookWebhookResponse The information about the response received after delivering an event payload.
//
// swagger:model hashicorp.cloud.webhook.WebhookResponse
type HashicorpCloudWebhookWebhookResponse struct {

	// The response body.
	Body string `json:"body,omitempty"`

	// The response headers.
	Headers map[string]string `json:"headers,omitempty"`

	// The response status code.
	StatusCode int32 `json:"status_code,omitempty"`
}

// Validate validates this hashicorp cloud webhook webhook response
func (m *HashicorpCloudWebhookWebhookResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud webhook webhook response based on context it is used
func (m *HashicorpCloudWebhookWebhookResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookWebhookResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
