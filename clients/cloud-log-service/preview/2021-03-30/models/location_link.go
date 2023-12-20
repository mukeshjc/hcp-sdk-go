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

// LocationLink Link is used to uniquely reference any resource within HashiCorp Cloud.
// This can be conceptually considered a "foreign key".
//
// swagger:model locationLink
type LocationLink struct {

	// description is a human-friendly description for this link. This is
	// used primarily for informational purposes such as error messages.
	Description string `json:"description,omitempty"`

	// id is the identifier for this resource.
	ID string `json:"id,omitempty"`

	// location is the location where this resource is.
	Location *CloudlocationLocation `json:"location,omitempty"`

	// type is the unique type of the resource. Each service publishes a
	// unique set of types. The type value is recommended to be formatted
	// in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	// in the future, but any string value will work.
	Type string `json:"type,omitempty"`

	// uuid is the unique UUID for this resource.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this location link
func (m *LocationLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationLink) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this location link based on the context it is used
func (m *LocationLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationLink) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationLink) UnmarshalBinary(b []byte) error {
	var res LocationLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}