// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListVersions2Reader is a Reader for the ListVersions2 structure.
type ListVersions2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersions2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVersions2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVersions2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVersions2OK creates a ListVersions2OK with default headers values
func NewListVersions2OK() *ListVersions2OK {
	return &ListVersions2OK{}
}

/*
ListVersions2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ListVersions2OK struct {
	Payload *models.HashicorpCloudConsul20210204ListVersionsResponse
}

// IsSuccess returns true when this list versions2 o k response has a 2xx status code
func (o *ListVersions2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list versions2 o k response has a 3xx status code
func (o *ListVersions2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list versions2 o k response has a 4xx status code
func (o *ListVersions2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list versions2 o k response has a 5xx status code
func (o *ListVersions2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list versions2 o k response a status code equal to that given
func (o *ListVersions2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVersions2OK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/versions][%d] listVersions2OK  %+v", 200, o.Payload)
}

func (o *ListVersions2OK) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/versions][%d] listVersions2OK  %+v", 200, o.Payload)
}

func (o *ListVersions2OK) GetPayload() *models.HashicorpCloudConsul20210204ListVersionsResponse {
	return o.Payload
}

func (o *ListVersions2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204ListVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersions2Default creates a ListVersions2Default with default headers values
func NewListVersions2Default(code int) *ListVersions2Default {
	return &ListVersions2Default{
		_statusCode: code,
	}
}

/*
ListVersions2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListVersions2Default struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list versions2 default response
func (o *ListVersions2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list versions2 default response has a 2xx status code
func (o *ListVersions2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list versions2 default response has a 3xx status code
func (o *ListVersions2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list versions2 default response has a 4xx status code
func (o *ListVersions2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list versions2 default response has a 5xx status code
func (o *ListVersions2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list versions2 default response a status code equal to that given
func (o *ListVersions2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVersions2Default) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/versions][%d] ListVersions2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersions2Default) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/versions][%d] ListVersions2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersions2Default) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListVersions2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}