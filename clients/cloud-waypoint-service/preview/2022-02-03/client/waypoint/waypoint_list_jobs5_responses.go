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

// WaypointListJobs5Reader is a Reader for the WaypointListJobs5 structure.
type WaypointListJobs5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListJobs5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListJobs5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListJobs5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListJobs5OK creates a WaypointListJobs5OK with default headers values
func NewWaypointListJobs5OK() *WaypointListJobs5OK {
	return &WaypointListJobs5OK{}
}

/*
WaypointListJobs5OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListJobs5OK struct {
	Payload *models.HashicorpWaypointListJobsResponse
}

// IsSuccess returns true when this waypoint list jobs5 o k response has a 2xx status code
func (o *WaypointListJobs5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list jobs5 o k response has a 3xx status code
func (o *WaypointListJobs5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list jobs5 o k response has a 4xx status code
func (o *WaypointListJobs5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list jobs5 o k response has a 5xx status code
func (o *WaypointListJobs5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list jobs5 o k response a status code equal to that given
func (o *WaypointListJobs5OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint list jobs5 o k response
func (o *WaypointListJobs5OK) Code() int {
	return 200
}

func (o *WaypointListJobs5OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/runner/by-id/{targetRunner.id.id}][%d] waypointListJobs5OK  %+v", 200, o.Payload)
}

func (o *WaypointListJobs5OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/runner/by-id/{targetRunner.id.id}][%d] waypointListJobs5OK  %+v", 200, o.Payload)
}

func (o *WaypointListJobs5OK) GetPayload() *models.HashicorpWaypointListJobsResponse {
	return o.Payload
}

func (o *WaypointListJobs5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListJobs5Default creates a WaypointListJobs5Default with default headers values
func NewWaypointListJobs5Default(code int) *WaypointListJobs5Default {
	return &WaypointListJobs5Default{
		_statusCode: code,
	}
}

/*
WaypointListJobs5Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListJobs5Default struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint list jobs5 default response has a 2xx status code
func (o *WaypointListJobs5Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list jobs5 default response has a 3xx status code
func (o *WaypointListJobs5Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list jobs5 default response has a 4xx status code
func (o *WaypointListJobs5Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list jobs5 default response has a 5xx status code
func (o *WaypointListJobs5Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list jobs5 default response a status code equal to that given
func (o *WaypointListJobs5Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint list jobs5 default response
func (o *WaypointListJobs5Default) Code() int {
	return o._statusCode
}

func (o *WaypointListJobs5Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/runner/by-id/{targetRunner.id.id}][%d] Waypoint_ListJobs5 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListJobs5Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/runner/by-id/{targetRunner.id.id}][%d] Waypoint_ListJobs5 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListJobs5Default) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListJobs5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
