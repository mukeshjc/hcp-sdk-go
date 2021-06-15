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

// DeleteSnapshotReader is a Reader for the DeleteSnapshot structure.
type DeleteSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSnapshotOK creates a DeleteSnapshotOK with default headers values
func NewDeleteSnapshotOK() *DeleteSnapshotOK {
	return &DeleteSnapshotOK{}
}

/*DeleteSnapshotOK handles this case with default header values.

A successful response.
*/
type DeleteSnapshotOK struct {
	Payload *models.HashicorpCloudVault20200420DeleteSnapshotResponse
}

func (o *DeleteSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] deleteSnapshotOK  %+v", 200, o.Payload)
}

func (o *DeleteSnapshotOK) GetPayload() *models.HashicorpCloudVault20200420DeleteSnapshotResponse {
	return o.Payload
}

func (o *DeleteSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420DeleteSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnapshotDefault creates a DeleteSnapshotDefault with default headers values
func NewDeleteSnapshotDefault(code int) *DeleteSnapshotDefault {
	return &DeleteSnapshotDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete snapshot default response
func (o *DeleteSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] DeleteSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeleteSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
