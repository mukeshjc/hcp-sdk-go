// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceGetServicePrincipal2Reader is a Reader for the ServicePrincipalsServiceGetServicePrincipal2 structure.
type ServicePrincipalsServiceGetServicePrincipal2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceGetServicePrincipal2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceGetServicePrincipal2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceGetServicePrincipal2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceGetServicePrincipal2OK creates a ServicePrincipalsServiceGetServicePrincipal2OK with default headers values
func NewServicePrincipalsServiceGetServicePrincipal2OK() *ServicePrincipalsServiceGetServicePrincipal2OK {
	return &ServicePrincipalsServiceGetServicePrincipal2OK{}
}

/*
ServicePrincipalsServiceGetServicePrincipal2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceGetServicePrincipal2OK struct {
	Payload *models.HashicorpCloudIamGetServicePrincipalResponse
}

// IsSuccess returns true when this service principals service get service principal2 o k response has a 2xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service get service principal2 o k response has a 3xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service get service principal2 o k response has a 4xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service get service principal2 o k response has a 5xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service get service principal2 o k response a status code equal to that given
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service get service principal2 o k response
func (o *ServicePrincipalsServiceGetServicePrincipal2OK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceGetServicePrincipal2OK) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name_1}][%d] servicePrincipalsServiceGetServicePrincipal2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceGetServicePrincipal2OK) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name_1}][%d] servicePrincipalsServiceGetServicePrincipal2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceGetServicePrincipal2OK) GetPayload() *models.HashicorpCloudIamGetServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceGetServicePrincipal2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetServicePrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceGetServicePrincipal2Default creates a ServicePrincipalsServiceGetServicePrincipal2Default with default headers values
func NewServicePrincipalsServiceGetServicePrincipal2Default(code int) *ServicePrincipalsServiceGetServicePrincipal2Default {
	return &ServicePrincipalsServiceGetServicePrincipal2Default{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceGetServicePrincipal2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceGetServicePrincipal2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service get service principal2 default response has a 2xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service get service principal2 default response has a 3xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service get service principal2 default response has a 4xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service get service principal2 default response has a 5xx status code
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service get service principal2 default response a status code equal to that given
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service get service principal2 default response
func (o *ServicePrincipalsServiceGetServicePrincipal2Default) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceGetServicePrincipal2Default) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name_1}][%d] ServicePrincipalsService_GetServicePrincipal2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceGetServicePrincipal2Default) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{resource_name_1}][%d] ServicePrincipalsService_GetServicePrincipal2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceGetServicePrincipal2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceGetServicePrincipal2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
