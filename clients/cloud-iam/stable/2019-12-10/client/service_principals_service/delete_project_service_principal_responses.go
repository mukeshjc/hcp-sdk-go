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

// DeleteProjectServicePrincipalReader is a Reader for the DeleteProjectServicePrincipal structure.
type DeleteProjectServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProjectServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectServicePrincipalOK creates a DeleteProjectServicePrincipalOK with default headers values
func NewDeleteProjectServicePrincipalOK() *DeleteProjectServicePrincipalOK {
	return &DeleteProjectServicePrincipalOK{}
}

/*
DeleteProjectServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProjectServicePrincipalOK struct {
	Payload models.HashicorpCloudIamDeleteProjectServicePrincipalResponse
}

// IsSuccess returns true when this delete project service principal o k response has a 2xx status code
func (o *DeleteProjectServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project service principal o k response has a 3xx status code
func (o *DeleteProjectServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project service principal o k response has a 4xx status code
func (o *DeleteProjectServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project service principal o k response has a 5xx status code
func (o *DeleteProjectServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project service principal o k response a status code equal to that given
func (o *DeleteProjectServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteProjectServicePrincipalOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] deleteProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectServicePrincipalOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] deleteProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectServicePrincipalOK) GetPayload() models.HashicorpCloudIamDeleteProjectServicePrincipalResponse {
	return o.Payload
}

func (o *DeleteProjectServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectServicePrincipalDefault creates a DeleteProjectServicePrincipalDefault with default headers values
func NewDeleteProjectServicePrincipalDefault(code int) *DeleteProjectServicePrincipalDefault {
	return &DeleteProjectServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
DeleteProjectServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProjectServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete project service principal default response
func (o *DeleteProjectServicePrincipalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete project service principal default response has a 2xx status code
func (o *DeleteProjectServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project service principal default response has a 3xx status code
func (o *DeleteProjectServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project service principal default response has a 4xx status code
func (o *DeleteProjectServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project service principal default response has a 5xx status code
func (o *DeleteProjectServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project service principal default response a status code equal to that given
func (o *DeleteProjectServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteProjectServicePrincipalDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] DeleteProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectServicePrincipalDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] DeleteProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteProjectServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}