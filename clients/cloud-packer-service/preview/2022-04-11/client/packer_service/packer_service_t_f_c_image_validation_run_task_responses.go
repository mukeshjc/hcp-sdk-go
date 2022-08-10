// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceTFCImageValidationRunTaskReader is a Reader for the PackerServiceTFCImageValidationRunTask structure.
type PackerServiceTFCImageValidationRunTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceTFCImageValidationRunTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceTFCImageValidationRunTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceTFCImageValidationRunTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceTFCImageValidationRunTaskOK creates a PackerServiceTFCImageValidationRunTaskOK with default headers values
func NewPackerServiceTFCImageValidationRunTaskOK() *PackerServiceTFCImageValidationRunTaskOK {
	return &PackerServiceTFCImageValidationRunTaskOK{}
}

/*PackerServiceTFCImageValidationRunTaskOK handles this case with default header values.

A successful response.
*/
type PackerServiceTFCImageValidationRunTaskOK struct {
	Payload models.HashicorpCloudPacker20220411TFCImageValidationRunTaskResponse
}

func (o *PackerServiceTFCImageValidationRunTaskOK) Error() string {
	return fmt.Sprintf("[POST /packer/2022-04-11/terraform-cloud/validation/{api_id}][%d] packerServiceTFCImageValidationRunTaskOK  %+v", 200, o.Payload)
}

func (o *PackerServiceTFCImageValidationRunTaskOK) GetPayload() models.HashicorpCloudPacker20220411TFCImageValidationRunTaskResponse {
	return o.Payload
}

func (o *PackerServiceTFCImageValidationRunTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceTFCImageValidationRunTaskDefault creates a PackerServiceTFCImageValidationRunTaskDefault with default headers values
func NewPackerServiceTFCImageValidationRunTaskDefault(code int) *PackerServiceTFCImageValidationRunTaskDefault {
	return &PackerServiceTFCImageValidationRunTaskDefault{
		_statusCode: code,
	}
}

/*PackerServiceTFCImageValidationRunTaskDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceTFCImageValidationRunTaskDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service t f c image validation run task default response
func (o *PackerServiceTFCImageValidationRunTaskDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceTFCImageValidationRunTaskDefault) Error() string {
	return fmt.Sprintf("[POST /packer/2022-04-11/terraform-cloud/validation/{api_id}][%d] PackerService_TFCImageValidationRunTask default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceTFCImageValidationRunTaskDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceTFCImageValidationRunTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
