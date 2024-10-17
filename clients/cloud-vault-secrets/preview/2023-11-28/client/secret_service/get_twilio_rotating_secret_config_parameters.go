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

// NewGetTwilioRotatingSecretConfigParams creates a new GetTwilioRotatingSecretConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTwilioRotatingSecretConfigParams() *GetTwilioRotatingSecretConfigParams {
	return &GetTwilioRotatingSecretConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTwilioRotatingSecretConfigParamsWithTimeout creates a new GetTwilioRotatingSecretConfigParams object
// with the ability to set a timeout on a request.
func NewGetTwilioRotatingSecretConfigParamsWithTimeout(timeout time.Duration) *GetTwilioRotatingSecretConfigParams {
	return &GetTwilioRotatingSecretConfigParams{
		timeout: timeout,
	}
}

// NewGetTwilioRotatingSecretConfigParamsWithContext creates a new GetTwilioRotatingSecretConfigParams object
// with the ability to set a context for a request.
func NewGetTwilioRotatingSecretConfigParamsWithContext(ctx context.Context) *GetTwilioRotatingSecretConfigParams {
	return &GetTwilioRotatingSecretConfigParams{
		Context: ctx,
	}
}

// NewGetTwilioRotatingSecretConfigParamsWithHTTPClient creates a new GetTwilioRotatingSecretConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTwilioRotatingSecretConfigParamsWithHTTPClient(client *http.Client) *GetTwilioRotatingSecretConfigParams {
	return &GetTwilioRotatingSecretConfigParams{
		HTTPClient: client,
	}
}

/*
GetTwilioRotatingSecretConfigParams contains all the parameters to send to the API endpoint

	for the get twilio rotating secret config operation.

	Typically these are written to a http.Request.
*/
type GetTwilioRotatingSecretConfigParams struct {

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

// WithDefaults hydrates default values in the get twilio rotating secret config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTwilioRotatingSecretConfigParams) WithDefaults() *GetTwilioRotatingSecretConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get twilio rotating secret config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTwilioRotatingSecretConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithTimeout(timeout time.Duration) *GetTwilioRotatingSecretConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithContext(ctx context.Context) *GetTwilioRotatingSecretConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithHTTPClient(client *http.Client) *GetTwilioRotatingSecretConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithAppName(appName string) *GetTwilioRotatingSecretConfigParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithName adds the name to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithName(name string) *GetTwilioRotatingSecretConfigParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithOrganizationID(organizationID string) *GetTwilioRotatingSecretConfigParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) WithProjectID(projectID string) *GetTwilioRotatingSecretConfigParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get twilio rotating secret config params
func (o *GetTwilioRotatingSecretConfigParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTwilioRotatingSecretConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
