// Code generated by go-swagger; DO NOT EDIT.

package terraform_cloud_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// TerraformCloudConfigListTFCProjectsReader is a Reader for the TerraformCloudConfigListTFCProjects structure.
type TerraformCloudConfigListTFCProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerraformCloudConfigListTFCProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTerraformCloudConfigListTFCProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTerraformCloudConfigListTFCProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTerraformCloudConfigListTFCProjectsOK creates a TerraformCloudConfigListTFCProjectsOK with default headers values
func NewTerraformCloudConfigListTFCProjectsOK() *TerraformCloudConfigListTFCProjectsOK {
	return &TerraformCloudConfigListTFCProjectsOK{}
}

/*
TerraformCloudConfigListTFCProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type TerraformCloudConfigListTFCProjectsOK struct {
	Payload *models.HashicorpCloudWaypointListTerraformCloudProjectsResponse
}

// IsSuccess returns true when this terraform cloud config list t f c projects o k response has a 2xx status code
func (o *TerraformCloudConfigListTFCProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this terraform cloud config list t f c projects o k response has a 3xx status code
func (o *TerraformCloudConfigListTFCProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this terraform cloud config list t f c projects o k response has a 4xx status code
func (o *TerraformCloudConfigListTFCProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this terraform cloud config list t f c projects o k response has a 5xx status code
func (o *TerraformCloudConfigListTFCProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this terraform cloud config list t f c projects o k response a status code equal to that given
func (o *TerraformCloudConfigListTFCProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the terraform cloud config list t f c projects o k response
func (o *TerraformCloudConfigListTFCProjectsOK) Code() int {
	return 200
}

func (o *TerraformCloudConfigListTFCProjectsOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/tfcprojects][%d] terraformCloudConfigListTFCProjectsOK  %+v", 200, o.Payload)
}

func (o *TerraformCloudConfigListTFCProjectsOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/tfcprojects][%d] terraformCloudConfigListTFCProjectsOK  %+v", 200, o.Payload)
}

func (o *TerraformCloudConfigListTFCProjectsOK) GetPayload() *models.HashicorpCloudWaypointListTerraformCloudProjectsResponse {
	return o.Payload
}

func (o *TerraformCloudConfigListTFCProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListTerraformCloudProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerraformCloudConfigListTFCProjectsDefault creates a TerraformCloudConfigListTFCProjectsDefault with default headers values
func NewTerraformCloudConfigListTFCProjectsDefault(code int) *TerraformCloudConfigListTFCProjectsDefault {
	return &TerraformCloudConfigListTFCProjectsDefault{
		_statusCode: code,
	}
}

/*
TerraformCloudConfigListTFCProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TerraformCloudConfigListTFCProjectsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this terraform cloud config list t f c projects default response has a 2xx status code
func (o *TerraformCloudConfigListTFCProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this terraform cloud config list t f c projects default response has a 3xx status code
func (o *TerraformCloudConfigListTFCProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this terraform cloud config list t f c projects default response has a 4xx status code
func (o *TerraformCloudConfigListTFCProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this terraform cloud config list t f c projects default response has a 5xx status code
func (o *TerraformCloudConfigListTFCProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this terraform cloud config list t f c projects default response a status code equal to that given
func (o *TerraformCloudConfigListTFCProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the terraform cloud config list t f c projects default response
func (o *TerraformCloudConfigListTFCProjectsDefault) Code() int {
	return o._statusCode
}

func (o *TerraformCloudConfigListTFCProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/tfcprojects][%d] TerraformCloudConfig_ListTFCProjects default  %+v", o._statusCode, o.Payload)
}

func (o *TerraformCloudConfigListTFCProjectsDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/tfcprojects][%d] TerraformCloudConfig_ListTFCProjects default  %+v", o._statusCode, o.Payload)
}

func (o *TerraformCloudConfigListTFCProjectsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *TerraformCloudConfigListTFCProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
