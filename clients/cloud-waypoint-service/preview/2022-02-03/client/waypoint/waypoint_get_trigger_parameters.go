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

// NewWaypointGetTriggerParams creates a new WaypointGetTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetTriggerParams() *WaypointGetTriggerParams {
	return &WaypointGetTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetTriggerParamsWithTimeout creates a new WaypointGetTriggerParams object
// with the ability to set a timeout on a request.
func NewWaypointGetTriggerParamsWithTimeout(timeout time.Duration) *WaypointGetTriggerParams {
	return &WaypointGetTriggerParams{
		timeout: timeout,
	}
}

// NewWaypointGetTriggerParamsWithContext creates a new WaypointGetTriggerParams object
// with the ability to set a context for a request.
func NewWaypointGetTriggerParamsWithContext(ctx context.Context) *WaypointGetTriggerParams {
	return &WaypointGetTriggerParams{
		Context: ctx,
	}
}

// NewWaypointGetTriggerParamsWithHTTPClient creates a new WaypointGetTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetTriggerParamsWithHTTPClient(client *http.Client) *WaypointGetTriggerParams {
	return &WaypointGetTriggerParams{
		HTTPClient: client,
	}
}

/*
WaypointGetTriggerParams contains all the parameters to send to the API endpoint

	for the waypoint get trigger operation.

	Typically these are written to a http.Request.
*/
type WaypointGetTriggerParams struct {

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

// WithDefaults hydrates default values in the waypoint get trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetTriggerParams) WithDefaults() *WaypointGetTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get trigger params
func (o *WaypointGetTriggerParams) WithTimeout(timeout time.Duration) *WaypointGetTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get trigger params
func (o *WaypointGetTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get trigger params
func (o *WaypointGetTriggerParams) WithContext(ctx context.Context) *WaypointGetTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get trigger params
func (o *WaypointGetTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get trigger params
func (o *WaypointGetTriggerParams) WithHTTPClient(client *http.Client) *WaypointGetTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get trigger params
func (o *WaypointGetTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get trigger params
func (o *WaypointGetTriggerParams) WithNamespaceID(namespaceID string) *WaypointGetTriggerParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get trigger params
func (o *WaypointGetTriggerParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint get trigger params
func (o *WaypointGetTriggerParams) WithRefID(refID string) *WaypointGetTriggerParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get trigger params
func (o *WaypointGetTriggerParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
