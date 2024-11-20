// Code generated by go-swagger; DO NOT EDIT.

package integration_connection_service

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

// NewUpdateIntegrationConnectionParams creates a new UpdateIntegrationConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIntegrationConnectionParams() *UpdateIntegrationConnectionParams {
	return &UpdateIntegrationConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntegrationConnectionParamsWithTimeout creates a new UpdateIntegrationConnectionParams object
// with the ability to set a timeout on a request.
func NewUpdateIntegrationConnectionParamsWithTimeout(timeout time.Duration) *UpdateIntegrationConnectionParams {
	return &UpdateIntegrationConnectionParams{
		timeout: timeout,
	}
}

// NewUpdateIntegrationConnectionParamsWithContext creates a new UpdateIntegrationConnectionParams object
// with the ability to set a context for a request.
func NewUpdateIntegrationConnectionParamsWithContext(ctx context.Context) *UpdateIntegrationConnectionParams {
	return &UpdateIntegrationConnectionParams{
		Context: ctx,
	}
}

// NewUpdateIntegrationConnectionParamsWithHTTPClient creates a new UpdateIntegrationConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIntegrationConnectionParamsWithHTTPClient(client *http.Client) *UpdateIntegrationConnectionParams {
	return &UpdateIntegrationConnectionParams{
		HTTPClient: client,
	}
}

/*
UpdateIntegrationConnectionParams contains all the parameters to send to the API endpoint

	for the update integration connection operation.

	Typically these are written to a http.Request.
*/
type UpdateIntegrationConnectionParams struct {

	// Body.
	Body UpdateIntegrationConnectionBody

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update integration connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationConnectionParams) WithDefaults() *UpdateIntegrationConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update integration connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update integration connection params
func (o *UpdateIntegrationConnectionParams) WithTimeout(timeout time.Duration) *UpdateIntegrationConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update integration connection params
func (o *UpdateIntegrationConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update integration connection params
func (o *UpdateIntegrationConnectionParams) WithContext(ctx context.Context) *UpdateIntegrationConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update integration connection params
func (o *UpdateIntegrationConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update integration connection params
func (o *UpdateIntegrationConnectionParams) WithHTTPClient(client *http.Client) *UpdateIntegrationConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update integration connection params
func (o *UpdateIntegrationConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update integration connection params
func (o *UpdateIntegrationConnectionParams) WithBody(body UpdateIntegrationConnectionBody) *UpdateIntegrationConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update integration connection params
func (o *UpdateIntegrationConnectionParams) SetBody(body UpdateIntegrationConnectionBody) {
	o.Body = body
}

// WithLocationProjectID adds the locationProjectID to the update integration connection params
func (o *UpdateIntegrationConnectionParams) WithLocationProjectID(locationProjectID string) *UpdateIntegrationConnectionParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update integration connection params
func (o *UpdateIntegrationConnectionParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntegrationConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}