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

// ActionUpdateActionConfigReader is a Reader for the ActionUpdateActionConfig structure.
type ActionUpdateActionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionUpdateActionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionUpdateActionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionUpdateActionConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionUpdateActionConfigOK creates a ActionUpdateActionConfigOK with default headers values
func NewActionUpdateActionConfigOK() *ActionUpdateActionConfigOK {
	return &ActionUpdateActionConfigOK{}
}

/*
ActionUpdateActionConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActionUpdateActionConfigOK struct {
	Payload *models.HashicorpCloudWaypointActionsdriverUpdateActionConfigResponse
}

// IsSuccess returns true when this action update action config o k response has a 2xx status code
func (o *ActionUpdateActionConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action update action config o k response has a 3xx status code
func (o *ActionUpdateActionConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action update action config o k response has a 4xx status code
func (o *ActionUpdateActionConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action update action config o k response has a 5xx status code
func (o *ActionUpdateActionConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action update action config o k response a status code equal to that given
func (o *ActionUpdateActionConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the action update action config o k response
func (o *ActionUpdateActionConfigOK) Code() int {
	return 200
}

func (o *ActionUpdateActionConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] actionUpdateActionConfigOK  %+v", 200, o.Payload)
}

func (o *ActionUpdateActionConfigOK) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] actionUpdateActionConfigOK  %+v", 200, o.Payload)
}

func (o *ActionUpdateActionConfigOK) GetPayload() *models.HashicorpCloudWaypointActionsdriverUpdateActionConfigResponse {
	return o.Payload
}

func (o *ActionUpdateActionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointActionsdriverUpdateActionConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionUpdateActionConfigDefault creates a ActionUpdateActionConfigDefault with default headers values
func NewActionUpdateActionConfigDefault(code int) *ActionUpdateActionConfigDefault {
	return &ActionUpdateActionConfigDefault{
		_statusCode: code,
	}
}

/*
ActionUpdateActionConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActionUpdateActionConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this action update action config default response has a 2xx status code
func (o *ActionUpdateActionConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this action update action config default response has a 3xx status code
func (o *ActionUpdateActionConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this action update action config default response has a 4xx status code
func (o *ActionUpdateActionConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this action update action config default response has a 5xx status code
func (o *ActionUpdateActionConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this action update action config default response a status code equal to that given
func (o *ActionUpdateActionConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the action update action config default response
func (o *ActionUpdateActionConfigDefault) Code() int {
	return o._statusCode
}

func (o *ActionUpdateActionConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] Action_UpdateActionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ActionUpdateActionConfigDefault) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] Action_UpdateActionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ActionUpdateActionConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ActionUpdateActionConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
