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

// ListDependenciesReader is a Reader for the ListDependencies structure.
type ListDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDependenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDependenciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDependenciesOK creates a ListDependenciesOK with default headers values
func NewListDependenciesOK() *ListDependenciesOK {
	return &ListDependenciesOK{}
}

/*ListDependenciesOK handles this case with default header values.

A successful response.
*/
type ListDependenciesOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListDependenciesResponse
}

func (o *ListDependenciesOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] listDependenciesOK  %+v", 200, o.Payload)
}

func (o *ListDependenciesOK) GetPayload() *models.HashicorpCloudNetwork20200907ListDependenciesResponse {
	return o.Payload
}

func (o *ListDependenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListDependenciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDependenciesDefault creates a ListDependenciesDefault with default headers values
func NewListDependenciesDefault(code int) *ListDependenciesDefault {
	return &ListDependenciesDefault{
		_statusCode: code,
	}
}

/*ListDependenciesDefault handles this case with default header values.

An unexpected error response.
*/
type ListDependenciesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list dependencies default response
func (o *ListDependenciesDefault) Code() int {
	return o._statusCode
}

func (o *ListDependenciesDefault) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] ListDependencies default  %+v", o._statusCode, o.Payload)
}

func (o *ListDependenciesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListDependenciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
