// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPackerChannelAssignment hashicorp cloud packer channel assignment
//
// swagger:model hashicorp.cloud.packer.ChannelAssignment
type HashicorpCloudPackerChannelAssignment struct {

	// The user who assigned this iteration.
	AuthorID string `json:"author_id,omitempty"`

	// The assigned iteration fingerprint.
	IterationFingerprint string `json:"iteration_fingerprint,omitempty"`

	// The assigned iteration ULID.
	IterationID string `json:"iteration_id,omitempty"`

	// The assigned iteration incremental version.
	IterationIncrementalVersion int32 `json:"iteration_incremental_version,omitempty"`

	// When the promotion happened.
	// Format: date-time
	PromotedAt strfmt.DateTime `json:"promoted_at,omitempty"`
}

// Validate validates this hashicorp cloud packer channel assignment
func (m *HashicorpCloudPackerChannelAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePromotedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerChannelAssignment) validatePromotedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("promoted_at", "body", "date-time", m.PromotedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud packer channel assignment based on context it is used
func (m *HashicorpCloudPackerChannelAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerChannelAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerChannelAssignment) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerChannelAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
