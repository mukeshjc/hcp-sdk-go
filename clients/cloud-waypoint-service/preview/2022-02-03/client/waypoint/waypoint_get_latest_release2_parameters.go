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

// NewWaypointGetLatestRelease2Params creates a new WaypointGetLatestRelease2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetLatestRelease2Params() *WaypointGetLatestRelease2Params {
	return &WaypointGetLatestRelease2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetLatestRelease2ParamsWithTimeout creates a new WaypointGetLatestRelease2Params object
// with the ability to set a timeout on a request.
func NewWaypointGetLatestRelease2ParamsWithTimeout(timeout time.Duration) *WaypointGetLatestRelease2Params {
	return &WaypointGetLatestRelease2Params{
		timeout: timeout,
	}
}

// NewWaypointGetLatestRelease2ParamsWithContext creates a new WaypointGetLatestRelease2Params object
// with the ability to set a context for a request.
func NewWaypointGetLatestRelease2ParamsWithContext(ctx context.Context) *WaypointGetLatestRelease2Params {
	return &WaypointGetLatestRelease2Params{
		Context: ctx,
	}
}

// NewWaypointGetLatestRelease2ParamsWithHTTPClient creates a new WaypointGetLatestRelease2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetLatestRelease2ParamsWithHTTPClient(client *http.Client) *WaypointGetLatestRelease2Params {
	return &WaypointGetLatestRelease2Params{
		HTTPClient: client,
	}
}

/*
WaypointGetLatestRelease2Params contains all the parameters to send to the API endpoint

	for the waypoint get latest release2 operation.

	Typically these are written to a http.Request.
*/
type WaypointGetLatestRelease2Params struct {

	// ApplicationApplication.
	ApplicationApplication string

	// ApplicationProject.
	ApplicationProject string

	/* LoadDetails.

	     Load additional details about the release. These will become available
	in the Preload section.

	     Default: "NONE"
	*/
	LoadDetails *string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// WorkspaceWorkspace.
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get latest release2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetLatestRelease2Params) WithDefaults() *WaypointGetLatestRelease2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get latest release2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetLatestRelease2Params) SetDefaults() {
	var (
		loadDetailsDefault = string("NONE")
	)

	val := WaypointGetLatestRelease2Params{
		LoadDetails: &loadDetailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithTimeout(timeout time.Duration) *WaypointGetLatestRelease2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithContext(ctx context.Context) *WaypointGetLatestRelease2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithHTTPClient(client *http.Client) *WaypointGetLatestRelease2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithApplicationApplication(applicationApplication string) *WaypointGetLatestRelease2Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithApplicationProject(applicationProject string) *WaypointGetLatestRelease2Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithLoadDetails adds the loadDetails to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithLoadDetails(loadDetails *string) *WaypointGetLatestRelease2Params {
	o.SetLoadDetails(loadDetails)
	return o
}

// SetLoadDetails adds the loadDetails to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetLoadDetails(loadDetails *string) {
	o.LoadDetails = loadDetails
}

// WithNamespaceID adds the namespaceID to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithNamespaceID(namespaceID string) *WaypointGetLatestRelease2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointGetLatestRelease2Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest release2 params
func (o *WaypointGetLatestRelease2Params) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetLatestRelease2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LoadDetails != nil {

		// query param load_details
		var qrLoadDetails string

		if o.LoadDetails != nil {
			qrLoadDetails = *o.LoadDetails
		}
		qLoadDetails := qrLoadDetails
		if qLoadDetails != "" {

			if err := r.SetQueryParam("load_details", qLoadDetails); err != nil {
				return err
			}
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
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
