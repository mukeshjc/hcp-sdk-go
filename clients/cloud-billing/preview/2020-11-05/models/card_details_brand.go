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

// CardDetailsBrand card details brand
//
// swagger:model CardDetailsBrand
type CardDetailsBrand string

func NewCardDetailsBrand(value CardDetailsBrand) *CardDetailsBrand {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CardDetailsBrand.
func (m CardDetailsBrand) Pointer() *CardDetailsBrand {
	return &m
}

const (

	// CardDetailsBrandUNKNOWN captures enum value "UNKNOWN"
	CardDetailsBrandUNKNOWN CardDetailsBrand = "UNKNOWN"

	// CardDetailsBrandAMERICANEXPRESS captures enum value "AMERICAN_EXPRESS"
	CardDetailsBrandAMERICANEXPRESS CardDetailsBrand = "AMERICAN_EXPRESS"

	// CardDetailsBrandDINERSCLUB captures enum value "DINERS_CLUB"
	CardDetailsBrandDINERSCLUB CardDetailsBrand = "DINERS_CLUB"

	// CardDetailsBrandDISCOVER captures enum value "DISCOVER"
	CardDetailsBrandDISCOVER CardDetailsBrand = "DISCOVER"

	// CardDetailsBrandJCB captures enum value "JCB"
	CardDetailsBrandJCB CardDetailsBrand = "JCB"

	// CardDetailsBrandMASTERCARD captures enum value "MASTERCARD"
	CardDetailsBrandMASTERCARD CardDetailsBrand = "MASTERCARD"

	// CardDetailsBrandUNIONPAY captures enum value "UNION_PAY"
	CardDetailsBrandUNIONPAY CardDetailsBrand = "UNION_PAY"

	// CardDetailsBrandVISA captures enum value "VISA"
	CardDetailsBrandVISA CardDetailsBrand = "VISA"
)

// for schema
var cardDetailsBrandEnum []interface{}

func init() {
	var res []CardDetailsBrand
	if err := json.Unmarshal([]byte(`["UNKNOWN","AMERICAN_EXPRESS","DINERS_CLUB","DISCOVER","JCB","MASTERCARD","UNION_PAY","VISA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cardDetailsBrandEnum = append(cardDetailsBrandEnum, v)
	}
}

func (m CardDetailsBrand) validateCardDetailsBrandEnum(path, location string, value CardDetailsBrand) error {
	if err := validate.EnumCase(path, location, value, cardDetailsBrandEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this card details brand
func (m CardDetailsBrand) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCardDetailsBrandEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this card details brand based on context it is used
func (m CardDetailsBrand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
