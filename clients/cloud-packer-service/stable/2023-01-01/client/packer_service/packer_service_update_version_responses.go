// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateVersionReader is a Reader for the PackerServiceUpdateVersion structure.
type PackerServiceUpdateVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateVersionOK creates a PackerServiceUpdateVersionOK with default headers values
func NewPackerServiceUpdateVersionOK() *PackerServiceUpdateVersionOK {
	return &PackerServiceUpdateVersionOK{}
}

/*
PackerServiceUpdateVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceUpdateVersionOK struct {
	Payload *models.HashicorpCloudPacker20230101UpdateVersionResponse
}

// IsSuccess returns true when this packer service update version o k response has a 2xx status code
func (o *PackerServiceUpdateVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service update version o k response has a 3xx status code
func (o *PackerServiceUpdateVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service update version o k response has a 4xx status code
func (o *PackerServiceUpdateVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service update version o k response has a 5xx status code
func (o *PackerServiceUpdateVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service update version o k response a status code equal to that given
func (o *PackerServiceUpdateVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service update version o k response
func (o *PackerServiceUpdateVersionOK) Code() int {
	return 200
}

func (o *PackerServiceUpdateVersionOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}][%d] packerServiceUpdateVersionOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateVersionOK) String() string {
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}][%d] packerServiceUpdateVersionOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateVersionOK) GetPayload() *models.HashicorpCloudPacker20230101UpdateVersionResponse {
	return o.Payload
}

func (o *PackerServiceUpdateVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101UpdateVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateVersionDefault creates a PackerServiceUpdateVersionDefault with default headers values
func NewPackerServiceUpdateVersionDefault(code int) *PackerServiceUpdateVersionDefault {
	return &PackerServiceUpdateVersionDefault{
		_statusCode: code,
	}
}

/*
PackerServiceUpdateVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceUpdateVersionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service update version default response has a 2xx status code
func (o *PackerServiceUpdateVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service update version default response has a 3xx status code
func (o *PackerServiceUpdateVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service update version default response has a 4xx status code
func (o *PackerServiceUpdateVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service update version default response has a 5xx status code
func (o *PackerServiceUpdateVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service update version default response a status code equal to that given
func (o *PackerServiceUpdateVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service update version default response
func (o *PackerServiceUpdateVersionDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceUpdateVersionDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}][%d] PackerService_UpdateVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateVersionDefault) String() string {
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}][%d] PackerService_UpdateVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateVersionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceUpdateVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}