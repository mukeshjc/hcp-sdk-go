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

// NewWaypointListTriggersParams creates a new WaypointListTriggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListTriggersParams() *WaypointListTriggersParams {
	return &WaypointListTriggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListTriggersParamsWithTimeout creates a new WaypointListTriggersParams object
// with the ability to set a timeout on a request.
func NewWaypointListTriggersParamsWithTimeout(timeout time.Duration) *WaypointListTriggersParams {
	return &WaypointListTriggersParams{
		timeout: timeout,
	}
}

// NewWaypointListTriggersParamsWithContext creates a new WaypointListTriggersParams object
// with the ability to set a context for a request.
func NewWaypointListTriggersParamsWithContext(ctx context.Context) *WaypointListTriggersParams {
	return &WaypointListTriggersParams{
		Context: ctx,
	}
}

// NewWaypointListTriggersParamsWithHTTPClient creates a new WaypointListTriggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListTriggersParamsWithHTTPClient(client *http.Client) *WaypointListTriggersParams {
	return &WaypointListTriggersParams{
		HTTPClient: client,
	}
}

/*
WaypointListTriggersParams contains all the parameters to send to the API endpoint

	for the waypoint list triggers operation.

	Typically these are written to a http.Request.
*/
type WaypointListTriggersParams struct {

	// ApplicationApplication.
	ApplicationApplication *string

	// ApplicationProject.
	ApplicationProject *string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// ProjectProject.
	ProjectProject *string

	/* Tags.

	   Will filter triggers by the requested labels if set.
	*/
	Tags []string

	// WorkspaceWorkspace.
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListTriggersParams) WithDefaults() *WaypointListTriggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListTriggersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithTimeout(timeout time.Duration) *WaypointListTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithContext(ctx context.Context) *WaypointListTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithHTTPClient(client *http.Client) *WaypointListTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithApplicationApplication(applicationApplication *string) *WaypointListTriggersParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithApplicationProject(applicationProject *string) *WaypointListTriggersParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithNamespaceID adds the namespaceID to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithNamespaceID(namespaceID string) *WaypointListTriggersParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectProject adds the projectProject to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithProjectProject(projectProject *string) *WaypointListTriggersParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithTags adds the tags to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithTags(tags []string) *WaypointListTriggersParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list triggers params
func (o *WaypointListTriggersParams) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointListTriggersParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list triggers params
func (o *WaypointListTriggersParams) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
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

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
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

// bindParamWaypointListTriggers binds the parameter tags
func (o *WaypointListTriggersParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: "multi"
	tagsIS := swag.JoinByFormat(tagsIC, "multi")

	return tagsIS
}
