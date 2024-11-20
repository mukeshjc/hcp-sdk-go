// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// ListGatewayPoolsReader is a Reader for the ListGatewayPools structure.
type ListGatewayPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGatewayPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGatewayPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGatewayPoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGatewayPoolsOK creates a ListGatewayPoolsOK with default headers values
func NewListGatewayPoolsOK() *ListGatewayPoolsOK {
	return &ListGatewayPoolsOK{}
}

/*
ListGatewayPoolsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGatewayPoolsOK struct {
	Payload *models.Secrets20231128ListGatewayPoolsResponse
}

// IsSuccess returns true when this list gateway pools o k response has a 2xx status code
func (o *ListGatewayPoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list gateway pools o k response has a 3xx status code
func (o *ListGatewayPoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gateway pools o k response has a 4xx status code
func (o *ListGatewayPoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gateway pools o k response has a 5xx status code
func (o *ListGatewayPoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list gateway pools o k response a status code equal to that given
func (o *ListGatewayPoolsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list gateway pools o k response
func (o *ListGatewayPoolsOK) Code() int {
	return 200
}

func (o *ListGatewayPoolsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools][%d] listGatewayPoolsOK  %+v", 200, o.Payload)
}

func (o *ListGatewayPoolsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools][%d] listGatewayPoolsOK  %+v", 200, o.Payload)
}

func (o *ListGatewayPoolsOK) GetPayload() *models.Secrets20231128ListGatewayPoolsResponse {
	return o.Payload
}

func (o *ListGatewayPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListGatewayPoolsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewayPoolsDefault creates a ListGatewayPoolsDefault with default headers values
func NewListGatewayPoolsDefault(code int) *ListGatewayPoolsDefault {
	return &ListGatewayPoolsDefault{
		_statusCode: code,
	}
}

/*
ListGatewayPoolsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListGatewayPoolsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list gateway pools default response has a 2xx status code
func (o *ListGatewayPoolsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list gateway pools default response has a 3xx status code
func (o *ListGatewayPoolsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list gateway pools default response has a 4xx status code
func (o *ListGatewayPoolsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list gateway pools default response has a 5xx status code
func (o *ListGatewayPoolsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list gateway pools default response a status code equal to that given
func (o *ListGatewayPoolsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list gateway pools default response
func (o *ListGatewayPoolsDefault) Code() int {
	return o._statusCode
}

func (o *ListGatewayPoolsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools][%d] ListGatewayPools default  %+v", o._statusCode, o.Payload)
}

func (o *ListGatewayPoolsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools][%d] ListGatewayPools default  %+v", o._statusCode, o.Payload)
}

func (o *ListGatewayPoolsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListGatewayPoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}