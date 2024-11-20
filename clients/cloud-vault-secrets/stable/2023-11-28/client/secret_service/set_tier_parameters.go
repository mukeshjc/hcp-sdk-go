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

// NewSetTierParams creates a new SetTierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetTierParams() *SetTierParams {
	return &SetTierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetTierParamsWithTimeout creates a new SetTierParams object
// with the ability to set a timeout on a request.
func NewSetTierParamsWithTimeout(timeout time.Duration) *SetTierParams {
	return &SetTierParams{
		timeout: timeout,
	}
}

// NewSetTierParamsWithContext creates a new SetTierParams object
// with the ability to set a context for a request.
func NewSetTierParamsWithContext(ctx context.Context) *SetTierParams {
	return &SetTierParams{
		Context: ctx,
	}
}

// NewSetTierParamsWithHTTPClient creates a new SetTierParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetTierParamsWithHTTPClient(client *http.Client) *SetTierParams {
	return &SetTierParams{
		HTTPClient: client,
	}
}

/*
SetTierParams contains all the parameters to send to the API endpoint

	for the set tier operation.

	Typically these are written to a http.Request.
*/
type SetTierParams struct {

	// Body.
	Body *models.SecretServiceSetTierBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTierParams) WithDefaults() *SetTierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set tier params
func (o *SetTierParams) WithTimeout(timeout time.Duration) *SetTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set tier params
func (o *SetTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set tier params
func (o *SetTierParams) WithContext(ctx context.Context) *SetTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set tier params
func (o *SetTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set tier params
func (o *SetTierParams) WithHTTPClient(client *http.Client) *SetTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set tier params
func (o *SetTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set tier params
func (o *SetTierParams) WithBody(body *models.SecretServiceSetTierBody) *SetTierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set tier params
func (o *SetTierParams) SetBody(body *models.SecretServiceSetTierBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the set tier params
func (o *SetTierParams) WithOrganizationID(organizationID string) *SetTierParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the set tier params
func (o *SetTierParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SetTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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