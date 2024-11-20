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

// UpdateMongoDBAtlasRotatingSecretReader is a Reader for the UpdateMongoDBAtlasRotatingSecret structure.
type UpdateMongoDBAtlasRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMongoDBAtlasRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMongoDBAtlasRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMongoDBAtlasRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMongoDBAtlasRotatingSecretOK creates a UpdateMongoDBAtlasRotatingSecretOK with default headers values
func NewUpdateMongoDBAtlasRotatingSecretOK() *UpdateMongoDBAtlasRotatingSecretOK {
	return &UpdateMongoDBAtlasRotatingSecretOK{}
}

/*
UpdateMongoDBAtlasRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateMongoDBAtlasRotatingSecretOK struct {
	Payload *models.Secrets20231128UpdateMongoDBAtlasRotatingSecretResponse
}

// IsSuccess returns true when this update mongo d b atlas rotating secret o k response has a 2xx status code
func (o *UpdateMongoDBAtlasRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mongo d b atlas rotating secret o k response has a 3xx status code
func (o *UpdateMongoDBAtlasRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mongo d b atlas rotating secret o k response has a 4xx status code
func (o *UpdateMongoDBAtlasRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mongo d b atlas rotating secret o k response has a 5xx status code
func (o *UpdateMongoDBAtlasRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update mongo d b atlas rotating secret o k response a status code equal to that given
func (o *UpdateMongoDBAtlasRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update mongo d b atlas rotating secret o k response
func (o *UpdateMongoDBAtlasRotatingSecretOK) Code() int {
	return 200
}

func (o *UpdateMongoDBAtlasRotatingSecretOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/mongodb-atlas/secret/{name}][%d] updateMongoDBAtlasRotatingSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateMongoDBAtlasRotatingSecretOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/mongodb-atlas/secret/{name}][%d] updateMongoDBAtlasRotatingSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateMongoDBAtlasRotatingSecretOK) GetPayload() *models.Secrets20231128UpdateMongoDBAtlasRotatingSecretResponse {
	return o.Payload
}

func (o *UpdateMongoDBAtlasRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateMongoDBAtlasRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMongoDBAtlasRotatingSecretDefault creates a UpdateMongoDBAtlasRotatingSecretDefault with default headers values
func NewUpdateMongoDBAtlasRotatingSecretDefault(code int) *UpdateMongoDBAtlasRotatingSecretDefault {
	return &UpdateMongoDBAtlasRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateMongoDBAtlasRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateMongoDBAtlasRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update mongo d b atlas rotating secret default response has a 2xx status code
func (o *UpdateMongoDBAtlasRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update mongo d b atlas rotating secret default response has a 3xx status code
func (o *UpdateMongoDBAtlasRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update mongo d b atlas rotating secret default response has a 4xx status code
func (o *UpdateMongoDBAtlasRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update mongo d b atlas rotating secret default response has a 5xx status code
func (o *UpdateMongoDBAtlasRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update mongo d b atlas rotating secret default response a status code equal to that given
func (o *UpdateMongoDBAtlasRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update mongo d b atlas rotating secret default response
func (o *UpdateMongoDBAtlasRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMongoDBAtlasRotatingSecretDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/mongodb-atlas/secret/{name}][%d] UpdateMongoDBAtlasRotatingSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMongoDBAtlasRotatingSecretDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/mongodb-atlas/secret/{name}][%d] UpdateMongoDBAtlasRotatingSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMongoDBAtlasRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateMongoDBAtlasRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}