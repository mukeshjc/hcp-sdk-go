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

// HashicorpCloudPacker20220411ListIterationsResponse hashicorp cloud packer 20220411 list iterations response
//
// swagger:model hashicorp.cloud.packer_20220411.ListIterationsResponse
type HashicorpCloudPacker20220411ListIterationsResponse struct {

	// The requested list of iterations.
	Iterations []*HashicorpCloudPacker20220411IterationforList `json:"iterations"`

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 list iterations response
func (m *HashicorpCloudPacker20220411ListIterationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIterations(formats); err != nil {
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

func (m *HashicorpCloudPacker20220411ListIterationsResponse) validateIterations(formats strfmt.Registry) error {
	if swag.IsZero(m.Iterations) { // not required
		return nil
	}

	for i := 0; i < len(m.Iterations); i++ {
		if swag.IsZero(m.Iterations[i]) { // not required
			continue
		}

		if m.Iterations[i] != nil {
			if err := m.Iterations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("iterations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("iterations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20220411ListIterationsResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud packer 20220411 list iterations response based on the context it is used
func (m *HashicorpCloudPacker20220411ListIterationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIterations(ctx, formats); err != nil {
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

func (m *HashicorpCloudPacker20220411ListIterationsResponse) contextValidateIterations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Iterations); i++ {

		if m.Iterations[i] != nil {
			if err := m.Iterations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("iterations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("iterations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20220411ListIterationsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
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
func (m *HashicorpCloudPacker20220411ListIterationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411ListIterationsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411ListIterationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}