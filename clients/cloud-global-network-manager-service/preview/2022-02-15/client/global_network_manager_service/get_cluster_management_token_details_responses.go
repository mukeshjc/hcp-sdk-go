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

// GetClusterManagementTokenDetailsReader is a Reader for the GetClusterManagementTokenDetails structure.
type GetClusterManagementTokenDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterManagementTokenDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterManagementTokenDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterManagementTokenDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterManagementTokenDetailsOK creates a GetClusterManagementTokenDetailsOK with default headers values
func NewGetClusterManagementTokenDetailsOK() *GetClusterManagementTokenDetailsOK {
	return &GetClusterManagementTokenDetailsOK{}
}

/*
GetClusterManagementTokenDetailsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetClusterManagementTokenDetailsOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215GetClusterManagementTokenDetailsResponse
}

// IsSuccess returns true when this get cluster management token details o k response has a 2xx status code
func (o *GetClusterManagementTokenDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster management token details o k response has a 3xx status code
func (o *GetClusterManagementTokenDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster management token details o k response has a 4xx status code
func (o *GetClusterManagementTokenDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster management token details o k response has a 5xx status code
func (o *GetClusterManagementTokenDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster management token details o k response a status code equal to that given
func (o *GetClusterManagementTokenDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterManagementTokenDetailsOK) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/management_token_details][%d] getClusterManagementTokenDetailsOK  %+v", 200, o.Payload)
}

func (o *GetClusterManagementTokenDetailsOK) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/management_token_details][%d] getClusterManagementTokenDetailsOK  %+v", 200, o.Payload)
}

func (o *GetClusterManagementTokenDetailsOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215GetClusterManagementTokenDetailsResponse {
	return o.Payload
}

func (o *GetClusterManagementTokenDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215GetClusterManagementTokenDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterManagementTokenDetailsDefault creates a GetClusterManagementTokenDetailsDefault with default headers values
func NewGetClusterManagementTokenDetailsDefault(code int) *GetClusterManagementTokenDetailsDefault {
	return &GetClusterManagementTokenDetailsDefault{
		_statusCode: code,
	}
}

/*
GetClusterManagementTokenDetailsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetClusterManagementTokenDetailsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the get cluster management token details default response
func (o *GetClusterManagementTokenDetailsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster management token details default response has a 2xx status code
func (o *GetClusterManagementTokenDetailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster management token details default response has a 3xx status code
func (o *GetClusterManagementTokenDetailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster management token details default response has a 4xx status code
func (o *GetClusterManagementTokenDetailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster management token details default response has a 5xx status code
func (o *GetClusterManagementTokenDetailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster management token details default response a status code equal to that given
func (o *GetClusterManagementTokenDetailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterManagementTokenDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/management_token_details][%d] GetClusterManagementTokenDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterManagementTokenDetailsDefault) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/management_token_details][%d] GetClusterManagementTokenDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterManagementTokenDetailsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GetClusterManagementTokenDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
