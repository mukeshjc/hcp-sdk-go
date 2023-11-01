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

// WaypointGetStatusReportReader is a Reader for the WaypointGetStatusReport structure.
type WaypointGetStatusReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetStatusReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetStatusReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetStatusReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetStatusReportOK creates a WaypointGetStatusReportOK with default headers values
func NewWaypointGetStatusReportOK() *WaypointGetStatusReportOK {
	return &WaypointGetStatusReportOK{}
}

/*
WaypointGetStatusReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetStatusReportOK struct {
	Payload *models.HashicorpWaypointStatusReport
}

// IsSuccess returns true when this waypoint get status report o k response has a 2xx status code
func (o *WaypointGetStatusReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get status report o k response has a 3xx status code
func (o *WaypointGetStatusReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get status report o k response has a 4xx status code
func (o *WaypointGetStatusReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get status report o k response has a 5xx status code
func (o *WaypointGetStatusReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get status report o k response a status code equal to that given
func (o *WaypointGetStatusReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get status report o k response
func (o *WaypointGetStatusReportOK) Code() int {
	return 200
}

func (o *WaypointGetStatusReportOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/by-id/{ref.id}][%d] waypointGetStatusReportOK  %+v", 200, o.Payload)
}

func (o *WaypointGetStatusReportOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/by-id/{ref.id}][%d] waypointGetStatusReportOK  %+v", 200, o.Payload)
}

func (o *WaypointGetStatusReportOK) GetPayload() *models.HashicorpWaypointStatusReport {
	return o.Payload
}

func (o *WaypointGetStatusReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointStatusReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetStatusReportDefault creates a WaypointGetStatusReportDefault with default headers values
func NewWaypointGetStatusReportDefault(code int) *WaypointGetStatusReportDefault {
	return &WaypointGetStatusReportDefault{
		_statusCode: code,
	}
}

/*
WaypointGetStatusReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetStatusReportDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get status report default response has a 2xx status code
func (o *WaypointGetStatusReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get status report default response has a 3xx status code
func (o *WaypointGetStatusReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get status report default response has a 4xx status code
func (o *WaypointGetStatusReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get status report default response has a 5xx status code
func (o *WaypointGetStatusReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get status report default response a status code equal to that given
func (o *WaypointGetStatusReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get status report default response
func (o *WaypointGetStatusReportDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetStatusReportDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/by-id/{ref.id}][%d] Waypoint_GetStatusReport default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetStatusReportDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/by-id/{ref.id}][%d] Waypoint_GetStatusReport default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetStatusReportDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetStatusReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
