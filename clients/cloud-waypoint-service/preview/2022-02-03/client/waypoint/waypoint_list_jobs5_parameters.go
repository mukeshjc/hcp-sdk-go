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
	"github.com/go-openapi/swag"
)

// NewWaypointListJobs5Params creates a new WaypointListJobs5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListJobs5Params() *WaypointListJobs5Params {
	return &WaypointListJobs5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListJobs5ParamsWithTimeout creates a new WaypointListJobs5Params object
// with the ability to set a timeout on a request.
func NewWaypointListJobs5ParamsWithTimeout(timeout time.Duration) *WaypointListJobs5Params {
	return &WaypointListJobs5Params{
		timeout: timeout,
	}
}

// NewWaypointListJobs5ParamsWithContext creates a new WaypointListJobs5Params object
// with the ability to set a context for a request.
func NewWaypointListJobs5ParamsWithContext(ctx context.Context) *WaypointListJobs5Params {
	return &WaypointListJobs5Params{
		Context: ctx,
	}
}

// NewWaypointListJobs5ParamsWithHTTPClient creates a new WaypointListJobs5Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListJobs5ParamsWithHTTPClient(client *http.Client) *WaypointListJobs5Params {
	return &WaypointListJobs5Params{
		HTTPClient: client,
	}
}

/*
WaypointListJobs5Params contains all the parameters to send to the API endpoint

	for the waypoint list jobs5 operation.

	Typically these are written to a http.Request.
*/
type WaypointListJobs5Params struct {

	// ApplicationApplication.
	ApplicationApplication *string

	// ApplicationProject.
	ApplicationProject *string

	// JobState.
	JobState []string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* PipelinePipelineID.

	   ID of the current pipeline (the one containing the current step).
	*/
	PipelinePipelineID *string

	/* PipelinePipelineName.

	   Name of the current pipeline.
	*/
	PipelinePipelineName *string

	/* PipelineRootPipelineID.

	   Name of the root pipeline (the one invoked in `waypoint pipeline run <pipeline>`).
	*/
	PipelineRootPipelineID *string

	/* PipelineRootPipelineName.

	   Name of the root pipeline.
	*/
	PipelineRootPipelineName *string

	/* PipelineRunSequence.

	   Run sequence for the root pipeline.

	   Format: uint64
	*/
	PipelineRunSequence *string

	/* PipelineStep.

	   Step name within the current pipeline.
	*/
	PipelineStep *string

	// ProjectProject.
	ProjectProject *string

	// TargetRunnerIDID.
	TargetRunnerIDID string

	// WorkspaceWorkspace.
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list jobs5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListJobs5Params) WithDefaults() *WaypointListJobs5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list jobs5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListJobs5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithTimeout(timeout time.Duration) *WaypointListJobs5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithContext(ctx context.Context) *WaypointListJobs5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithHTTPClient(client *http.Client) *WaypointListJobs5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithApplicationApplication(applicationApplication *string) *WaypointListJobs5Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithApplicationProject(applicationProject *string) *WaypointListJobs5Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithJobState adds the jobState to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithJobState(jobState []string) *WaypointListJobs5Params {
	o.SetJobState(jobState)
	return o
}

// SetJobState adds the jobState to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetJobState(jobState []string) {
	o.JobState = jobState
}

// WithNamespaceID adds the namespaceID to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithNamespaceID(namespaceID string) *WaypointListJobs5Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointListJobs5Params {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPaginationPageSize(paginationPageSize *int64) *WaypointListJobs5Params {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointListJobs5Params {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPipelinePipelineID adds the pipelinePipelineID to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelinePipelineID(pipelinePipelineID *string) *WaypointListJobs5Params {
	o.SetPipelinePipelineID(pipelinePipelineID)
	return o
}

// SetPipelinePipelineID adds the pipelinePipelineId to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelinePipelineID(pipelinePipelineID *string) {
	o.PipelinePipelineID = pipelinePipelineID
}

// WithPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelinePipelineName(pipelinePipelineName *string) *WaypointListJobs5Params {
	o.SetPipelinePipelineName(pipelinePipelineName)
	return o
}

// SetPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelinePipelineName(pipelinePipelineName *string) {
	o.PipelinePipelineName = pipelinePipelineName
}

// WithPipelineRootPipelineID adds the pipelineRootPipelineID to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelineRootPipelineID(pipelineRootPipelineID *string) *WaypointListJobs5Params {
	o.SetPipelineRootPipelineID(pipelineRootPipelineID)
	return o
}

// SetPipelineRootPipelineID adds the pipelineRootPipelineId to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelineRootPipelineID(pipelineRootPipelineID *string) {
	o.PipelineRootPipelineID = pipelineRootPipelineID
}

// WithPipelineRootPipelineName adds the pipelineRootPipelineName to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelineRootPipelineName(pipelineRootPipelineName *string) *WaypointListJobs5Params {
	o.SetPipelineRootPipelineName(pipelineRootPipelineName)
	return o
}

// SetPipelineRootPipelineName adds the pipelineRootPipelineName to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelineRootPipelineName(pipelineRootPipelineName *string) {
	o.PipelineRootPipelineName = pipelineRootPipelineName
}

// WithPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelineRunSequence(pipelineRunSequence *string) *WaypointListJobs5Params {
	o.SetPipelineRunSequence(pipelineRunSequence)
	return o
}

// SetPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelineRunSequence(pipelineRunSequence *string) {
	o.PipelineRunSequence = pipelineRunSequence
}

// WithPipelineStep adds the pipelineStep to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithPipelineStep(pipelineStep *string) *WaypointListJobs5Params {
	o.SetPipelineStep(pipelineStep)
	return o
}

// SetPipelineStep adds the pipelineStep to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetPipelineStep(pipelineStep *string) {
	o.PipelineStep = pipelineStep
}

// WithProjectProject adds the projectProject to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithProjectProject(projectProject *string) *WaypointListJobs5Params {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithTargetRunnerIDID adds the targetRunnerIDID to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithTargetRunnerIDID(targetRunnerIDID string) *WaypointListJobs5Params {
	o.SetTargetRunnerIDID(targetRunnerIDID)
	return o
}

// SetTargetRunnerIDID adds the targetRunnerIdId to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetTargetRunnerIDID(targetRunnerIDID string) {
	o.TargetRunnerIDID = targetRunnerIDID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointListJobs5Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs5 params
