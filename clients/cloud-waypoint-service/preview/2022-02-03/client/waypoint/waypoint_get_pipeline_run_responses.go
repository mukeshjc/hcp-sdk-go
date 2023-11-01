// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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

// WaypointGetPipelineRunReader is a Reader for the WaypointGetPipelineRun structure.
type WaypointGetPipelineRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetPipelineRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetPipelineRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetPipelineRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetPipelineRunOK creates a WaypointGetPipelineRunOK with default headers values
func NewWaypointGetPipelineRunOK() *WaypointGetPipelineRunOK {
	return &WaypointGetPipelineRunOK{}
}

/*
WaypointGetPipelineRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetPipelineRunOK struct {
	Payload *models.HashicorpWaypointGetPipelineRunResponse
}

// IsSuccess returns true when this waypoint get pipeline run o k response has a 2xx status code
func (o *WaypointGetPipelineRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get pipeline run o k response has a 3xx status code
func (o *WaypointGetPipelineRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get pipeline run o k response has a 4xx status code
func (o *WaypointGetPipelineRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get pipeline run o k response has a 5xx status code
func (o *WaypointGetPipelineRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get pipeline run o k response a status code equal to that given
func (o *WaypointGetPipelineRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get pipeline run o k response
func (o *WaypointGetPipelineRunOK) Code() int {
	return 200
}

func (o *WaypointGetPipelineRunOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run/{sequence}][%d] waypointGetPipelineRunOK  %+v", 200, o.Payload)
}

func (o *WaypointGetPipelineRunOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run/{sequence}][%d] waypointGetPipelineRunOK  %+v", 200, o.Payload)
}

func (o *WaypointGetPipelineRunOK) GetPayload() *models.HashicorpWaypointGetPipelineRunResponse {
	return o.Payload
}

func (o *WaypointGetPipelineRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetPipelineRunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetPipelineRunDefault creates a WaypointGetPipelineRunDefault with default headers values
func NewWaypointGetPipelineRunDefault(code int) *WaypointGetPipelineRunDefault {
	return &WaypointGetPipelineRunDefault{
		_statusCode: code,
	}
}

/*
WaypointGetPipelineRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetPipelineRunDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get pipeline run default response has a 2xx status code
func (o *WaypointGetPipelineRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get pipeline run default response has a 3xx status code
func (o *WaypointGetPipelineRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get pipeline run default response has a 4xx status code
func (o *WaypointGetPipelineRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get pipeline run default response has a 5xx status code
func (o *WaypointGetPipelineRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get pipeline run default response a status code equal to that given
func (o *WaypointGetPipelineRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get pipeline run default response
func (o *WaypointGetPipelineRunDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetPipelineRunDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run/{sequence}][%d] Waypoint_GetPipelineRun default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetPipelineRunDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run/{sequence}][%d] Waypoint_GetPipelineRun default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetPipelineRunDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetPipelineRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
