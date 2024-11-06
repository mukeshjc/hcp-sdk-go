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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// NewUpdatePostgresIntegrationParams creates a new UpdatePostgresIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePostgresIntegrationParams() *UpdatePostgresIntegrationParams {
	return &UpdatePostgresIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePostgresIntegrationParamsWithTimeout creates a new UpdatePostgresIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpdatePostgresIntegrationParamsWithTimeout(timeout time.Duration) *UpdatePostgresIntegrationParams {
	return &UpdatePostgresIntegrationParams{
		timeout: timeout,
	}
}

// NewUpdatePostgresIntegrationParamsWithContext creates a new UpdatePostgresIntegrationParams object
// with the ability to set a context for a request.
func NewUpdatePostgresIntegrationParamsWithContext(ctx context.Context) *UpdatePostgresIntegrationParams {
	return &UpdatePostgresIntegrationParams{
		Context: ctx,
	}
}

// NewUpdatePostgresIntegrationParamsWithHTTPClient creates a new UpdatePostgresIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePostgresIntegrationParamsWithHTTPClient(client *http.Client) *UpdatePostgresIntegrationParams {
	return &UpdatePostgresIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpdatePostgresIntegrationParams contains all the parameters to send to the API endpoint

	for the update postgres integration operation.

	Typically these are written to a http.Request.
*/
type UpdatePostgresIntegrationParams struct {

	// Body.
	Body *models.SecretServiceUpdatePostgresIntegrationBody

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

// WithDefaults hydrates default values in the update postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePostgresIntegrationParams) WithDefaults() *UpdatePostgresIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePostgresIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithTimeout(timeout time.Duration) *UpdatePostgresIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithContext(ctx context.Context) *UpdatePostgresIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithHTTPClient(client *http.Client) *UpdatePostgresIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithBody(body *models.SecretServiceUpdatePostgresIntegrationBody) *UpdatePostgresIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetBody(body *models.SecretServiceUpdatePostgresIntegrationBody) {
	o.Body = body
}

// WithName adds the name to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithName(name string) *UpdatePostgresIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithOrganizationID(organizationID string) *UpdatePostgresIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) WithProjectID(projectID string) *UpdatePostgresIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update postgres integration params
func (o *UpdatePostgresIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePostgresIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
