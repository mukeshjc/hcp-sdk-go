// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateHVNRouteReader is a Reader for the CreateHVNRoute structure.
type CreateHVNRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHVNRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHVNRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateHVNRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHVNRouteOK creates a CreateHVNRouteOK with default headers values
func NewCreateHVNRouteOK() *CreateHVNRouteOK {
	return &CreateHVNRouteOK{}
}

/*CreateHVNRouteOK handles this case with default header values.

A successful response.
*/
type CreateHVNRouteOK struct {
	Payload *models.HashicorpCloudNetwork20200907CreateHVNRouteResponse
}

func (o *CreateHVNRouteOK) Error() string {
	return fmt.Sprintf("[POST /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] createHVNRouteOK  %+v", 200, o.Payload)
}

func (o *CreateHVNRouteOK) GetPayload() *models.HashicorpCloudNetwork20200907CreateHVNRouteResponse {
	return o.Payload
}

func (o *CreateHVNRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907CreateHVNRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHVNRouteDefault creates a CreateHVNRouteDefault with default headers values
func NewCreateHVNRouteDefault(code int) *CreateHVNRouteDefault {
	return &CreateHVNRouteDefault{
		_statusCode: code,
	}
}

/*CreateHVNRouteDefault handles this case with default header values.

An unexpected error response.
*/
type CreateHVNRouteDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create h v n route default response
func (o *CreateHVNRouteDefault) Code() int {
	return o._statusCode
}

func (o *CreateHVNRouteDefault) Error() string {
	return fmt.Sprintf("[POST /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] CreateHVNRoute default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHVNRouteDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateHVNRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}