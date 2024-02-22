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

// PackerServiceGetBucketReader is a Reader for the PackerServiceGetBucket structure.
type PackerServiceGetBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetBucketOK creates a PackerServiceGetBucketOK with default headers values
func NewPackerServiceGetBucketOK() *PackerServiceGetBucketOK {
	return &PackerServiceGetBucketOK{}
}

/*
PackerServiceGetBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceGetBucketOK struct {
	Payload *models.HashicorpCloudPacker20230101GetBucketResponse
}

// IsSuccess returns true when this packer service get bucket o k response has a 2xx status code
func (o *PackerServiceGetBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service get bucket o k response has a 3xx status code
func (o *PackerServiceGetBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service get bucket o k response has a 4xx status code
func (o *PackerServiceGetBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service get bucket o k response has a 5xx status code
func (o *PackerServiceGetBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service get bucket o k response a status code equal to that given
func (o *PackerServiceGetBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service get bucket o k response
func (o *PackerServiceGetBucketOK) Code() int {
	return 200
}

func (o *PackerServiceGetBucketOK) Error() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}][%d] packerServiceGetBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetBucketOK) String() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}][%d] packerServiceGetBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetBucketOK) GetPayload() *models.HashicorpCloudPacker20230101GetBucketResponse {
	return o.Payload
}

func (o *PackerServiceGetBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101GetBucketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetBucketDefault creates a PackerServiceGetBucketDefault with default headers values
func NewPackerServiceGetBucketDefault(code int) *PackerServiceGetBucketDefault {
	return &PackerServiceGetBucketDefault{
		_statusCode: code,
	}
}

/*
PackerServiceGetBucketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceGetBucketDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service get bucket default response has a 2xx status code
func (o *PackerServiceGetBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service get bucket default response has a 3xx status code
func (o *PackerServiceGetBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service get bucket default response has a 4xx status code
func (o *PackerServiceGetBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service get bucket default response has a 5xx status code
func (o *PackerServiceGetBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service get bucket default response a status code equal to that given
func (o *PackerServiceGetBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service get bucket default response
func (o *PackerServiceGetBucketDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceGetBucketDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}][%d] PackerService_GetBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetBucketDefault) String() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}][%d] PackerService_GetBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetBucketDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceGetBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}