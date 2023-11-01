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

// NewWaypointGetProjectTemplateParams creates a new WaypointGetProjectTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetProjectTemplateParams() *WaypointGetProjectTemplateParams {
	return &WaypointGetProjectTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetProjectTemplateParamsWithTimeout creates a new WaypointGetProjectTemplateParams object
// with the ability to set a timeout on a request.
func NewWaypointGetProjectTemplateParamsWithTimeout(timeout time.Duration) *WaypointGetProjectTemplateParams {
	return &WaypointGetProjectTemplateParams{
		timeout: timeout,
	}
}

// NewWaypointGetProjectTemplateParamsWithContext creates a new WaypointGetProjectTemplateParams object
// with the ability to set a context for a request.
func NewWaypointGetProjectTemplateParamsWithContext(ctx context.Context) *WaypointGetProjectTemplateParams {
	return &WaypointGetProjectTemplateParams{
		Context: ctx,
	}
}

// NewWaypointGetProjectTemplateParamsWithHTTPClient creates a new WaypointGetProjectTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetProjectTemplateParamsWithHTTPClient(client *http.Client) *WaypointGetProjectTemplateParams {
	return &WaypointGetProjectTemplateParams{
		HTTPClient: client,
	}
}

/*
WaypointGetProjectTemplateParams contains all the parameters to send to the API endpoint

	for the waypoint get project template operation.

	Typically these are written to a http.Request.
*/
type WaypointGetProjectTemplateParams struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* ProjectTemplateID.

	   ID of the ProjectTemplate
	*/
	ProjectTemplateID string

	/* ProjectTemplateName.

	   Name of the ProjectTemplate.
	*/
	ProjectTemplateName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get project template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetProjectTemplateParams) WithDefaults() *WaypointGetProjectTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get project template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetProjectTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithTimeout(timeout time.Duration) *WaypointGetProjectTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithContext(ctx context.Context) *WaypointGetProjectTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithHTTPClient(client *http.Client) *WaypointGetProjectTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithNamespaceID(namespaceID string) *WaypointGetProjectTemplateParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectTemplateID adds the projectTemplateID to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithProjectTemplateID(projectTemplateID string) *WaypointGetProjectTemplateParams {
	o.SetProjectTemplateID(projectTemplateID)
	return o
}

// SetProjectTemplateID adds the projectTemplateId to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetProjectTemplateID(projectTemplateID string) {
	o.ProjectTemplateID = projectTemplateID
}

// WithProjectTemplateName adds the projectTemplateName to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) WithProjectTemplateName(projectTemplateName *string) *WaypointGetProjectTemplateParams {
	o.SetProjectTemplateName(projectTemplateName)
	return o
}

// SetProjectTemplateName adds the projectTemplateName to the waypoint get project template params
func (o *WaypointGetProjectTemplateParams) SetProjectTemplateName(projectTemplateName *string) {
	o.ProjectTemplateName = projectTemplateName
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetProjectTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param project_template.id
	if err := r.SetPathParam("project_template.id", o.ProjectTemplateID); err != nil {
		return err
	}

	if o.ProjectTemplateName != nil {

		// query param project_template.name
		var qrProjectTemplateName string

		if o.ProjectTemplateName != nil {
			qrProjectTemplateName = *o.ProjectTemplateName
		}
		qProjectTemplateName := qrProjectTemplateName
		if qProjectTemplateName != "" {

			if err := r.SetQueryParam("project_template.name", qProjectTemplateName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
