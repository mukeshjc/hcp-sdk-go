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

// WaypointGenerateRunnerTokenReader is a Reader for the WaypointGenerateRunnerToken structure.
type WaypointGenerateRunnerTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGenerateRunnerTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGenerateRunnerTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGenerateRunnerTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGenerateRunnerTokenOK creates a WaypointGenerateRunnerTokenOK with default headers values
func NewWaypointGenerateRunnerTokenOK() *WaypointGenerateRunnerTokenOK {
	return &WaypointGenerateRunnerTokenOK{}
}

/*
WaypointGenerateRunnerTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGenerateRunnerTokenOK struct {
	Payload *models.HashicorpWaypointNewTokenResponse
}

// IsSuccess returns true when this waypoint generate runner token o k response has a 2xx status code
func (o *WaypointGenerateRunnerTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint generate runner token o k response has a 3xx status code
func (o *WaypointGenerateRunnerTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint generate runner token o k response has a 4xx status code
func (o *WaypointGenerateRunnerTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint generate runner token o k response has a 5xx status code
func (o *WaypointGenerateRunnerTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint generate runner token o k response a status code equal to that given
func (o *WaypointGenerateRunnerTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint generate runner token o k response
func (o *WaypointGenerateRunnerTokenOK) Code() int {
	return 200
}

func (o *WaypointGenerateRunnerTokenOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/runner][%d] waypointGenerateRunnerTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointGenerateRunnerTokenOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/runner][%d] waypointGenerateRunnerTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointGenerateRunnerTokenOK) GetPayload() *models.HashicorpWaypointNewTokenResponse {
	return o.Payload
}

func (o *WaypointGenerateRunnerTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointNewTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGenerateRunnerTokenDefault creates a WaypointGenerateRunnerTokenDefault with default headers values
func NewWaypointGenerateRunnerTokenDefault(code int) *WaypointGenerateRunnerTokenDefault {
	return &WaypointGenerateRunnerTokenDefault{
		_statusCode: code,
	}
}

/*
WaypointGenerateRunnerTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGenerateRunnerTokenDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint generate runner token default response has a 2xx status code
func (o *WaypointGenerateRunnerTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint generate runner token default response has a 3xx status code
func (o *WaypointGenerateRunnerTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint generate runner token default response has a 4xx status code
func (o *WaypointGenerateRunnerTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint generate runner token default response has a 5xx status code
func (o *WaypointGenerateRunnerTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint generate runner token default response a status code equal to that given
func (o *WaypointGenerateRunnerTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint generate runner token default response
func (o *WaypointGenerateRunnerTokenDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGenerateRunnerTokenDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/runner][%d] Waypoint_GenerateRunnerToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGenerateRunnerTokenDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/token/runner][%d] Waypoint_GenerateRunnerToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGenerateRunnerTokenDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGenerateRunnerTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}