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

// NewWaypointGetJobParams creates a new WaypointGetJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetJobParams() *WaypointGetJobParams {
	return &WaypointGetJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetJobParamsWithTimeout creates a new WaypointGetJobParams object
// with the ability to set a timeout on a request.
func NewWaypointGetJobParamsWithTimeout(timeout time.Duration) *WaypointGetJobParams {
	return &WaypointGetJobParams{
		timeout: timeout,
	}
}

// NewWaypointGetJobParamsWithContext creates a new WaypointGetJobParams object
// with the ability to set a context for a request.
func NewWaypointGetJobParamsWithContext(ctx context.Context) *WaypointGetJobParams {
	return &WaypointGetJobParams{
		Context: ctx,
	}
}

// NewWaypointGetJobParamsWithHTTPClient creates a new WaypointGetJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetJobParamsWithHTTPClient(client *http.Client) *WaypointGetJobParams {
	return &WaypointGetJobParams{
		HTTPClient: client,
	}
}

/*
WaypointGetJobParams contains all the parameters to send to the API endpoint

	for the waypoint get job operation.

	Typically these are written to a http.Request.
*/
type WaypointGetJobParams struct {

	/* JobID.

	   ID of the job to request.
	*/
	JobID string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetJobParams) WithDefaults() *WaypointGetJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get job params
func (o *WaypointGetJobParams) WithTimeout(timeout time.Duration) *WaypointGetJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get job params
func (o *WaypointGetJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get job params
func (o *WaypointGetJobParams) WithContext(ctx context.Context) *WaypointGetJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get job params
func (o *WaypointGetJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get job params
func (o *WaypointGetJobParams) WithHTTPClient(client *http.Client) *WaypointGetJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get job params
func (o *WaypointGetJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the waypoint get job params
func (o *WaypointGetJobParams) WithJobID(jobID string) *WaypointGetJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the waypoint get job params
func (o *WaypointGetJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithNamespaceID adds the namespaceID to the waypoint get job params
func (o *WaypointGetJobParams) WithNamespaceID(namespaceID string) *WaypointGetJobParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get job params
func (o *WaypointGetJobParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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