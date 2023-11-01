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

// NewWaypointDeleteUserParams creates a new WaypointDeleteUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointDeleteUserParams() *WaypointDeleteUserParams {
	return &WaypointDeleteUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDeleteUserParamsWithTimeout creates a new WaypointDeleteUserParams object
// with the ability to set a timeout on a request.
func NewWaypointDeleteUserParamsWithTimeout(timeout time.Duration) *WaypointDeleteUserParams {
	return &WaypointDeleteUserParams{
		timeout: timeout,
	}
}

// NewWaypointDeleteUserParamsWithContext creates a new WaypointDeleteUserParams object
// with the ability to set a context for a request.
func NewWaypointDeleteUserParamsWithContext(ctx context.Context) *WaypointDeleteUserParams {
	return &WaypointDeleteUserParams{
		Context: ctx,
	}
}

// NewWaypointDeleteUserParamsWithHTTPClient creates a new WaypointDeleteUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointDeleteUserParamsWithHTTPClient(client *http.Client) *WaypointDeleteUserParams {
	return &WaypointDeleteUserParams{
		HTTPClient: client,
	}
}

/*
WaypointDeleteUserParams contains all the parameters to send to the API endpoint

	for the waypoint delete user operation.

	Typically these are written to a http.Request.
*/
type WaypointDeleteUserParams struct {

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// UserIDID.
	UserIDID string

	// UserUsernameUsername.
	UserUsernameUsername *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteUserParams) WithDefaults() *WaypointDeleteUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithTimeout(timeout time.Duration) *WaypointDeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithContext(ctx context.Context) *WaypointDeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithHTTPClient(client *http.Client) *WaypointDeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithNamespaceID(namespaceID string) *WaypointDeleteUserParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithUserIDID adds the userIDID to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithUserIDID(userIDID string) *WaypointDeleteUserParams {
	o.SetUserIDID(userIDID)
	return o
}

// SetUserIDID adds the userIdId to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetUserIDID(userIDID string) {
	o.UserIDID = userIDID
}

// WithUserUsernameUsername adds the userUsernameUsername to the waypoint delete user params
func (o *WaypointDeleteUserParams) WithUserUsernameUsername(userUsernameUsername *string) *WaypointDeleteUserParams {
	o.SetUserUsernameUsername(userUsernameUsername)
	return o
}

// SetUserUsernameUsername adds the userUsernameUsername to the waypoint delete user params
func (o *WaypointDeleteUserParams) SetUserUsernameUsername(userUsernameUsername *string) {
	o.UserUsernameUsername = userUsernameUsername
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param user.id.id
	if err := r.SetPathParam("user.id.id", o.UserIDID); err != nil {
		return err
	}

	if o.UserUsernameUsername != nil {

		// query param user.username.username
		var qrUserUsernameUsername string

		if o.UserUsernameUsername != nil {
			qrUserUsernameUsername = *o.UserUsernameUsername
		}
		qUserUsernameUsername := qrUserUsernameUsername
		if qUserUsernameUsername != "" {

			if err := r.SetQueryParam("user.username.username", qUserUsernameUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
