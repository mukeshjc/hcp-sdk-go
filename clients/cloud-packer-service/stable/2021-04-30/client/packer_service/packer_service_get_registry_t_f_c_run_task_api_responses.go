// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceGetRegistryTFCRunTaskAPIReader is a Reader for the PackerServiceGetRegistryTFCRunTaskAPI structure.
type PackerServiceGetRegistryTFCRunTaskAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetRegistryTFCRunTaskAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetRegistryTFCRunTaskAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetRegistryTFCRunTaskAPIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetRegistryTFCRunTaskAPIOK creates a PackerServiceGetRegistryTFCRunTaskAPIOK with default headers values
func NewPackerServiceGetRegistryTFCRunTaskAPIOK() *PackerServiceGetRegistryTFCRunTaskAPIOK {
	return &PackerServiceGetRegistryTFCRunTaskAPIOK{}
}

/*
PackerServiceGetRegistryTFCRunTaskAPIOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceGetRegistryTFCRunTaskAPIOK struct {
	Payload *models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse
}

// IsSuccess returns true when this packer service get registry t f c run task Api o k response has a 2xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service get registry t f c run task Api o k response has a 3xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service get registry t f c run task Api o k response has a 4xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service get registry t f c run task Api o k response has a 5xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service get registry t f c run task Api o k response a status code equal to that given
func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] packerServiceGetRegistryTFCRunTaskApiOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] packerServiceGetRegistryTFCRunTaskApiOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) GetPayload() *models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse {
	return o.Payload
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetRegistryTFCRunTaskAPIDefault creates a PackerServiceGetRegistryTFCRunTaskAPIDefault with default headers values
func NewPackerServiceGetRegistryTFCRunTaskAPIDefault(code int) *PackerServiceGetRegistryTFCRunTaskAPIDefault {
	return &PackerServiceGetRegistryTFCRunTaskAPIDefault{
		_statusCode: code,
	}
}

/*
PackerServiceGetRegistryTFCRunTaskAPIDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceGetRegistryTFCRunTaskAPIDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service get registry t f c run task API default response
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service get registry t f c run task API default response has a 2xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service get registry t f c run task API default response has a 3xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service get registry t f c run task API default response has a 4xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service get registry t f c run task API default response has a 5xx status code
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service get registry t f c run task API default response a status code equal to that given
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] PackerService_GetRegistryTFCRunTaskAPI default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] PackerService_GetRegistryTFCRunTaskAPI default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
