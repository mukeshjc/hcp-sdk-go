// Code generated by go-swagger; DO NOT EDIT.

package registry_service

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

// NewActivateRegistryParams creates a new ActivateRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActivateRegistryParams() *ActivateRegistryParams {
	return &ActivateRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActivateRegistryParamsWithTimeout creates a new ActivateRegistryParams object
// with the ability to set a timeout on a request.
func NewActivateRegistryParamsWithTimeout(timeout time.Duration) *ActivateRegistryParams {
	return &ActivateRegistryParams{
		timeout: timeout,
	}
}

// NewActivateRegistryParamsWithContext creates a new ActivateRegistryParams object
// with the ability to set a context for a request.
func NewActivateRegistryParamsWithContext(ctx context.Context) *ActivateRegistryParams {
	return &ActivateRegistryParams{
		Context: ctx,
	}
}

// NewActivateRegistryParamsWithHTTPClient creates a new ActivateRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewActivateRegistryParamsWithHTTPClient(client *http.Client) *ActivateRegistryParams {
	return &ActivateRegistryParams{
		HTTPClient: client,
	}
}

/*
ActivateRegistryParams contains all the parameters to send to the API endpoint

	for the activate registry operation.

	Typically these are written to a http.Request.
*/
type ActivateRegistryParams struct {

	/* Registry.

	   The name of the Registry to activate.
	*/
	Registry string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the activate registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivateRegistryParams) WithDefaults() *ActivateRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the activate registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivateRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the activate registry params
func (o *ActivateRegistryParams) WithTimeout(timeout time.Duration) *ActivateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activate registry params
func (o *ActivateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activate registry params
func (o *ActivateRegistryParams) WithContext(ctx context.Context) *ActivateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activate registry params
func (o *ActivateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activate registry params
func (o *ActivateRegistryParams) WithHTTPClient(client *http.Client) *ActivateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activate registry params
func (o *ActivateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistry adds the registry to the activate registry params
func (o *ActivateRegistryParams) WithRegistry(registry string) *ActivateRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the activate registry params
func (o *ActivateRegistryParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *ActivateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
