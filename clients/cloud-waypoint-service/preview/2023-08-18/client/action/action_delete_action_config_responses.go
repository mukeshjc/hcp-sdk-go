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
)

// ActionDeleteActionConfigReader is a Reader for the ActionDeleteActionConfig structure.
type ActionDeleteActionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionDeleteActionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionDeleteActionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionDeleteActionConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionDeleteActionConfigOK creates a ActionDeleteActionConfigOK with default headers values
func NewActionDeleteActionConfigOK() *ActionDeleteActionConfigOK {
	return &ActionDeleteActionConfigOK{}
}

/*
ActionDeleteActionConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActionDeleteActionConfigOK struct {
	Payload interface{}
}

// IsSuccess returns true when this action delete action config o k response has a 2xx status code
func (o *ActionDeleteActionConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action delete action config o k response has a 3xx status code
func (o *ActionDeleteActionConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action delete action config o k response has a 4xx status code
func (o *ActionDeleteActionConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action delete action config o k response has a 5xx status code
func (o *ActionDeleteActionConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action delete action config o k response a status code equal to that given
func (o *ActionDeleteActionConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the action delete action config o k response
func (o *ActionDeleteActionConfigOK) Code() int {
	return 200
}

func (o *ActionDeleteActionConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] actionDeleteActionConfigOK  %+v", 200, o.Payload)
}

func (o *ActionDeleteActionConfigOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] actionDeleteActionConfigOK  %+v", 200, o.Payload)
}

func (o *ActionDeleteActionConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ActionDeleteActionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionDeleteActionConfigDefault creates a ActionDeleteActionConfigDefault with default headers values
func NewActionDeleteActionConfigDefault(code int) *ActionDeleteActionConfigDefault {
	return &ActionDeleteActionConfigDefault{
		_statusCode: code,
	}
}

/*
ActionDeleteActionConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActionDeleteActionConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this action delete action config default response has a 2xx status code
func (o *ActionDeleteActionConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this action delete action config default response has a 3xx status code
func (o *ActionDeleteActionConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this action delete action config default response has a 4xx status code
func (o *ActionDeleteActionConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this action delete action config default response has a 5xx status code
func (o *ActionDeleteActionConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this action delete action config default response a status code equal to that given
func (o *ActionDeleteActionConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the action delete action config default response
func (o *ActionDeleteActionConfigDefault) Code() int {
	return o._statusCode
}

func (o *ActionDeleteActionConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] Action_DeleteActionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ActionDeleteActionConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/actionconfig][%d] Action_DeleteActionConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ActionDeleteActionConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ActionDeleteActionConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}