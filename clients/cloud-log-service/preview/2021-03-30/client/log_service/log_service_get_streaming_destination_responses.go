// Code generated by go-swagger; DO NOT EDIT.

package log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// LogServiceGetStreamingDestinationReader is a Reader for the LogServiceGetStreamingDestination structure.
type LogServiceGetStreamingDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceGetStreamingDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceGetStreamingDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceGetStreamingDestinationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceGetStreamingDestinationOK creates a LogServiceGetStreamingDestinationOK with default headers values
func NewLogServiceGetStreamingDestinationOK() *LogServiceGetStreamingDestinationOK {
	return &LogServiceGetStreamingDestinationOK{}
}

/*
LogServiceGetStreamingDestinationOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceGetStreamingDestinationOK struct {
	Payload *models.LogService20210330GetStreamingDestinationResponse
}

// IsSuccess returns true when this log service get streaming destination o k response has a 2xx status code
func (o *LogServiceGetStreamingDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service get streaming destination o k response has a 3xx status code
func (o *LogServiceGetStreamingDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service get streaming destination o k response has a 4xx status code
func (o *LogServiceGetStreamingDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service get streaming destination o k response has a 5xx status code
func (o *LogServiceGetStreamingDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service get streaming destination o k response a status code equal to that given
func (o *LogServiceGetStreamingDestinationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service get streaming destination o k response
func (o *LogServiceGetStreamingDestinationOK) Code() int {
	return 200
}

func (o *LogServiceGetStreamingDestinationOK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/{destination_id}][%d] logServiceGetStreamingDestinationOK  %+v", 200, o.Payload)
}

func (o *LogServiceGetStreamingDestinationOK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/{destination_id}][%d] logServiceGetStreamingDestinationOK  %+v", 200, o.Payload)
}

func (o *LogServiceGetStreamingDestinationOK) GetPayload() *models.LogService20210330GetStreamingDestinationResponse {
	return o.Payload
}

func (o *LogServiceGetStreamingDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330GetStreamingDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceGetStreamingDestinationDefault creates a LogServiceGetStreamingDestinationDefault with default headers values
func NewLogServiceGetStreamingDestinationDefault(code int) *LogServiceGetStreamingDestinationDefault {
	return &LogServiceGetStreamingDestinationDefault{
		_statusCode: code,
	}
}

/*
LogServiceGetStreamingDestinationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceGetStreamingDestinationDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service get streaming destination default response has a 2xx status code
func (o *LogServiceGetStreamingDestinationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service get streaming destination default response has a 3xx status code
func (o *LogServiceGetStreamingDestinationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service get streaming destination default response has a 4xx status code
func (o *LogServiceGetStreamingDestinationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service get streaming destination default response has a 5xx status code
func (o *LogServiceGetStreamingDestinationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service get streaming destination default response a status code equal to that given
func (o *LogServiceGetStreamingDestinationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service get streaming destination default response
func (o *LogServiceGetStreamingDestinationDefault) Code() int {
	return o._statusCode
}

func (o *LogServiceGetStreamingDestinationDefault) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/{destination_id}][%d] LogService_GetStreamingDestination default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceGetStreamingDestinationDefault) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/{destination_id}][%d] LogService_GetStreamingDestination default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceGetStreamingDestinationDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceGetStreamingDestinationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}