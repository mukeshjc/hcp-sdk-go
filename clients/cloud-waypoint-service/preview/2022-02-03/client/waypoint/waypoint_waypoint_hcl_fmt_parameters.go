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

// NewWaypointWaypointHclFmtParams creates a new WaypointWaypointHclFmtParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointWaypointHclFmtParams() *WaypointWaypointHclFmtParams {
	return &WaypointWaypointHclFmtParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointWaypointHclFmtParamsWithTimeout creates a new WaypointWaypointHclFmtParams object
// with the ability to set a timeout on a request.
func NewWaypointWaypointHclFmtParamsWithTimeout(timeout time.Duration) *WaypointWaypointHclFmtParams {
	return &WaypointWaypointHclFmtParams{
		timeout: timeout,
	}
}

// NewWaypointWaypointHclFmtParamsWithContext creates a new WaypointWaypointHclFmtParams object
// with the ability to set a context for a request.
func NewWaypointWaypointHclFmtParamsWithContext(ctx context.Context) *WaypointWaypointHclFmtParams {
	return &WaypointWaypointHclFmtParams{
		Context: ctx,
	}
}

// NewWaypointWaypointHclFmtParamsWithHTTPClient creates a new WaypointWaypointHclFmtParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointWaypointHclFmtParamsWithHTTPClient(client *http.Client) *WaypointWaypointHclFmtParams {
	return &WaypointWaypointHclFmtParams{
		HTTPClient: client,
	}
}

/*
WaypointWaypointHclFmtParams contains all the parameters to send to the API endpoint

	for the waypoint waypoint hcl fmt operation.

	Typically these are written to a http.Request.
*/
type WaypointWaypointHclFmtParams struct {

	// Body.
	Body *models.HashicorpWaypointWaypointHclFmtRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint waypoint hcl fmt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointWaypointHclFmtParams) WithDefaults() *WaypointWaypointHclFmtParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint waypoint hcl fmt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointWaypointHclFmtParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) WithTimeout(timeout time.Duration) *WaypointWaypointHclFmtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) WithContext(ctx context.Context) *WaypointWaypointHclFmtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) WithHTTPClient(client *http.Client) *WaypointWaypointHclFmtParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) WithBody(body *models.HashicorpWaypointWaypointHclFmtRequest) *WaypointWaypointHclFmtParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) SetBody(body *models.HashicorpWaypointWaypointHclFmtRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) WithNamespaceID(namespaceID string) *WaypointWaypointHclFmtParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint waypoint hcl fmt params
func (o *WaypointWaypointHclFmtParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointWaypointHclFmtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
