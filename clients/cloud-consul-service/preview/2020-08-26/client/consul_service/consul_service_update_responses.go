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

// ConsulServiceUpdateReader is a Reader for the ConsulServiceUpdate structure.
type ConsulServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsulServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsulServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsulServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsulServiceUpdateOK creates a ConsulServiceUpdateOK with default headers values
func NewConsulServiceUpdateOK() *ConsulServiceUpdateOK {
	return &ConsulServiceUpdateOK{}
}

/*ConsulServiceUpdateOK handles this case with default header values.

A successful response.
*/
type ConsulServiceUpdateOK struct {
	Payload *models.HashicorpCloudConsul20200826UpdateResponse
}

func (o *ConsulServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters/{cluster.id}][%d] consulServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ConsulServiceUpdateOK) GetPayload() *models.HashicorpCloudConsul20200826UpdateResponse {
	return o.Payload
}

func (o *ConsulServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826UpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsulServiceUpdateDefault creates a ConsulServiceUpdateDefault with default headers values
func NewConsulServiceUpdateDefault(code int) *ConsulServiceUpdateDefault {
	return &ConsulServiceUpdateDefault{
		_statusCode: code,
	}
}

/*ConsulServiceUpdateDefault handles this case with default header values.

An unexpected error response.
*/
type ConsulServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the consul service update default response
func (o *ConsulServiceUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ConsulServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters/{cluster.id}][%d] ConsulService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ConsulServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ConsulServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
