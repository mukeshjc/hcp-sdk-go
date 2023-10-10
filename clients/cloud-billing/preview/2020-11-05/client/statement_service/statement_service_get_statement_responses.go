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

// StatementServiceGetStatementReader is a Reader for the StatementServiceGetStatement structure.
type StatementServiceGetStatementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatementServiceGetStatementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatementServiceGetStatementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatementServiceGetStatementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatementServiceGetStatementOK creates a StatementServiceGetStatementOK with default headers values
func NewStatementServiceGetStatementOK() *StatementServiceGetStatementOK {
	return &StatementServiceGetStatementOK{}
}

/*
StatementServiceGetStatementOK describes a response with status code 200, with default header values.

A successful response.
*/
type StatementServiceGetStatementOK struct {
	Payload *models.Billing20201105GetStatementResponse
}

// IsSuccess returns true when this statement service get statement o k response has a 2xx status code
func (o *StatementServiceGetStatementOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this statement service get statement o k response has a 3xx status code
func (o *StatementServiceGetStatementOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this statement service get statement o k response has a 4xx status code
func (o *StatementServiceGetStatementOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this statement service get statement o k response has a 5xx status code
func (o *StatementServiceGetStatementOK) IsServerError() bool {
	return false
}

// IsCode returns true when this statement service get statement o k response a status code equal to that given
func (o *StatementServiceGetStatementOK) IsCode(code int) bool {
	return code == 200
}

func (o *StatementServiceGetStatementOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}][%d] statementServiceGetStatementOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetStatementOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}][%d] statementServiceGetStatementOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetStatementOK) GetPayload() *models.Billing20201105GetStatementResponse {
	return o.Payload
}

func (o *StatementServiceGetStatementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetStatementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementServiceGetStatementDefault creates a StatementServiceGetStatementDefault with default headers values
func NewStatementServiceGetStatementDefault(code int) *StatementServiceGetStatementDefault {
	return &StatementServiceGetStatementDefault{
		_statusCode: code,
	}
}

/*
StatementServiceGetStatementDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StatementServiceGetStatementDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the statement service get statement default response
func (o *StatementServiceGetStatementDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this statement service get statement default response has a 2xx status code
func (o *StatementServiceGetStatementDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this statement service get statement default response has a 3xx status code
func (o *StatementServiceGetStatementDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this statement service get statement default response has a 4xx status code
func (o *StatementServiceGetStatementDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this statement service get statement default response has a 5xx status code
func (o *StatementServiceGetStatementDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this statement service get statement default response a status code equal to that given
func (o *StatementServiceGetStatementDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StatementServiceGetStatementDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}][%d] StatementService_GetStatement default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetStatementDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}][%d] StatementService_GetStatement default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetStatementDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *StatementServiceGetStatementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}