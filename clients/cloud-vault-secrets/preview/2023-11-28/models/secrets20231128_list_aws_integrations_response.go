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

// Secrets20231128ListAwsIntegrationsResponse secrets 20231128 list aws integrations response
//
// swagger:model secrets_20231128ListAwsIntegrationsResponse
type Secrets20231128ListAwsIntegrationsResponse struct {

	// integrations
	Integrations []*Secrets20231128AwsIntegration `json:"integrations"`

	// pagination
	Pagination *CommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this secrets 20231128 list aws integrations response
func (m *Secrets20231128ListAwsIntegrationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrations(formats); err != nil {
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

func (m *Secrets20231128ListAwsIntegrationsResponse) validateIntegrations(formats strfmt.Registry) error {
	if swag.IsZero(m.Integrations) { // not required
		return nil
	}

	for i := 0; i < len(m.Integrations); i++ {
		if swag.IsZero(m.Integrations[i]) { // not required
			continue
		}

		if m.Integrations[i] != nil {
			if err := m.Integrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128ListAwsIntegrationsResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this secrets 20231128 list aws integrations response based on the context it is used
func (m *Secrets20231128ListAwsIntegrationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegrations(ctx, formats); err != nil {
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

func (m *Secrets20231128ListAwsIntegrationsResponse) contextValidateIntegrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Integrations); i++ {

		if m.Integrations[i] != nil {

			if swag.IsZero(m.Integrations[i]) { // not required
				return nil
			}

			if err := m.Integrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128ListAwsIntegrationsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Secrets20231128ListAwsIntegrationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128ListAwsIntegrationsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128ListAwsIntegrationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
