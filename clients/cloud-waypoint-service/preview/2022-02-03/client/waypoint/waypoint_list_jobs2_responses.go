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

// WaypointListJobs2Reader is a Reader for the WaypointListJobs2 structure.
type WaypointListJobs2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListJobs2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListJobs2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListJobs2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListJobs2OK creates a WaypointListJobs2OK with default headers values
func NewWaypointListJobs2OK() *WaypointListJobs2OK {
	return &WaypointListJobs2OK{}
}

/*
WaypointListJobs2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListJobs2OK struct {
	Payload *models.HashicorpWaypointListJobsResponse
}

// IsSuccess returns true when this waypoint list jobs2 o k response has a 2xx status code
func (o *WaypointListJobs2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list jobs2 o k response has a 3xx status code
func (o *WaypointListJobs2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list jobs2 o k response has a 4xx status code
func (o *WaypointListJobs2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list jobs2 o k response has a 5xx status code
func (o *WaypointListJobs2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list jobs2 o k response a status code equal to that given
func (o *WaypointListJobs2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint list jobs2 o k response
func (o *WaypointListJobs2OK) Code() int {
	return 200
}

func (o *WaypointListJobs2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/workspace/{workspace.workspace}/state/{jobState}][%d] waypointListJobs2OK  %+v", 200, o.Payload)
}

func (o *WaypointListJobs2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/workspace/{workspace.workspace}/state/{jobState}][%d] waypointListJobs2OK  %+v", 200, o.Payload)
}

func (o *WaypointListJobs2OK) GetPayload() *models.HashicorpWaypointListJobsResponse {
	return o.Payload
}

func (o *WaypointListJobs2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListJobs2Default creates a WaypointListJobs2Default with default headers values
func NewWaypointListJobs2Default(code int) *WaypointListJobs2Default {
	return &WaypointListJobs2Default{
		_statusCode: code,
	}
}

/*
WaypointListJobs2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListJobs2Default struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint list jobs2 default response has a 2xx status code
func (o *WaypointListJobs2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list jobs2 default response has a 3xx status code
func (o *WaypointListJobs2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list jobs2 default response has a 4xx status code
func (o *WaypointListJobs2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list jobs2 default response has a 5xx status code
func (o *WaypointListJobs2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list jobs2 default response a status code equal to that given
func (o *WaypointListJobs2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint list jobs2 default response
func (o *WaypointListJobs2Default) Code() int {
	return o._statusCode
}

func (o *WaypointListJobs2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/workspace/{workspace.workspace}/state/{jobState}][%d] Waypoint_ListJobs2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListJobs2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/workspace/{workspace.workspace}/state/{jobState}][%d] Waypoint_ListJobs2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListJobs2Default) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListJobs2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
