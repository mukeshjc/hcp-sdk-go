// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogService20210330DataPlaneAuditFilter DataPlaneAuditFilter represents a filter that regulates what resources and resource types should have data plane audit logs streamed.
//
// swagger:model log_service_20210330DataPlaneAuditFilter
type LogService20210330DataPlaneAuditFilter struct {

	// id represents the id of the destination filter
	ID string `json:"id,omitempty"`

	// include_resources represents resources that should have data plane audit logs streamed.
	IncludeResources []*LogService20210330AuditFilterScope `json:"include_resources"`
}

// Validate validates this log service 20210330 data plane audit filter
func (m *LogService20210330DataPlaneAuditFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncludeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330DataPlaneAuditFilter) validateIncludeResources(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeResources) { // not required
		return nil
	}

	for i := 0; i < len(m.IncludeResources); i++ {
		if swag.IsZero(m.IncludeResources[i]) { // not required
			continue
		}

		if m.IncludeResources[i] != nil {
			if err := m.IncludeResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("include_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("include_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this log service 20210330 data plane audit filter based on the context it is used
func (m *LogService20210330DataPlaneAuditFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncludeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330DataPlaneAuditFilter) contextValidateIncludeResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncludeResources); i++ {

		if m.IncludeResources[i] != nil {

			if swag.IsZero(m.IncludeResources[i]) { // not required
				return nil
			}

			if err := m.IncludeResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("include_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("include_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330DataPlaneAuditFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330DataPlaneAuditFilter) UnmarshalBinary(b []byte) error {
	var res LogService20210330DataPlaneAuditFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}