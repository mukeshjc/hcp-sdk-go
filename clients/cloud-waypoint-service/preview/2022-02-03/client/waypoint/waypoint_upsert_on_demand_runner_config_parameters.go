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

// NewWaypointUpsertOnDemandRunnerConfigParams creates a new WaypointUpsertOnDemandRunnerConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUpsertOnDemandRunnerConfigParams() *WaypointUpsertOnDemandRunnerConfigParams {
	return &WaypointUpsertOnDemandRunnerConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUpsertOnDemandRunnerConfigParamsWithTimeout creates a new WaypointUpsertOnDemandRunnerConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointUpsertOnDemandRunnerConfigParamsWithTimeout(timeout time.Duration) *WaypointUpsertOnDemandRunnerConfigParams {
	return &WaypointUpsertOnDemandRunnerConfigParams{
		timeout: timeout,
	}
}

// NewWaypointUpsertOnDemandRunnerConfigParamsWithContext creates a new WaypointUpsertOnDemandRunnerConfigParams object
// with the ability to set a context for a request.
func NewWaypointUpsertOnDemandRunnerConfigParamsWithContext(ctx context.Context) *WaypointUpsertOnDemandRunnerConfigParams {
	return &WaypointUpsertOnDemandRunnerConfigParams{
		Context: ctx,
	}
}

// NewWaypointUpsertOnDemandRunnerConfigParamsWithHTTPClient creates a new WaypointUpsertOnDemandRunnerConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUpsertOnDemandRunnerConfigParamsWithHTTPClient(client *http.Client) *WaypointUpsertOnDemandRunnerConfigParams {
	return &WaypointUpsertOnDemandRunnerConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointUpsertOnDemandRunnerConfigParams contains all the parameters to send to the API endpoint

	for the waypoint upsert on demand runner config operation.

	Typically these are written to a http.Request.
*/
type WaypointUpsertOnDemandRunnerConfigParams struct {

	// Body.
	Body *models.HashicorpWaypointUpsertOnDemandRunnerConfigRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint upsert on demand runner config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithDefaults() *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint upsert on demand runner config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithTimeout(timeout time.Duration) *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithContext(ctx context.Context) *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithHTTPClient(client *http.Client) *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithBody(body *models.HashicorpWaypointUpsertOnDemandRunnerConfigRequest) *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetBody(body *models.HashicorpWaypointUpsertOnDemandRunnerConfigRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) WithNamespaceID(namespaceID string) *WaypointUpsertOnDemandRunnerConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint upsert on demand runner config params
func (o *WaypointUpsertOnDemandRunnerConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUpsertOnDemandRunnerConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
