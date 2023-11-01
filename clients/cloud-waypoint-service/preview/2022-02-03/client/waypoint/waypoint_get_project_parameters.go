// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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

// NewWaypointGetProjectParams creates a new WaypointGetProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetProjectParams() *WaypointGetProjectParams {
	return &WaypointGetProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetProjectParamsWithTimeout creates a new WaypointGetProjectParams object
// with the ability to set a timeout on a request.
func NewWaypointGetProjectParamsWithTimeout(timeout time.Duration) *WaypointGetProjectParams {
	return &WaypointGetProjectParams{
		timeout: timeout,
	}
}

// NewWaypointGetProjectParamsWithContext creates a new WaypointGetProjectParams object
// with the ability to set a context for a request.
func NewWaypointGetProjectParamsWithContext(ctx context.Context) *WaypointGetProjectParams {
	return &WaypointGetProjectParams{
		Context: ctx,
	}
}

// NewWaypointGetProjectParamsWithHTTPClient creates a new WaypointGetProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetProjectParamsWithHTTPClient(client *http.Client) *WaypointGetProjectParams {
	return &WaypointGetProjectParams{
		HTTPClient: client,
	}
}

/*
WaypointGetProjectParams contains all the parameters to send to the API endpoint

	for the waypoint get project operation.

	Typically these are written to a http.Request.
*/
type WaypointGetProjectParams struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// ProjectProject.
	ProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetProjectParams) WithDefaults() *WaypointGetProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get project params
func (o *WaypointGetProjectParams) WithTimeout(timeout time.Duration) *WaypointGetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get project params
func (o *WaypointGetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get project params
func (o *WaypointGetProjectParams) WithContext(ctx context.Context) *WaypointGetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get project params
func (o *WaypointGetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get project params
func (o *WaypointGetProjectParams) WithHTTPClient(client *http.Client) *WaypointGetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get project params
func (o *WaypointGetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get project params
func (o *WaypointGetProjectParams) WithNamespaceID(namespaceID string) *WaypointGetProjectParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get project params
func (o *WaypointGetProjectParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectProject adds the projectProject to the waypoint get project params
func (o *WaypointGetProjectParams) WithProjectProject(projectProject string) *WaypointGetProjectParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint get project params
func (o *WaypointGetProjectParams) SetProjectProject(projectProject string) {
	o.ProjectProject = projectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param project.project
	if err := r.SetPathParam("project.project", o.ProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
