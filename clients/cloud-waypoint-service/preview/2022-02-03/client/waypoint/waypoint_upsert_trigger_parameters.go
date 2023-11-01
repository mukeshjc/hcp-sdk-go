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

// NewWaypointUpsertTriggerParams creates a new WaypointUpsertTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUpsertTriggerParams() *WaypointUpsertTriggerParams {
	return &WaypointUpsertTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUpsertTriggerParamsWithTimeout creates a new WaypointUpsertTriggerParams object
// with the ability to set a timeout on a request.
func NewWaypointUpsertTriggerParamsWithTimeout(timeout time.Duration) *WaypointUpsertTriggerParams {
	return &WaypointUpsertTriggerParams{
		timeout: timeout,
	}
}

// NewWaypointUpsertTriggerParamsWithContext creates a new WaypointUpsertTriggerParams object
// with the ability to set a context for a request.
func NewWaypointUpsertTriggerParamsWithContext(ctx context.Context) *WaypointUpsertTriggerParams {
	return &WaypointUpsertTriggerParams{
		Context: ctx,
	}
}

// NewWaypointUpsertTriggerParamsWithHTTPClient creates a new WaypointUpsertTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUpsertTriggerParamsWithHTTPClient(client *http.Client) *WaypointUpsertTriggerParams {
	return &WaypointUpsertTriggerParams{
		HTTPClient: client,
	}
}

/*
WaypointUpsertTriggerParams contains all the parameters to send to the API endpoint

	for the waypoint upsert trigger operation.

	Typically these are written to a http.Request.
*/
type WaypointUpsertTriggerParams struct {

	// Body.
	Body *models.HashicorpWaypointUpsertTriggerRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint upsert trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpsertTriggerParams) WithDefaults() *WaypointUpsertTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint upsert trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpsertTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) WithTimeout(timeout time.Duration) *WaypointUpsertTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) WithContext(ctx context.Context) *WaypointUpsertTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) WithHTTPClient(client *http.Client) *WaypointUpsertTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) WithBody(body *models.HashicorpWaypointUpsertTriggerRequest) *WaypointUpsertTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) SetBody(body *models.HashicorpWaypointUpsertTriggerRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) WithNamespaceID(namespaceID string) *WaypointUpsertTriggerParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint upsert trigger params
func (o *WaypointUpsertTriggerParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUpsertTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
