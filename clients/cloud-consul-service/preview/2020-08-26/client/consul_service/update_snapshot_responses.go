// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/models"
)

// UpdateSnapshotReader is a Reader for the UpdateSnapshot structure.
type UpdateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSnapshotOK creates a UpdateSnapshotOK with default headers values
func NewUpdateSnapshotOK() *UpdateSnapshotOK {
	return &UpdateSnapshotOK{}
}

/*UpdateSnapshotOK handles this case with default header values.

A successful response.
*/
type UpdateSnapshotOK struct {
	Payload *models.HashicorpCloudConsul20200826UpdateSnapshotResponse
}

func (o *UpdateSnapshotOK) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.id}][%d] updateSnapshotOK  %+v", 200, o.Payload)
}

func (o *UpdateSnapshotOK) GetPayload() *models.HashicorpCloudConsul20200826UpdateSnapshotResponse {
	return o.Payload
}

func (o *UpdateSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826UpdateSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnapshotDefault creates a UpdateSnapshotDefault with default headers values
func NewUpdateSnapshotDefault(code int) *UpdateSnapshotDefault {
	return &UpdateSnapshotDefault{
		_statusCode: code,
	}
}

/*UpdateSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateSnapshotDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the update snapshot default response
func (o *UpdateSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSnapshotDefault) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.id}][%d] UpdateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSnapshotDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
