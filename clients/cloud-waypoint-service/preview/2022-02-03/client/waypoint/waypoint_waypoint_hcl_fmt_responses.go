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

// WaypointWaypointHclFmtReader is a Reader for the WaypointWaypointHclFmt structure.
type WaypointWaypointHclFmtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointWaypointHclFmtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointWaypointHclFmtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointWaypointHclFmtDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointWaypointHclFmtOK creates a WaypointWaypointHclFmtOK with default headers values
func NewWaypointWaypointHclFmtOK() *WaypointWaypointHclFmtOK {
	return &WaypointWaypointHclFmtOK{}
}

/*
WaypointWaypointHclFmtOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointWaypointHclFmtOK struct {
	Payload *models.HashicorpWaypointWaypointHclFmtResponse
}

// IsSuccess returns true when this waypoint waypoint hcl fmt o k response has a 2xx status code
func (o *WaypointWaypointHclFmtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint waypoint hcl fmt o k response has a 3xx status code
func (o *WaypointWaypointHclFmtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint waypoint hcl fmt o k response has a 4xx status code
func (o *WaypointWaypointHclFmtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint waypoint hcl fmt o k response has a 5xx status code
func (o *WaypointWaypointHclFmtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint waypoint hcl fmt o k response a status code equal to that given
func (o *WaypointWaypointHclFmtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint waypoint hcl fmt o k response
func (o *WaypointWaypointHclFmtOK) Code() int {
	return 200
}

func (o *WaypointWaypointHclFmtOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/hcl/format][%d] waypointWaypointHclFmtOK  %+v", 200, o.Payload)
}

func (o *WaypointWaypointHclFmtOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/hcl/format][%d] waypointWaypointHclFmtOK  %+v", 200, o.Payload)
}

func (o *WaypointWaypointHclFmtOK) GetPayload() *models.HashicorpWaypointWaypointHclFmtResponse {
	return o.Payload
}

func (o *WaypointWaypointHclFmtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointWaypointHclFmtResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointWaypointHclFmtDefault creates a WaypointWaypointHclFmtDefault with default headers values
func NewWaypointWaypointHclFmtDefault(code int) *WaypointWaypointHclFmtDefault {
	return &WaypointWaypointHclFmtDefault{
		_statusCode: code,
	}
}

/*
WaypointWaypointHclFmtDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointWaypointHclFmtDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint waypoint hcl fmt default response has a 2xx status code
func (o *WaypointWaypointHclFmtDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint waypoint hcl fmt default response has a 3xx status code
func (o *WaypointWaypointHclFmtDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint waypoint hcl fmt default response has a 4xx status code
func (o *WaypointWaypointHclFmtDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint waypoint hcl fmt default response has a 5xx status code
func (o *WaypointWaypointHclFmtDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint waypoint hcl fmt default response a status code equal to that given
func (o *WaypointWaypointHclFmtDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint waypoint hcl fmt default response
func (o *WaypointWaypointHclFmtDefault) Code() int {
	return o._statusCode
}

func (o *WaypointWaypointHclFmtDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/hcl/format][%d] Waypoint_WaypointHclFmt default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointWaypointHclFmtDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/hcl/format][%d] Waypoint_WaypointHclFmt default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointWaypointHclFmtDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointWaypointHclFmtDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}