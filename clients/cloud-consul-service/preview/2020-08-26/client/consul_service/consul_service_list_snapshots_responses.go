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

// ConsulServiceListSnapshotsReader is a Reader for the ConsulServiceListSnapshots structure.
type ConsulServiceListSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsulServiceListSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsulServiceListSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsulServiceListSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsulServiceListSnapshotsOK creates a ConsulServiceListSnapshotsOK with default headers values
func NewConsulServiceListSnapshotsOK() *ConsulServiceListSnapshotsOK {
	return &ConsulServiceListSnapshotsOK{}
}

/*ConsulServiceListSnapshotsOK handles this case with default header values.

A successful response.
*/
type ConsulServiceListSnapshotsOK struct {
	Payload *models.HashicorpCloudConsul20200826ListSnapshotsResponse
}

func (o *ConsulServiceListSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] consulServiceListSnapshotsOK  %+v", 200, o.Payload)
}

func (o *ConsulServiceListSnapshotsOK) GetPayload() *models.HashicorpCloudConsul20200826ListSnapshotsResponse {
	return o.Payload
}

func (o *ConsulServiceListSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826ListSnapshotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsulServiceListSnapshotsDefault creates a ConsulServiceListSnapshotsDefault with default headers values
func NewConsulServiceListSnapshotsDefault(code int) *ConsulServiceListSnapshotsDefault {
	return &ConsulServiceListSnapshotsDefault{
		_statusCode: code,
	}
}

/*ConsulServiceListSnapshotsDefault handles this case with default header values.

An unexpected error response.
*/
type ConsulServiceListSnapshotsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the consul service list snapshots default response
func (o *ConsulServiceListSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *ConsulServiceListSnapshotsDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] ConsulService_ListSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *ConsulServiceListSnapshotsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ConsulServiceListSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
