// Code generated by go-swagger; DO NOT EDIT.

package auth_config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// AuthConfigServiceDeleteAuthConnectionFromOrganizationReader is a Reader for the AuthConfigServiceDeleteAuthConnectionFromOrganization structure.
type AuthConfigServiceDeleteAuthConnectionFromOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthConfigServiceDeleteAuthConnectionFromOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthConfigServiceDeleteAuthConnectionFromOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthConfigServiceDeleteAuthConnectionFromOrganizationOK creates a AuthConfigServiceDeleteAuthConnectionFromOrganizationOK with default headers values
func NewAuthConfigServiceDeleteAuthConnectionFromOrganizationOK() *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK {
	return &AuthConfigServiceDeleteAuthConnectionFromOrganizationOK{}
}

/*
AuthConfigServiceDeleteAuthConnectionFromOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthConfigServiceDeleteAuthConnectionFromOrganizationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this auth config service delete auth connection from organization o k response has a 2xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth config service delete auth connection from organization o k response has a 3xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth config service delete auth connection from organization o k response has a 4xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth config service delete auth connection from organization o k response has a 5xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth config service delete auth connection from organization o k response a status code equal to that given
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auth config service delete auth connection from organization o k response
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) Code() int {
	return 200
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceDeleteAuthConnectionFromOrganizationOK  %+v", 200, o.Payload)
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceDeleteAuthConnectionFromOrganizationOK  %+v", 200, o.Payload)
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthConfigServiceDeleteAuthConnectionFromOrganizationDefault creates a AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault with default headers values
func NewAuthConfigServiceDeleteAuthConnectionFromOrganizationDefault(code int) *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault {
	return &AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault{
		_statusCode: code,
	}
}

/*
AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this auth config service delete auth connection from organization default response has a 2xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auth config service delete auth connection from organization default response has a 3xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auth config service delete auth connection from organization default response has a 4xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auth config service delete auth connection from organization default response has a 5xx status code
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auth config service delete auth connection from organization default response a status code equal to that given
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auth config service delete auth connection from organization default response
func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_DeleteAuthConnectionFromOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_DeleteAuthConnectionFromOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AuthConfigServiceDeleteAuthConnectionFromOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
