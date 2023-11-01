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

// Billing20201105InvoiceLineItem InvoiceLineItem represents a line item on an invoice.
//
// swagger:model billing_20201105InvoiceLineItem
type Billing20201105InvoiceLineItem struct {

	// amount is cost of this line item in USD.
	Amount string `json:"amount,omitempty"`

	// last_usage_recorded_at contains the timestamp of the last "usage_until" from
	// inner line items that are contained in the aggregation.
	// This can also be used to group items that belongs to the same configuration,
	// given that they all get their recording at the same point in time.
	// Format: date-time
	LastUsageRecordedAt strfmt.DateTime `json:"last_usage_recorded_at,omitempty"`

	// product_key is the unique identifier of the associated product.
	ProductKey string `json:"product_key,omitempty"`

	// product_name is the name of the product associated with this line item.
	ProductName string `json:"product_name,omitempty"`

	// quantity is how many units of this product this item contains.
	Quantity string `json:"quantity,omitempty"`

	// standalone_item when set to true will indicate that this item should not be grouped
	// with others by its last_usage_recorded_at.
	// It usually will be set to true for items that are visually independent of any
	// package or selected configuration.
	StandaloneItem bool `json:"standalone_item,omitempty"`

	// unit is the unit on which the product price is calculated.
	Unit *Billing20201105Unit `json:"unit,omitempty"`

	// unit_display_key is the human representation label of the unit that was used in this product.
	UnitDisplayKey string `json:"unit_display_key,omitempty"`

	// unit_price is the price per unit of the product associated.
	UnitPrice string `json:"unit_price,omitempty"`
}

// Validate validates this billing 20201105 invoice line item
func (m *Billing20201105InvoiceLineItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUsageRecordedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105InvoiceLineItem) validateLastUsageRecordedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUsageRecordedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_usage_recorded_at", "body", "date-time", m.LastUsageRecordedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Billing20201105InvoiceLineItem) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 invoice line item based on the context it is used
func (m *Billing20201105InvoiceLineItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105InvoiceLineItem) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105InvoiceLineItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105InvoiceLineItem) UnmarshalBinary(b []byte) error {
	var res Billing20201105InvoiceLineItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
