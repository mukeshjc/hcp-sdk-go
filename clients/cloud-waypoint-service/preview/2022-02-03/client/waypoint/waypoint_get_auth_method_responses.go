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

// WaypointGetAuthMethodReader is a Reader for the WaypointGetAuthMethod structure.
type WaypointGetAuthMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetAuthMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetAuthMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetAuthMethodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetAuthMethodOK creates a WaypointGetAuthMethodOK with default headers values
func NewWaypointGetAuthMethodOK() *WaypointGetAuthMethodOK {
	return &WaypointGetAuthMethodOK{}
}

/*
WaypointGetAuthMethodOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetAuthMethodOK struct {
	Payload *models.HashicorpWaypointGetAuthMethodResponse
}

// IsSuccess returns true when this waypoint get auth method o k response has a 2xx status code
func (o *WaypointGetAuthMethodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get auth method o k response has a 3xx status code
func (o *WaypointGetAuthMethodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get auth method o k response has a 4xx status code
func (o *WaypointGetAuthMethodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get auth method o k response has a 5xx status code
func (o *WaypointGetAuthMethodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get auth method o k response a status code equal to that given
func (o *WaypointGetAuthMethodOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get auth method o k response
func (o *WaypointGetAuthMethodOK) Code() int {
	return 200
}

func (o *WaypointGetAuthMethodOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/auth-method/{auth_method.name}][%d] waypointGetAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointGetAuthMethodOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/auth-method/{auth_method.name}][%d] waypointGetAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointGetAuthMethodOK) GetPayload() *models.HashicorpWaypointGetAuthMethodResponse {
	return o.Payload
}

func (o *WaypointGetAuthMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetAuthMethodResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetAuthMethodDefault creates a WaypointGetAuthMethodDefault with default headers values
func NewWaypointGetAuthMethodDefault(code int) *WaypointGetAuthMethodDefault {
	return &WaypointGetAuthMethodDefault{
		_statusCode: code,
	}
}

/*
WaypointGetAuthMethodDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetAuthMethodDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get auth method default response has a 2xx status code
func (o *WaypointGetAuthMethodDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get auth method default response has a 3xx status code
func (o *WaypointGetAuthMethodDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get auth method default response has a 4xx status code
func (o *WaypointGetAuthMethodDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get auth method default response has a 5xx status code
func (o *WaypointGetAuthMethodDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get auth method default response a status code equal to that given
func (o *WaypointGetAuthMethodDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get auth method default response
func (o *WaypointGetAuthMethodDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetAuthMethodDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/auth-method/{auth_method.name}][%d] Waypoint_GetAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetAuthMethodDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/auth-method/{auth_method.name}][%d] Waypoint_GetAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetAuthMethodDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetAuthMethodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
