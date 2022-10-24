// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// ListVersionsReader is a Reader for the ListVersions structure.
type ListVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListVersionsOK creates a ListVersionsOK with default headers values
func NewListVersionsOK() *ListVersionsOK {
	return &ListVersionsOK{}
}

/*
ListVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListVersionsOK struct {
	Payload *models.HashicorpCloudVagrantListVersionsResponse
}

// IsSuccess returns true when this list versions o k response has a 2xx status code
func (o *ListVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list versions o k response has a 3xx status code
func (o *ListVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list versions o k response has a 4xx status code
func (o *ListVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list versions o k response has a 5xx status code
func (o *ListVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list versions o k response a status code equal to that given
func (o *ListVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVersionsOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions][%d] listVersionsOK  %+v", 200, o.Payload)
}

func (o *ListVersionsOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions][%d] listVersionsOK  %+v", 200, o.Payload)
}

func (o *ListVersionsOK) GetPayload() *models.HashicorpCloudVagrantListVersionsResponse {
	return o.Payload
}

func (o *ListVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrantListVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}