// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Billing20201105RemoveOnDemandPaymentMethodResponse billing 20201105 remove on demand payment method response
//
// swagger:model billing_20201105RemoveOnDemandPaymentMethodResponse
type Billing20201105RemoveOnDemandPaymentMethodResponse struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this billing 20201105 remove on demand payment method response
func (m *Billing20201105RemoveOnDemandPaymentMethodResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing 20201105 remove on demand payment method response based on context it is used
func (m *Billing20201105RemoveOnDemandPaymentMethodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105RemoveOnDemandPaymentMethodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105RemoveOnDemandPaymentMethodResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105RemoveOnDemandPaymentMethodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}