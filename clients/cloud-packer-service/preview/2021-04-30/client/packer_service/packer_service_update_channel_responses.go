// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateChannelReader is a Reader for the PackerServiceUpdateChannel structure.
type PackerServiceUpdateChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateChannelOK creates a PackerServiceUpdateChannelOK with default headers values
func NewPackerServiceUpdateChannelOK() *PackerServiceUpdateChannelOK {
	return &PackerServiceUpdateChannelOK{}
}

/*PackerServiceUpdateChannelOK handles this case with default header values.

A successful response.
*/
type PackerServiceUpdateChannelOK struct {
	Payload *models.HashicorpCloudPackerUpdateChannelResponse
}

func (o *PackerServiceUpdateChannelOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceUpdateChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateChannelOK) GetPayload() *models.HashicorpCloudPackerUpdateChannelResponse {
	return o.Payload
}

func (o *PackerServiceUpdateChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateChannelDefault creates a PackerServiceUpdateChannelDefault with default headers values
func NewPackerServiceUpdateChannelDefault(code int) *PackerServiceUpdateChannelDefault {
	return &PackerServiceUpdateChannelDefault{
		_statusCode: code,
	}
}

/*PackerServiceUpdateChannelDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceUpdateChannelDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service update channel default response
func (o *PackerServiceUpdateChannelDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceUpdateChannelDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_UpdateChannel default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateChannelDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceUpdateChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}