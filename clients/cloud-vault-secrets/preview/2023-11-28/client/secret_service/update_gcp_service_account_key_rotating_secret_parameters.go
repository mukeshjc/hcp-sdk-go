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

// NewUpdateGcpServiceAccountKeyRotatingSecretParams creates a new UpdateGcpServiceAccountKeyRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGcpServiceAccountKeyRotatingSecretParams() *UpdateGcpServiceAccountKeyRotatingSecretParams {
	return &UpdateGcpServiceAccountKeyRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithTimeout creates a new UpdateGcpServiceAccountKeyRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithTimeout(timeout time.Duration) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	return &UpdateGcpServiceAccountKeyRotatingSecretParams{
		timeout: timeout,
	}
}

// NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithContext creates a new UpdateGcpServiceAccountKeyRotatingSecretParams object
// with the ability to set a context for a request.
func NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithContext(ctx context.Context) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	return &UpdateGcpServiceAccountKeyRotatingSecretParams{
		Context: ctx,
	}
}

// NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithHTTPClient creates a new UpdateGcpServiceAccountKeyRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGcpServiceAccountKeyRotatingSecretParamsWithHTTPClient(client *http.Client) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	return &UpdateGcpServiceAccountKeyRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateGcpServiceAccountKeyRotatingSecretParams contains all the parameters to send to the API endpoint

	for the update gcp service account key rotating secret operation.

	Typically these are written to a http.Request.
*/
type UpdateGcpServiceAccountKeyRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody

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

// WithDefaults hydrates default values in the update gcp service account key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithDefaults() *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update gcp service account key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithTimeout(timeout time.Duration) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithContext(ctx context.Context) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithHTTPClient(client *http.Client) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithAppName(appName string) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithBody(body *models.SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetBody(body *models.SecretServiceUpdateGcpServiceAccountKeyRotatingSecretBody) {
	o.Body = body
}

// WithName adds the name to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithName(name string) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithOrganizationID(organizationID string) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WithProjectID(projectID string) *UpdateGcpServiceAccountKeyRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update gcp service account key rotating secret params
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGcpServiceAccountKeyRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}
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
