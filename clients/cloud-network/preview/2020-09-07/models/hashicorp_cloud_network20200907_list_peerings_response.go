// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudNetwork20200907ListPeeringsResponse ListPeeringsResponse is a response type for ListPeerings endpoint
//
// swagger:model hashicorp.cloud.network_20200907.ListPeeringsResponse
type HashicorpCloudNetwork20200907ListPeeringsResponse struct {

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// Peerings is a list of peerings
	Peerings []*HashicorpCloudNetwork20200907Peering `json:"peerings"`
}

// Validate validates this hashicorp cloud network 20200907 list peerings response
func (m *HashicorpCloudNetwork20200907ListPeeringsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907ListPeeringsResponse) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907ListPeeringsResponse) validatePeerings(formats strfmt.Registry) error {

	if swag.IsZero(m.Peerings) { // not required
		return nil
	}

	for i := 0; i < len(m.Peerings); i++ {
		if swag.IsZero(m.Peerings[i]) { // not required
			continue
		}

		if m.Peerings[i] != nil {
			if err := m.Peerings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peerings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ListPeeringsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ListPeeringsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907ListPeeringsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
