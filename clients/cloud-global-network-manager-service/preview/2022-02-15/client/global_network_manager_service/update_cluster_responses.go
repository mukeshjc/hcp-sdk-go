// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/*
UpdateClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateClusterOK struct {
	Payload models.HashicorpCloudGlobalNetworkManager20220215UpdateClusterResponse
}

// IsSuccess returns true when this update cluster o k response has a 2xx status code
func (o *UpdateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cluster o k response has a 3xx status code
func (o *UpdateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster o k response has a 4xx status code
func (o *UpdateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster o k response has a 5xx status code
func (o *UpdateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster o k response a status code equal to that given
func (o *UpdateClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[PATCH /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] updateClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterOK) String() string {
	return fmt.Sprintf("[PATCH /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] updateClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterOK) GetPayload() models.HashicorpCloudGlobalNetworkManager20220215UpdateClusterResponse {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterDefault creates a UpdateClusterDefault with default headers values
func NewUpdateClusterDefault(code int) *UpdateClusterDefault {
	return &UpdateClusterDefault{
		_statusCode: code,
	}
}

/*
UpdateClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateClusterDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the update cluster default response
func (o *UpdateClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update cluster default response has a 2xx status code
func (o *UpdateClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update cluster default response has a 3xx status code
func (o *UpdateClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update cluster default response has a 4xx status code
func (o *UpdateClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update cluster default response has a 5xx status code
func (o *UpdateClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update cluster default response a status code equal to that given
func (o *UpdateClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateClusterDefault) Error() string {
	return fmt.Sprintf("[PATCH /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] UpdateCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateClusterDefault) String() string {
	return fmt.Sprintf("[PATCH /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] UpdateCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateClusterDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *UpdateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateClusterBody UpdateClusterRequest contains the individual fields of the cluster to update.
// Fields are added to the request object as they are supported.
//
// The request does not contain the `Cluster` object itself since there are
// cluster fields we want to support updateing but do not want to expose
// publically in other APIs that also use `Cluster` i.e. we do not want to
// add `Cluster.management_token` and start returning the management_token in
// the GetCluster API.
swagger:model UpdateClusterBody
*/
type UpdateClusterBody struct {

	// consul_access_level is a cluster field to update
	ConsulAccessLevel *models.HashicorpCloudGlobalNetworkManager20220215ClusterConsulAccessLevel `json:"consul_access_level,omitempty"`

	// location
	Location *UpdateClusterParamsBodyLocation `json:"location,omitempty"`

	// management_token is a cluster-related field to update
	ManagementToken string `json:"management_token,omitempty"`

	// update_mask is the mask of fields to update
	UpdateMask string `json:"update_mask,omitempty"`
}

// Validate validates this update cluster body
func (o *UpdateClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConsulAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateClusterBody) validateConsulAccessLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.ConsulAccessLevel) { // not required
		return nil
	}

	if o.ConsulAccessLevel != nil {
		if err := o.ConsulAccessLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "consul_access_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "consul_access_level")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateClusterBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update cluster body based on the context it is used
func (o *UpdateClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConsulAccessLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateClusterBody) contextValidateConsulAccessLevel(ctx context.Context, formats strfmt.Registry) error {

	if o.ConsulAccessLevel != nil {
		if err := o.ConsulAccessLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "consul_access_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "consul_access_level")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateClusterBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {
		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateClusterBody) UnmarshalBinary(b []byte) error {
	var res UpdateClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateClusterParamsBodyLocation location is the location of the cluster to update with an optional provider
// and region
swagger:model UpdateClusterParamsBodyLocation
*/
type UpdateClusterParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this update cluster params body location
func (o *UpdateClusterParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateClusterParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update cluster params body location based on the context it is used
func (o *UpdateClusterParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateClusterParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateClusterParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateClusterParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpdateClusterParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}