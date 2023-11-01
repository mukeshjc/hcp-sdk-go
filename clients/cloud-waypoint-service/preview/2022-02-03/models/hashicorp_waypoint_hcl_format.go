// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointHclFormat HCL files can be in either HCL or JSON syntax. We need to know ahead of
// time in order to parse it properly. We could perform heuristics but
// we prefer to be explicit.
//
// swagger:model hashicorp.waypoint.Hcl.Format
type HashicorpWaypointHclFormat string

func NewHashicorpWaypointHclFormat(value HashicorpWaypointHclFormat) *HashicorpWaypointHclFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointHclFormat.
func (m HashicorpWaypointHclFormat) Pointer() *HashicorpWaypointHclFormat {
	return &m
}

const (

	// HashicorpWaypointHclFormatHCL captures enum value "HCL"
	HashicorpWaypointHclFormatHCL HashicorpWaypointHclFormat = "HCL"

	// HashicorpWaypointHclFormatJSON captures enum value "JSON"
	HashicorpWaypointHclFormatJSON HashicorpWaypointHclFormat = "JSON"
)

// for schema
var hashicorpWaypointHclFormatEnum []interface{}

func init() {
	var res []HashicorpWaypointHclFormat
	if err := json.Unmarshal([]byte(`["HCL","JSON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointHclFormatEnum = append(hashicorpWaypointHclFormatEnum, v)
	}
}

func (m HashicorpWaypointHclFormat) validateHashicorpWaypointHclFormatEnum(path, location string, value HashicorpWaypointHclFormat) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointHclFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint hcl format
func (m HashicorpWaypointHclFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointHclFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint hcl format based on context it is used
func (m HashicorpWaypointHclFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}