func (o *WaypointListJobs5Params) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListJobs5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationApplication != nil {

		// query param application.application
		var qrApplicationApplication string

		if o.ApplicationApplication != nil {
			qrApplicationApplication = *o.ApplicationApplication
		}
		qApplicationApplication := qrApplicationApplication
		if qApplicationApplication != "" {

			if err := r.SetQueryParam("application.application", qApplicationApplication); err != nil {
				return err
			}
		}
	}

	if o.ApplicationProject != nil {

		// query param application.project
		var qrApplicationProject string

		if o.ApplicationProject != nil {
			qrApplicationProject = *o.ApplicationProject
		}
		qApplicationProject := qrApplicationProject
		if qApplicationProject != "" {

			if err := r.SetQueryParam("application.project", qApplicationProject); err != nil {
				return err
			}
		}
	}

	if o.JobState != nil {

		// binding items for jobState
		joinedJobState := o.bindParamJobState(reg)

		// query array param jobState
		if err := r.SetQueryParam("jobState", joinedJobState...); err != nil {
			return err
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.PipelinePipelineID != nil {

		// query param pipeline.pipeline_id
		var qrPipelinePipelineID string

		if o.PipelinePipelineID != nil {
			qrPipelinePipelineID = *o.PipelinePipelineID
		}
		qPipelinePipelineID := qrPipelinePipelineID
		if qPipelinePipelineID != "" {

			if err := r.SetQueryParam("pipeline.pipeline_id", qPipelinePipelineID); err != nil {
				return err
			}
		}
	}

	if o.PipelinePipelineName != nil {

		// query param pipeline.pipeline_name
		var qrPipelinePipelineName string

		if o.PipelinePipelineName != nil {
			qrPipelinePipelineName = *o.PipelinePipelineName
		}
		qPipelinePipelineName := qrPipelinePipelineName
		if qPipelinePipelineName != "" {

			if err := r.SetQueryParam("pipeline.pipeline_name", qPipelinePipelineName); err != nil {
				return err
			}
		}
	}

	if o.PipelineRootPipelineID != nil {

		// query param pipeline.root_pipeline_id
		var qrPipelineRootPipelineID string

		if o.PipelineRootPipelineID != nil {
			qrPipelineRootPipelineID = *o.PipelineRootPipelineID
		}
		qPipelineRootPipelineID := qrPipelineRootPipelineID
		if qPipelineRootPipelineID != "" {

			if err := r.SetQueryParam("pipeline.root_pipeline_id", qPipelineRootPipelineID); err != nil {
				return err
			}
		}
	}

	if o.PipelineRootPipelineName != nil {

		// query param pipeline.root_pipeline_name
		var qrPipelineRootPipelineName string

		if o.PipelineRootPipelineName != nil {
			qrPipelineRootPipelineName = *o.PipelineRootPipelineName
		}
		qPipelineRootPipelineName := qrPipelineRootPipelineName
		if qPipelineRootPipelineName != "" {

			if err := r.SetQueryParam("pipeline.root_pipeline_name", qPipelineRootPipelineName); err != nil {
				return err
			}
		}
	}

	if o.PipelineRunSequence != nil {

		// query param pipeline.run_sequence
		var qrPipelineRunSequence string

		if o.PipelineRunSequence != nil {
			qrPipelineRunSequence = *o.PipelineRunSequence
		}
		qPipelineRunSequence := qrPipelineRunSequence
		if qPipelineRunSequence != "" {

			if err := r.SetQueryParam("pipeline.run_sequence", qPipelineRunSequence); err != nil {
				return err
			}
		}
	}

	if o.PipelineStep != nil {

		// query param pipeline.step
		var qrPipelineStep string

		if o.PipelineStep != nil {
			qrPipelineStep = *o.PipelineStep
		}
		qPipelineStep := qrPipelineStep
		if qPipelineStep != "" {

			if err := r.SetQueryParam("pipeline.step", qPipelineStep); err != nil {
				return err
			}
		}
	}

	if o.ProjectProject != nil {

		// query param project.project
		var qrProjectProject string

		if o.ProjectProject != nil {
			qrProjectProject = *o.ProjectProject
		}
		qProjectProject := qrProjectProject
		if qProjectProject != "" {

			if err := r.SetQueryParam("project.project", qProjectProject); err != nil {
				return err
			}
		}
	}

	// path param targetRunner.id.id
	if err := r.SetPathParam("targetRunner.id.id", o.TargetRunnerIDID); err != nil {
		return err
	}

	if o.WorkspaceWorkspace != nil {

		// query param workspace.workspace
		var qrWorkspaceWorkspace string

		if o.WorkspaceWorkspace != nil {
			qrWorkspaceWorkspace = *o.WorkspaceWorkspace
		}
		qWorkspaceWorkspace := qrWorkspaceWorkspace
		if qWorkspaceWorkspace != "" {

			if err := r.SetQueryParam("workspace.workspace", qWorkspaceWorkspace); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWaypointListJobs5 binds the parameter jobState
func (o *WaypointListJobs5Params) bindParamJobState(formats strfmt.Registry) []string {
	jobStateIR := o.JobState

	var jobStateIC []string
	for _, jobStateIIR := range jobStateIR { // explode []string

		jobStateIIV := jobStateIIR // string as string
		jobStateIC = append(jobStateIC, jobStateIIV)
	}

	// items.CollectionFormat: "multi"
	jobStateIS := swag.JoinByFormat(jobStateIC, "multi")

	return jobStateIS
}
