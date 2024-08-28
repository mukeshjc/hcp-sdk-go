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

// StreamingServiceGetDestinationReader is a Reader for the StreamingServiceGetDestination structure.
type StreamingServiceGetDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamingServiceGetDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamingServiceGetDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStreamingServiceGetDestinationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStreamingServiceGetDestinationOK creates a StreamingServiceGetDestinationOK with default headers values
func NewStreamingServiceGetDestinationOK() *StreamingServiceGetDestinationOK {
	return &StreamingServiceGetDestinationOK{}
}

/*
StreamingServiceGetDestinationOK describes a response with status code 200, with default header values.

A successful response.
*/
type StreamingServiceGetDestinationOK struct {
	Payload *models.LogService20210330GetDestinationResponse
}

// IsSuccess returns true when this streaming service get destination o k response has a 2xx status code
func (o *StreamingServiceGetDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this streaming service get destination o k response has a 3xx status code
func (o *StreamingServiceGetDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this streaming service get destination o k response has a 4xx status code
func (o *StreamingServiceGetDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this streaming service get destination o k response has a 5xx status code
func (o *StreamingServiceGetDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this streaming service get destination o k response a status code equal to that given
func (o *StreamingServiceGetDestinationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the streaming service get destination o k response
func (o *StreamingServiceGetDestinationOK) Code() int {
	return 200
}

func (o *StreamingServiceGetDestinationOK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{organization_id}/destinations/{destination_id}][%d] streamingServiceGetDestinationOK  %+v", 200, o.Payload)
}

func (o *StreamingServiceGetDestinationOK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{organization_id}/destinations/{destination_id}][%d] streamingServiceGetDestinationOK  %+v", 200, o.Payload)
}

func (o *StreamingServiceGetDestinationOK) GetPayload() *models.LogService20210330GetDestinationResponse {
	return o.Payload
}

func (o *StreamingServiceGetDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330GetDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamingServiceGetDestinationDefault creates a StreamingServiceGetDestinationDefault with default headers values
func NewStreamingServiceGetDestinationDefault(code int) *StreamingServiceGetDestinationDefault {
	return &StreamingServiceGetDestinationDefault{
		_statusCode: code,
	}
}

/*
StreamingServiceGetDestinationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StreamingServiceGetDestinationDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this streaming service get destination default response has a 2xx status code
func (o *StreamingServiceGetDestinationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this streaming service get destination default response has a 3xx status code
func (o *StreamingServiceGetDestinationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this streaming service get destination default response has a 4xx status code
func (o *StreamingServiceGetDestinationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this streaming service get destination default response has a 5xx status code
func (o *StreamingServiceGetDestinationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this streaming service get destination default response a status code equal to that given
func (o *StreamingServiceGetDestinationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the streaming service get destination default response
func (o *StreamingServiceGetDestinationDefault) Code() int {
	return o._statusCode
}

func (o *StreamingServiceGetDestinationDefault) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{organization_id}/destinations/{destination_id}][%d] StreamingService_GetDestination default  %+v", o._statusCode, o.Payload)
}

func (o *StreamingServiceGetDestinationDefault) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{organization_id}/destinations/{destination_id}][%d] StreamingService_GetDestination default  %+v", o._statusCode, o.Payload)
}

func (o *StreamingServiceGetDestinationDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *StreamingServiceGetDestinationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}