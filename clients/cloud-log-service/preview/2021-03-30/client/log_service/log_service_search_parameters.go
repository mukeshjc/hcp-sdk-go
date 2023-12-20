// Code generated by go-swagger; DO NOT EDIT.

package log_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// NewLogServiceSearchParams creates a new LogServiceSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceSearchParams() *LogServiceSearchParams {
	return &LogServiceSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceSearchParamsWithTimeout creates a new LogServiceSearchParams object
// with the ability to set a timeout on a request.
func NewLogServiceSearchParamsWithTimeout(timeout time.Duration) *LogServiceSearchParams {
	return &LogServiceSearchParams{
		timeout: timeout,
	}
}

// NewLogServiceSearchParamsWithContext creates a new LogServiceSearchParams object
// with the ability to set a context for a request.
func NewLogServiceSearchParamsWithContext(ctx context.Context) *LogServiceSearchParams {
	return &LogServiceSearchParams{
		Context: ctx,
	}
}

// NewLogServiceSearchParamsWithHTTPClient creates a new LogServiceSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceSearchParamsWithHTTPClient(client *http.Client) *LogServiceSearchParams {
	return &LogServiceSearchParams{
		HTTPClient: client,
	}
}

/*
LogServiceSearchParams contains all the parameters to send to the API endpoint

	for the log service search operation.

	Typically these are written to a http.Request.
*/
type LogServiceSearchParams struct {

	// Body.
	Body *models.LogService20210330SearchRequest

	/* OrganizationID.

	   organization_id is the organization id where the query would be ran.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log service search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceSearchParams) WithDefaults() *LogServiceSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service search params
func (o *LogServiceSearchParams) WithTimeout(timeout time.Duration) *LogServiceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service search params
func (o *LogServiceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service search params
func (o *LogServiceSearchParams) WithContext(ctx context.Context) *LogServiceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service search params
func (o *LogServiceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service search params
func (o *LogServiceSearchParams) WithHTTPClient(client *http.Client) *LogServiceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service search params
func (o *LogServiceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the log service search params
func (o *LogServiceSearchParams) WithBody(body *models.LogService20210330SearchRequest) *LogServiceSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the log service search params
func (o *LogServiceSearchParams) SetBody(body *models.LogService20210330SearchRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the log service search params
func (o *LogServiceSearchParams) WithOrganizationID(organizationID string) *LogServiceSearchParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the log service search params
func (o *LogServiceSearchParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
