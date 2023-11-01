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

// NewWaypointDeleteAuthMethodParams creates a new WaypointDeleteAuthMethodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointDeleteAuthMethodParams() *WaypointDeleteAuthMethodParams {
	return &WaypointDeleteAuthMethodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDeleteAuthMethodParamsWithTimeout creates a new WaypointDeleteAuthMethodParams object
// with the ability to set a timeout on a request.
func NewWaypointDeleteAuthMethodParamsWithTimeout(timeout time.Duration) *WaypointDeleteAuthMethodParams {
	return &WaypointDeleteAuthMethodParams{
		timeout: timeout,
	}
}

// NewWaypointDeleteAuthMethodParamsWithContext creates a new WaypointDeleteAuthMethodParams object
// with the ability to set a context for a request.
func NewWaypointDeleteAuthMethodParamsWithContext(ctx context.Context) *WaypointDeleteAuthMethodParams {
	return &WaypointDeleteAuthMethodParams{
		Context: ctx,
	}
}

// NewWaypointDeleteAuthMethodParamsWithHTTPClient creates a new WaypointDeleteAuthMethodParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointDeleteAuthMethodParamsWithHTTPClient(client *http.Client) *WaypointDeleteAuthMethodParams {
	return &WaypointDeleteAuthMethodParams{
		HTTPClient: client,
	}
}

/*
WaypointDeleteAuthMethodParams contains all the parameters to send to the API endpoint

	for the waypoint delete auth method operation.

	Typically these are written to a http.Request.
*/
type WaypointDeleteAuthMethodParams struct {

	// AuthMethodName.
	AuthMethodName string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint delete auth method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteAuthMethodParams) WithDefaults() *WaypointDeleteAuthMethodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint delete auth method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteAuthMethodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) WithTimeout(timeout time.Duration) *WaypointDeleteAuthMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) WithContext(ctx context.Context) *WaypointDeleteAuthMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) WithHTTPClient(client *http.Client) *WaypointDeleteAuthMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthMethodName adds the authMethodName to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) WithAuthMethodName(authMethodName string) *WaypointDeleteAuthMethodParams {
	o.SetAuthMethodName(authMethodName)
	return o
}

// SetAuthMethodName adds the authMethodName to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) SetAuthMethodName(authMethodName string) {
	o.AuthMethodName = authMethodName
}

// WithNamespaceID adds the namespaceID to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) WithNamespaceID(namespaceID string) *WaypointDeleteAuthMethodParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint delete auth method params
func (o *WaypointDeleteAuthMethodParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDeleteAuthMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param auth_method.name
	if err := r.SetPathParam("auth_method.name", o.AuthMethodName); err != nil {
		return err
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
