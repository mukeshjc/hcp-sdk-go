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

// WaypointConvertInviteTokenReader is a Reader for the WaypointConvertInviteToken structure.
type WaypointConvertInviteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointConvertInviteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointConvertInviteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointConvertInviteTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointConvertInviteTokenOK creates a WaypointConvertInviteTokenOK with default headers values
func NewWaypointConvertInviteTokenOK() *WaypointConvertInviteTokenOK {
	return &WaypointConvertInviteTokenOK{}
}

/*
WaypointConvertInviteTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointConvertInviteTokenOK struct {
	Payload *models.HashicorpWaypointNewTokenResponse
}

// IsSuccess returns true when this waypoint convert invite token o k response has a 2xx status code
func (o *WaypointConvertInviteTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint convert invite token o k response has a 3xx status code
func (o *WaypointConvertInviteTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint convert invite token o k response has a 4xx status code
func (o *WaypointConvertInviteTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint convert invite token o k response has a 5xx status code
func (o *WaypointConvertInviteTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint convert invite token o k response a status code equal to that given
func (o *WaypointConvertInviteTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint convert invite token o k response
func (o *WaypointConvertInviteTokenOK) Code() int {
	return 200
}

func (o *WaypointConvertInviteTokenOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/exchange][%d] waypointConvertInviteTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointConvertInviteTokenOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/exchange][%d] waypointConvertInviteTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointConvertInviteTokenOK) GetPayload() *models.HashicorpWaypointNewTokenResponse {
	return o.Payload
}

func (o *WaypointConvertInviteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointNewTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointConvertInviteTokenDefault creates a WaypointConvertInviteTokenDefault with default headers values
func NewWaypointConvertInviteTokenDefault(code int) *WaypointConvertInviteTokenDefault {
	return &WaypointConvertInviteTokenDefault{
		_statusCode: code,
	}
}

/*
WaypointConvertInviteTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointConvertInviteTokenDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint convert invite token default response has a 2xx status code
func (o *WaypointConvertInviteTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint convert invite token default response has a 3xx status code
func (o *WaypointConvertInviteTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint convert invite token default response has a 4xx status code
func (o *WaypointConvertInviteTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint convert invite token default response has a 5xx status code
func (o *WaypointConvertInviteTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint convert invite token default response a status code equal to that given
func (o *WaypointConvertInviteTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint convert invite token default response
func (o *WaypointConvertInviteTokenDefault) Code() int {
	return o._statusCode
}

func (o *WaypointConvertInviteTokenDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/exchange][%d] Waypoint_ConvertInviteToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointConvertInviteTokenDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/exchange][%d] Waypoint_ConvertInviteToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointConvertInviteTokenDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointConvertInviteTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
