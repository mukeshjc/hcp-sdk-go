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

// StatementServiceListStatementsReader is a Reader for the StatementServiceListStatements structure.
type StatementServiceListStatementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatementServiceListStatementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatementServiceListStatementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatementServiceListStatementsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatementServiceListStatementsOK creates a StatementServiceListStatementsOK with default headers values
func NewStatementServiceListStatementsOK() *StatementServiceListStatementsOK {
	return &StatementServiceListStatementsOK{}
}

/*
StatementServiceListStatementsOK describes a response with status code 200, with default header values.

A successful response.
*/
type StatementServiceListStatementsOK struct {
	Payload *models.Billing20201105ListStatementsResponse
}

// IsSuccess returns true when this statement service list statements o k response has a 2xx status code
func (o *StatementServiceListStatementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this statement service list statements o k response has a 3xx status code
func (o *StatementServiceListStatementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this statement service list statements o k response has a 4xx status code
func (o *StatementServiceListStatementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this statement service list statements o k response has a 5xx status code
func (o *StatementServiceListStatementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this statement service list statements o k response a status code equal to that given
func (o *StatementServiceListStatementsOK) IsCode(code int) bool {
	return code == 200
}

func (o *StatementServiceListStatementsOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements][%d] statementServiceListStatementsOK  %+v", 200, o.Payload)
}

func (o *StatementServiceListStatementsOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements][%d] statementServiceListStatementsOK  %+v", 200, o.Payload)
}

func (o *StatementServiceListStatementsOK) GetPayload() *models.Billing20201105ListStatementsResponse {
	return o.Payload
}

func (o *StatementServiceListStatementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105ListStatementsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementServiceListStatementsDefault creates a StatementServiceListStatementsDefault with default headers values
func NewStatementServiceListStatementsDefault(code int) *StatementServiceListStatementsDefault {
	return &StatementServiceListStatementsDefault{
		_statusCode: code,
	}
}

/*
StatementServiceListStatementsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StatementServiceListStatementsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the statement service list statements default response
func (o *StatementServiceListStatementsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this statement service list statements default response has a 2xx status code
func (o *StatementServiceListStatementsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this statement service list statements default response has a 3xx status code
func (o *StatementServiceListStatementsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this statement service list statements default response has a 4xx status code
func (o *StatementServiceListStatementsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this statement service list statements default response has a 5xx status code
func (o *StatementServiceListStatementsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this statement service list statements default response a status code equal to that given
func (o *StatementServiceListStatementsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StatementServiceListStatementsDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements][%d] StatementService_ListStatements default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceListStatementsDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements][%d] StatementService_ListStatements default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceListStatementsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *StatementServiceListStatementsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
