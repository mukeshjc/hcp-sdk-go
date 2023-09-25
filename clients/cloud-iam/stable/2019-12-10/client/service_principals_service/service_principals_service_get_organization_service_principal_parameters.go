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

// NewServicePrincipalsServiceGetOrganizationServicePrincipalParams creates a new ServicePrincipalsServiceGetOrganizationServicePrincipalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceGetOrganizationServicePrincipalParams() *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	return &ServicePrincipalsServiceGetOrganizationServicePrincipalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithTimeout creates a new ServicePrincipalsServiceGetOrganizationServicePrincipalParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	return &ServicePrincipalsServiceGetOrganizationServicePrincipalParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithContext creates a new ServicePrincipalsServiceGetOrganizationServicePrincipalParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithContext(ctx context.Context) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	return &ServicePrincipalsServiceGetOrganizationServicePrincipalParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithHTTPClient creates a new ServicePrincipalsServiceGetOrganizationServicePrincipalParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceGetOrganizationServicePrincipalParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	return &ServicePrincipalsServiceGetOrganizationServicePrincipalParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceGetOrganizationServicePrincipalParams contains all the parameters to send to the API endpoint

	for the service principals service get organization service principal operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceGetOrganizationServicePrincipalParams struct {

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization that
	owns the requested service principal or project.
	*/
	OrganizationID string

	/* PrincipalID.

	   principal_id is the ID of the service principal being requested.
	*/
	PrincipalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service get organization service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithDefaults() *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service get organization service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithContext(ctx context.Context) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithOrganizationID(organizationID string) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPrincipalID adds the principalID to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WithPrincipalID(principalID string) *ServicePrincipalsServiceGetOrganizationServicePrincipalParams {
	o.SetPrincipalID(principalID)
	return o
}

// SetPrincipalID adds the principalId to the service principals service get organization service principal params
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) SetPrincipalID(principalID string) {
	o.PrincipalID = principalID
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceGetOrganizationServicePrincipalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param principal_id
	if err := r.SetPathParam("principal_id", o.PrincipalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}