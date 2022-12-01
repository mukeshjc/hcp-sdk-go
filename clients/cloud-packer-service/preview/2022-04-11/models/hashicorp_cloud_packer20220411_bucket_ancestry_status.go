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

// HashicorpCloudPacker20220411BucketAncestryStatus  - UNDETERMINED: HCP Packer cannot determine whether the child image is out of date.
// This status occurs when the parent iteration was not assigned to an image channel when Packer built the image,
// the image channel no longer exists in the image bucket, or HCP Packer is not tracking the parent.
//   - UP_TO_DATE: The child image is built from the iteration currently assigned to the image channel.
//   - OUT_OF_DATE: The child image is built from a different iteration than the one currently assigned to the image channel.
//
// swagger:model hashicorp.cloud.packer_20220411.BucketAncestry.Status
type HashicorpCloudPacker20220411BucketAncestryStatus string

func NewHashicorpCloudPacker20220411BucketAncestryStatus(value HashicorpCloudPacker20220411BucketAncestryStatus) *HashicorpCloudPacker20220411BucketAncestryStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPacker20220411BucketAncestryStatus.
func (m HashicorpCloudPacker20220411BucketAncestryStatus) Pointer() *HashicorpCloudPacker20220411BucketAncestryStatus {
	return &m
}

const (

	// HashicorpCloudPacker20220411BucketAncestryStatusUNDETERMINED captures enum value "UNDETERMINED"
	HashicorpCloudPacker20220411BucketAncestryStatusUNDETERMINED HashicorpCloudPacker20220411BucketAncestryStatus = "UNDETERMINED"

	// HashicorpCloudPacker20220411BucketAncestryStatusUPTODATE captures enum value "UP_TO_DATE"
	HashicorpCloudPacker20220411BucketAncestryStatusUPTODATE HashicorpCloudPacker20220411BucketAncestryStatus = "UP_TO_DATE"

	// HashicorpCloudPacker20220411BucketAncestryStatusOUTOFDATE captures enum value "OUT_OF_DATE"
	HashicorpCloudPacker20220411BucketAncestryStatusOUTOFDATE HashicorpCloudPacker20220411BucketAncestryStatus = "OUT_OF_DATE"
)

// for schema
var hashicorpCloudPacker20220411BucketAncestryStatusEnum []interface{}

func init() {
	var res []HashicorpCloudPacker20220411BucketAncestryStatus
	if err := json.Unmarshal([]byte(`["UNDETERMINED","UP_TO_DATE","OUT_OF_DATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPacker20220411BucketAncestryStatusEnum = append(hashicorpCloudPacker20220411BucketAncestryStatusEnum, v)
	}
}

func (m HashicorpCloudPacker20220411BucketAncestryStatus) validateHashicorpCloudPacker20220411BucketAncestryStatusEnum(path, location string, value HashicorpCloudPacker20220411BucketAncestryStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPacker20220411BucketAncestryStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer 20220411 bucket ancestry status
func (m HashicorpCloudPacker20220411BucketAncestryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPacker20220411BucketAncestryStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20220411 bucket ancestry status based on context it is used
func (m HashicorpCloudPacker20220411BucketAncestryStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
