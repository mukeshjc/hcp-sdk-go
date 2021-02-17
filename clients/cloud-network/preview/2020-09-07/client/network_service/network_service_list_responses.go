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

// NetworkServiceListReader is a Reader for the NetworkServiceList structure.
type NetworkServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkServiceListOK creates a NetworkServiceListOK with default headers values
func NewNetworkServiceListOK() *NetworkServiceListOK {
	return &NetworkServiceListOK{}
}

/*NetworkServiceListOK handles this case with default header values.

A successful response.
*/
type NetworkServiceListOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListResponse
}

func (o *NetworkServiceListOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks][%d] networkServiceListOK  %+v", 200, o.Payload)
}

func (o *NetworkServiceListOK) GetPayload() *models.HashicorpCloudNetwork20200907ListResponse {
	return o.Payload
}

func (o *NetworkServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkServiceListDefault creates a NetworkServiceListDefault with default headers values
func NewNetworkServiceListDefault(code int) *NetworkServiceListDefault {
	return &NetworkServiceListDefault{
		_statusCode: code,
	}
}

/*NetworkServiceListDefault handles this case with default header values.

An unexpected error response.
*/
type NetworkServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the network service list default response
func (o *NetworkServiceListDefault) Code() int {
	return o._statusCode
}

func (o *NetworkServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks][%d] NetworkService_List default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NetworkServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
