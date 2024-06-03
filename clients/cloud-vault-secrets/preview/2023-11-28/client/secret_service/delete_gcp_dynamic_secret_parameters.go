// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewDeleteGcpDynamicSecretParams creates a new DeleteGcpDynamicSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGcpDynamicSecretParams() *DeleteGcpDynamicSecretParams {
	return &DeleteGcpDynamicSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGcpDynamicSecretParamsWithTimeout creates a new DeleteGcpDynamicSecretParams object
// with the ability to set a timeout on a request.
func NewDeleteGcpDynamicSecretParamsWithTimeout(timeout time.Duration) *DeleteGcpDynamicSecretParams {
	return &DeleteGcpDynamicSecretParams{
		timeout: timeout,
	}
}

// NewDeleteGcpDynamicSecretParamsWithContext creates a new DeleteGcpDynamicSecretParams object
// with the ability to set a context for a request.
func NewDeleteGcpDynamicSecretParamsWithContext(ctx context.Context) *DeleteGcpDynamicSecretParams {
	return &DeleteGcpDynamicSecretParams{
		Context: ctx,
	}
}

// NewDeleteGcpDynamicSecretParamsWithHTTPClient creates a new DeleteGcpDynamicSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGcpDynamicSecretParamsWithHTTPClient(client *http.Client) *DeleteGcpDynamicSecretParams {
	return &DeleteGcpDynamicSecretParams{
		HTTPClient: client,
	}
}

/*
DeleteGcpDynamicSecretParams contains all the parameters to send to the API endpoint

	for the delete gcp dynamic secret operation.

	Typically these are written to a http.Request.
*/
type DeleteGcpDynamicSecretParams struct {

	// AppName.
	AppName string

	// Name.
	Name string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete gcp dynamic secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGcpDynamicSecretParams) WithDefaults() *DeleteGcpDynamicSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete gcp dynamic secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGcpDynamicSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithTimeout(timeout time.Duration) *DeleteGcpDynamicSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithContext(ctx context.Context) *DeleteGcpDynamicSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithHTTPClient(client *http.Client) *DeleteGcpDynamicSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithAppName(appName string) *DeleteGcpDynamicSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithName adds the name to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithName(name string) *DeleteGcpDynamicSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithOrganizationID(organizationID string) *DeleteGcpDynamicSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) WithProjectID(projectID string) *DeleteGcpDynamicSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete gcp dynamic secret params
func (o *DeleteGcpDynamicSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGcpDynamicSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
