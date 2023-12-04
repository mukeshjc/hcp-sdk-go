// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListTGWAttachmentsReader is a Reader for the ListTGWAttachments structure.
type ListTGWAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTGWAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTGWAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListTGWAttachmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTGWAttachmentsOK creates a ListTGWAttachmentsOK with default headers values
func NewListTGWAttachmentsOK() *ListTGWAttachmentsOK {
	return &ListTGWAttachmentsOK{}
}

/*
ListTGWAttachmentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListTGWAttachmentsOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListTGWAttachmentsResponse
}

// IsSuccess returns true when this list t g w attachments o k response has a 2xx status code
func (o *ListTGWAttachmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list t g w attachments o k response has a 3xx status code
func (o *ListTGWAttachmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list t g w attachments o k response has a 4xx status code
func (o *ListTGWAttachmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list t g w attachments o k response has a 5xx status code
func (o *ListTGWAttachmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list t g w attachments o k response a status code equal to that given
func (o *ListTGWAttachmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list t g w attachments o k response
func (o *ListTGWAttachmentsOK) Code() int {
	return 200
}

func (o *ListTGWAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] listTGWAttachmentsOK  %+v", 200, o.Payload)
}

func (o *ListTGWAttachmentsOK) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] listTGWAttachmentsOK  %+v", 200, o.Payload)
}

func (o *ListTGWAttachmentsOK) GetPayload() *models.HashicorpCloudNetwork20200907ListTGWAttachmentsResponse {
	return o.Payload
}

func (o *ListTGWAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListTGWAttachmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTGWAttachmentsDefault creates a ListTGWAttachmentsDefault with default headers values
func NewListTGWAttachmentsDefault(code int) *ListTGWAttachmentsDefault {
	return &ListTGWAttachmentsDefault{
		_statusCode: code,
	}
}

/*
ListTGWAttachmentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListTGWAttachmentsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this list t g w attachments default response has a 2xx status code
func (o *ListTGWAttachmentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list t g w attachments default response has a 3xx status code
func (o *ListTGWAttachmentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list t g w attachments default response has a 4xx status code
func (o *ListTGWAttachmentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list t g w attachments default response has a 5xx status code
func (o *ListTGWAttachmentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list t g w attachments default response a status code equal to that given
func (o *ListTGWAttachmentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list t g w attachments default response
func (o *ListTGWAttachmentsDefault) Code() int {
	return o._statusCode
}

func (o *ListTGWAttachmentsDefault) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] ListTGWAttachments default  %+v", o._statusCode, o.Payload)
}

func (o *ListTGWAttachmentsDefault) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] ListTGWAttachments default  %+v", o._statusCode, o.Payload)
}

func (o *ListTGWAttachmentsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListTGWAttachmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
