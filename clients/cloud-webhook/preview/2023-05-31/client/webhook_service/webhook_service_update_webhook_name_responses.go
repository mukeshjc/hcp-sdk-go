// Code generated by go-swagger; DO NOT EDIT.

package webhook_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-webhook/preview/2023-05-31/models"
)

// WebhookServiceUpdateWebhookNameReader is a Reader for the WebhookServiceUpdateWebhookName structure.
type WebhookServiceUpdateWebhookNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookServiceUpdateWebhookNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebhookServiceUpdateWebhookNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebhookServiceUpdateWebhookNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebhookServiceUpdateWebhookNameOK creates a WebhookServiceUpdateWebhookNameOK with default headers values
func NewWebhookServiceUpdateWebhookNameOK() *WebhookServiceUpdateWebhookNameOK {
	return &WebhookServiceUpdateWebhookNameOK{}
}

/*
WebhookServiceUpdateWebhookNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type WebhookServiceUpdateWebhookNameOK struct {
	Payload *models.HashicorpCloudWebhookUpdateWebhookNameResponse
}

// IsSuccess returns true when this webhook service update webhook name o k response has a 2xx status code
func (o *WebhookServiceUpdateWebhookNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webhook service update webhook name o k response has a 3xx status code
func (o *WebhookServiceUpdateWebhookNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook service update webhook name o k response has a 4xx status code
func (o *WebhookServiceUpdateWebhookNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook service update webhook name o k response has a 5xx status code
func (o *WebhookServiceUpdateWebhookNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook service update webhook name o k response a status code equal to that given
func (o *WebhookServiceUpdateWebhookNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webhook service update webhook name o k response
func (o *WebhookServiceUpdateWebhookNameOK) Code() int {
	return 200
}

func (o *WebhookServiceUpdateWebhookNameOK) Error() string {
	return fmt.Sprintf("[PATCH /2023-05-31/{resource_name}/name][%d] webhookServiceUpdateWebhookNameOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceUpdateWebhookNameOK) String() string {
	return fmt.Sprintf("[PATCH /2023-05-31/{resource_name}/name][%d] webhookServiceUpdateWebhookNameOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceUpdateWebhookNameOK) GetPayload() *models.HashicorpCloudWebhookUpdateWebhookNameResponse {
	return o.Payload
}

func (o *WebhookServiceUpdateWebhookNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWebhookUpdateWebhookNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookServiceUpdateWebhookNameDefault creates a WebhookServiceUpdateWebhookNameDefault with default headers values
func NewWebhookServiceUpdateWebhookNameDefault(code int) *WebhookServiceUpdateWebhookNameDefault {
	return &WebhookServiceUpdateWebhookNameDefault{
		_statusCode: code,
	}
}

/*
WebhookServiceUpdateWebhookNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WebhookServiceUpdateWebhookNameDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this webhook service update webhook name default response has a 2xx status code
func (o *WebhookServiceUpdateWebhookNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webhook service update webhook name default response has a 3xx status code
func (o *WebhookServiceUpdateWebhookNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webhook service update webhook name default response has a 4xx status code
func (o *WebhookServiceUpdateWebhookNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webhook service update webhook name default response has a 5xx status code
func (o *WebhookServiceUpdateWebhookNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webhook service update webhook name default response a status code equal to that given
func (o *WebhookServiceUpdateWebhookNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webhook service update webhook name default response
func (o *WebhookServiceUpdateWebhookNameDefault) Code() int {
	return o._statusCode
}

func (o *WebhookServiceUpdateWebhookNameDefault) Error() string {
	return fmt.Sprintf("[PATCH /2023-05-31/{resource_name}/name][%d] WebhookService_UpdateWebhookName default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceUpdateWebhookNameDefault) String() string {
	return fmt.Sprintf("[PATCH /2023-05-31/{resource_name}/name][%d] WebhookService_UpdateWebhookName default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceUpdateWebhookNameDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WebhookServiceUpdateWebhookNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
