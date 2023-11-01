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

// NewWaypointDeleteAddOnDefinition2Params creates a new WaypointDeleteAddOnDefinition2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointDeleteAddOnDefinition2Params() *WaypointDeleteAddOnDefinition2Params {
	return &WaypointDeleteAddOnDefinition2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDeleteAddOnDefinition2ParamsWithTimeout creates a new WaypointDeleteAddOnDefinition2Params object
// with the ability to set a timeout on a request.
func NewWaypointDeleteAddOnDefinition2ParamsWithTimeout(timeout time.Duration) *WaypointDeleteAddOnDefinition2Params {
	return &WaypointDeleteAddOnDefinition2Params{
		timeout: timeout,
	}
}

// NewWaypointDeleteAddOnDefinition2ParamsWithContext creates a new WaypointDeleteAddOnDefinition2Params object
// with the ability to set a context for a request.
func NewWaypointDeleteAddOnDefinition2ParamsWithContext(ctx context.Context) *WaypointDeleteAddOnDefinition2Params {
	return &WaypointDeleteAddOnDefinition2Params{
		Context: ctx,
	}
}

// NewWaypointDeleteAddOnDefinition2ParamsWithHTTPClient creates a new WaypointDeleteAddOnDefinition2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointDeleteAddOnDefinition2ParamsWithHTTPClient(client *http.Client) *WaypointDeleteAddOnDefinition2Params {
	return &WaypointDeleteAddOnDefinition2Params{
		HTTPClient: client,
	}
}

/*
WaypointDeleteAddOnDefinition2Params contains all the parameters to send to the API endpoint

	for the waypoint delete add on definition2 operation.

	Typically these are written to a http.Request.
*/
type WaypointDeleteAddOnDefinition2Params struct {

	// AddOnDefinitionID.
	AddOnDefinitionID *string

	// AddOnDefinitionName.
	AddOnDefinitionName string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint delete add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteAddOnDefinition2Params) WithDefaults() *WaypointDeleteAddOnDefinition2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint delete add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointDeleteAddOnDefinition2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithTimeout(timeout time.Duration) *WaypointDeleteAddOnDefinition2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithContext(ctx context.Context) *WaypointDeleteAddOnDefinition2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithHTTPClient(client *http.Client) *WaypointDeleteAddOnDefinition2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnDefinitionID adds the addOnDefinitionID to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithAddOnDefinitionID(addOnDefinitionID *string) *WaypointDeleteAddOnDefinition2Params {
	o.SetAddOnDefinitionID(addOnDefinitionID)
	return o
}

// SetAddOnDefinitionID adds the addOnDefinitionId to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetAddOnDefinitionID(addOnDefinitionID *string) {
	o.AddOnDefinitionID = addOnDefinitionID
}

// WithAddOnDefinitionName adds the addOnDefinitionName to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithAddOnDefinitionName(addOnDefinitionName string) *WaypointDeleteAddOnDefinition2Params {
	o.SetAddOnDefinitionName(addOnDefinitionName)
	return o
}

// SetAddOnDefinitionName adds the addOnDefinitionName to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetAddOnDefinitionName(addOnDefinitionName string) {
	o.AddOnDefinitionName = addOnDefinitionName
}

// WithNamespaceID adds the namespaceID to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) WithNamespaceID(namespaceID string) *WaypointDeleteAddOnDefinition2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint delete add on definition2 params
func (o *WaypointDeleteAddOnDefinition2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDeleteAddOnDefinition2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddOnDefinitionID != nil {

		// query param add_on_definition.id
		var qrAddOnDefinitionID string

		if o.AddOnDefinitionID != nil {
			qrAddOnDefinitionID = *o.AddOnDefinitionID
		}
		qAddOnDefinitionID := qrAddOnDefinitionID
		if qAddOnDefinitionID != "" {

			if err := r.SetQueryParam("add_on_definition.id", qAddOnDefinitionID); err != nil {
				return err
			}
		}
	}

	// path param add_on_definition.name
	if err := r.SetPathParam("add_on_definition.name", o.AddOnDefinitionName); err != nil {
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
