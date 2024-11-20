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

// UpdateAwsDynamicSecretReader is a Reader for the UpdateAwsDynamicSecret structure.
type UpdateAwsDynamicSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsDynamicSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAwsDynamicSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAwsDynamicSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAwsDynamicSecretOK creates a UpdateAwsDynamicSecretOK with default headers values
func NewUpdateAwsDynamicSecretOK() *UpdateAwsDynamicSecretOK {
	return &UpdateAwsDynamicSecretOK{}
}

/*
UpdateAwsDynamicSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAwsDynamicSecretOK struct {
	Payload *models.Secrets20231128UpdateAwsDynamicSecretResponse
}

// IsSuccess returns true when this update aws dynamic secret o k response has a 2xx status code
func (o *UpdateAwsDynamicSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update aws dynamic secret o k response has a 3xx status code
func (o *UpdateAwsDynamicSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update aws dynamic secret o k response has a 4xx status code
func (o *UpdateAwsDynamicSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update aws dynamic secret o k response has a 5xx status code
func (o *UpdateAwsDynamicSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update aws dynamic secret o k response a status code equal to that given
func (o *UpdateAwsDynamicSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update aws dynamic secret o k response
func (o *UpdateAwsDynamicSecretOK) Code() int {
	return 200
}

func (o *UpdateAwsDynamicSecretOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret/{name}][%d] updateAwsDynamicSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateAwsDynamicSecretOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret/{name}][%d] updateAwsDynamicSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateAwsDynamicSecretOK) GetPayload() *models.Secrets20231128UpdateAwsDynamicSecretResponse {
	return o.Payload
}

func (o *UpdateAwsDynamicSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateAwsDynamicSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsDynamicSecretDefault creates a UpdateAwsDynamicSecretDefault with default headers values
func NewUpdateAwsDynamicSecretDefault(code int) *UpdateAwsDynamicSecretDefault {
	return &UpdateAwsDynamicSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateAwsDynamicSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAwsDynamicSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update aws dynamic secret default response has a 2xx status code
func (o *UpdateAwsDynamicSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update aws dynamic secret default response has a 3xx status code
func (o *UpdateAwsDynamicSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update aws dynamic secret default response has a 4xx status code
func (o *UpdateAwsDynamicSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update aws dynamic secret default response has a 5xx status code
func (o *UpdateAwsDynamicSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update aws dynamic secret default response a status code equal to that given
func (o *UpdateAwsDynamicSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update aws dynamic secret default response
func (o *UpdateAwsDynamicSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAwsDynamicSecretDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret/{name}][%d] UpdateAwsDynamicSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAwsDynamicSecretDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret/{name}][%d] UpdateAwsDynamicSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAwsDynamicSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateAwsDynamicSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}