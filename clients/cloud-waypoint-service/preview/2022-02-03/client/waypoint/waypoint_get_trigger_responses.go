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

// WaypointGetTriggerReader is a Reader for the WaypointGetTrigger structure.
type WaypointGetTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetTriggerOK creates a WaypointGetTriggerOK with default headers values
func NewWaypointGetTriggerOK() *WaypointGetTriggerOK {
	return &WaypointGetTriggerOK{}
}

/*
WaypointGetTriggerOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetTriggerOK struct {
	Payload *models.HashicorpWaypointGetTriggerResponse
}

// IsSuccess returns true when this waypoint get trigger o k response has a 2xx status code
func (o *WaypointGetTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get trigger o k response has a 3xx status code
func (o *WaypointGetTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get trigger o k response has a 4xx status code
func (o *WaypointGetTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get trigger o k response has a 5xx status code
func (o *WaypointGetTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get trigger o k response a status code equal to that given
func (o *WaypointGetTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get trigger o k response
func (o *WaypointGetTriggerOK) Code() int {
	return 200
}

func (o *WaypointGetTriggerOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/trigger/{ref.id}][%d] waypointGetTriggerOK  %+v", 200, o.Payload)
}

func (o *WaypointGetTriggerOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/trigger/{ref.id}][%d] waypointGetTriggerOK  %+v", 200, o.Payload)
}

func (o *WaypointGetTriggerOK) GetPayload() *models.HashicorpWaypointGetTriggerResponse {
	return o.Payload
}

func (o *WaypointGetTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetTriggerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetTriggerDefault creates a WaypointGetTriggerDefault with default headers values
func NewWaypointGetTriggerDefault(code int) *WaypointGetTriggerDefault {
	return &WaypointGetTriggerDefault{
		_statusCode: code,
	}
}

/*
WaypointGetTriggerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetTriggerDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get trigger default response has a 2xx status code
func (o *WaypointGetTriggerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get trigger default response has a 3xx status code
func (o *WaypointGetTriggerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get trigger default response has a 4xx status code
func (o *WaypointGetTriggerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get trigger default response has a 5xx status code
func (o *WaypointGetTriggerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get trigger default response a status code equal to that given
func (o *WaypointGetTriggerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get trigger default response
func (o *WaypointGetTriggerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetTriggerDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/trigger/{ref.id}][%d] Waypoint_GetTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetTriggerDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/trigger/{ref.id}][%d] Waypoint_GetTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetTriggerDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}