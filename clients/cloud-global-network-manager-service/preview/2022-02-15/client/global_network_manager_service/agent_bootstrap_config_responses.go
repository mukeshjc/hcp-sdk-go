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

// AgentBootstrapConfigReader is a Reader for the AgentBootstrapConfig structure.
type AgentBootstrapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentBootstrapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentBootstrapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAgentBootstrapConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentBootstrapConfigOK creates a AgentBootstrapConfigOK with default headers values
func NewAgentBootstrapConfigOK() *AgentBootstrapConfigOK {
	return &AgentBootstrapConfigOK{}
}

/*
AgentBootstrapConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type AgentBootstrapConfigOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215AgentBootstrapResponse
}

// IsSuccess returns true when this agent bootstrap config o k response has a 2xx status code
func (o *AgentBootstrapConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent bootstrap config o k response has a 3xx status code
func (o *AgentBootstrapConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent bootstrap config o k response has a 4xx status code
func (o *AgentBootstrapConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent bootstrap config o k response has a 5xx status code
func (o *AgentBootstrapConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent bootstrap config o k response a status code equal to that given
func (o *AgentBootstrapConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *AgentBootstrapConfigOK) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/bootstrap_config][%d] agentBootstrapConfigOK  %+v", 200, o.Payload)
}

func (o *AgentBootstrapConfigOK) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/bootstrap_config][%d] agentBootstrapConfigOK  %+v", 200, o.Payload)
}

func (o *AgentBootstrapConfigOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215AgentBootstrapResponse {
	return o.Payload
}

func (o *AgentBootstrapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215AgentBootstrapResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentBootstrapConfigDefault creates a AgentBootstrapConfigDefault with default headers values
func NewAgentBootstrapConfigDefault(code int) *AgentBootstrapConfigDefault {
	return &AgentBootstrapConfigDefault{
		_statusCode: code,
	}
}

/*
AgentBootstrapConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AgentBootstrapConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the agent bootstrap config default response
func (o *AgentBootstrapConfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this agent bootstrap config default response has a 2xx status code
func (o *AgentBootstrapConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this agent bootstrap config default response has a 3xx status code
func (o *AgentBootstrapConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this agent bootstrap config default response has a 4xx status code
func (o *AgentBootstrapConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this agent bootstrap config default response has a 5xx status code
func (o *AgentBootstrapConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this agent bootstrap config default response a status code equal to that given
func (o *AgentBootstrapConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AgentBootstrapConfigDefault) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/bootstrap_config][%d] AgentBootstrapConfig default  %+v", o._statusCode, o.Payload)
}

func (o *AgentBootstrapConfigDefault) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/bootstrap_config][%d] AgentBootstrapConfig default  %+v", o._statusCode, o.Payload)
}

func (o *AgentBootstrapConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AgentBootstrapConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}