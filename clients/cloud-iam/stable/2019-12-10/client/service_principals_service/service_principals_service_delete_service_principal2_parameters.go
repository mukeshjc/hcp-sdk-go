// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewServicePrincipalsServiceDeleteServicePrincipal2Params creates a new ServicePrincipalsServiceDeleteServicePrincipal2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceDeleteServicePrincipal2Params() *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipal2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithTimeout creates a new ServicePrincipalsServiceDeleteServicePrincipal2Params object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipal2Params{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithContext creates a new ServicePrincipalsServiceDeleteServicePrincipal2Params object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipal2Params{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithHTTPClient creates a new ServicePrincipalsServiceDeleteServicePrincipal2Params object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceDeleteServicePrincipal2ParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipal2Params{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceDeleteServicePrincipal2Params contains all the parameters to send to the API endpoint

	for the service principals service delete service principal2 operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceDeleteServicePrincipal2Params struct {

	/* ResourceName1.

	   resource_name is the resource name of the service principal to delete.
	*/
	ResourceName1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service delete service principal2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WithDefaults() *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service delete service principal2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName1 adds the resourceName1 to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WithResourceName1(resourceName1 string) *ServicePrincipalsServiceDeleteServicePrincipal2Params {
	o.SetResourceName1(resourceName1)
	return o
}

// SetResourceName1 adds the resourceName1 to the service principals service delete service principal2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) SetResourceName1(resourceName1 string) {
	o.ResourceName1 = resourceName1
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceDeleteServicePrincipal2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name_1
	if err := r.SetPathParam("resource_name_1", o.ResourceName1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
