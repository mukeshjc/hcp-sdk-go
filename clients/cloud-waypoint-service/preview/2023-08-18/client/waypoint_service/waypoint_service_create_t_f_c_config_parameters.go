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

// NewWaypointServiceCreateTFCConfigParams creates a new WaypointServiceCreateTFCConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateTFCConfigParams() *WaypointServiceCreateTFCConfigParams {
	return &WaypointServiceCreateTFCConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateTFCConfigParamsWithTimeout creates a new WaypointServiceCreateTFCConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateTFCConfigParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateTFCConfigParams {
	return &WaypointServiceCreateTFCConfigParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateTFCConfigParamsWithContext creates a new WaypointServiceCreateTFCConfigParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateTFCConfigParamsWithContext(ctx context.Context) *WaypointServiceCreateTFCConfigParams {
	return &WaypointServiceCreateTFCConfigParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateTFCConfigParamsWithHTTPClient creates a new WaypointServiceCreateTFCConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateTFCConfigParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateTFCConfigParams {
	return &WaypointServiceCreateTFCConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateTFCConfigParams contains all the parameters to send to the API endpoint

	for the waypoint service create t f c config operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateTFCConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointCreateTFCConfigRequest

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service create t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateTFCConfigParams) WithDefaults() *WaypointServiceCreateTFCConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateTFCConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateTFCConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) WithContext(ctx context.Context) *WaypointServiceCreateTFCConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateTFCConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) WithBody(body *models.HashicorpCloudWaypointCreateTFCConfigRequest) *WaypointServiceCreateTFCConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) SetBody(body *models.HashicorpCloudWaypointCreateTFCConfigRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) WithNamespaceID(namespaceID string) *WaypointServiceCreateTFCConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service create t f c config params
func (o *WaypointServiceCreateTFCConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateTFCConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
