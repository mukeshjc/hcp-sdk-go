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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// NewUpdateMongoDBAtlasRotatingSecretParams creates a new UpdateMongoDBAtlasRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMongoDBAtlasRotatingSecretParams() *UpdateMongoDBAtlasRotatingSecretParams {
	return &UpdateMongoDBAtlasRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMongoDBAtlasRotatingSecretParamsWithTimeout creates a new UpdateMongoDBAtlasRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateMongoDBAtlasRotatingSecretParamsWithTimeout(timeout time.Duration) *UpdateMongoDBAtlasRotatingSecretParams {
	return &UpdateMongoDBAtlasRotatingSecretParams{
		timeout: timeout,
	}
}

// NewUpdateMongoDBAtlasRotatingSecretParamsWithContext creates a new UpdateMongoDBAtlasRotatingSecretParams object
// with the ability to set a context for a request.
func NewUpdateMongoDBAtlasRotatingSecretParamsWithContext(ctx context.Context) *UpdateMongoDBAtlasRotatingSecretParams {
	return &UpdateMongoDBAtlasRotatingSecretParams{
		Context: ctx,
	}
}

// NewUpdateMongoDBAtlasRotatingSecretParamsWithHTTPClient creates a new UpdateMongoDBAtlasRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMongoDBAtlasRotatingSecretParamsWithHTTPClient(client *http.Client) *UpdateMongoDBAtlasRotatingSecretParams {
	return &UpdateMongoDBAtlasRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateMongoDBAtlasRotatingSecretParams contains all the parameters to send to the API endpoint

	for the update mongo d b atlas rotating secret operation.

	Typically these are written to a http.Request.
*/
type UpdateMongoDBAtlasRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceUpdateMongoDBAtlasRotatingSecretBody

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

// WithDefaults hydrates default values in the update mongo d b atlas rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithDefaults() *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mongo d b atlas rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithTimeout(timeout time.Duration) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithContext(ctx context.Context) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithHTTPClient(client *http.Client) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithAppName(appName string) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithBody(body *models.SecretServiceUpdateMongoDBAtlasRotatingSecretBody) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetBody(body *models.SecretServiceUpdateMongoDBAtlasRotatingSecretBody) {
	o.Body = body
}

// WithName adds the name to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithName(name string) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithOrganizationID(organizationID string) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) WithProjectID(projectID string) *UpdateMongoDBAtlasRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update mongo d b atlas rotating secret params
func (o *UpdateMongoDBAtlasRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMongoDBAtlasRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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