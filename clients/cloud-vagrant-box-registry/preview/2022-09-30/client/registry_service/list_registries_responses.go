// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// ListRegistriesReader is a Reader for the ListRegistries structure.
type ListRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRegistriesOK creates a ListRegistriesOK with default headers values
func NewListRegistriesOK() *ListRegistriesOK {
	return &ListRegistriesOK{}
}

/*
ListRegistriesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRegistriesOK struct {
	Payload *models.HashicorpCloudVagrantListRegistriesResponse
}

// IsSuccess returns true when this list registries o k response has a 2xx status code
func (o *ListRegistriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list registries o k response has a 3xx status code
func (o *ListRegistriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registries o k response has a 4xx status code
func (o *ListRegistriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list registries o k response has a 5xx status code
func (o *ListRegistriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list registries o k response a status code equal to that given
func (o *ListRegistriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry][%d] listRegistriesOK  %+v", 200, o.Payload)
}

func (o *ListRegistriesOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry][%d] listRegistriesOK  %+v", 200, o.Payload)
}

func (o *ListRegistriesOK) GetPayload() *models.HashicorpCloudVagrantListRegistriesResponse {
	return o.Payload
}

func (o *ListRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrantListRegistriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}