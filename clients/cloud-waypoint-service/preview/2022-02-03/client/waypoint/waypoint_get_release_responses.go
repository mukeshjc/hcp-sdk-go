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

// WaypointGetReleaseReader is a Reader for the WaypointGetRelease structure.
type WaypointGetReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetReleaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetReleaseOK creates a WaypointGetReleaseOK with default headers values
func NewWaypointGetReleaseOK() *WaypointGetReleaseOK {
	return &WaypointGetReleaseOK{}
}

/*
WaypointGetReleaseOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetReleaseOK struct {
	Payload *models.HashicorpWaypointRelease
}

// IsSuccess returns true when this waypoint get release o k response has a 2xx status code
func (o *WaypointGetReleaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get release o k response has a 3xx status code
func (o *WaypointGetReleaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get release o k response has a 4xx status code
func (o *WaypointGetReleaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get release o k response has a 5xx status code
func (o *WaypointGetReleaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get release o k response a status code equal to that given
func (o *WaypointGetReleaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get release o k response
func (o *WaypointGetReleaseOK) Code() int {
	return 200
}

func (o *WaypointGetReleaseOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/{ref.id}][%d] waypointGetReleaseOK  %+v", 200, o.Payload)
}

func (o *WaypointGetReleaseOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/{ref.id}][%d] waypointGetReleaseOK  %+v", 200, o.Payload)
}

func (o *WaypointGetReleaseOK) GetPayload() *models.HashicorpWaypointRelease {
	return o.Payload
}

func (o *WaypointGetReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetReleaseDefault creates a WaypointGetReleaseDefault with default headers values
func NewWaypointGetReleaseDefault(code int) *WaypointGetReleaseDefault {
	return &WaypointGetReleaseDefault{
		_statusCode: code,
	}
}

/*
WaypointGetReleaseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetReleaseDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get release default response has a 2xx status code
func (o *WaypointGetReleaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get release default response has a 3xx status code
func (o *WaypointGetReleaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get release default response has a 4xx status code
func (o *WaypointGetReleaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get release default response has a 5xx status code
func (o *WaypointGetReleaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get release default response a status code equal to that given
func (o *WaypointGetReleaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get release default response
func (o *WaypointGetReleaseDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetReleaseDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/{ref.id}][%d] Waypoint_GetRelease default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetReleaseDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/release/{ref.id}][%d] Waypoint_GetRelease default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetReleaseDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetReleaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
