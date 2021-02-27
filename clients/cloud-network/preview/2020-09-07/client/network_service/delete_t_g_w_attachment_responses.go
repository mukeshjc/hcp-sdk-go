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

// DeleteTGWAttachmentReader is a Reader for the DeleteTGWAttachment structure.
type DeleteTGWAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTGWAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTGWAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteTGWAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTGWAttachmentOK creates a DeleteTGWAttachmentOK with default headers values
func NewDeleteTGWAttachmentOK() *DeleteTGWAttachmentOK {
	return &DeleteTGWAttachmentOK{}
}

/*DeleteTGWAttachmentOK handles this case with default header values.

A successful response.
*/
type DeleteTGWAttachmentOK struct {
	Payload *models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse
}

func (o *DeleteTGWAttachmentOK) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments/{id}][%d] deleteTGWAttachmentOK  %+v", 200, o.Payload)
}

func (o *DeleteTGWAttachmentOK) GetPayload() *models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse {
	return o.Payload
}

func (o *DeleteTGWAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTGWAttachmentDefault creates a DeleteTGWAttachmentDefault with default headers values
func NewDeleteTGWAttachmentDefault(code int) *DeleteTGWAttachmentDefault {
	return &DeleteTGWAttachmentDefault{
		_statusCode: code,
	}
}

/*DeleteTGWAttachmentDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteTGWAttachmentDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete t g w attachment default response
func (o *DeleteTGWAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTGWAttachmentDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments/{id}][%d] DeleteTGWAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTGWAttachmentDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeleteTGWAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
