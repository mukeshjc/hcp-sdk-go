// Code generated by go-swagger; DO NOT EDIT.

package project_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ProjectServiceSetIamPolicyReader is a Reader for the ProjectServiceSetIamPolicy structure.
type ProjectServiceSetIamPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceSetIamPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceSetIamPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceSetIamPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceSetIamPolicyOK creates a ProjectServiceSetIamPolicyOK with default headers values
func NewProjectServiceSetIamPolicyOK() *ProjectServiceSetIamPolicyOK {
	return &ProjectServiceSetIamPolicyOK{}
}

/*
ProjectServiceSetIamPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceSetIamPolicyOK struct {
	Payload *models.HashicorpCloudResourcemanagerProjectSetIamPolicyResponse
}

// IsSuccess returns true when this project service set iam policy o k response has a 2xx status code
func (o *ProjectServiceSetIamPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project service set iam policy o k response has a 3xx status code
func (o *ProjectServiceSetIamPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project service set iam policy o k response has a 4xx status code
func (o *ProjectServiceSetIamPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project service set iam policy o k response has a 5xx status code
func (o *ProjectServiceSetIamPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project service set iam policy o k response a status code equal to that given
func (o *ProjectServiceSetIamPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectServiceSetIamPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/iam-policy][%d] projectServiceSetIamPolicyOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetIamPolicyOK) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/iam-policy][%d] projectServiceSetIamPolicyOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetIamPolicyOK) GetPayload() *models.HashicorpCloudResourcemanagerProjectSetIamPolicyResponse {
	return o.Payload
}

func (o *ProjectServiceSetIamPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerProjectSetIamPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceSetIamPolicyDefault creates a ProjectServiceSetIamPolicyDefault with default headers values
func NewProjectServiceSetIamPolicyDefault(code int) *ProjectServiceSetIamPolicyDefault {
	return &ProjectServiceSetIamPolicyDefault{
		_statusCode: code,
	}
}

/*
ProjectServiceSetIamPolicyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceSetIamPolicyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the project service set iam policy default response
func (o *ProjectServiceSetIamPolicyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this project service set iam policy default response has a 2xx status code
func (o *ProjectServiceSetIamPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this project service set iam policy default response has a 3xx status code
func (o *ProjectServiceSetIamPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this project service set iam policy default response has a 4xx status code
func (o *ProjectServiceSetIamPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this project service set iam policy default response has a 5xx status code
func (o *ProjectServiceSetIamPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this project service set iam policy default response a status code equal to that given
func (o *ProjectServiceSetIamPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProjectServiceSetIamPolicyDefault) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/iam-policy][%d] ProjectService_SetIamPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetIamPolicyDefault) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/iam-policy][%d] ProjectService_SetIamPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetIamPolicyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ProjectServiceSetIamPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ProjectServiceSetIamPolicyBody ProjectSetIamPolicyRequest see ProjectService.SetIamPolicy
swagger:model ProjectServiceSetIamPolicyBody
*/
type ProjectServiceSetIamPolicyBody struct {

	// Policy is the IAM policy to be updated for the project. The policy
	// will be completely replaced and therefore Policy.Etag must be specified
	// in order to prevent concurrent updates.
	Policy *models.HashicorpCloudResourcemanagerPolicy `json:"policy,omitempty"`
}

// Validate validates this project service set iam policy body
func (o *ProjectServiceSetIamPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectServiceSetIamPolicyBody) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.Policy) { // not required
		return nil
	}

	if o.Policy != nil {
		if err := o.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project service set iam policy body based on the context it is used
func (o *ProjectServiceSetIamPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectServiceSetIamPolicyBody) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if o.Policy != nil {
		if err := o.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ProjectServiceSetIamPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectServiceSetIamPolicyBody) UnmarshalBinary(b []byte) error {
	var res ProjectServiceSetIamPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
