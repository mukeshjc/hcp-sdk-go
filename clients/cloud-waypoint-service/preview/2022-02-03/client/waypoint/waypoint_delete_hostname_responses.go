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

// WaypointDeleteHostnameReader is a Reader for the WaypointDeleteHostname structure.
type WaypointDeleteHostnameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointDeleteHostnameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointDeleteHostnameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointDeleteHostnameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointDeleteHostnameOK creates a WaypointDeleteHostnameOK with default headers values
func NewWaypointDeleteHostnameOK() *WaypointDeleteHostnameOK {
	return &WaypointDeleteHostnameOK{}
}

/*
WaypointDeleteHostnameOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointDeleteHostnameOK struct {
	Payload interface{}
}

// IsSuccess returns true when this waypoint delete hostname o k response has a 2xx status code
func (o *WaypointDeleteHostnameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint delete hostname o k response has a 3xx status code
func (o *WaypointDeleteHostnameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint delete hostname o k response has a 4xx status code
func (o *WaypointDeleteHostnameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint delete hostname o k response has a 5xx status code
func (o *WaypointDeleteHostnameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint delete hostname o k response a status code equal to that given
func (o *WaypointDeleteHostnameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint delete hostname o k response
func (o *WaypointDeleteHostnameOK) Code() int {
	return 200
}

func (o *WaypointDeleteHostnameOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/hostname/{hostname}][%d] waypointDeleteHostnameOK  %+v", 200, o.Payload)
}

func (o *WaypointDeleteHostnameOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/hostname/{hostname}][%d] waypointDeleteHostnameOK  %+v", 200, o.Payload)
}

func (o *WaypointDeleteHostnameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *WaypointDeleteHostnameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointDeleteHostnameDefault creates a WaypointDeleteHostnameDefault with default headers values
func NewWaypointDeleteHostnameDefault(code int) *WaypointDeleteHostnameDefault {
	return &WaypointDeleteHostnameDefault{
		_statusCode: code,
	}
}

/*
WaypointDeleteHostnameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointDeleteHostnameDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint delete hostname default response has a 2xx status code
func (o *WaypointDeleteHostnameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint delete hostname default response has a 3xx status code
func (o *WaypointDeleteHostnameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint delete hostname default response has a 4xx status code
func (o *WaypointDeleteHostnameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint delete hostname default response has a 5xx status code
func (o *WaypointDeleteHostnameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint delete hostname default response a status code equal to that given
func (o *WaypointDeleteHostnameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint delete hostname default response
func (o *WaypointDeleteHostnameDefault) Code() int {
	return o._statusCode
}

func (o *WaypointDeleteHostnameDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/hostname/{hostname}][%d] Waypoint_DeleteHostname default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDeleteHostnameDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/hostname/{hostname}][%d] Waypoint_DeleteHostname default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDeleteHostnameDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointDeleteHostnameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
