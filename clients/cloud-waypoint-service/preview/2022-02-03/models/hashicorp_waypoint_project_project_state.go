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

// HashicorpWaypointProjectProjectState ProjectState is set to active when the project is created and throughout its lifetime.
// When a project destroy operation begins, the state is set to destroying.
//
// swagger:model hashicorp.waypoint.Project.ProjectState
type HashicorpWaypointProjectProjectState string

func NewHashicorpWaypointProjectProjectState(value HashicorpWaypointProjectProjectState) *HashicorpWaypointProjectProjectState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointProjectProjectState.
func (m HashicorpWaypointProjectProjectState) Pointer() *HashicorpWaypointProjectProjectState {
	return &m
}

const (

	// HashicorpWaypointProjectProjectStateACTIVE captures enum value "ACTIVE"
	HashicorpWaypointProjectProjectStateACTIVE HashicorpWaypointProjectProjectState = "ACTIVE"

	// HashicorpWaypointProjectProjectStateDESTROYING captures enum value "DESTROYING"
	HashicorpWaypointProjectProjectStateDESTROYING HashicorpWaypointProjectProjectState = "DESTROYING"
)

// for schema
var hashicorpWaypointProjectProjectStateEnum []interface{}

func init() {
	var res []HashicorpWaypointProjectProjectState
	if err := json.Unmarshal([]byte(`["ACTIVE","DESTROYING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointProjectProjectStateEnum = append(hashicorpWaypointProjectProjectStateEnum, v)
	}
}

func (m HashicorpWaypointProjectProjectState) validateHashicorpWaypointProjectProjectStateEnum(path, location string, value HashicorpWaypointProjectProjectState) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointProjectProjectStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint project project state
func (m HashicorpWaypointProjectProjectState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointProjectProjectStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint project project state based on context it is used
func (m HashicorpWaypointProjectProjectState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
