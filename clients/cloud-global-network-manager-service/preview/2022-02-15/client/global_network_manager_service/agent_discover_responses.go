// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// AgentDiscoverReader is a Reader for the AgentDiscover structure.
type AgentDiscoverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentDiscoverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentDiscoverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAgentDiscoverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentDiscoverOK creates a AgentDiscoverOK with default headers values
func NewAgentDiscoverOK() *AgentDiscoverOK {
	return &AgentDiscoverOK{}
}

/*
AgentDiscoverOK describes a response with status code 200, with default header values.

A successful response.
*/
type AgentDiscoverOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215AgentDiscoverResponse
}

// IsSuccess returns true when this agent discover o k response has a 2xx status code
func (o *AgentDiscoverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent discover o k response has a 3xx status code
func (o *AgentDiscoverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent discover o k response has a 4xx status code
func (o *AgentDiscoverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent discover o k response has a 5xx status code
func (o *AgentDiscoverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent discover o k response a status code equal to that given
func (o *AgentDiscoverOK) IsCode(code int) bool {
	return code == 200
}

func (o *AgentDiscoverOK) Error() string {
	return fmt.Sprintf("[POST /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/discover][%d] agentDiscoverOK  %+v", 200, o.Payload)
}

func (o *AgentDiscoverOK) String() string {
	return fmt.Sprintf("[POST /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/discover][%d] agentDiscoverOK  %+v", 200, o.Payload)
}

func (o *AgentDiscoverOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215AgentDiscoverResponse {
	return o.Payload
}

func (o *AgentDiscoverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215AgentDiscoverResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentDiscoverDefault creates a AgentDiscoverDefault with default headers values
func NewAgentDiscoverDefault(code int) *AgentDiscoverDefault {
	return &AgentDiscoverDefault{
		_statusCode: code,
	}
}

/*
AgentDiscoverDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AgentDiscoverDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the agent discover default response
func (o *AgentDiscoverDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this agent discover default response has a 2xx status code
func (o *AgentDiscoverDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this agent discover default response has a 3xx status code
func (o *AgentDiscoverDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this agent discover default response has a 4xx status code
func (o *AgentDiscoverDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this agent discover default response has a 5xx status code
func (o *AgentDiscoverDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this agent discover default response a status code equal to that given
func (o *AgentDiscoverDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AgentDiscoverDefault) Error() string {
	return fmt.Sprintf("[POST /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/discover][%d] AgentDiscover default  %+v", o._statusCode, o.Payload)
}

func (o *AgentDiscoverDefault) String() string {
	return fmt.Sprintf("[POST /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/discover][%d] AgentDiscover default  %+v", o._statusCode, o.Payload)
}

func (o *AgentDiscoverDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AgentDiscoverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
