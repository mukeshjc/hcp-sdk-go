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

// NewWaypointGetJobStreamParams creates a new WaypointGetJobStreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetJobStreamParams() *WaypointGetJobStreamParams {
	return &WaypointGetJobStreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetJobStreamParamsWithTimeout creates a new WaypointGetJobStreamParams object
// with the ability to set a timeout on a request.
func NewWaypointGetJobStreamParamsWithTimeout(timeout time.Duration) *WaypointGetJobStreamParams {
	return &WaypointGetJobStreamParams{
		timeout: timeout,
	}
}

// NewWaypointGetJobStreamParamsWithContext creates a new WaypointGetJobStreamParams object
// with the ability to set a context for a request.
func NewWaypointGetJobStreamParamsWithContext(ctx context.Context) *WaypointGetJobStreamParams {
	return &WaypointGetJobStreamParams{
		Context: ctx,
	}
}

// NewWaypointGetJobStreamParamsWithHTTPClient creates a new WaypointGetJobStreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetJobStreamParamsWithHTTPClient(client *http.Client) *WaypointGetJobStreamParams {
	return &WaypointGetJobStreamParams{
		HTTPClient: client,
	}
}

/*
WaypointGetJobStreamParams contains all the parameters to send to the API endpoint

	for the waypoint get job stream operation.

	Typically these are written to a http.Request.
*/
type WaypointGetJobStreamParams struct {

	// JobID.
	JobID string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get job stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetJobStreamParams) WithDefaults() *WaypointGetJobStreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get job stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetJobStreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) WithTimeout(timeout time.Duration) *WaypointGetJobStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) WithContext(ctx context.Context) *WaypointGetJobStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) WithHTTPClient(client *http.Client) *WaypointGetJobStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) WithJobID(jobID string) *WaypointGetJobStreamParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithNamespaceID adds the namespaceID to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) WithNamespaceID(namespaceID string) *WaypointGetJobStreamParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get job stream params
func (o *WaypointGetJobStreamParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetJobStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
