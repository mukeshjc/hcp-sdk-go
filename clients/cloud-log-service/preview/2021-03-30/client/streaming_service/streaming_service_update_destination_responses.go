// Code generated by go-swagger; DO NOT EDIT.

package streaming_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// StreamingServiceUpdateDestinationReader is a Reader for the StreamingServiceUpdateDestination structure.
type StreamingServiceUpdateDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamingServiceUpdateDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamingServiceUpdateDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStreamingServiceUpdateDestinationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStreamingServiceUpdateDestinationOK creates a StreamingServiceUpdateDestinationOK with default headers values
func NewStreamingServiceUpdateDestinationOK() *StreamingServiceUpdateDestinationOK {
	return &StreamingServiceUpdateDestinationOK{}
}

/*
StreamingServiceUpdateDestinationOK describes a response with status code 200, with default header values.

A successful response.
*/
type StreamingServiceUpdateDestinationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this streaming service update destination o k response has a 2xx status code
func (o *StreamingServiceUpdateDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this streaming service update destination o k response has a 3xx status code
func (o *StreamingServiceUpdateDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this streaming service update destination o k response has a 4xx status code
func (o *StreamingServiceUpdateDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this streaming service update destination o k response has a 5xx status code
func (o *StreamingServiceUpdateDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this streaming service update destination o k response a status code equal to that given
func (o *StreamingServiceUpdateDestinationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the streaming service update destination o k response
func (o *StreamingServiceUpdateDestinationOK) Code() int {
	return 200
}

func (o *StreamingServiceUpdateDestinationOK) Error() string {
	return fmt.Sprintf("[PATCH /logs/2021-03-30/organizations/{organization_id}/destinations/{id}][%d] streamingServiceUpdateDestinationOK  %+v", 200, o.Payload)
}

func (o *StreamingServiceUpdateDestinationOK) String() string {
	return fmt.Sprintf("[PATCH /logs/2021-03-30/organizations/{organization_id}/destinations/{id}][%d] streamingServiceUpdateDestinationOK  %+v", 200, o.Payload)
}

func (o *StreamingServiceUpdateDestinationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StreamingServiceUpdateDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamingServiceUpdateDestinationDefault creates a StreamingServiceUpdateDestinationDefault with default headers values
func NewStreamingServiceUpdateDestinationDefault(code int) *StreamingServiceUpdateDestinationDefault {
	return &StreamingServiceUpdateDestinationDefault{
		_statusCode: code,
	}
}

/*
StreamingServiceUpdateDestinationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StreamingServiceUpdateDestinationDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this streaming service update destination default response has a 2xx status code
func (o *StreamingServiceUpdateDestinationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this streaming service update destination default response has a 3xx status code
func (o *StreamingServiceUpdateDestinationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this streaming service update destination default response has a 4xx status code
func (o *StreamingServiceUpdateDestinationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this streaming service update destination default response has a 5xx status code
func (o *StreamingServiceUpdateDestinationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this streaming service update destination default response a status code equal to that given
func (o *StreamingServiceUpdateDestinationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the streaming service update destination default response
func (o *StreamingServiceUpdateDestinationDefault) Code() int {
	return o._statusCode
}

func (o *StreamingServiceUpdateDestinationDefault) Error() string {
	return fmt.Sprintf("[PATCH /logs/2021-03-30/organizations/{organization_id}/destinations/{id}][%d] StreamingService_UpdateDestination default  %+v", o._statusCode, o.Payload)
}

func (o *StreamingServiceUpdateDestinationDefault) String() string {
	return fmt.Sprintf("[PATCH /logs/2021-03-30/organizations/{organization_id}/destinations/{id}][%d] StreamingService_UpdateDestination default  %+v", o._statusCode, o.Payload)
}

func (o *StreamingServiceUpdateDestinationDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *StreamingServiceUpdateDestinationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}