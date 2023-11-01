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

// NewSSOManagementServiceVerifySSODomainOwnershipParams creates a new SSOManagementServiceVerifySSODomainOwnershipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSOManagementServiceVerifySSODomainOwnershipParams() *SSOManagementServiceVerifySSODomainOwnershipParams {
	return &SSOManagementServiceVerifySSODomainOwnershipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSOManagementServiceVerifySSODomainOwnershipParamsWithTimeout creates a new SSOManagementServiceVerifySSODomainOwnershipParams object
// with the ability to set a timeout on a request.
func NewSSOManagementServiceVerifySSODomainOwnershipParamsWithTimeout(timeout time.Duration) *SSOManagementServiceVerifySSODomainOwnershipParams {
	return &SSOManagementServiceVerifySSODomainOwnershipParams{
		timeout: timeout,
	}
}

// NewSSOManagementServiceVerifySSODomainOwnershipParamsWithContext creates a new SSOManagementServiceVerifySSODomainOwnershipParams object
// with the ability to set a context for a request.
func NewSSOManagementServiceVerifySSODomainOwnershipParamsWithContext(ctx context.Context) *SSOManagementServiceVerifySSODomainOwnershipParams {
	return &SSOManagementServiceVerifySSODomainOwnershipParams{
		Context: ctx,
	}
}

// NewSSOManagementServiceVerifySSODomainOwnershipParamsWithHTTPClient creates a new SSOManagementServiceVerifySSODomainOwnershipParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSOManagementServiceVerifySSODomainOwnershipParamsWithHTTPClient(client *http.Client) *SSOManagementServiceVerifySSODomainOwnershipParams {
	return &SSOManagementServiceVerifySSODomainOwnershipParams{
		HTTPClient: client,
	}
}

/*
SSOManagementServiceVerifySSODomainOwnershipParams contains all the parameters to send to the API endpoint

	for the s s o management service verify s s o domain ownership operation.

	Typically these are written to a http.Request.
*/
type SSOManagementServiceVerifySSODomainOwnershipParams struct {

	// Body.
	Body SSOManagementServiceVerifySSODomainOwnershipBody

	/* OrganizationID.

	     organization_id is the ID of the organization for which domain ownership
	will be verified.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s s o management service verify s s o domain ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithDefaults() *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s s o management service verify s s o domain ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithTimeout(timeout time.Duration) *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithContext(ctx context.Context) *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithHTTPClient(client *http.Client) *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithBody(body SSOManagementServiceVerifySSODomainOwnershipBody) *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetBody(body SSOManagementServiceVerifySSODomainOwnershipBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WithOrganizationID(organizationID string) *SSOManagementServiceVerifySSODomainOwnershipParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the s s o management service verify s s o domain ownership params
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SSOManagementServiceVerifySSODomainOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
