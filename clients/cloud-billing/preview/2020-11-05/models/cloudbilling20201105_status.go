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

// Cloudbilling20201105Status Status represents the status of a given statement.
//
// swagger:model cloudbilling_20201105Status
type Cloudbilling20201105Status string

func NewCloudbilling20201105Status(value Cloudbilling20201105Status) *Cloudbilling20201105Status {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Cloudbilling20201105Status.
func (m Cloudbilling20201105Status) Pointer() *Cloudbilling20201105Status {
	return &m
}

const (

	// Cloudbilling20201105StatusOPEN captures enum value "OPEN"
	Cloudbilling20201105StatusOPEN Cloudbilling20201105Status = "OPEN"

	// Cloudbilling20201105StatusPAID captures enum value "PAID"
	Cloudbilling20201105StatusPAID Cloudbilling20201105Status = "PAID"

	// Cloudbilling20201105StatusOVERDUE captures enum value "OVERDUE"
	Cloudbilling20201105StatusOVERDUE Cloudbilling20201105Status = "OVERDUE"
)

// for schema
var cloudbilling20201105StatusEnum []interface{}

func init() {
	var res []Cloudbilling20201105Status
	if err := json.Unmarshal([]byte(`["OPEN","PAID","OVERDUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudbilling20201105StatusEnum = append(cloudbilling20201105StatusEnum, v)
	}
}

func (m Cloudbilling20201105Status) validateCloudbilling20201105StatusEnum(path, location string, value Cloudbilling20201105Status) error {
	if err := validate.EnumCase(path, location, value, cloudbilling20201105StatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cloudbilling 20201105 status
func (m Cloudbilling20201105Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCloudbilling20201105StatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cloudbilling 20201105 status based on context it is used
func (m Cloudbilling20201105Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
