// Code generated by go-swagger; DO NOT EDIT.

package product_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/models"
)

// ProductServiceListAvailablePackagesReader is a Reader for the ProductServiceListAvailablePackages structure.
type ProductServiceListAvailablePackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductServiceListAvailablePackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductServiceListAvailablePackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProductServiceListAvailablePackagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProductServiceListAvailablePackagesOK creates a ProductServiceListAvailablePackagesOK with default headers values
func NewProductServiceListAvailablePackagesOK() *ProductServiceListAvailablePackagesOK {
	return &ProductServiceListAvailablePackagesOK{}
}

/*
ProductServiceListAvailablePackagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProductServiceListAvailablePackagesOK struct {
	Payload *models.Billing20201105ListAvailablePackagesResponse
}

// IsSuccess returns true when this product service list available packages o k response has a 2xx status code
func (o *ProductServiceListAvailablePackagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this product service list available packages o k response has a 3xx status code
func (o *ProductServiceListAvailablePackagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this product service list available packages o k response has a 4xx status code
func (o *ProductServiceListAvailablePackagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this product service list available packages o k response has a 5xx status code
func (o *ProductServiceListAvailablePackagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this product service list available packages o k response a status code equal to that given
func (o *ProductServiceListAvailablePackagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProductServiceListAvailablePackagesOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/packages][%d] productServiceListAvailablePackagesOK  %+v", 200, o.Payload)
}

func (o *ProductServiceListAvailablePackagesOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/packages][%d] productServiceListAvailablePackagesOK  %+v", 200, o.Payload)
}

func (o *ProductServiceListAvailablePackagesOK) GetPayload() *models.Billing20201105ListAvailablePackagesResponse {
	return o.Payload
}

func (o *ProductServiceListAvailablePackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105ListAvailablePackagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductServiceListAvailablePackagesDefault creates a ProductServiceListAvailablePackagesDefault with default headers values
func NewProductServiceListAvailablePackagesDefault(code int) *ProductServiceListAvailablePackagesDefault {
	return &ProductServiceListAvailablePackagesDefault{
		_statusCode: code,
	}
}

/*
ProductServiceListAvailablePackagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProductServiceListAvailablePackagesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the product service list available packages default response
func (o *ProductServiceListAvailablePackagesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this product service list available packages default response has a 2xx status code
func (o *ProductServiceListAvailablePackagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this product service list available packages default response has a 3xx status code
func (o *ProductServiceListAvailablePackagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this product service list available packages default response has a 4xx status code
func (o *ProductServiceListAvailablePackagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this product service list available packages default response has a 5xx status code
func (o *ProductServiceListAvailablePackagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this product service list available packages default response a status code equal to that given
func (o *ProductServiceListAvailablePackagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProductServiceListAvailablePackagesDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/packages][%d] ProductService_ListAvailablePackages default  %+v", o._statusCode, o.Payload)
}

func (o *ProductServiceListAvailablePackagesDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/packages][%d] ProductService_ListAvailablePackages default  %+v", o._statusCode, o.Payload)
}

func (o *ProductServiceListAvailablePackagesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ProductServiceListAvailablePackagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
