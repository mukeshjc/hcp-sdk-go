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

// NewWaypointListPipelineRunsParams creates a new WaypointListPipelineRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListPipelineRunsParams() *WaypointListPipelineRunsParams {
	return &WaypointListPipelineRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListPipelineRunsParamsWithTimeout creates a new WaypointListPipelineRunsParams object
// with the ability to set a timeout on a request.
func NewWaypointListPipelineRunsParamsWithTimeout(timeout time.Duration) *WaypointListPipelineRunsParams {
	return &WaypointListPipelineRunsParams{
		timeout: timeout,
	}
}

// NewWaypointListPipelineRunsParamsWithContext creates a new WaypointListPipelineRunsParams object
// with the ability to set a context for a request.
func NewWaypointListPipelineRunsParamsWithContext(ctx context.Context) *WaypointListPipelineRunsParams {
	return &WaypointListPipelineRunsParams{
		Context: ctx,
	}
}

// NewWaypointListPipelineRunsParamsWithHTTPClient creates a new WaypointListPipelineRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListPipelineRunsParamsWithHTTPClient(client *http.Client) *WaypointListPipelineRunsParams {
	return &WaypointListPipelineRunsParams{
		HTTPClient: client,
	}
}

/*
WaypointListPipelineRunsParams contains all the parameters to send to the API endpoint

	for the waypoint list pipeline runs operation.

	Typically these are written to a http.Request.
*/
type WaypointListPipelineRunsParams struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* PipelineID.

	   Reference a single pipeline by ID.
	*/
	PipelineID *string

	/* PipelineOwnerPipelineName.

	   the name of the defined pipeline config
	*/
	PipelineOwnerPipelineName string

	// PipelineOwnerProjectProject.
	PipelineOwnerProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list pipeline runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListPipelineRunsParams) WithDefaults() *WaypointListPipelineRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list pipeline runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListPipelineRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithTimeout(timeout time.Duration) *WaypointListPipelineRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithContext(ctx context.Context) *WaypointListPipelineRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithHTTPClient(client *http.Client) *WaypointListPipelineRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithNamespaceID(namespaceID string) *WaypointListPipelineRunsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPipelineID adds the pipelineID to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithPipelineID(pipelineID *string) *WaypointListPipelineRunsParams {
	o.SetPipelineID(pipelineID)
	return o
}

// SetPipelineID adds the pipelineId to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetPipelineID(pipelineID *string) {
	o.PipelineID = pipelineID
}

// WithPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithPipelineOwnerPipelineName(pipelineOwnerPipelineName string) *WaypointListPipelineRunsParams {
	o.SetPipelineOwnerPipelineName(pipelineOwnerPipelineName)
	return o
}

// SetPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetPipelineOwnerPipelineName(pipelineOwnerPipelineName string) {
	o.PipelineOwnerPipelineName = pipelineOwnerPipelineName
}

// WithPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) WithPipelineOwnerProjectProject(pipelineOwnerProjectProject string) *WaypointListPipelineRunsParams {
	o.SetPipelineOwnerProjectProject(pipelineOwnerProjectProject)
	return o
}

// SetPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint list pipeline runs params
func (o *WaypointListPipelineRunsParams) SetPipelineOwnerProjectProject(pipelineOwnerProjectProject string) {
	o.PipelineOwnerProjectProject = pipelineOwnerProjectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListPipelineRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.PipelineID != nil {

		// query param pipeline.id
		var qrPipelineID string

		if o.PipelineID != nil {
			qrPipelineID = *o.PipelineID
		}
		qPipelineID := qrPipelineID
		if qPipelineID != "" {

			if err := r.SetQueryParam("pipeline.id", qPipelineID); err != nil {
				return err
			}
		}
	}

	// path param pipeline.owner.pipeline_name
	if err := r.SetPathParam("pipeline.owner.pipeline_name", o.PipelineOwnerPipelineName); err != nil {
		return err
	}

	// path param pipeline.owner.project.project
	if err := r.SetPathParam("pipeline.owner.project.project", o.PipelineOwnerProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
