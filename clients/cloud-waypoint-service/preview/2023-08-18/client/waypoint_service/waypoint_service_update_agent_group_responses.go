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

// WaypointServiceUpdateAgentGroupReader is a Reader for the WaypointServiceUpdateAgentGroup structure.
type WaypointServiceUpdateAgentGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateAgentGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateAgentGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateAgentGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateAgentGroupOK creates a WaypointServiceUpdateAgentGroupOK with default headers values
func NewWaypointServiceUpdateAgentGroupOK() *WaypointServiceUpdateAgentGroupOK {
	return &WaypointServiceUpdateAgentGroupOK{}
}

/*
WaypointServiceUpdateAgentGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateAgentGroupOK struct {
	Payload *models.HashicorpCloudWaypointUpdateAgentGroupResponse
}

// IsSuccess returns true when this waypoint service update agent group o k response has a 2xx status code
func (o *WaypointServiceUpdateAgentGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update agent group o k response has a 3xx status code
func (o *WaypointServiceUpdateAgentGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update agent group o k response has a 4xx status code
func (o *WaypointServiceUpdateAgentGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update agent group o k response has a 5xx status code
func (o *WaypointServiceUpdateAgentGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update agent group o k response a status code equal to that given
func (o *WaypointServiceUpdateAgentGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update agent group o k response
func (o *WaypointServiceUpdateAgentGroupOK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateAgentGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] waypointServiceUpdateAgentGroupOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUpdateAgentGroupOK) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] waypointServiceUpdateAgentGroupOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUpdateAgentGroupOK) GetPayload() *models.HashicorpCloudWaypointUpdateAgentGroupResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateAgentGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUpdateAgentGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateAgentGroupDefault creates a WaypointServiceUpdateAgentGroupDefault with default headers values
func NewWaypointServiceUpdateAgentGroupDefault(code int) *WaypointServiceUpdateAgentGroupDefault {
	return &WaypointServiceUpdateAgentGroupDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateAgentGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateAgentGroupDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update agent group default response has a 2xx status code
func (o *WaypointServiceUpdateAgentGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update agent group default response has a 3xx status code
func (o *WaypointServiceUpdateAgentGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update agent group default response has a 4xx status code
func (o *WaypointServiceUpdateAgentGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update agent group default response has a 5xx status code
func (o *WaypointServiceUpdateAgentGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update agent group default response a status code equal to that given
func (o *WaypointServiceUpdateAgentGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update agent group default response
func (o *WaypointServiceUpdateAgentGroupDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateAgentGroupDefault) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] WaypointService_UpdateAgentGroup default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUpdateAgentGroupDefault) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] WaypointService_UpdateAgentGroup default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUpdateAgentGroupDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateAgentGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}