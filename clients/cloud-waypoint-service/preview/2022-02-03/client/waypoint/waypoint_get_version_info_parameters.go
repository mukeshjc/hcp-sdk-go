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

// NewWaypointGetVersionInfoParams creates a new WaypointGetVersionInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetVersionInfoParams() *WaypointGetVersionInfoParams {
	return &WaypointGetVersionInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetVersionInfoParamsWithTimeout creates a new WaypointGetVersionInfoParams object
// with the ability to set a timeout on a request.
func NewWaypointGetVersionInfoParamsWithTimeout(timeout time.Duration) *WaypointGetVersionInfoParams {
	return &WaypointGetVersionInfoParams{
		timeout: timeout,
	}
}

// NewWaypointGetVersionInfoParamsWithContext creates a new WaypointGetVersionInfoParams object
// with the ability to set a context for a request.
func NewWaypointGetVersionInfoParamsWithContext(ctx context.Context) *WaypointGetVersionInfoParams {
	return &WaypointGetVersionInfoParams{
		Context: ctx,
	}
}

// NewWaypointGetVersionInfoParamsWithHTTPClient creates a new WaypointGetVersionInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetVersionInfoParamsWithHTTPClient(client *http.Client) *WaypointGetVersionInfoParams {
	return &WaypointGetVersionInfoParams{
		HTTPClient: client,
	}
}

/*
WaypointGetVersionInfoParams contains all the parameters to send to the API endpoint

	for the waypoint get version info operation.

	Typically these are written to a http.Request.
*/
type WaypointGetVersionInfoParams struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get version info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetVersionInfoParams) WithDefaults() *WaypointGetVersionInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get version info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetVersionInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) WithTimeout(timeout time.Duration) *WaypointGetVersionInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) WithContext(ctx context.Context) *WaypointGetVersionInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) WithHTTPClient(client *http.Client) *WaypointGetVersionInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) WithNamespaceID(namespaceID string) *WaypointGetVersionInfoParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get version info params
func (o *WaypointGetVersionInfoParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetVersionInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
