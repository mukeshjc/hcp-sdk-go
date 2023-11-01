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

// NewWaypointRunTriggerParams creates a new WaypointRunTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointRunTriggerParams() *WaypointRunTriggerParams {
	return &WaypointRunTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointRunTriggerParamsWithTimeout creates a new WaypointRunTriggerParams object
// with the ability to set a timeout on a request.
func NewWaypointRunTriggerParamsWithTimeout(timeout time.Duration) *WaypointRunTriggerParams {
	return &WaypointRunTriggerParams{
		timeout: timeout,
	}
}

// NewWaypointRunTriggerParamsWithContext creates a new WaypointRunTriggerParams object
// with the ability to set a context for a request.
func NewWaypointRunTriggerParamsWithContext(ctx context.Context) *WaypointRunTriggerParams {
	return &WaypointRunTriggerParams{
		Context: ctx,
	}
}

// NewWaypointRunTriggerParamsWithHTTPClient creates a new WaypointRunTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointRunTriggerParamsWithHTTPClient(client *http.Client) *WaypointRunTriggerParams {
	return &WaypointRunTriggerParams{
		HTTPClient: client,
	}
}

/*
WaypointRunTriggerParams contains all the parameters to send to the API endpoint

	for the waypoint run trigger operation.

	Typically these are written to a http.Request.
*/
type WaypointRunTriggerParams struct {

	// Body.
	Body *models.HashicorpWaypointRunTriggerRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// RefID.
	RefID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint run trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointRunTriggerParams) WithDefaults() *WaypointRunTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint run trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointRunTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithTimeout(timeout time.Duration) *WaypointRunTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithContext(ctx context.Context) *WaypointRunTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithHTTPClient(client *http.Client) *WaypointRunTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithBody(body *models.HashicorpWaypointRunTriggerRequest) *WaypointRunTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetBody(body *models.HashicorpWaypointRunTriggerRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithNamespaceID(namespaceID string) *WaypointRunTriggerParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint run trigger params
func (o *WaypointRunTriggerParams) WithRefID(refID string) *WaypointRunTriggerParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint run trigger params
func (o *WaypointRunTriggerParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointRunTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ref.id
	if err := r.SetPathParam("ref.id", o.RefID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
