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

// PackerServiceListBucketsReader is a Reader for the PackerServiceListBuckets structure.
type PackerServiceListBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceListBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceListBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceListBucketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceListBucketsOK creates a PackerServiceListBucketsOK with default headers values
func NewPackerServiceListBucketsOK() *PackerServiceListBucketsOK {
	return &PackerServiceListBucketsOK{}
}

/*
PackerServiceListBucketsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceListBucketsOK struct {
	Payload *models.HashicorpCloudPacker20230101ListBucketsResponse
}

// IsSuccess returns true when this packer service list buckets o k response has a 2xx status code
func (o *PackerServiceListBucketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service list buckets o k response has a 3xx status code
func (o *PackerServiceListBucketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service list buckets o k response has a 4xx status code
func (o *PackerServiceListBucketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service list buckets o k response has a 5xx status code
func (o *PackerServiceListBucketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service list buckets o k response a status code equal to that given
func (o *PackerServiceListBucketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service list buckets o k response
func (o *PackerServiceListBucketsOK) Code() int {
	return 200
}

func (o *PackerServiceListBucketsOK) Error() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets][%d] packerServiceListBucketsOK  %+v", 200, o.Payload)
}

func (o *PackerServiceListBucketsOK) String() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets][%d] packerServiceListBucketsOK  %+v", 200, o.Payload)
}

func (o *PackerServiceListBucketsOK) GetPayload() *models.HashicorpCloudPacker20230101ListBucketsResponse {
	return o.Payload
}

func (o *PackerServiceListBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101ListBucketsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceListBucketsDefault creates a PackerServiceListBucketsDefault with default headers values
func NewPackerServiceListBucketsDefault(code int) *PackerServiceListBucketsDefault {
	return &PackerServiceListBucketsDefault{
		_statusCode: code,
	}
}

/*
PackerServiceListBucketsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceListBucketsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service list buckets default response has a 2xx status code
func (o *PackerServiceListBucketsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service list buckets default response has a 3xx status code
func (o *PackerServiceListBucketsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service list buckets default response has a 4xx status code
func (o *PackerServiceListBucketsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service list buckets default response has a 5xx status code
func (o *PackerServiceListBucketsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service list buckets default response a status code equal to that given
func (o *PackerServiceListBucketsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service list buckets default response
func (o *PackerServiceListBucketsDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceListBucketsDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets][%d] PackerService_ListBuckets default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceListBucketsDefault) String() string {
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets][%d] PackerService_ListBuckets default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceListBucketsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceListBucketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}