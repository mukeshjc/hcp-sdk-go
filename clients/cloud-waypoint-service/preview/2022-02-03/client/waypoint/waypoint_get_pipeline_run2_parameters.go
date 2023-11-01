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

// NewWaypointGetPipelineRun2Params creates a new WaypointGetPipelineRun2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetPipelineRun2Params() *WaypointGetPipelineRun2Params {
	return &WaypointGetPipelineRun2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetPipelineRun2ParamsWithTimeout creates a new WaypointGetPipelineRun2Params object
// with the ability to set a timeout on a request.
func NewWaypointGetPipelineRun2ParamsWithTimeout(timeout time.Duration) *WaypointGetPipelineRun2Params {
	return &WaypointGetPipelineRun2Params{
		timeout: timeout,
	}
}

// NewWaypointGetPipelineRun2ParamsWithContext creates a new WaypointGetPipelineRun2Params object
// with the ability to set a context for a request.
func NewWaypointGetPipelineRun2ParamsWithContext(ctx context.Context) *WaypointGetPipelineRun2Params {
	return &WaypointGetPipelineRun2Params{
		Context: ctx,
	}
}

// NewWaypointGetPipelineRun2ParamsWithHTTPClient creates a new WaypointGetPipelineRun2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetPipelineRun2ParamsWithHTTPClient(client *http.Client) *WaypointGetPipelineRun2Params {
	return &WaypointGetPipelineRun2Params{
		HTTPClient: client,
	}
}

/*
WaypointGetPipelineRun2Params contains all the parameters to send to the API endpoint

	for the waypoint get pipeline run2 operation.

	Typically these are written to a http.Request.
*/
type WaypointGetPipelineRun2Params struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* PipelineID.

	   Reference a single pipeline by ID.
	*/
	PipelineID string

	/* PipelineOwnerPipelineName.

	   the name of the defined pipeline config.
	*/
	PipelineOwnerPipelineName *string

	// PipelineOwnerProjectProject.
	PipelineOwnerProjectProject *string

	// Sequence.
	//
	// Format: uint64
	Sequence string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get pipeline run2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetPipelineRun2Params) WithDefaults() *WaypointGetPipelineRun2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get pipeline run2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetPipelineRun2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithTimeout(timeout time.Duration) *WaypointGetPipelineRun2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithContext(ctx context.Context) *WaypointGetPipelineRun2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithHTTPClient(client *http.Client) *WaypointGetPipelineRun2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithNamespaceID(namespaceID string) *WaypointGetPipelineRun2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPipelineID adds the pipelineID to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithPipelineID(pipelineID string) *WaypointGetPipelineRun2Params {
	o.SetPipelineID(pipelineID)
	return o
}

// SetPipelineID adds the pipelineId to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetPipelineID(pipelineID string) {
	o.PipelineID = pipelineID
}

// WithPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithPipelineOwnerPipelineName(pipelineOwnerPipelineName *string) *WaypointGetPipelineRun2Params {
	o.SetPipelineOwnerPipelineName(pipelineOwnerPipelineName)
	return o
}

// SetPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetPipelineOwnerPipelineName(pipelineOwnerPipelineName *string) {
	o.PipelineOwnerPipelineName = pipelineOwnerPipelineName
}

// WithPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithPipelineOwnerProjectProject(pipelineOwnerProjectProject *string) *WaypointGetPipelineRun2Params {
	o.SetPipelineOwnerProjectProject(pipelineOwnerProjectProject)
	return o
}

// SetPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetPipelineOwnerProjectProject(pipelineOwnerProjectProject *string) {
	o.PipelineOwnerProjectProject = pipelineOwnerProjectProject
}

// WithSequence adds the sequence to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) WithSequence(sequence string) *WaypointGetPipelineRun2Params {
	o.SetSequence(sequence)
	return o
}

// SetSequence adds the sequence to the waypoint get pipeline run2 params
func (o *WaypointGetPipelineRun2Params) SetSequence(sequence string) {
	o.Sequence = sequence
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetPipelineRun2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param pipeline.id
	if err := r.SetPathParam("pipeline.id", o.PipelineID); err != nil {
		return err
	}

	if o.PipelineOwnerPipelineName != nil {

		// query param pipeline.owner.pipeline_name
		var qrPipelineOwnerPipelineName string

		if o.PipelineOwnerPipelineName != nil {
			qrPipelineOwnerPipelineName = *o.PipelineOwnerPipelineName
		}
		qPipelineOwnerPipelineName := qrPipelineOwnerPipelineName
		if qPipelineOwnerPipelineName != "" {

			if err := r.SetQueryParam("pipeline.owner.pipeline_name", qPipelineOwnerPipelineName); err != nil {
				return err
			}
		}
	}

	if o.PipelineOwnerProjectProject != nil {

		// query param pipeline.owner.project.project
		var qrPipelineOwnerProjectProject string

		if o.PipelineOwnerProjectProject != nil {
			qrPipelineOwnerProjectProject = *o.PipelineOwnerProjectProject
		}
		qPipelineOwnerProjectProject := qrPipelineOwnerProjectProject
		if qPipelineOwnerProjectProject != "" {

			if err := r.SetQueryParam("pipeline.owner.project.project", qPipelineOwnerProjectProject); err != nil {
				return err
			}
		}
	}

	// path param sequence
	if err := r.SetPathParam("sequence", o.Sequence); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
