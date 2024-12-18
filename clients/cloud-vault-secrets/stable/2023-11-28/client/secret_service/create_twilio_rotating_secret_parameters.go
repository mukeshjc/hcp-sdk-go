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

// NewCreateTwilioRotatingSecretParams creates a new CreateTwilioRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTwilioRotatingSecretParams() *CreateTwilioRotatingSecretParams {
	return &CreateTwilioRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTwilioRotatingSecretParamsWithTimeout creates a new CreateTwilioRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewCreateTwilioRotatingSecretParamsWithTimeout(timeout time.Duration) *CreateTwilioRotatingSecretParams {
	return &CreateTwilioRotatingSecretParams{
		timeout: timeout,
	}
}

// NewCreateTwilioRotatingSecretParamsWithContext creates a new CreateTwilioRotatingSecretParams object
// with the ability to set a context for a request.
func NewCreateTwilioRotatingSecretParamsWithContext(ctx context.Context) *CreateTwilioRotatingSecretParams {
	return &CreateTwilioRotatingSecretParams{
		Context: ctx,
	}
}

// NewCreateTwilioRotatingSecretParamsWithHTTPClient creates a new CreateTwilioRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTwilioRotatingSecretParamsWithHTTPClient(client *http.Client) *CreateTwilioRotatingSecretParams {
	return &CreateTwilioRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
CreateTwilioRotatingSecretParams contains all the parameters to send to the API endpoint

	for the create twilio rotating secret operation.

	Typically these are written to a http.Request.
*/
type CreateTwilioRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceCreateTwilioRotatingSecretBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create twilio rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTwilioRotatingSecretParams) WithDefaults() *CreateTwilioRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create twilio rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTwilioRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithTimeout(timeout time.Duration) *CreateTwilioRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithContext(ctx context.Context) *CreateTwilioRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithHTTPClient(client *http.Client) *CreateTwilioRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithAppName(appName string) *CreateTwilioRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithBody(body *models.SecretServiceCreateTwilioRotatingSecretBody) *CreateTwilioRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetBody(body *models.SecretServiceCreateTwilioRotatingSecretBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithOrganizationID(organizationID string) *CreateTwilioRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) WithProjectID(projectID string) *CreateTwilioRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create twilio rotating secret params
func (o *CreateTwilioRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTwilioRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
