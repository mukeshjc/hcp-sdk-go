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
)

// WaypointCancelTaskReader is a Reader for the WaypointCancelTask structure.
type WaypointCancelTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointCancelTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointCancelTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointCancelTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointCancelTaskOK creates a WaypointCancelTaskOK with default headers values
func NewWaypointCancelTaskOK() *WaypointCancelTaskOK {
	return &WaypointCancelTaskOK{}
}

/*
WaypointCancelTaskOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointCancelTaskOK struct {
	Payload interface{}
}

// IsSuccess returns true when this waypoint cancel task o k response has a 2xx status code
func (o *WaypointCancelTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint cancel task o k response has a 3xx status code
func (o *WaypointCancelTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint cancel task o k response has a 4xx status code
func (o *WaypointCancelTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint cancel task o k response has a 5xx status code
func (o *WaypointCancelTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint cancel task o k response a status code equal to that given
func (o *WaypointCancelTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint cancel task o k response
func (o *WaypointCancelTaskOK) Code() int {
	return 200
}

func (o *WaypointCancelTaskOK) Error() string {
	return fmt.Sprintf("[PUT /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}/cancel][%d] waypointCancelTaskOK  %+v", 200, o.Payload)
}

func (o *WaypointCancelTaskOK) String() string {
	return fmt.Sprintf("[PUT /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}/cancel][%d] waypointCancelTaskOK  %+v", 200, o.Payload)
}

func (o *WaypointCancelTaskOK) GetPayload() interface{} {
	return o.Payload
}

func (o *WaypointCancelTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointCancelTaskDefault creates a WaypointCancelTaskDefault with default headers values
func NewWaypointCancelTaskDefault(code int) *WaypointCancelTaskDefault {
	return &WaypointCancelTaskDefault{
		_statusCode: code,
	}
}

/*
WaypointCancelTaskDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointCancelTaskDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint cancel task default response has a 2xx status code
func (o *WaypointCancelTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint cancel task default response has a 3xx status code
func (o *WaypointCancelTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint cancel task default response has a 4xx status code
func (o *WaypointCancelTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint cancel task default response has a 5xx status code
func (o *WaypointCancelTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint cancel task default response a status code equal to that given
func (o *WaypointCancelTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint cancel task default response
func (o *WaypointCancelTaskDefault) Code() int {
	return o._statusCode
}

func (o *WaypointCancelTaskDefault) Error() string {
	return fmt.Sprintf("[PUT /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}/cancel][%d] Waypoint_CancelTask default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointCancelTaskDefault) String() string {
	return fmt.Sprintf("[PUT /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}/cancel][%d] Waypoint_CancelTask default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointCancelTaskDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointCancelTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
