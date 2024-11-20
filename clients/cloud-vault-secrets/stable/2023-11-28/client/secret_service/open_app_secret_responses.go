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

// OpenAppSecretReader is a Reader for the OpenAppSecret structure.
type OpenAppSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenAppSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenAppSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOpenAppSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOpenAppSecretOK creates a OpenAppSecretOK with default headers values
func NewOpenAppSecretOK() *OpenAppSecretOK {
	return &OpenAppSecretOK{}
}

/*
OpenAppSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type OpenAppSecretOK struct {
	Payload *models.Secrets20231128OpenAppSecretResponse
}

// IsSuccess returns true when this open app secret o k response has a 2xx status code
func (o *OpenAppSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this open app secret o k response has a 3xx status code
func (o *OpenAppSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open app secret o k response has a 4xx status code
func (o *OpenAppSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this open app secret o k response has a 5xx status code
func (o *OpenAppSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this open app secret o k response a status code equal to that given
func (o *OpenAppSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the open app secret o k response
func (o *OpenAppSecretOK) Code() int {
	return 200
}

func (o *OpenAppSecretOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:open][%d] openAppSecretOK  %+v", 200, o.Payload)
}

func (o *OpenAppSecretOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:open][%d] openAppSecretOK  %+v", 200, o.Payload)
}

func (o *OpenAppSecretOK) GetPayload() *models.Secrets20231128OpenAppSecretResponse {
	return o.Payload
}

func (o *OpenAppSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128OpenAppSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenAppSecretDefault creates a OpenAppSecretDefault with default headers values
func NewOpenAppSecretDefault(code int) *OpenAppSecretDefault {
	return &OpenAppSecretDefault{
		_statusCode: code,
	}
}

/*
OpenAppSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OpenAppSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this open app secret default response has a 2xx status code
func (o *OpenAppSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this open app secret default response has a 3xx status code
func (o *OpenAppSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this open app secret default response has a 4xx status code
func (o *OpenAppSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this open app secret default response has a 5xx status code
func (o *OpenAppSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this open app secret default response a status code equal to that given
func (o *OpenAppSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the open app secret default response
func (o *OpenAppSecretDefault) Code() int {
	return o._statusCode
}

func (o *OpenAppSecretDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:open][%d] OpenAppSecret default  %+v", o._statusCode, o.Payload)
}

func (o *OpenAppSecretDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:open][%d] OpenAppSecret default  %+v", o._statusCode, o.Payload)
}

func (o *OpenAppSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *OpenAppSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}