// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-04-20/models"
)

// DisableCORSReader is a Reader for the DisableCORS structure.
type DisableCORSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableCORSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableCORSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableCORSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableCORSOK creates a DisableCORSOK with default headers values
func NewDisableCORSOK() *DisableCORSOK {
	return &DisableCORSOK{}
}

/*DisableCORSOK handles this case with default header values.

A successful response.
*/
type DisableCORSOK struct {
	Payload models.HashicorpCloudVault20200420DisableCORSResponse
}

func (o *DisableCORSOK) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/disable-cors][%d] disableCORSOK  %+v", 200, o.Payload)
}

func (o *DisableCORSOK) GetPayload() models.HashicorpCloudVault20200420DisableCORSResponse {
	return o.Payload
}

func (o *DisableCORSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableCORSDefault creates a DisableCORSDefault with default headers values
func NewDisableCORSDefault(code int) *DisableCORSDefault {
	return &DisableCORSDefault{
		_statusCode: code,
	}
}

/*DisableCORSDefault handles this case with default header values.

An unexpected error response.
*/
type DisableCORSDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the disable c o r s default response
func (o *DisableCORSDefault) Code() int {
	return o._statusCode
}

func (o *DisableCORSDefault) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/disable-cors][%d] DisableCORS default  %+v", o._statusCode, o.Payload)
}

func (o *DisableCORSDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DisableCORSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
