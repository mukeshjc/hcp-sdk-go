// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-boundary-service/stable/2021-12-21/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// BoundaryServiceListReader is a Reader for the BoundaryServiceList structure.
type BoundaryServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BoundaryServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBoundaryServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBoundaryServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBoundaryServiceListOK creates a BoundaryServiceListOK with default headers values
func NewBoundaryServiceListOK() *BoundaryServiceListOK {
	return &BoundaryServiceListOK{}
}

/*
BoundaryServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type BoundaryServiceListOK struct {
	Payload *models.HashicorpCloudBoundary20211221ListResponse
}

// IsSuccess returns true when this boundary service list o k response has a 2xx status code
func (o *BoundaryServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this boundary service list o k response has a 3xx status code
func (o *BoundaryServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this boundary service list o k response has a 4xx status code
func (o *BoundaryServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this boundary service list o k response has a 5xx status code
func (o *BoundaryServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this boundary service list o k response a status code equal to that given
func (o *BoundaryServiceListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the boundary service list o k response
func (o *BoundaryServiceListOK) Code() int {
	return 200
}

func (o *BoundaryServiceListOK) Error() string {
	return fmt.Sprintf("[GET /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] boundaryServiceListOK  %+v", 200, o.Payload)
}

func (o *BoundaryServiceListOK) String() string {
	return fmt.Sprintf("[GET /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] boundaryServiceListOK  %+v", 200, o.Payload)
}

func (o *BoundaryServiceListOK) GetPayload() *models.HashicorpCloudBoundary20211221ListResponse {
	return o.Payload
}

func (o *BoundaryServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudBoundary20211221ListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBoundaryServiceListDefault creates a BoundaryServiceListDefault with default headers values
func NewBoundaryServiceListDefault(code int) *BoundaryServiceListDefault {
	return &BoundaryServiceListDefault{
		_statusCode: code,
	}
}

/*
BoundaryServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BoundaryServiceListDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this boundary service list default response has a 2xx status code
func (o *BoundaryServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this boundary service list default response has a 3xx status code
func (o *BoundaryServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this boundary service list default response has a 4xx status code
func (o *BoundaryServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this boundary service list default response has a 5xx status code
func (o *BoundaryServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this boundary service list default response a status code equal to that given
func (o *BoundaryServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the boundary service list default response
func (o *BoundaryServiceListDefault) Code() int {
	return o._statusCode
}

func (o *BoundaryServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] BoundaryService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BoundaryServiceListDefault) String() string {
	return fmt.Sprintf("[GET /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] BoundaryService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BoundaryServiceListDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *BoundaryServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
