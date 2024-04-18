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

// NewWaypointServiceUpdateAgentGroupParams creates a new WaypointServiceUpdateAgentGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateAgentGroupParams() *WaypointServiceUpdateAgentGroupParams {
	return &WaypointServiceUpdateAgentGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateAgentGroupParamsWithTimeout creates a new WaypointServiceUpdateAgentGroupParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateAgentGroupParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateAgentGroupParams {
	return &WaypointServiceUpdateAgentGroupParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateAgentGroupParamsWithContext creates a new WaypointServiceUpdateAgentGroupParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateAgentGroupParamsWithContext(ctx context.Context) *WaypointServiceUpdateAgentGroupParams {
	return &WaypointServiceUpdateAgentGroupParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateAgentGroupParamsWithHTTPClient creates a new WaypointServiceUpdateAgentGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateAgentGroupParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateAgentGroupParams {
	return &WaypointServiceUpdateAgentGroupParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateAgentGroupParams contains all the parameters to send to the API endpoint

	for the waypoint service update agent group operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateAgentGroupParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateAgentGroupBody

	/* Name.

	     The name of the group to update
	We don’t allow groups to be renamed so this is a stable identifier.
	*/
	Name string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update agent group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAgentGroupParams) WithDefaults() *WaypointServiceUpdateAgentGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update agent group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAgentGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateAgentGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithContext(ctx context.Context) *WaypointServiceUpdateAgentGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateAgentGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateAgentGroupBody) *WaypointServiceUpdateAgentGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateAgentGroupBody) {
	o.Body = body
}

// WithName adds the name to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithName(name string) *WaypointServiceUpdateAgentGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetName(name string) {
	o.Name = name
}

// WithNamespaceID adds the namespaceID to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) WithNamespaceID(namespaceID string) *WaypointServiceUpdateAgentGroupParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update agent group params
func (o *WaypointServiceUpdateAgentGroupParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateAgentGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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