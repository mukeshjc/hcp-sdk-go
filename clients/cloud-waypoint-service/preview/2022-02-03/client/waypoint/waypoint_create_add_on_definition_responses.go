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

// WaypointCreateAddOnDefinitionReader is a Reader for the WaypointCreateAddOnDefinition structure.
type WaypointCreateAddOnDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointCreateAddOnDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointCreateAddOnDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointCreateAddOnDefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointCreateAddOnDefinitionOK creates a WaypointCreateAddOnDefinitionOK with default headers values
func NewWaypointCreateAddOnDefinitionOK() *WaypointCreateAddOnDefinitionOK {
	return &WaypointCreateAddOnDefinitionOK{}
}

/*
WaypointCreateAddOnDefinitionOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointCreateAddOnDefinitionOK struct {
	Payload *models.HashicorpWaypointCreateAddOnDefinitionResponse
}

// IsSuccess returns true when this waypoint create add on definition o k response has a 2xx status code
func (o *WaypointCreateAddOnDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint create add on definition o k response has a 3xx status code
func (o *WaypointCreateAddOnDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint create add on definition o k response has a 4xx status code
func (o *WaypointCreateAddOnDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint create add on definition o k response has a 5xx status code
func (o *WaypointCreateAddOnDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint create add on definition o k response a status code equal to that given
func (o *WaypointCreateAddOnDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint create add on definition o k response
func (o *WaypointCreateAddOnDefinitionOK) Code() int {
	return 200
}

func (o *WaypointCreateAddOnDefinitionOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/add-on-definition][%d] waypointCreateAddOnDefinitionOK  %+v", 200, o.Payload)
}

func (o *WaypointCreateAddOnDefinitionOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/add-on-definition][%d] waypointCreateAddOnDefinitionOK  %+v", 200, o.Payload)
}

func (o *WaypointCreateAddOnDefinitionOK) GetPayload() *models.HashicorpWaypointCreateAddOnDefinitionResponse {
	return o.Payload
}

func (o *WaypointCreateAddOnDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointCreateAddOnDefinitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointCreateAddOnDefinitionDefault creates a WaypointCreateAddOnDefinitionDefault with default headers values
func NewWaypointCreateAddOnDefinitionDefault(code int) *WaypointCreateAddOnDefinitionDefault {
	return &WaypointCreateAddOnDefinitionDefault{
		_statusCode: code,
	}
}

/*
WaypointCreateAddOnDefinitionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointCreateAddOnDefinitionDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint create add on definition default response has a 2xx status code
func (o *WaypointCreateAddOnDefinitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint create add on definition default response has a 3xx status code
func (o *WaypointCreateAddOnDefinitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint create add on definition default response has a 4xx status code
func (o *WaypointCreateAddOnDefinitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint create add on definition default response has a 5xx status code
func (o *WaypointCreateAddOnDefinitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint create add on definition default response a status code equal to that given
func (o *WaypointCreateAddOnDefinitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint create add on definition default response
func (o *WaypointCreateAddOnDefinitionDefault) Code() int {
	return o._statusCode
}

func (o *WaypointCreateAddOnDefinitionDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/add-on-definition][%d] Waypoint_CreateAddOnDefinition default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointCreateAddOnDefinitionDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/add-on-definition][%d] Waypoint_CreateAddOnDefinition default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointCreateAddOnDefinitionDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointCreateAddOnDefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
