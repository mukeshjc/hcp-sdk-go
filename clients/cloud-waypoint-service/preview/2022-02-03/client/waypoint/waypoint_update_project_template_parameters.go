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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// NewWaypointUpdateProjectTemplateParams creates a new WaypointUpdateProjectTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUpdateProjectTemplateParams() *WaypointUpdateProjectTemplateParams {
	return &WaypointUpdateProjectTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUpdateProjectTemplateParamsWithTimeout creates a new WaypointUpdateProjectTemplateParams object
// with the ability to set a timeout on a request.
func NewWaypointUpdateProjectTemplateParamsWithTimeout(timeout time.Duration) *WaypointUpdateProjectTemplateParams {
	return &WaypointUpdateProjectTemplateParams{
		timeout: timeout,
	}
}

// NewWaypointUpdateProjectTemplateParamsWithContext creates a new WaypointUpdateProjectTemplateParams object
// with the ability to set a context for a request.
func NewWaypointUpdateProjectTemplateParamsWithContext(ctx context.Context) *WaypointUpdateProjectTemplateParams {
	return &WaypointUpdateProjectTemplateParams{
		Context: ctx,
	}
}

// NewWaypointUpdateProjectTemplateParamsWithHTTPClient creates a new WaypointUpdateProjectTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUpdateProjectTemplateParamsWithHTTPClient(client *http.Client) *WaypointUpdateProjectTemplateParams {
	return &WaypointUpdateProjectTemplateParams{
		HTTPClient: client,
	}
}

/*
WaypointUpdateProjectTemplateParams contains all the parameters to send to the API endpoint

	for the waypoint update project template operation.

	Typically these are written to a http.Request.
*/
type WaypointUpdateProjectTemplateParams struct {

	// Body.
	Body *models.HashicorpWaypointUpdateProjectTemplateRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* ProjectTemplateID.

	   Unique ID of the ProjectTemplate
	*/
	ProjectTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint update project template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpdateProjectTemplateParams) WithDefaults() *WaypointUpdateProjectTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint update project template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpdateProjectTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithTimeout(timeout time.Duration) *WaypointUpdateProjectTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithContext(ctx context.Context) *WaypointUpdateProjectTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithHTTPClient(client *http.Client) *WaypointUpdateProjectTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithBody(body *models.HashicorpWaypointUpdateProjectTemplateRequest) *WaypointUpdateProjectTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetBody(body *models.HashicorpWaypointUpdateProjectTemplateRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithNamespaceID(namespaceID string) *WaypointUpdateProjectTemplateParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectTemplateID adds the projectTemplateID to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) WithProjectTemplateID(projectTemplateID string) *WaypointUpdateProjectTemplateParams {
	o.SetProjectTemplateID(projectTemplateID)
	return o
}

// SetProjectTemplateID adds the projectTemplateId to the waypoint update project template params
func (o *WaypointUpdateProjectTemplateParams) SetProjectTemplateID(projectTemplateID string) {
	o.ProjectTemplateID = projectTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUpdateProjectTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param project_template.id
	if err := r.SetPathParam("project_template.id", o.ProjectTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
