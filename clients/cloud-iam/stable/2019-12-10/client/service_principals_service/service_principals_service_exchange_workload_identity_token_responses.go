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

// ServicePrincipalsServiceExchangeWorkloadIdentityTokenReader is a Reader for the ServicePrincipalsServiceExchangeWorkloadIdentityToken structure.
type ServicePrincipalsServiceExchangeWorkloadIdentityTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceExchangeWorkloadIdentityTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenOK creates a ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK with default headers values
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenOK() *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK{}
}

/*
ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK struct {
	Payload *models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse
}

// IsSuccess returns true when this service principals service exchange workload identity token o k response has a 2xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service exchange workload identity token o k response has a 3xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service exchange workload identity token o k response has a 4xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service exchange workload identity token o k response has a 5xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service exchange workload identity token o k response a status code equal to that given
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] servicePrincipalsServiceExchangeWorkloadIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] servicePrincipalsServiceExchangeWorkloadIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) GetPayload() *models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault creates a ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault with default headers values
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault(code int) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service exchange workload identity token default response
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service exchange workload identity token default response has a 2xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service exchange workload identity token default response has a 3xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service exchange workload identity token default response has a 4xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service exchange workload identity token default response has a 5xx status code
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service exchange workload identity token default response a status code equal to that given
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] ServicePrincipalsService_ExchangeWorkloadIdentityToken default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] ServicePrincipalsService_ExchangeWorkloadIdentityToken default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}