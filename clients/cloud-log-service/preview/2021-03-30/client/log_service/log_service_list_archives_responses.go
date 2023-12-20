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

// LogServiceListArchivesReader is a Reader for the LogServiceListArchives structure.
type LogServiceListArchivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceListArchivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceListArchivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceListArchivesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceListArchivesOK creates a LogServiceListArchivesOK with default headers values
func NewLogServiceListArchivesOK() *LogServiceListArchivesOK {
	return &LogServiceListArchivesOK{}
}

/*
LogServiceListArchivesOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceListArchivesOK struct {
	Payload *models.LogService20210330ListArchivesResponse
}

// IsSuccess returns true when this log service list archives o k response has a 2xx status code
func (o *LogServiceListArchivesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service list archives o k response has a 3xx status code
func (o *LogServiceListArchivesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service list archives o k response has a 4xx status code
func (o *LogServiceListArchivesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service list archives o k response has a 5xx status code
func (o *LogServiceListArchivesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service list archives o k response a status code equal to that given
func (o *LogServiceListArchivesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service list archives o k response
func (o *LogServiceListArchivesOK) Code() int {
	return 200
}

func (o *LogServiceListArchivesOK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/resources/{resource.id}/archives][%d] logServiceListArchivesOK  %+v", 200, o.Payload)
}

func (o *LogServiceListArchivesOK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/resources/{resource.id}/archives][%d] logServiceListArchivesOK  %+v", 200, o.Payload)
}

func (o *LogServiceListArchivesOK) GetPayload() *models.LogService20210330ListArchivesResponse {
	return o.Payload
}

func (o *LogServiceListArchivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330ListArchivesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceListArchivesDefault creates a LogServiceListArchivesDefault with default headers values
func NewLogServiceListArchivesDefault(code int) *LogServiceListArchivesDefault {
	return &LogServiceListArchivesDefault{
		_statusCode: code,
	}
}

/*
LogServiceListArchivesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceListArchivesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service list archives default response has a 2xx status code
func (o *LogServiceListArchivesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service list archives default response has a 3xx status code
func (o *LogServiceListArchivesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service list archives default response has a 4xx status code
func (o *LogServiceListArchivesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service list archives default response has a 5xx status code
func (o *LogServiceListArchivesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service list archives default response a status code equal to that given
func (o *LogServiceListArchivesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service list archives default response
func (o *LogServiceListArchivesDefault) Code() int {
	return o._statusCode
}

func (o *LogServiceListArchivesDefault) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/resources/{resource.id}/archives][%d] LogService_ListArchives default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListArchivesDefault) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/resources/{resource.id}/archives][%d] LogService_ListArchives default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListArchivesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceListArchivesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
