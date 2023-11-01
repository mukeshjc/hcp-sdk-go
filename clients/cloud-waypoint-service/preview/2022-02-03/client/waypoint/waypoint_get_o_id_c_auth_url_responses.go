// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// WaypointGetOIDCAuthURLReader is a Reader for the WaypointGetOIDCAuthURL structure.
type WaypointGetOIDCAuthURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetOIDCAuthURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetOIDCAuthURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetOIDCAuthURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetOIDCAuthURLOK creates a WaypointGetOIDCAuthURLOK with default headers values
func NewWaypointGetOIDCAuthURLOK() *WaypointGetOIDCAuthURLOK {
	return &WaypointGetOIDCAuthURLOK{}
}

/*
WaypointGetOIDCAuthURLOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetOIDCAuthURLOK struct {
	Payload *models.HashicorpWaypointGetOIDCAuthURLResponse
}

// IsSuccess returns true when this waypoint get o Id c auth Url o k response has a 2xx status code
func (o *WaypointGetOIDCAuthURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get o Id c auth Url o k response has a 3xx status code
func (o *WaypointGetOIDCAuthURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get o Id c auth Url o k response has a 4xx status code
func (o *WaypointGetOIDCAuthURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get o Id c auth Url o k response has a 5xx status code
func (o *WaypointGetOIDCAuthURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get o Id c auth Url o k response a status code equal to that given
func (o *WaypointGetOIDCAuthURLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get o Id c auth Url o k response
func (o *WaypointGetOIDCAuthURLOK) Code() int {
	return 200
}

func (o *WaypointGetOIDCAuthURLOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/oidc/{auth_method.name}/url][%d] waypointGetOIdCAuthUrlOK  %+v", 200, o.Payload)
}

func (o *WaypointGetOIDCAuthURLOK) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/oidc/{auth_method.name}/url][%d] waypointGetOIdCAuthUrlOK  %+v", 200, o.Payload)
}

func (o *WaypointGetOIDCAuthURLOK) GetPayload() *models.HashicorpWaypointGetOIDCAuthURLResponse {
	return o.Payload
}

func (o *WaypointGetOIDCAuthURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetOIDCAuthURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetOIDCAuthURLDefault creates a WaypointGetOIDCAuthURLDefault with default headers values
func NewWaypointGetOIDCAuthURLDefault(code int) *WaypointGetOIDCAuthURLDefault {
	return &WaypointGetOIDCAuthURLDefault{
		_statusCode: code,
	}
}

/*
WaypointGetOIDCAuthURLDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetOIDCAuthURLDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get o ID c auth URL default response has a 2xx status code
func (o *WaypointGetOIDCAuthURLDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get o ID c auth URL default response has a 3xx status code
func (o *WaypointGetOIDCAuthURLDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get o ID c auth URL default response has a 4xx status code
func (o *WaypointGetOIDCAuthURLDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get o ID c auth URL default response has a 5xx status code
func (o *WaypointGetOIDCAuthURLDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get o ID c auth URL default response a status code equal to that given
func (o *WaypointGetOIDCAuthURLDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get o ID c auth URL default response
func (o *WaypointGetOIDCAuthURLDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetOIDCAuthURLDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/oidc/{auth_method.name}/url][%d] Waypoint_GetOIDCAuthURL default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetOIDCAuthURLDefault) String() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/oidc/{auth_method.name}/url][%d] Waypoint_GetOIDCAuthURL default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetOIDCAuthURLDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetOIDCAuthURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
