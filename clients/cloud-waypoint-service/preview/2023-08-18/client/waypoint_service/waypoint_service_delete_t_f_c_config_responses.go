// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// WaypointServiceDeleteTFCConfigReader is a Reader for the WaypointServiceDeleteTFCConfig structure.
type WaypointServiceDeleteTFCConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceDeleteTFCConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceDeleteTFCConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceDeleteTFCConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceDeleteTFCConfigOK creates a WaypointServiceDeleteTFCConfigOK with default headers values
func NewWaypointServiceDeleteTFCConfigOK() *WaypointServiceDeleteTFCConfigOK {
	return &WaypointServiceDeleteTFCConfigOK{}
}

/*
WaypointServiceDeleteTFCConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceDeleteTFCConfigOK struct {
	Payload models.HashicorpCloudWaypointDeleteTFCConfigResponse
}

// IsSuccess returns true when this waypoint service delete t f c config o k response has a 2xx status code
func (o *WaypointServiceDeleteTFCConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service delete t f c config o k response has a 3xx status code
func (o *WaypointServiceDeleteTFCConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service delete t f c config o k response has a 4xx status code
func (o *WaypointServiceDeleteTFCConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service delete t f c config o k response has a 5xx status code
func (o *WaypointServiceDeleteTFCConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service delete t f c config o k response a status code equal to that given
func (o *WaypointServiceDeleteTFCConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service delete t f c config o k response
func (o *WaypointServiceDeleteTFCConfigOK) Code() int {
	return 200
}

func (o *WaypointServiceDeleteTFCConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/tfcconfig][%d] waypointServiceDeleteTFCConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceDeleteTFCConfigOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/tfcconfig][%d] waypointServiceDeleteTFCConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceDeleteTFCConfigOK) GetPayload() models.HashicorpCloudWaypointDeleteTFCConfigResponse {
	return o.Payload
}

func (o *WaypointServiceDeleteTFCConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceDeleteTFCConfigDefault creates a WaypointServiceDeleteTFCConfigDefault with default headers values
func NewWaypointServiceDeleteTFCConfigDefault(code int) *WaypointServiceDeleteTFCConfigDefault {
	return &WaypointServiceDeleteTFCConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceDeleteTFCConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceDeleteTFCConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint service delete t f c config default response has a 2xx status code
func (o *WaypointServiceDeleteTFCConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service delete t f c config default response has a 3xx status code
func (o *WaypointServiceDeleteTFCConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service delete t f c config default response has a 4xx status code
func (o *WaypointServiceDeleteTFCConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service delete t f c config default response has a 5xx status code
func (o *WaypointServiceDeleteTFCConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service delete t f c config default response a status code equal to that given
func (o *WaypointServiceDeleteTFCConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service delete t f c config default response
func (o *WaypointServiceDeleteTFCConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceDeleteTFCConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/tfcconfig][%d] WaypointService_DeleteTFCConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceDeleteTFCConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/tfcconfig][%d] WaypointService_DeleteTFCConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceDeleteTFCConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointServiceDeleteTFCConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
