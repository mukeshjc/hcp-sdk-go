// Code generated by go-swagger; DO NOT EDIT.

package action

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

// ActionListActionRunsReader is a Reader for the ActionListActionRuns structure.
type ActionListActionRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionListActionRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionListActionRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionListActionRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionListActionRunsOK creates a ActionListActionRunsOK with default headers values
func NewActionListActionRunsOK() *ActionListActionRunsOK {
	return &ActionListActionRunsOK{}
}

/*
ActionListActionRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActionListActionRunsOK struct {
	Payload *models.HashicorpCloudWaypointActionsdriverListActionRunsResponse
}

// IsSuccess returns true when this action list action runs o k response has a 2xx status code
func (o *ActionListActionRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action list action runs o k response has a 3xx status code
func (o *ActionListActionRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action list action runs o k response has a 4xx status code
func (o *ActionListActionRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action list action runs o k response has a 5xx status code
func (o *ActionListActionRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action list action runs o k response a status code equal to that given
func (o *ActionListActionRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the action list action runs o k response
func (o *ActionListActionRunsOK) Code() int {
	return 200
}

func (o *ActionListActionRunsOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/actionruns][%d] actionListActionRunsOK  %+v", 200, o.Payload)
}

func (o *ActionListActionRunsOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/actionruns][%d] actionListActionRunsOK  %+v", 200, o.Payload)
}

func (o *ActionListActionRunsOK) GetPayload() *models.HashicorpCloudWaypointActionsdriverListActionRunsResponse {
	return o.Payload
}

func (o *ActionListActionRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointActionsdriverListActionRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionListActionRunsDefault creates a ActionListActionRunsDefault with default headers values
func NewActionListActionRunsDefault(code int) *ActionListActionRunsDefault {
	return &ActionListActionRunsDefault{
		_statusCode: code,
	}
}

/*
ActionListActionRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActionListActionRunsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this action list action runs default response has a 2xx status code
func (o *ActionListActionRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this action list action runs default response has a 3xx status code
func (o *ActionListActionRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this action list action runs default response has a 4xx status code
func (o *ActionListActionRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this action list action runs default response has a 5xx status code
func (o *ActionListActionRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this action list action runs default response a status code equal to that given
func (o *ActionListActionRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the action list action runs default response
func (o *ActionListActionRunsDefault) Code() int {
	return o._statusCode
}

func (o *ActionListActionRunsDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/actionruns][%d] Action_ListActionRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ActionListActionRunsDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/actionruns][%d] Action_ListActionRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ActionListActionRunsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ActionListActionRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}