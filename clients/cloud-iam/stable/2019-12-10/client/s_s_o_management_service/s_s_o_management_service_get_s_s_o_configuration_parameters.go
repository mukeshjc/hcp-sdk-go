// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

// NewSSOManagementServiceGetSSOConfigurationParams creates a new SSOManagementServiceGetSSOConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSOManagementServiceGetSSOConfigurationParams() *SSOManagementServiceGetSSOConfigurationParams {
	return &SSOManagementServiceGetSSOConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSOManagementServiceGetSSOConfigurationParamsWithTimeout creates a new SSOManagementServiceGetSSOConfigurationParams object
// with the ability to set a timeout on a request.
func NewSSOManagementServiceGetSSOConfigurationParamsWithTimeout(timeout time.Duration) *SSOManagementServiceGetSSOConfigurationParams {
	return &SSOManagementServiceGetSSOConfigurationParams{
		timeout: timeout,
	}
}

// NewSSOManagementServiceGetSSOConfigurationParamsWithContext creates a new SSOManagementServiceGetSSOConfigurationParams object
// with the ability to set a context for a request.
func NewSSOManagementServiceGetSSOConfigurationParamsWithContext(ctx context.Context) *SSOManagementServiceGetSSOConfigurationParams {
	return &SSOManagementServiceGetSSOConfigurationParams{
		Context: ctx,
	}
}

// NewSSOManagementServiceGetSSOConfigurationParamsWithHTTPClient creates a new SSOManagementServiceGetSSOConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSOManagementServiceGetSSOConfigurationParamsWithHTTPClient(client *http.Client) *SSOManagementServiceGetSSOConfigurationParams {
	return &SSOManagementServiceGetSSOConfigurationParams{
		HTTPClient: client,
	}
}

/*
SSOManagementServiceGetSSOConfigurationParams contains all the parameters to send to the API endpoint

	for the s s o management service get s s o configuration operation.

	Typically these are written to a http.Request.
*/
type SSOManagementServiceGetSSOConfigurationParams struct {

	/* OrganizationID.

	   OrganizationId is the identifier of the organization.
	*/
	OrganizationID string

	/* Type.

	   Type is the type of Single Sign-On.
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s s o management service get s s o configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceGetSSOConfigurationParams) WithDefaults() *SSOManagementServiceGetSSOConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s s o management service get s s o configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceGetSSOConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) WithTimeout(timeout time.Duration) *SSOManagementServiceGetSSOConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) WithContext(ctx context.Context) *SSOManagementServiceGetSSOConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) WithHTTPClient(client *http.Client) *SSOManagementServiceGetSSOConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) WithOrganizationID(organizationID string) *SSOManagementServiceGetSSOConfigurationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithType adds the typeVar to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) WithType(typeVar string) *SSOManagementServiceGetSSOConfigurationParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the s s o management service get s s o configuration params
func (o *SSOManagementServiceGetSSOConfigurationParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SSOManagementServiceGetSSOConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}