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

// Billing20201105ContractBillingMethodDetails ContractBillingMethodDetails contains both the information used to register a
// Billing Account billed via Contract, and additional information for display
// purposes.
//
// swagger:model billing_20201105ContractBillingMethodDetails
type Billing20201105ContractBillingMethodDetails struct {

	// billing_method contains the information used to register a Contract-backed
	// Billing Account.
	BillingMethod *Billing20201105ContractBillingMethod `json:"billing_method,omitempty"`
}

// Validate validates this billing 20201105 contract billing method details
func (m *Billing20201105ContractBillingMethodDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ContractBillingMethodDetails) validateBillingMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingMethod) { // not required
		return nil
	}

	if m.BillingMethod != nil {
		if err := m.BillingMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_method")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 contract billing method details based on the context it is used
func (m *Billing20201105ContractBillingMethodDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ContractBillingMethodDetails) contextValidateBillingMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingMethod != nil {
		if err := m.BillingMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_method")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ContractBillingMethodDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ContractBillingMethodDetails) UnmarshalBinary(b []byte) error {
	var res Billing20201105ContractBillingMethodDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
