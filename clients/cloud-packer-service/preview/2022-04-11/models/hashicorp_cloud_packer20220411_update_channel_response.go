// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411UpdateChannelResponse hashicorp cloud packer 20220411 update channel response
//
// swagger:model hashicorp.cloud.packer_20220411.UpdateChannelResponse
type HashicorpCloudPacker20220411UpdateChannelResponse struct {

	// channel
	Channel *HashicorpCloudPacker20220411Channel `json:"channel,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 update channel response
func (m *HashicorpCloudPacker20220411UpdateChannelResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411UpdateChannelResponse) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411UpdateChannelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411UpdateChannelResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411UpdateChannelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
