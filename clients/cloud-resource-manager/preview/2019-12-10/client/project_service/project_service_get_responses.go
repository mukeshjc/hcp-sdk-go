// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ProjectServiceGetReader is a Reader for the ProjectServiceGet structure.
type ProjectServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceGetOK creates a ProjectServiceGetOK with default headers values
func NewProjectServiceGetOK() *ProjectServiceGetOK {
	return &ProjectServiceGetOK{}
}

/*ProjectServiceGetOK handles this case with default header values.

A successful response.
*/
type ProjectServiceGetOK struct {
	Payload *models.HashicorpCloudResourcemanagerProjectGetResponse
}

func (o *ProjectServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/projects/{id}][%d] projectServiceGetOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceGetOK) GetPayload() *models.HashicorpCloudResourcemanagerProjectGetResponse {
	return o.Payload
}

func (o *ProjectServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerProjectGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceGetDefault creates a ProjectServiceGetDefault with default headers values
func NewProjectServiceGetDefault(code int) *ProjectServiceGetDefault {
	return &ProjectServiceGetDefault{
		_statusCode: code,
	}
}

/*ProjectServiceGetDefault handles this case with default header values.

An unexpected error response.
*/
type ProjectServiceGetDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the project service get default response
func (o *ProjectServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/projects/{id}][%d] ProjectService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceGetDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ProjectServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
