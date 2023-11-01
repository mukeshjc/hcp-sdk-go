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

// WaypointUpsertTriggerReader is a Reader for the WaypointUpsertTrigger structure.
type WaypointUpsertTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpsertTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpsertTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpsertTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpsertTriggerOK creates a WaypointUpsertTriggerOK with default headers values
func NewWaypointUpsertTriggerOK() *WaypointUpsertTriggerOK {
	return &WaypointUpsertTriggerOK{}
}

/*
WaypointUpsertTriggerOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUpsertTriggerOK struct {
	Payload *models.HashicorpWaypointUpsertTriggerResponse
}

// IsSuccess returns true when this waypoint upsert trigger o k response has a 2xx status code
func (o *WaypointUpsertTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint upsert trigger o k response has a 3xx status code
func (o *WaypointUpsertTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint upsert trigger o k response has a 4xx status code
func (o *WaypointUpsertTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint upsert trigger o k response has a 5xx status code
func (o *WaypointUpsertTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint upsert trigger o k response a status code equal to that given
func (o *WaypointUpsertTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint upsert trigger o k response
func (o *WaypointUpsertTriggerOK) Code() int {
	return 200
}

func (o *WaypointUpsertTriggerOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] waypointUpsertTriggerOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertTriggerOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] waypointUpsertTriggerOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertTriggerOK) GetPayload() *models.HashicorpWaypointUpsertTriggerResponse {
	return o.Payload
}

func (o *WaypointUpsertTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpsertTriggerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpsertTriggerDefault creates a WaypointUpsertTriggerDefault with default headers values
func NewWaypointUpsertTriggerDefault(code int) *WaypointUpsertTriggerDefault {
	return &WaypointUpsertTriggerDefault{
		_statusCode: code,
	}
}

/*
WaypointUpsertTriggerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUpsertTriggerDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint upsert trigger default response has a 2xx status code
func (o *WaypointUpsertTriggerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint upsert trigger default response has a 3xx status code
func (o *WaypointUpsertTriggerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint upsert trigger default response has a 4xx status code
func (o *WaypointUpsertTriggerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint upsert trigger default response has a 5xx status code
func (o *WaypointUpsertTriggerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint upsert trigger default response a status code equal to that given
func (o *WaypointUpsertTriggerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint upsert trigger default response
func (o *WaypointUpsertTriggerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointUpsertTriggerDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] Waypoint_UpsertTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertTriggerDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] Waypoint_UpsertTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertTriggerDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpsertTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
