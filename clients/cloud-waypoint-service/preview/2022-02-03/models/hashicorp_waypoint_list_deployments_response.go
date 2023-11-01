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

// HashicorpWaypointListDeploymentsResponse hashicorp waypoint list deployments response
//
// swagger:model hashicorp.waypoint.ListDeploymentsResponse
type HashicorpWaypointListDeploymentsResponse struct {

	// deployments is the list of deployments.
	Deployments []*HashicorpWaypointDeployment `json:"deployments"`
}

// Validate validates this hashicorp waypoint list deployments response
func (m *HashicorpWaypointListDeploymentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListDeploymentsResponse) validateDeployments(formats strfmt.Registry) error {
	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint list deployments response based on the context it is used
func (m *HashicorpWaypointListDeploymentsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListDeploymentsResponse) contextValidateDeployments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deployments); i++ {

		if m.Deployments[i] != nil {

			if swag.IsZero(m.Deployments[i]) { // not required
				return nil
			}

			if err := m.Deployments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointListDeploymentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointListDeploymentsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointListDeploymentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
