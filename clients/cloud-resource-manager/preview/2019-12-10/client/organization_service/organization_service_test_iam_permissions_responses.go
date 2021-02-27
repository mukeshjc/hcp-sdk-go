// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// OrganizationServiceTestIamPermissionsReader is a Reader for the OrganizationServiceTestIamPermissions structure.
type OrganizationServiceTestIamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceTestIamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceTestIamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceTestIamPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceTestIamPermissionsOK creates a OrganizationServiceTestIamPermissionsOK with default headers values
func NewOrganizationServiceTestIamPermissionsOK() *OrganizationServiceTestIamPermissionsOK {
	return &OrganizationServiceTestIamPermissionsOK{}
}

/*OrganizationServiceTestIamPermissionsOK handles this case with default header values.

A successful response.
*/
type OrganizationServiceTestIamPermissionsOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse
}

func (o *OrganizationServiceTestIamPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] organizationServiceTestIamPermissionsOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse {
	return o.Payload
}

func (o *OrganizationServiceTestIamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceTestIamPermissionsDefault creates a OrganizationServiceTestIamPermissionsDefault with default headers values
func NewOrganizationServiceTestIamPermissionsDefault(code int) *OrganizationServiceTestIamPermissionsDefault {
	return &OrganizationServiceTestIamPermissionsDefault{
		_statusCode: code,
	}
}

/*OrganizationServiceTestIamPermissionsDefault handles this case with default header values.

An unexpected error response.
*/
type OrganizationServiceTestIamPermissionsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service test iam permissions default response
func (o *OrganizationServiceTestIamPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceTestIamPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] OrganizationService_TestIamPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceTestIamPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
