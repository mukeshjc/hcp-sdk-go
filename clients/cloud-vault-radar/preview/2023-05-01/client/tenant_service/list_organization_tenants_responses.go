// Code generated by go-swagger; DO NOT EDIT.

package tenant_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// ListOrganizationTenantsReader is a Reader for the ListOrganizationTenants structure.
type ListOrganizationTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOrganizationTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationTenantsOK creates a ListOrganizationTenantsOK with default headers values
func NewListOrganizationTenantsOK() *ListOrganizationTenantsOK {
	return &ListOrganizationTenantsOK{}
}

/*
ListOrganizationTenantsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListOrganizationTenantsOK struct {
	Payload *models.VaultRadar20230501ListOrganizationTenantsResponse
}

// IsSuccess returns true when this list organization tenants o k response has a 2xx status code
func (o *ListOrganizationTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organization tenants o k response has a 3xx status code
func (o *ListOrganizationTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organization tenants o k response has a 4xx status code
func (o *ListOrganizationTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organization tenants o k response has a 5xx status code
func (o *ListOrganizationTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list organization tenants o k response a status code equal to that given
func (o *ListOrganizationTenantsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list organization tenants o k response
func (o *ListOrganizationTenantsOK) Code() int {
	return 200
}

func (o *ListOrganizationTenantsOK) Error() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/tenants/organizations/{location.organization_id}][%d] listOrganizationTenantsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationTenantsOK) String() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/tenants/organizations/{location.organization_id}][%d] listOrganizationTenantsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationTenantsOK) GetPayload() *models.VaultRadar20230501ListOrganizationTenantsResponse {
	return o.Payload
}

func (o *ListOrganizationTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultRadar20230501ListOrganizationTenantsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationTenantsDefault creates a ListOrganizationTenantsDefault with default headers values
func NewListOrganizationTenantsDefault(code int) *ListOrganizationTenantsDefault {
	return &ListOrganizationTenantsDefault{
		_statusCode: code,
	}
}

/*
ListOrganizationTenantsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListOrganizationTenantsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this list organization tenants default response has a 2xx status code
func (o *ListOrganizationTenantsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list organization tenants default response has a 3xx status code
func (o *ListOrganizationTenantsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list organization tenants default response has a 4xx status code
func (o *ListOrganizationTenantsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list organization tenants default response has a 5xx status code
func (o *ListOrganizationTenantsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list organization tenants default response a status code equal to that given
func (o *ListOrganizationTenantsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list organization tenants default response
func (o *ListOrganizationTenantsDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationTenantsDefault) Error() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/tenants/organizations/{location.organization_id}][%d] ListOrganizationTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationTenantsDefault) String() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/tenants/organizations/{location.organization_id}][%d] ListOrganizationTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationTenantsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ListOrganizationTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}