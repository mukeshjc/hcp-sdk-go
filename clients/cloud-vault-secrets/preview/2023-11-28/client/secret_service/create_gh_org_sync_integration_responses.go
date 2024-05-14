// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// CreateGhOrgSyncIntegrationReader is a Reader for the CreateGhOrgSyncIntegration structure.
type CreateGhOrgSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGhOrgSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGhOrgSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateGhOrgSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGhOrgSyncIntegrationOK creates a CreateGhOrgSyncIntegrationOK with default headers values
func NewCreateGhOrgSyncIntegrationOK() *CreateGhOrgSyncIntegrationOK {
	return &CreateGhOrgSyncIntegrationOK{}
}

/*
CreateGhOrgSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateGhOrgSyncIntegrationOK struct {
	Payload *models.Secrets20231128CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create gh org sync integration o k response has a 2xx status code
func (o *CreateGhOrgSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create gh org sync integration o k response has a 3xx status code
func (o *CreateGhOrgSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gh org sync integration o k response has a 4xx status code
func (o *CreateGhOrgSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gh org sync integration o k response has a 5xx status code
func (o *CreateGhOrgSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create gh org sync integration o k response a status code equal to that given
func (o *CreateGhOrgSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create gh org sync integration o k response
func (o *CreateGhOrgSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateGhOrgSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/gh-org][%d] createGhOrgSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateGhOrgSyncIntegrationOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/gh-org][%d] createGhOrgSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateGhOrgSyncIntegrationOK) GetPayload() *models.Secrets20231128CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateGhOrgSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGhOrgSyncIntegrationDefault creates a CreateGhOrgSyncIntegrationDefault with default headers values
func NewCreateGhOrgSyncIntegrationDefault(code int) *CreateGhOrgSyncIntegrationDefault {
	return &CreateGhOrgSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateGhOrgSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateGhOrgSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create gh org sync integration default response has a 2xx status code
func (o *CreateGhOrgSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create gh org sync integration default response has a 3xx status code
func (o *CreateGhOrgSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create gh org sync integration default response has a 4xx status code
func (o *CreateGhOrgSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create gh org sync integration default response has a 5xx status code
func (o *CreateGhOrgSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create gh org sync integration default response a status code equal to that given
func (o *CreateGhOrgSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create gh org sync integration default response
func (o *CreateGhOrgSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateGhOrgSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/gh-org][%d] CreateGhOrgSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGhOrgSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/gh-org][%d] CreateGhOrgSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGhOrgSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateGhOrgSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateGhOrgSyncIntegrationBody create gh org sync integration body
swagger:model CreateGhOrgSyncIntegrationBody
*/
type CreateGhOrgSyncIntegrationBody struct {

	// gh org connection details
	GhOrgConnectionDetails *models.Secrets20231128GhOrgConnectionDetailsRequest `json:"gh_org_connection_details,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create gh org sync integration body
func (o *CreateGhOrgSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGhOrgConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGhOrgSyncIntegrationBody) validateGhOrgConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.GhOrgConnectionDetails) { // not required
		return nil
	}

	if o.GhOrgConnectionDetails != nil {
		if err := o.GhOrgConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gh_org_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gh_org_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create gh org sync integration body based on the context it is used
func (o *CreateGhOrgSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGhOrgConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGhOrgSyncIntegrationBody) contextValidateGhOrgConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.GhOrgConnectionDetails != nil {

		if swag.IsZero(o.GhOrgConnectionDetails) { // not required
			return nil
		}

		if err := o.GhOrgConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gh_org_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gh_org_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateGhOrgSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGhOrgSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateGhOrgSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
