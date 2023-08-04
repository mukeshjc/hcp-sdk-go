// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetWorkloadIdentityProviderReader is a Reader for the GetWorkloadIdentityProvider structure.
type GetWorkloadIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkloadIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkloadIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkloadIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkloadIdentityProviderOK creates a GetWorkloadIdentityProviderOK with default headers values
func NewGetWorkloadIdentityProviderOK() *GetWorkloadIdentityProviderOK {
	return &GetWorkloadIdentityProviderOK{}
}

/*
GetWorkloadIdentityProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetWorkloadIdentityProviderOK struct {
	Payload *models.HashicorpCloudIamGetWorkloadIdentityProviderResponse
}

// IsSuccess returns true when this get workload identity provider o k response has a 2xx status code
func (o *GetWorkloadIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workload identity provider o k response has a 3xx status code
func (o *GetWorkloadIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workload identity provider o k response has a 4xx status code
func (o *GetWorkloadIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workload identity provider o k response has a 5xx status code
func (o *GetWorkloadIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workload identity provider o k response a status code equal to that given
func (o *GetWorkloadIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkloadIdentityProviderOK) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name}][%d] getWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *GetWorkloadIdentityProviderOK) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name}][%d] getWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *GetWorkloadIdentityProviderOK) GetPayload() *models.HashicorpCloudIamGetWorkloadIdentityProviderResponse {
	return o.Payload
}

func (o *GetWorkloadIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetWorkloadIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadIdentityProviderDefault creates a GetWorkloadIdentityProviderDefault with default headers values
func NewGetWorkloadIdentityProviderDefault(code int) *GetWorkloadIdentityProviderDefault {
	return &GetWorkloadIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
GetWorkloadIdentityProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetWorkloadIdentityProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the get workload identity provider default response
func (o *GetWorkloadIdentityProviderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get workload identity provider default response has a 2xx status code
func (o *GetWorkloadIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get workload identity provider default response has a 3xx status code
func (o *GetWorkloadIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get workload identity provider default response has a 4xx status code
func (o *GetWorkloadIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get workload identity provider default response has a 5xx status code
func (o *GetWorkloadIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get workload identity provider default response a status code equal to that given
func (o *GetWorkloadIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetWorkloadIdentityProviderDefault) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name}][%d] GetWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadIdentityProviderDefault) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name}][%d] GetWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadIdentityProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GetWorkloadIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}