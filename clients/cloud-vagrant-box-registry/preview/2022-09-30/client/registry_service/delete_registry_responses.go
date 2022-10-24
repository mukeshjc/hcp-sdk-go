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

// DeleteRegistryReader is a Reader for the DeleteRegistry structure.
type DeleteRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegistryOK creates a DeleteRegistryOK with default headers values
func NewDeleteRegistryOK() *DeleteRegistryOK {
	return &DeleteRegistryOK{}
}

/*
DeleteRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRegistryOK struct {
	Payload models.HashicorpCloudVagrantDeleteRegistryResponse
}

// IsSuccess returns true when this delete registry o k response has a 2xx status code
func (o *DeleteRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry o k response has a 3xx status code
func (o *DeleteRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry o k response has a 4xx status code
func (o *DeleteRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry o k response has a 5xx status code
func (o *DeleteRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry o k response a status code equal to that given
func (o *DeleteRegistryOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRegistryOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}][%d] deleteRegistryOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryOK) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}][%d] deleteRegistryOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryOK) GetPayload() models.HashicorpCloudVagrantDeleteRegistryResponse {
	return o.Payload
}

func (o *DeleteRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}