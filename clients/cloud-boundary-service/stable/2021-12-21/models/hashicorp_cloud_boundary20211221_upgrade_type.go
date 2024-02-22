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

// HashicorpCloudBoundary20211221UpgradeType UpgradeType is the upgrade method for the cluster
//
// - UPGRADE_TYPE_UNSPECIFIED: UPGRADE_TYPE_UNSPECIFIED is an unspecified upgrade type
//   - UPGRADE_TYPE_AUTOMATIC: UPGRADE_TYPE_AUTOMATIC is the upgrade type where we pick the schedule
//   - UPGRADE_TYPE_SCHEDULED: UPGRADE_TYPE_SCHEDULED is the upgrade type where the user picks the schedule
//   - UPGRADE_TYPE_MANUAL: UPGRADE_TYPE_MANUAL is the upgrade type where the user can manually trigger the upgrade before a set window,
//
// after which the cluster will be automatically upgraded.
//
// swagger:model hashicorp.cloud.boundary_20211221.UpgradeType
type HashicorpCloudBoundary20211221UpgradeType string

func NewHashicorpCloudBoundary20211221UpgradeType(value HashicorpCloudBoundary20211221UpgradeType) *HashicorpCloudBoundary20211221UpgradeType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudBoundary20211221UpgradeType.
func (m HashicorpCloudBoundary20211221UpgradeType) Pointer() *HashicorpCloudBoundary20211221UpgradeType {
	return &m
}

const (

	// HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEUNSPECIFIED captures enum value "UPGRADE_TYPE_UNSPECIFIED"
	HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEUNSPECIFIED HashicorpCloudBoundary20211221UpgradeType = "UPGRADE_TYPE_UNSPECIFIED"

	// HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEAUTOMATIC captures enum value "UPGRADE_TYPE_AUTOMATIC"
	HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEAUTOMATIC HashicorpCloudBoundary20211221UpgradeType = "UPGRADE_TYPE_AUTOMATIC"

	// HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPESCHEDULED captures enum value "UPGRADE_TYPE_SCHEDULED"
	HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPESCHEDULED HashicorpCloudBoundary20211221UpgradeType = "UPGRADE_TYPE_SCHEDULED"

	// HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEMANUAL captures enum value "UPGRADE_TYPE_MANUAL"
	HashicorpCloudBoundary20211221UpgradeTypeUPGRADETYPEMANUAL HashicorpCloudBoundary20211221UpgradeType = "UPGRADE_TYPE_MANUAL"
)

// for schema
var hashicorpCloudBoundary20211221UpgradeTypeEnum []interface{}

func init() {
	var res []HashicorpCloudBoundary20211221UpgradeType
	if err := json.Unmarshal([]byte(`["UPGRADE_TYPE_UNSPECIFIED","UPGRADE_TYPE_AUTOMATIC","UPGRADE_TYPE_SCHEDULED","UPGRADE_TYPE_MANUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudBoundary20211221UpgradeTypeEnum = append(hashicorpCloudBoundary20211221UpgradeTypeEnum, v)
	}
}

func (m HashicorpCloudBoundary20211221UpgradeType) validateHashicorpCloudBoundary20211221UpgradeTypeEnum(path, location string, value HashicorpCloudBoundary20211221UpgradeType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudBoundary20211221UpgradeTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud boundary 20211221 upgrade type
func (m HashicorpCloudBoundary20211221UpgradeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudBoundary20211221UpgradeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud boundary 20211221 upgrade type based on context it is used
func (m HashicorpCloudBoundary20211221UpgradeType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
