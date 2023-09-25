// Code generated by go-swagger; DO NOT EDIT.

package invitations_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// InvitationsServiceListOrganizationInvitationsReader is a Reader for the InvitationsServiceListOrganizationInvitations structure.
type InvitationsServiceListOrganizationInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvitationsServiceListOrganizationInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvitationsServiceListOrganizationInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInvitationsServiceListOrganizationInvitationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInvitationsServiceListOrganizationInvitationsOK creates a InvitationsServiceListOrganizationInvitationsOK with default headers values
func NewInvitationsServiceListOrganizationInvitationsOK() *InvitationsServiceListOrganizationInvitationsOK {
	return &InvitationsServiceListOrganizationInvitationsOK{}
}

/*
InvitationsServiceListOrganizationInvitationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type InvitationsServiceListOrganizationInvitationsOK struct {
	Payload *models.HashicorpCloudIamListOrganizationInvitationsResponse
}

// IsSuccess returns true when this invitations service list organization invitations o k response has a 2xx status code
func (o *InvitationsServiceListOrganizationInvitationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invitations service list organization invitations o k response has a 3xx status code
func (o *InvitationsServiceListOrganizationInvitationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invitations service list organization invitations o k response has a 4xx status code
func (o *InvitationsServiceListOrganizationInvitationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invitations service list organization invitations o k response has a 5xx status code
func (o *InvitationsServiceListOrganizationInvitationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invitations service list organization invitations o k response a status code equal to that given
func (o *InvitationsServiceListOrganizationInvitationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *InvitationsServiceListOrganizationInvitationsOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/invitations][%d] invitationsServiceListOrganizationInvitationsOK  %+v", 200, o.Payload)
}

func (o *InvitationsServiceListOrganizationInvitationsOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/invitations][%d] invitationsServiceListOrganizationInvitationsOK  %+v", 200, o.Payload)
}

func (o *InvitationsServiceListOrganizationInvitationsOK) GetPayload() *models.HashicorpCloudIamListOrganizationInvitationsResponse {
	return o.Payload
}

func (o *InvitationsServiceListOrganizationInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListOrganizationInvitationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvitationsServiceListOrganizationInvitationsDefault creates a InvitationsServiceListOrganizationInvitationsDefault with default headers values
func NewInvitationsServiceListOrganizationInvitationsDefault(code int) *InvitationsServiceListOrganizationInvitationsDefault {
	return &InvitationsServiceListOrganizationInvitationsDefault{
		_statusCode: code,
	}
}

/*
InvitationsServiceListOrganizationInvitationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InvitationsServiceListOrganizationInvitationsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the invitations service list organization invitations default response
func (o *InvitationsServiceListOrganizationInvitationsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this invitations service list organization invitations default response has a 2xx status code
func (o *InvitationsServiceListOrganizationInvitationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this invitations service list organization invitations default response has a 3xx status code
func (o *InvitationsServiceListOrganizationInvitationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this invitations service list organization invitations default response has a 4xx status code
func (o *InvitationsServiceListOrganizationInvitationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this invitations service list organization invitations default response has a 5xx status code
func (o *InvitationsServiceListOrganizationInvitationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this invitations service list organization invitations default response a status code equal to that given
func (o *InvitationsServiceListOrganizationInvitationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InvitationsServiceListOrganizationInvitationsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/invitations][%d] InvitationsService_ListOrganizationInvitations default  %+v", o._statusCode, o.Payload)
}

func (o *InvitationsServiceListOrganizationInvitationsDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/invitations][%d] InvitationsService_ListOrganizationInvitations default  %+v", o._statusCode, o.Payload)
}

func (o *InvitationsServiceListOrganizationInvitationsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *InvitationsServiceListOrganizationInvitationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}