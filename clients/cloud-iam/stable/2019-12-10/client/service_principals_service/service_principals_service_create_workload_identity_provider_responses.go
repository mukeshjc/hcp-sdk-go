// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceCreateWorkloadIdentityProviderReader is a Reader for the ServicePrincipalsServiceCreateWorkloadIdentityProvider structure.
type ServicePrincipalsServiceCreateWorkloadIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceCreateWorkloadIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceCreateWorkloadIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderOK creates a ServicePrincipalsServiceCreateWorkloadIdentityProviderOK with default headers values
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderOK() *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderOK{}
}

/*
ServicePrincipalsServiceCreateWorkloadIdentityProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceCreateWorkloadIdentityProviderOK struct {
	Payload *models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse
}

// IsSuccess returns true when this service principals service create workload identity provider o k response has a 2xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service create workload identity provider o k response has a 3xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service create workload identity provider o k response has a 4xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service create workload identity provider o k response has a 5xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service create workload identity provider o k response a status code equal to that given
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service create workload identity provider o k response
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] servicePrincipalsServiceCreateWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] servicePrincipalsServiceCreateWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) GetPayload() *models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderDefault creates a ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault with default headers values
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderDefault(code int) *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service create workload identity provider default response has a 2xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service create workload identity provider default response has a 3xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service create workload identity provider default response has a 4xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service create workload identity provider default response has a 5xx status code
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service create workload identity provider default response a status code equal to that given
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service create workload identity provider default response
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] ServicePrincipalsService_CreateWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] ServicePrincipalsService_CreateWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServicePrincipalsServiceCreateWorkloadIdentityProviderBody CreateWorkloadIdentityProviderRequest is the request message used when
// creating a workload identity provider.
swagger:model ServicePrincipalsServiceCreateWorkloadIdentityProviderBody
*/
type ServicePrincipalsServiceCreateWorkloadIdentityProviderBody struct {

	// name is the name of the workload identity provider to be created.
	Name string `json:"name,omitempty"`

	// provider is the workload identity provider to create.
	Provider *models.HashicorpCloudIamWorkloadIdentityProvider `json:"provider,omitempty"`
}

// Validate validates this service principals service create workload identity provider body
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(o.Provider) { // not required
		return nil
	}

	if o.Provider != nil {
		if err := o.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "provider")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service principals service create workload identity provider body based on the context it is used
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if o.Provider != nil {

		if swag.IsZero(o.Provider) { // not required
			return nil
		}

		if err := o.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderBody) UnmarshalBinary(b []byte) error {
	var res ServicePrincipalsServiceCreateWorkloadIdentityProviderBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
