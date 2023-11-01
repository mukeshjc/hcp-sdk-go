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

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParams creates a new ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParams() *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithTimeout creates a new ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithContext creates a new ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithContext(ctx context.Context) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithHTTPClient creates a new ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the service principals service update workload identity provider operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams struct {

	/* Provider.

	   provider is the workload identity provider to update.
	*/
	Provider ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody

	/* ProviderResourceName.

	     resource_name is the resource name for this workload identity
	provider.
	*/
	ProviderResourceName string

	/* UpdateMask.

	   update_mask describing which resource attribute(s) to update.
	*/
	UpdateMask *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service update workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithDefaults() *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service update workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithContext(ctx context.Context) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithProvider(provider ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetProvider(provider ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) {
	o.Provider = provider
}

// WithProviderResourceName adds the providerResourceName to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithProviderResourceName(providerResourceName string) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetProviderResourceName(providerResourceName)
	return o
}

// SetProviderResourceName adds the providerResourceName to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetProviderResourceName(providerResourceName string) {
	o.ProviderResourceName = providerResourceName
}

// WithUpdateMask adds the updateMask to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WithUpdateMask(updateMask *string) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams {
	o.SetUpdateMask(updateMask)
	return o
}

// SetUpdateMask adds the updateMask to the service principals service update workload identity provider params
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) SetUpdateMask(updateMask *string) {
	o.UpdateMask = updateMask
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Provider); err != nil {
		return err
	}

	// path param provider.resource_name
	if err := r.SetPathParam("provider.resource_name", o.ProviderResourceName); err != nil {
		return err
	}

	if o.UpdateMask != nil {

		// query param update_mask
		var qrUpdateMask string

		if o.UpdateMask != nil {
			qrUpdateMask = *o.UpdateMask
		}
		qUpdateMask := qrUpdateMask
		if qUpdateMask != "" {

			if err := r.SetQueryParam("update_mask", qUpdateMask); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
