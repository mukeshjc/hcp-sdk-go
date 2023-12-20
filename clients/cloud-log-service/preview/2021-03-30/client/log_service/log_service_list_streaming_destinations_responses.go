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

// LogServiceListStreamingDestinationsReader is a Reader for the LogServiceListStreamingDestinations structure.
type LogServiceListStreamingDestinationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceListStreamingDestinationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceListStreamingDestinationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceListStreamingDestinationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceListStreamingDestinationsOK creates a LogServiceListStreamingDestinationsOK with default headers values
func NewLogServiceListStreamingDestinationsOK() *LogServiceListStreamingDestinationsOK {
	return &LogServiceListStreamingDestinationsOK{}
}

/*
LogServiceListStreamingDestinationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceListStreamingDestinationsOK struct {
	Payload *models.LogService20210330ListStreamingDestinationsResponse
}

// IsSuccess returns true when this log service list streaming destinations o k response has a 2xx status code
func (o *LogServiceListStreamingDestinationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service list streaming destinations o k response has a 3xx status code
func (o *LogServiceListStreamingDestinationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service list streaming destinations o k response has a 4xx status code
func (o *LogServiceListStreamingDestinationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service list streaming destinations o k response has a 5xx status code
func (o *LogServiceListStreamingDestinationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service list streaming destinations o k response a status code equal to that given
func (o *LogServiceListStreamingDestinationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service list streaming destinations o k response
func (o *LogServiceListStreamingDestinationsOK) Code() int {
	return 200
}

func (o *LogServiceListStreamingDestinationsOK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations][%d] logServiceListStreamingDestinationsOK  %+v", 200, o.Payload)
}

func (o *LogServiceListStreamingDestinationsOK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations][%d] logServiceListStreamingDestinationsOK  %+v", 200, o.Payload)
}

func (o *LogServiceListStreamingDestinationsOK) GetPayload() *models.LogService20210330ListStreamingDestinationsResponse {
	return o.Payload
}

func (o *LogServiceListStreamingDestinationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330ListStreamingDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceListStreamingDestinationsDefault creates a LogServiceListStreamingDestinationsDefault with default headers values
func NewLogServiceListStreamingDestinationsDefault(code int) *LogServiceListStreamingDestinationsDefault {
	return &LogServiceListStreamingDestinationsDefault{
		_statusCode: code,
	}
}

/*
LogServiceListStreamingDestinationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceListStreamingDestinationsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service list streaming destinations default response has a 2xx status code
func (o *LogServiceListStreamingDestinationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service list streaming destinations default response has a 3xx status code
func (o *LogServiceListStreamingDestinationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service list streaming destinations default response has a 4xx status code
func (o *LogServiceListStreamingDestinationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service list streaming destinations default response has a 5xx status code
func (o *LogServiceListStreamingDestinationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service list streaming destinations default response a status code equal to that given
func (o *LogServiceListStreamingDestinationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service list streaming destinations default response
func (o *LogServiceListStreamingDestinationsDefault) Code() int {
	return o._statusCode
}

func (o *LogServiceListStreamingDestinationsDefault) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations][%d] LogService_ListStreamingDestinations default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListStreamingDestinationsDefault) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/projects/{location.project_id}/resources/destinations][%d] LogService_ListStreamingDestinations default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListStreamingDestinationsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceListStreamingDestinationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
