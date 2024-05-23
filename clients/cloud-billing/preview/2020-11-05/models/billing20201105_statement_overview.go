// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Billing20201105StatementOverview StatementOverview contains a few key details from a Statement.
// TODO: This will replace the public StatementSummary object.
//
// swagger:model billing_20201105StatementOverview
type Billing20201105StatementOverview struct {

	// billing_period_end is the end date for the statement billing period
	// Format: date-time
	BillingPeriodEnd strfmt.DateTime `json:"billing_period_end,omitempty"`

	// billing_period_start is the start date for the statement billing period
	// Format: date-time
	BillingPeriodStart strfmt.DateTime `json:"billing_period_start,omitempty"`

	// id is the id associated with the statement
	ID string `json:"id,omitempty"`

	// state indicates the life cycle step that the Statement is currently in.
	State *Billing20201105StatementState `json:"state,omitempty"`

	// total is the sum total in dollars for the statement
	Total string `json:"total,omitempty"`
}

// Validate validates this billing 20201105 statement overview
func (m *Billing20201105StatementOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriodEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingPeriodStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105StatementOverview) validateBillingPeriodEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriodEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("billing_period_end", "body", "date-time", m.BillingPeriodEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Billing20201105StatementOverview) validateBillingPeriodStart(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriodStart) { // not required
		return nil
	}

	if err := validate.FormatOf("billing_period_start", "body", "date-time", m.BillingPeriodStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Billing20201105StatementOverview) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 statement overview based on the context it is used
func (m *Billing20201105StatementOverview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105StatementOverview) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105StatementOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105StatementOverview) UnmarshalBinary(b []byte) error {
	var res Billing20201105StatementOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}