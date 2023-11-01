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

// WaypointGetRunnerReader is a Reader for the WaypointGetRunner structure.
type WaypointGetRunnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetRunnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetRunnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetRunnerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetRunnerOK creates a WaypointGetRunnerOK with default headers values
func NewWaypointGetRunnerOK() *WaypointGetRunnerOK {
	return &WaypointGetRunnerOK{}
}

/*
WaypointGetRunnerOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetRunnerOK struct {
	Payload *models.HashicorpWaypointRunner
}

// IsSuccess returns true when this waypoint get runner o k response has a 2xx status code
func (o *WaypointGetRunnerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get runner o k response has a 3xx status code
func (o *WaypointGetRunnerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get runner o k response has a 4xx status code
func (o *WaypointGetRunnerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get runner o k response has a 5xx status code
func (o *WaypointGetRunnerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get runner o k response a status code equal to that given
func (o *WaypointGetRunnerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get runner o k response
func (o *WaypointGetRunnerOK) Code() int {
	return 200
}

func (o *WaypointGetRunnerOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] waypointGetRunnerOK  %+v", 200, o.Payload)
}

func (o *WaypointGetRunnerOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] waypointGetRunnerOK  %+v", 200, o.Payload)
}

func (o *WaypointGetRunnerOK) GetPayload() *models.HashicorpWaypointRunner {
	return o.Payload
}

func (o *WaypointGetRunnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRunner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetRunnerDefault creates a WaypointGetRunnerDefault with default headers values
func NewWaypointGetRunnerDefault(code int) *WaypointGetRunnerDefault {
	return &WaypointGetRunnerDefault{
		_statusCode: code,
	}
}

/*
WaypointGetRunnerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetRunnerDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get runner default response has a 2xx status code
func (o *WaypointGetRunnerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get runner default response has a 3xx status code
func (o *WaypointGetRunnerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get runner default response has a 4xx status code
func (o *WaypointGetRunnerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get runner default response has a 5xx status code
func (o *WaypointGetRunnerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get runner default response a status code equal to that given
func (o *WaypointGetRunnerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get runner default response
func (o *WaypointGetRunnerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetRunnerDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] Waypoint_GetRunner default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetRunnerDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] Waypoint_GetRunner default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetRunnerDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetRunnerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
