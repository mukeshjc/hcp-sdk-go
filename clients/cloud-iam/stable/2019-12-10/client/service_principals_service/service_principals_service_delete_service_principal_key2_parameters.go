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

// NewServicePrincipalsServiceDeleteServicePrincipalKey2Params creates a new ServicePrincipalsServiceDeleteServicePrincipalKey2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceDeleteServicePrincipalKey2Params() *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithTimeout creates a new ServicePrincipalsServiceDeleteServicePrincipalKey2Params object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2Params{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithContext creates a new ServicePrincipalsServiceDeleteServicePrincipalKey2Params object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2Params{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithHTTPClient creates a new ServicePrincipalsServiceDeleteServicePrincipalKey2Params object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKey2ParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2Params{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKey2Params contains all the parameters to send to the API endpoint

	for the service principals service delete service principal key2 operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKey2Params struct {

	/* ResourceName3.

	     resource_name is the resource name of the service principal key to
	delete.
	*/
	ResourceName3 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service delete service principal key2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WithDefaults() *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service delete service principal key2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName3 adds the resourceName3 to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WithResourceName3(resourceName3 string) *ServicePrincipalsServiceDeleteServicePrincipalKey2Params {
	o.SetResourceName3(resourceName3)
	return o
}

// SetResourceName3 adds the resourceName3 to the service principals service delete service principal key2 params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) SetResourceName3(resourceName3 string) {
	o.ResourceName3 = resourceName3
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name_3
	if err := r.SetPathParam("resource_name_3", o.ResourceName3); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
