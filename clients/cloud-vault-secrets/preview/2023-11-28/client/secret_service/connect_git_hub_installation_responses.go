// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// ConnectGitHubInstallationReader is a Reader for the ConnectGitHubInstallation structure.
type ConnectGitHubInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectGitHubInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectGitHubInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectGitHubInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectGitHubInstallationOK creates a ConnectGitHubInstallationOK with default headers values
func NewConnectGitHubInstallationOK() *ConnectGitHubInstallationOK {
	return &ConnectGitHubInstallationOK{}
}

/*
ConnectGitHubInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConnectGitHubInstallationOK struct {
	Payload *models.Secrets20231128ConnectGitHubInstallationResponse
}

// IsSuccess returns true when this connect git hub installation o k response has a 2xx status code
func (o *ConnectGitHubInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this connect git hub installation o k response has a 3xx status code
func (o *ConnectGitHubInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect git hub installation o k response has a 4xx status code
func (o *ConnectGitHubInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this connect git hub installation o k response has a 5xx status code
func (o *ConnectGitHubInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this connect git hub installation o k response a status code equal to that given
func (o *ConnectGitHubInstallationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the connect git hub installation o k response
func (o *ConnectGitHubInstallationOK) Code() int {
	return 200
}

func (o *ConnectGitHubInstallationOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/github/installations][%d] connectGitHubInstallationOK  %+v", 200, o.Payload)
}

func (o *ConnectGitHubInstallationOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/github/installations][%d] connectGitHubInstallationOK  %+v", 200, o.Payload)
}

func (o *ConnectGitHubInstallationOK) GetPayload() *models.Secrets20231128ConnectGitHubInstallationResponse {
	return o.Payload
}

func (o *ConnectGitHubInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ConnectGitHubInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectGitHubInstallationDefault creates a ConnectGitHubInstallationDefault with default headers values
func NewConnectGitHubInstallationDefault(code int) *ConnectGitHubInstallationDefault {
	return &ConnectGitHubInstallationDefault{
		_statusCode: code,
	}
}

/*
ConnectGitHubInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConnectGitHubInstallationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this connect git hub installation default response has a 2xx status code
func (o *ConnectGitHubInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this connect git hub installation default response has a 3xx status code
func (o *ConnectGitHubInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this connect git hub installation default response has a 4xx status code
func (o *ConnectGitHubInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this connect git hub installation default response has a 5xx status code
func (o *ConnectGitHubInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this connect git hub installation default response a status code equal to that given
func (o *ConnectGitHubInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the connect git hub installation default response
func (o *ConnectGitHubInstallationDefault) Code() int {
	return o._statusCode
}

func (o *ConnectGitHubInstallationDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/github/installations][%d] ConnectGitHubInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectGitHubInstallationDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/github/installations][%d] ConnectGitHubInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectGitHubInstallationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ConnectGitHubInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ConnectGitHubInstallationBody connect git hub installation body
swagger:model ConnectGitHubInstallationBody
*/
type ConnectGitHubInstallationBody struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this connect git hub installation body
func (o *ConnectGitHubInstallationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connect git hub installation body based on context it is used
func (o *ConnectGitHubInstallationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectGitHubInstallationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectGitHubInstallationBody) UnmarshalBinary(b []byte) error {
	var res ConnectGitHubInstallationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
