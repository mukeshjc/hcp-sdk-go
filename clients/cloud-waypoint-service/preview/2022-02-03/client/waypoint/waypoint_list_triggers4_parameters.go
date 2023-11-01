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

// NewWaypointListTriggers4Params creates a new WaypointListTriggers4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListTriggers4Params() *WaypointListTriggers4Params {
	return &WaypointListTriggers4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListTriggers4ParamsWithTimeout creates a new WaypointListTriggers4Params object
// with the ability to set a timeout on a request.
func NewWaypointListTriggers4ParamsWithTimeout(timeout time.Duration) *WaypointListTriggers4Params {
	return &WaypointListTriggers4Params{
		timeout: timeout,
	}
}

// NewWaypointListTriggers4ParamsWithContext creates a new WaypointListTriggers4Params object
// with the ability to set a context for a request.
func NewWaypointListTriggers4ParamsWithContext(ctx context.Context) *WaypointListTriggers4Params {
	return &WaypointListTriggers4Params{
		Context: ctx,
	}
}

// NewWaypointListTriggers4ParamsWithHTTPClient creates a new WaypointListTriggers4Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListTriggers4ParamsWithHTTPClient(client *http.Client) *WaypointListTriggers4Params {
	return &WaypointListTriggers4Params{
		HTTPClient: client,
	}
}

/*
WaypointListTriggers4Params contains all the parameters to send to the API endpoint

	for the waypoint list triggers4 operation.

	Typically these are written to a http.Request.
*/
type WaypointListTriggers4Params struct {

	// ApplicationApplication.
	ApplicationApplication string

	// ApplicationProject.
	ApplicationProject string

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
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list triggers4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListTriggers4Params) WithDefaults() *WaypointListTriggers4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list triggers4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListTriggers4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithTimeout(timeout time.Duration) *WaypointListTriggers4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithContext(ctx context.Context) *WaypointListTriggers4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithHTTPClient(client *http.Client) *WaypointListTriggers4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithApplicationApplication(applicationApplication string) *WaypointListTriggers4Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithApplicationProject(applicationProject string) *WaypointListTriggers4Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithNamespaceID adds the namespaceID to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithNamespaceID(namespaceID string) *WaypointListTriggers4Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectProject adds the projectProject to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithProjectProject(projectProject *string) *WaypointListTriggers4Params {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithTags adds the tags to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithTags(tags []string) *WaypointListTriggers4Params {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetTags(tags []string) {
	o.Tags = tags
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointListTriggers4Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list triggers4 params
func (o *WaypointListTriggers4Params) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListTriggers4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.application
	if err := r.SetPathParam("application.application", o.ApplicationApplication); err != nil {
		return err
	}

	// path param application.project
	if err := r.SetPathParam("application.project", o.ApplicationProject); err != nil {
		return err
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

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWaypointListTriggers4 binds the parameter tags
func (o *WaypointListTriggers4Params) bindParamTags(formats strfmt.Registry) []string {
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
