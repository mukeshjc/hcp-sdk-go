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

// DeleteGcpIntegrationReader is a Reader for the DeleteGcpIntegration structure.
type DeleteGcpIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGcpIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGcpIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteGcpIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGcpIntegrationOK creates a DeleteGcpIntegrationOK with default headers values
func NewDeleteGcpIntegrationOK() *DeleteGcpIntegrationOK {
	return &DeleteGcpIntegrationOK{}
}

/*
DeleteGcpIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteGcpIntegrationOK struct {
	Payload models.Secrets20231128DeleteGcpIntegrationResponse
}

// IsSuccess returns true when this delete gcp integration o k response has a 2xx status code
func (o *DeleteGcpIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete gcp integration o k response has a 3xx status code
func (o *DeleteGcpIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gcp integration o k response has a 4xx status code
func (o *DeleteGcpIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gcp integration o k response has a 5xx status code
func (o *DeleteGcpIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gcp integration o k response a status code equal to that given
func (o *DeleteGcpIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete gcp integration o k response
func (o *DeleteGcpIntegrationOK) Code() int {
	return 200
}

func (o *DeleteGcpIntegrationOK) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/gcp/config/{name}][%d] deleteGcpIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteGcpIntegrationOK) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/gcp/config/{name}][%d] deleteGcpIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteGcpIntegrationOK) GetPayload() models.Secrets20231128DeleteGcpIntegrationResponse {
	return o.Payload
}

func (o *DeleteGcpIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGcpIntegrationDefault creates a DeleteGcpIntegrationDefault with default headers values
func NewDeleteGcpIntegrationDefault(code int) *DeleteGcpIntegrationDefault {
	return &DeleteGcpIntegrationDefault{
		_statusCode: code,
	}
}

/*
DeleteGcpIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteGcpIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete gcp integration default response has a 2xx status code
func (o *DeleteGcpIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete gcp integration default response has a 3xx status code
func (o *DeleteGcpIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete gcp integration default response has a 4xx status code
func (o *DeleteGcpIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete gcp integration default response has a 5xx status code
func (o *DeleteGcpIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete gcp integration default response a status code equal to that given
func (o *DeleteGcpIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete gcp integration default response
func (o *DeleteGcpIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGcpIntegrationDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/gcp/config/{name}][%d] DeleteGcpIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGcpIntegrationDefault) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/gcp/config/{name}][%d] DeleteGcpIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGcpIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteGcpIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}