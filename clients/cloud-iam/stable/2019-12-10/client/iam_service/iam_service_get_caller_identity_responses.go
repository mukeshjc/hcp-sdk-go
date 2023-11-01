// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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

// IamServiceGetCallerIdentityReader is a Reader for the IamServiceGetCallerIdentity structure.
type IamServiceGetCallerIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamServiceGetCallerIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamServiceGetCallerIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamServiceGetCallerIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamServiceGetCallerIdentityOK creates a IamServiceGetCallerIdentityOK with default headers values
func NewIamServiceGetCallerIdentityOK() *IamServiceGetCallerIdentityOK {
	return &IamServiceGetCallerIdentityOK{}
}

/*
IamServiceGetCallerIdentityOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamServiceGetCallerIdentityOK struct {
	Payload *models.HashicorpCloudIamGetCallerIdentityResponse
}

// IsSuccess returns true when this iam service get caller identity o k response has a 2xx status code
func (o *IamServiceGetCallerIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam service get caller identity o k response has a 3xx status code
func (o *IamServiceGetCallerIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam service get caller identity o k response has a 4xx status code
func (o *IamServiceGetCallerIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam service get caller identity o k response has a 5xx status code
func (o *IamServiceGetCallerIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam service get caller identity o k response a status code equal to that given
func (o *IamServiceGetCallerIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iam service get caller identity o k response
func (o *IamServiceGetCallerIdentityOK) Code() int {
	return 200
}

func (o *IamServiceGetCallerIdentityOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/caller-identity][%d] iamServiceGetCallerIdentityOK  %+v", 200, o.Payload)
}

func (o *IamServiceGetCallerIdentityOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/caller-identity][%d] iamServiceGetCallerIdentityOK  %+v", 200, o.Payload)
}

func (o *IamServiceGetCallerIdentityOK) GetPayload() *models.HashicorpCloudIamGetCallerIdentityResponse {
	return o.Payload
}

func (o *IamServiceGetCallerIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetCallerIdentityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamServiceGetCallerIdentityDefault creates a IamServiceGetCallerIdentityDefault with default headers values
func NewIamServiceGetCallerIdentityDefault(code int) *IamServiceGetCallerIdentityDefault {
	return &IamServiceGetCallerIdentityDefault{
		_statusCode: code,
	}
}

/*
IamServiceGetCallerIdentityDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IamServiceGetCallerIdentityDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this iam service get caller identity default response has a 2xx status code
func (o *IamServiceGetCallerIdentityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam service get caller identity default response has a 3xx status code
func (o *IamServiceGetCallerIdentityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam service get caller identity default response has a 4xx status code
func (o *IamServiceGetCallerIdentityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam service get caller identity default response has a 5xx status code
func (o *IamServiceGetCallerIdentityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam service get caller identity default response a status code equal to that given
func (o *IamServiceGetCallerIdentityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iam service get caller identity default response
func (o *IamServiceGetCallerIdentityDefault) Code() int {
	return o._statusCode
}

func (o *IamServiceGetCallerIdentityDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/caller-identity][%d] IamService_GetCallerIdentity default  %+v", o._statusCode, o.Payload)
}

func (o *IamServiceGetCallerIdentityDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/caller-identity][%d] IamService_GetCallerIdentity default  %+v", o._statusCode, o.Payload)
}

func (o *IamServiceGetCallerIdentityDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *IamServiceGetCallerIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
