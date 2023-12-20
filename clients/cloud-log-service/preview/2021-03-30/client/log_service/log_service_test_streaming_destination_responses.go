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

// LogServiceTestStreamingDestinationReader is a Reader for the LogServiceTestStreamingDestination structure.
type LogServiceTestStreamingDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceTestStreamingDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceTestStreamingDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceTestStreamingDestinationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceTestStreamingDestinationOK creates a LogServiceTestStreamingDestinationOK with default headers values
func NewLogServiceTestStreamingDestinationOK() *LogServiceTestStreamingDestinationOK {
	return &LogServiceTestStreamingDestinationOK{}
}

/*
LogServiceTestStreamingDestinationOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceTestStreamingDestinationOK struct {
	Payload *models.LogService20210330TestStreamingDestinationResponse
}

// IsSuccess returns true when this log service test streaming destination o k response has a 2xx status code
func (o *LogServiceTestStreamingDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service test streaming destination o k response has a 3xx status code
func (o *LogServiceTestStreamingDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service test streaming destination o k response has a 4xx status code
func (o *LogServiceTestStreamingDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service test streaming destination o k response has a 5xx status code
func (o *LogServiceTestStreamingDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service test streaming destination o k response a status code equal to that given
func (o *LogServiceTestStreamingDestinationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service test streaming destination o k response
func (o *LogServiceTestStreamingDestinationOK) Code() int {
	return 200
}

func (o *LogServiceTestStreamingDestinationOK) Error() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/test][%d] logServiceTestStreamingDestinationOK  %+v", 200, o.Payload)
}

func (o *LogServiceTestStreamingDestinationOK) String() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/test][%d] logServiceTestStreamingDestinationOK  %+v", 200, o.Payload)
}

func (o *LogServiceTestStreamingDestinationOK) GetPayload() *models.LogService20210330TestStreamingDestinationResponse {
	return o.Payload
}

func (o *LogServiceTestStreamingDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330TestStreamingDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceTestStreamingDestinationDefault creates a LogServiceTestStreamingDestinationDefault with default headers values
func NewLogServiceTestStreamingDestinationDefault(code int) *LogServiceTestStreamingDestinationDefault {
	return &LogServiceTestStreamingDestinationDefault{
		_statusCode: code,
	}
}

/*
LogServiceTestStreamingDestinationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceTestStreamingDestinationDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service test streaming destination default response has a 2xx status code
func (o *LogServiceTestStreamingDestinationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service test streaming destination default response has a 3xx status code
func (o *LogServiceTestStreamingDestinationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service test streaming destination default response has a 4xx status code
func (o *LogServiceTestStreamingDestinationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service test streaming destination default response has a 5xx status code
func (o *LogServiceTestStreamingDestinationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service test streaming destination default response a status code equal to that given
func (o *LogServiceTestStreamingDestinationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service test streaming destination default response
func (o *LogServiceTestStreamingDestinationDefault) Code() int {
	return o._statusCode
}

func (o *LogServiceTestStreamingDestinationDefault) Error() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/test][%d] LogService_TestStreamingDestination default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceTestStreamingDestinationDefault) String() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations/test][%d] LogService_TestStreamingDestination default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceTestStreamingDestinationDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceTestStreamingDestinationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
