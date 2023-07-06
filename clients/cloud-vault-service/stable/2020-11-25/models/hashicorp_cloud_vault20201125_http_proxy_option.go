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

// HashicorpCloudVault20201125HTTPProxyOption HttpProxyOption specifies whether a cluster has HTTP Proxy enabled or disabled.
// We use an enum instead of bool for readability and extensibility.
//
// swagger:model hashicorp.cloud.vault_20201125.HttpProxyOption
type HashicorpCloudVault20201125HTTPProxyOption string

func NewHashicorpCloudVault20201125HTTPProxyOption(value HashicorpCloudVault20201125HTTPProxyOption) *HashicorpCloudVault20201125HTTPProxyOption {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVault20201125HTTPProxyOption.
func (m HashicorpCloudVault20201125HTTPProxyOption) Pointer() *HashicorpCloudVault20201125HTTPProxyOption {
	return &m
}

const (

	// HashicorpCloudVault20201125HTTPProxyOptionHTTPPROXYOPTIONINVALID captures enum value "HTTP_PROXY_OPTION_INVALID"
	HashicorpCloudVault20201125HTTPProxyOptionHTTPPROXYOPTIONINVALID HashicorpCloudVault20201125HTTPProxyOption = "HTTP_PROXY_OPTION_INVALID"

	// HashicorpCloudVault20201125HTTPProxyOptionENABLED captures enum value "ENABLED"
	HashicorpCloudVault20201125HTTPProxyOptionENABLED HashicorpCloudVault20201125HTTPProxyOption = "ENABLED"

	// HashicorpCloudVault20201125HTTPProxyOptionDISABLED captures enum value "DISABLED"
	HashicorpCloudVault20201125HTTPProxyOptionDISABLED HashicorpCloudVault20201125HTTPProxyOption = "DISABLED"
)

// for schema
var hashicorpCloudVault20201125HttpProxyOptionEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125HTTPProxyOption
	if err := json.Unmarshal([]byte(`["HTTP_PROXY_OPTION_INVALID","ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125HttpProxyOptionEnum = append(hashicorpCloudVault20201125HttpProxyOptionEnum, v)
	}
}

func (m HashicorpCloudVault20201125HTTPProxyOption) validateHashicorpCloudVault20201125HTTPProxyOptionEnum(path, location string, value HashicorpCloudVault20201125HTTPProxyOption) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125HttpProxyOptionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 Http proxy option
func (m HashicorpCloudVault20201125HTTPProxyOption) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125HTTPProxyOptionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 Http proxy option based on context it is used
func (m HashicorpCloudVault20201125HTTPProxyOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}