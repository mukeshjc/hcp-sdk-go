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

// WebhookServiceListWebhooksReader is a Reader for the WebhookServiceListWebhooks structure.
type WebhookServiceListWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookServiceListWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebhookServiceListWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebhookServiceListWebhooksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebhookServiceListWebhooksOK creates a WebhookServiceListWebhooksOK with default headers values
func NewWebhookServiceListWebhooksOK() *WebhookServiceListWebhooksOK {
	return &WebhookServiceListWebhooksOK{}
}

/*
WebhookServiceListWebhooksOK describes a response with status code 200, with default header values.

A successful response.
*/
type WebhookServiceListWebhooksOK struct {
	Payload *models.HashicorpCloudWebhookListWebhooksResponse
}

// IsSuccess returns true when this webhook service list webhooks o k response has a 2xx status code
func (o *WebhookServiceListWebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webhook service list webhooks o k response has a 3xx status code
func (o *WebhookServiceListWebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook service list webhooks o k response has a 4xx status code
func (o *WebhookServiceListWebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook service list webhooks o k response has a 5xx status code
func (o *WebhookServiceListWebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook service list webhooks o k response a status code equal to that given
func (o *WebhookServiceListWebhooksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webhook service list webhooks o k response
func (o *WebhookServiceListWebhooksOK) Code() int {
	return 200
}

func (o *WebhookServiceListWebhooksOK) Error() string {
	return fmt.Sprintf("[GET /2023-05-31/webhook/{parent_resource_name}/webhooks][%d] webhookServiceListWebhooksOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceListWebhooksOK) String() string {
	return fmt.Sprintf("[GET /2023-05-31/webhook/{parent_resource_name}/webhooks][%d] webhookServiceListWebhooksOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceListWebhooksOK) GetPayload() *models.HashicorpCloudWebhookListWebhooksResponse {
	return o.Payload
}

func (o *WebhookServiceListWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWebhookListWebhooksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookServiceListWebhooksDefault creates a WebhookServiceListWebhooksDefault with default headers values
func NewWebhookServiceListWebhooksDefault(code int) *WebhookServiceListWebhooksDefault {
	return &WebhookServiceListWebhooksDefault{
		_statusCode: code,
	}
}

/*
WebhookServiceListWebhooksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WebhookServiceListWebhooksDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this webhook service list webhooks default response has a 2xx status code
func (o *WebhookServiceListWebhooksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webhook service list webhooks default response has a 3xx status code
func (o *WebhookServiceListWebhooksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webhook service list webhooks default response has a 4xx status code
func (o *WebhookServiceListWebhooksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webhook service list webhooks default response has a 5xx status code
func (o *WebhookServiceListWebhooksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webhook service list webhooks default response a status code equal to that given
func (o *WebhookServiceListWebhooksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webhook service list webhooks default response
func (o *WebhookServiceListWebhooksDefault) Code() int {
	return o._statusCode
}

func (o *WebhookServiceListWebhooksDefault) Error() string {
	return fmt.Sprintf("[GET /2023-05-31/webhook/{parent_resource_name}/webhooks][%d] WebhookService_ListWebhooks default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceListWebhooksDefault) String() string {
	return fmt.Sprintf("[GET /2023-05-31/webhook/{parent_resource_name}/webhooks][%d] WebhookService_ListWebhooks default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceListWebhooksDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WebhookServiceListWebhooksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
