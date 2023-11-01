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

// WaypointQueueJobReader is a Reader for the WaypointQueueJob structure.
type WaypointQueueJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointQueueJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointQueueJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointQueueJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointQueueJobOK creates a WaypointQueueJobOK with default headers values
func NewWaypointQueueJobOK() *WaypointQueueJobOK {
	return &WaypointQueueJobOK{}
}

/*
WaypointQueueJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointQueueJobOK struct {
	Payload *models.HashicorpWaypointQueueJobResponse
}

// IsSuccess returns true when this waypoint queue job o k response has a 2xx status code
func (o *WaypointQueueJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint queue job o k response has a 3xx status code
func (o *WaypointQueueJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint queue job o k response has a 4xx status code
func (o *WaypointQueueJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint queue job o k response has a 5xx status code
func (o *WaypointQueueJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint queue job o k response a status code equal to that given
func (o *WaypointQueueJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint queue job o k response
func (o *WaypointQueueJobOK) Code() int {
	return 200
}

func (o *WaypointQueueJobOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/jobs/queue][%d] waypointQueueJobOK  %+v", 200, o.Payload)
}

func (o *WaypointQueueJobOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/jobs/queue][%d] waypointQueueJobOK  %+v", 200, o.Payload)
}

func (o *WaypointQueueJobOK) GetPayload() *models.HashicorpWaypointQueueJobResponse {
	return o.Payload
}

func (o *WaypointQueueJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointQueueJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointQueueJobDefault creates a WaypointQueueJobDefault with default headers values
func NewWaypointQueueJobDefault(code int) *WaypointQueueJobDefault {
	return &WaypointQueueJobDefault{
		_statusCode: code,
	}
}

/*
WaypointQueueJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointQueueJobDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint queue job default response has a 2xx status code
func (o *WaypointQueueJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint queue job default response has a 3xx status code
func (o *WaypointQueueJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint queue job default response has a 4xx status code
func (o *WaypointQueueJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint queue job default response has a 5xx status code
func (o *WaypointQueueJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint queue job default response a status code equal to that given
func (o *WaypointQueueJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint queue job default response
func (o *WaypointQueueJobDefault) Code() int {
	return o._statusCode
}

func (o *WaypointQueueJobDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/jobs/queue][%d] Waypoint_QueueJob default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointQueueJobDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/jobs/queue][%d] Waypoint_QueueJob default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointQueueJobDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointQueueJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
