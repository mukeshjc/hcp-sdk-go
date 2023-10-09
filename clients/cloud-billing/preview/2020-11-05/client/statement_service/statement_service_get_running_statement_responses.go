// Code generated by go-swagger; DO NOT EDIT.

package statement_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/models"
)

// StatementServiceGetRunningStatementReader is a Reader for the StatementServiceGetRunningStatement structure.
type StatementServiceGetRunningStatementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatementServiceGetRunningStatementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatementServiceGetRunningStatementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatementServiceGetRunningStatementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatementServiceGetRunningStatementOK creates a StatementServiceGetRunningStatementOK with default headers values
func NewStatementServiceGetRunningStatementOK() *StatementServiceGetRunningStatementOK {
	return &StatementServiceGetRunningStatementOK{}
}

/*
StatementServiceGetRunningStatementOK describes a response with status code 200, with default header values.

A successful response.
*/
type StatementServiceGetRunningStatementOK struct {
	Payload *models.Billing20201105GetRunningStatementResponse
}

// IsSuccess returns true when this statement service get running statement o k response has a 2xx status code
func (o *StatementServiceGetRunningStatementOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this statement service get running statement o k response has a 3xx status code
func (o *StatementServiceGetRunningStatementOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this statement service get running statement o k response has a 4xx status code
func (o *StatementServiceGetRunningStatementOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this statement service get running statement o k response has a 5xx status code
func (o *StatementServiceGetRunningStatementOK) IsServerError() bool {
	return false
}

// IsCode returns true when this statement service get running statement o k response a status code equal to that given
func (o *StatementServiceGetRunningStatementOK) IsCode(code int) bool {
	return code == 200
}

func (o *StatementServiceGetRunningStatementOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/running-statement][%d] statementServiceGetRunningStatementOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetRunningStatementOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/running-statement][%d] statementServiceGetRunningStatementOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetRunningStatementOK) GetPayload() *models.Billing20201105GetRunningStatementResponse {
	return o.Payload
}

func (o *StatementServiceGetRunningStatementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetRunningStatementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementServiceGetRunningStatementDefault creates a StatementServiceGetRunningStatementDefault with default headers values
func NewStatementServiceGetRunningStatementDefault(code int) *StatementServiceGetRunningStatementDefault {
	return &StatementServiceGetRunningStatementDefault{
		_statusCode: code,
	}
}

/*
StatementServiceGetRunningStatementDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StatementServiceGetRunningStatementDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the statement service get running statement default response
func (o *StatementServiceGetRunningStatementDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this statement service get running statement default response has a 2xx status code
func (o *StatementServiceGetRunningStatementDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this statement service get running statement default response has a 3xx status code
func (o *StatementServiceGetRunningStatementDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this statement service get running statement default response has a 4xx status code
func (o *StatementServiceGetRunningStatementDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this statement service get running statement default response has a 5xx status code
func (o *StatementServiceGetRunningStatementDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this statement service get running statement default response a status code equal to that given
func (o *StatementServiceGetRunningStatementDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StatementServiceGetRunningStatementDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/running-statement][%d] StatementService_GetRunningStatement default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetRunningStatementDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/running-statement][%d] StatementService_GetRunningStatement default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetRunningStatementDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *StatementServiceGetRunningStatementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
