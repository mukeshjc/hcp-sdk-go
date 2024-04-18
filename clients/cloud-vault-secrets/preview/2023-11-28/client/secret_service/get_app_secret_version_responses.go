// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// GetAppSecretVersionReader is a Reader for the GetAppSecretVersion structure.
type GetAppSecretVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSecretVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSecretVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppSecretVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppSecretVersionOK creates a GetAppSecretVersionOK with default headers values
func NewGetAppSecretVersionOK() *GetAppSecretVersionOK {
	return &GetAppSecretVersionOK{}
}

/*
GetAppSecretVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAppSecretVersionOK struct {
	Payload *models.Secrets20231128GetAppSecretVersionResponse
}

// IsSuccess returns true when this get app secret version o k response has a 2xx status code
func (o *GetAppSecretVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app secret version o k response has a 3xx status code
func (o *GetAppSecretVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secret version o k response has a 4xx status code
func (o *GetAppSecretVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app secret version o k response has a 5xx status code
func (o *GetAppSecretVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secret version o k response a status code equal to that given
func (o *GetAppSecretVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app secret version o k response
func (o *GetAppSecretVersionOK) Code() int {
	return 200
}

func (o *GetAppSecretVersionOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}/versions/{version}][%d] getAppSecretVersionOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretVersionOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}/versions/{version}][%d] getAppSecretVersionOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretVersionOK) GetPayload() *models.Secrets20231128GetAppSecretVersionResponse {
	return o.Payload
}

func (o *GetAppSecretVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetAppSecretVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretVersionDefault creates a GetAppSecretVersionDefault with default headers values
func NewGetAppSecretVersionDefault(code int) *GetAppSecretVersionDefault {
	return &GetAppSecretVersionDefault{
		_statusCode: code,
	}
}

/*
GetAppSecretVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAppSecretVersionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get app secret version default response has a 2xx status code
func (o *GetAppSecretVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get app secret version default response has a 3xx status code
func (o *GetAppSecretVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get app secret version default response has a 4xx status code
func (o *GetAppSecretVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get app secret version default response has a 5xx status code
func (o *GetAppSecretVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get app secret version default response a status code equal to that given
func (o *GetAppSecretVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get app secret version default response
func (o *GetAppSecretVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetAppSecretVersionDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}/versions/{version}][%d] GetAppSecretVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSecretVersionDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}/versions/{version}][%d] GetAppSecretVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSecretVersionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetAppSecretVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}