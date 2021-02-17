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
)

// NetworkServiceDeletePeeringReader is a Reader for the NetworkServiceDeletePeering structure.
type NetworkServiceDeletePeeringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkServiceDeletePeeringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkServiceDeletePeeringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkServiceDeletePeeringDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkServiceDeletePeeringOK creates a NetworkServiceDeletePeeringOK with default headers values
func NewNetworkServiceDeletePeeringOK() *NetworkServiceDeletePeeringOK {
	return &NetworkServiceDeletePeeringOK{}
}

/*NetworkServiceDeletePeeringOK handles this case with default header values.

A successful response.
*/
type NetworkServiceDeletePeeringOK struct {
	Payload *models.HashicorpCloudNetwork20200907DeletePeeringResponse
}

func (o *NetworkServiceDeletePeeringOK) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] networkServiceDeletePeeringOK  %+v", 200, o.Payload)
}

func (o *NetworkServiceDeletePeeringOK) GetPayload() *models.HashicorpCloudNetwork20200907DeletePeeringResponse {
	return o.Payload
}

func (o *NetworkServiceDeletePeeringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907DeletePeeringResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkServiceDeletePeeringDefault creates a NetworkServiceDeletePeeringDefault with default headers values
func NewNetworkServiceDeletePeeringDefault(code int) *NetworkServiceDeletePeeringDefault {
	return &NetworkServiceDeletePeeringDefault{
		_statusCode: code,
	}
}

/*NetworkServiceDeletePeeringDefault handles this case with default header values.

An unexpected error response.
*/
type NetworkServiceDeletePeeringDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the network service delete peering default response
func (o *NetworkServiceDeletePeeringDefault) Code() int {
	return o._statusCode
}

func (o *NetworkServiceDeletePeeringDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] NetworkService_DeletePeering default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkServiceDeletePeeringDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NetworkServiceDeletePeeringDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
