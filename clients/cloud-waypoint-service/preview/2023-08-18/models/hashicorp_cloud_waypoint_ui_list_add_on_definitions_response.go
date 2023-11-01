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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudWaypointUIListAddOnDefinitionsResponse hashicorp cloud waypoint UI list add on definitions response
//
// swagger:model hashicorp.cloud.waypoint.UI.ListAddOnDefinitionsResponse
type HashicorpCloudWaypointUIListAddOnDefinitionsResponse struct {

	// add on definition bundles
	AddOnDefinitionBundles []*HashicorpCloudWaypointUIAddOnDefinitionBundle `json:"add_on_definition_bundles"`

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud waypoint UI list add on definitions response
func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOnDefinitionBundles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) validateAddOnDefinitionBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOnDefinitionBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.AddOnDefinitionBundles); i++ {
		if swag.IsZero(m.AddOnDefinitionBundles[i]) { // not required
			continue
		}

		if m.AddOnDefinitionBundles[i] != nil {
			if err := m.AddOnDefinitionBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("add_on_definition_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("add_on_definition_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint UI list add on definitions response based on the context it is used
func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOnDefinitionBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) contextValidateAddOnDefinitionBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddOnDefinitionBundles); i++ {

		if m.AddOnDefinitionBundles[i] != nil {

			if swag.IsZero(m.AddOnDefinitionBundles[i]) { // not required
				return nil
			}

			if err := m.AddOnDefinitionBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("add_on_definition_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("add_on_definition_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIListAddOnDefinitionsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUIListAddOnDefinitionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
