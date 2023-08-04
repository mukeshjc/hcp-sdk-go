// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

// VerifyDomainOwnershipReader is a Reader for the VerifyDomainOwnership structure.
type VerifyDomainOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyDomainOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyDomainOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVerifyDomainOwnershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVerifyDomainOwnershipOK creates a VerifyDomainOwnershipOK with default headers values
func NewVerifyDomainOwnershipOK() *VerifyDomainOwnershipOK {
	return &VerifyDomainOwnershipOK{}
}

/*
VerifyDomainOwnershipOK describes a response with status code 200, with default header values.

A successful response.
*/
type VerifyDomainOwnershipOK struct {
	Payload *models.HashicorpCloudIamVerifyDomainOwnershipResponse
}

// IsSuccess returns true when this verify domain ownership o k response has a 2xx status code
func (o *VerifyDomainOwnershipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this verify domain ownership o k response has a 3xx status code
func (o *VerifyDomainOwnershipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify domain ownership o k response has a 4xx status code
func (o *VerifyDomainOwnershipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this verify domain ownership o k response has a 5xx status code
func (o *VerifyDomainOwnershipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this verify domain ownership o k response a status code equal to that given
func (o *VerifyDomainOwnershipOK) IsCode(code int) bool {
	return code == 200
}

func (o *VerifyDomainOwnershipOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-domain-ownership][%d] verifyDomainOwnershipOK  %+v", 200, o.Payload)
}

func (o *VerifyDomainOwnershipOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-domain-ownership][%d] verifyDomainOwnershipOK  %+v", 200, o.Payload)
}

func (o *VerifyDomainOwnershipOK) GetPayload() *models.HashicorpCloudIamVerifyDomainOwnershipResponse {
	return o.Payload
}

func (o *VerifyDomainOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamVerifyDomainOwnershipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyDomainOwnershipDefault creates a VerifyDomainOwnershipDefault with default headers values
func NewVerifyDomainOwnershipDefault(code int) *VerifyDomainOwnershipDefault {
	return &VerifyDomainOwnershipDefault{
		_statusCode: code,
	}
}

/*
VerifyDomainOwnershipDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VerifyDomainOwnershipDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the verify domain ownership default response
func (o *VerifyDomainOwnershipDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this verify domain ownership default response has a 2xx status code
func (o *VerifyDomainOwnershipDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this verify domain ownership default response has a 3xx status code
func (o *VerifyDomainOwnershipDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this verify domain ownership default response has a 4xx status code
func (o *VerifyDomainOwnershipDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this verify domain ownership default response has a 5xx status code
func (o *VerifyDomainOwnershipDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this verify domain ownership default response a status code equal to that given
func (o *VerifyDomainOwnershipDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VerifyDomainOwnershipDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-domain-ownership][%d] VerifyDomainOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *VerifyDomainOwnershipDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-domain-ownership][%d] VerifyDomainOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *VerifyDomainOwnershipDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *VerifyDomainOwnershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}