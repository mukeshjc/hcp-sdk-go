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

// HashicorpCloudWebhookCreateWebhookResponse hashicorp cloud webhook create webhook response
//
// swagger:model hashicorp.cloud.webhook.CreateWebhookResponse
type HashicorpCloudWebhookCreateWebhookResponse struct {

	// The newly create webhook.
	Webhook *HashicorpCloudWebhookWebhook `json:"webhook,omitempty"`
}

// Validate validates this hashicorp cloud webhook create webhook response
func (m *HashicorpCloudWebhookCreateWebhookResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookCreateWebhookResponse) validateWebhook(formats strfmt.Registry) error {
	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud webhook create webhook response based on the context it is used
func (m *HashicorpCloudWebhookCreateWebhookResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebhook(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookCreateWebhookResponse) contextValidateWebhook(ctx context.Context, formats strfmt.Registry) error {

	if m.Webhook != nil {

		if swag.IsZero(m.Webhook) { // not required
			return nil
		}

		if err := m.Webhook.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookCreateWebhookResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookCreateWebhookResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookCreateWebhookResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
