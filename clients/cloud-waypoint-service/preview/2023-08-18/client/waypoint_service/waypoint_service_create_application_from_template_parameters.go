// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// NewWaypointServiceCreateApplicationFromTemplateParams creates a new WaypointServiceCreateApplicationFromTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateApplicationFromTemplateParams() *WaypointServiceCreateApplicationFromTemplateParams {
	return &WaypointServiceCreateApplicationFromTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateApplicationFromTemplateParamsWithTimeout creates a new WaypointServiceCreateApplicationFromTemplateParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateApplicationFromTemplateParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateApplicationFromTemplateParams {
	return &WaypointServiceCreateApplicationFromTemplateParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateApplicationFromTemplateParamsWithContext creates a new WaypointServiceCreateApplicationFromTemplateParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateApplicationFromTemplateParamsWithContext(ctx context.Context) *WaypointServiceCreateApplicationFromTemplateParams {
	return &WaypointServiceCreateApplicationFromTemplateParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateApplicationFromTemplateParamsWithHTTPClient creates a new WaypointServiceCreateApplicationFromTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateApplicationFromTemplateParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateApplicationFromTemplateParams {
	return &WaypointServiceCreateApplicationFromTemplateParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateApplicationFromTemplateParams contains all the parameters to send to the API endpoint

	for the waypoint service create application from template operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateApplicationFromTemplateParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointCreateApplicationFromTemplateRequest

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service create application from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithDefaults() *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create application from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithContext(ctx context.Context) *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithBody(body *models.HashicorpCloudWaypointCreateApplicationFromTemplateRequest) *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetBody(body *models.HashicorpCloudWaypointCreateApplicationFromTemplateRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) WithNamespaceID(namespaceID string) *WaypointServiceCreateApplicationFromTemplateParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service create application from template params
func (o *WaypointServiceCreateApplicationFromTemplateParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateApplicationFromTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}