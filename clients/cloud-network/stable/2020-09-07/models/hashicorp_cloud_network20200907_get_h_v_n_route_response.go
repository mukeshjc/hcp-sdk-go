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

// HashicorpCloudNetwork20200907GetHVNRouteResponse hashicorp cloud network 20200907 get h v n route response
//
// swagger:model hashicorp.cloud.network_20200907.GetHVNRouteResponse
type HashicorpCloudNetwork20200907GetHVNRouteResponse struct {

	// route
	Route *HashicorpCloudNetwork20200907HVNRoute `json:"route,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 get h v n route response
func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 get h v n route response based on the context it is used
func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.Route != nil {

		if swag.IsZero(m.Route) { // not required
			return nil
		}

		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907GetHVNRouteResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907GetHVNRouteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
