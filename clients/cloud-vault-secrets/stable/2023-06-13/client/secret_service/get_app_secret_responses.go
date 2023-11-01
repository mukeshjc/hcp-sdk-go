// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// GetAppSecretReader is a Reader for the GetAppSecret structure.
type GetAppSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppSecretOK creates a GetAppSecretOK with default headers values
func NewGetAppSecretOK() *GetAppSecretOK {
	return &GetAppSecretOK{}
}

/*
GetAppSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAppSecretOK struct {
	Payload *models.Secrets20230613GetAppSecretResponse
}

// IsSuccess returns true when this get app secret o k response has a 2xx status code
func (o *GetAppSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app secret o k response has a 3xx status code
func (o *GetAppSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secret o k response has a 4xx status code
func (o *GetAppSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app secret o k response has a 5xx status code
func (o *GetAppSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secret o k response a status code equal to that given
func (o *GetAppSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAppSecretOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/secrets/{secret_name}][%d] getAppSecretOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/secrets/{secret_name}][%d] getAppSecretOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretOK) GetPayload() *models.Secrets20230613GetAppSecretResponse {
	return o.Payload
}

func (o *GetAppSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetAppSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretDefault creates a GetAppSecretDefault with default headers values
func NewGetAppSecretDefault(code int) *GetAppSecretDefault {
	return &GetAppSecretDefault{
		_statusCode: code,
	}
}

/*
GetAppSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAppSecretDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the get app secret default response
func (o *GetAppSecretDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get app secret default response has a 2xx status code
func (o *GetAppSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get app secret default response has a 3xx status code
func (o *GetAppSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get app secret default response has a 4xx status code
func (o *GetAppSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get app secret default response has a 5xx status code
func (o *GetAppSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get app secret default response a status code equal to that given
func (o *GetAppSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAppSecretDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/secrets/{secret_name}][%d] GetAppSecret default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSecretDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/secrets/{secret_name}][%d] GetAppSecret default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSecretDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetAppSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
