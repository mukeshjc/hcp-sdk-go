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

// NewWaypointGenerateLoginTokenParams creates a new WaypointGenerateLoginTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGenerateLoginTokenParams() *WaypointGenerateLoginTokenParams {
	return &WaypointGenerateLoginTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGenerateLoginTokenParamsWithTimeout creates a new WaypointGenerateLoginTokenParams object
// with the ability to set a timeout on a request.
func NewWaypointGenerateLoginTokenParamsWithTimeout(timeout time.Duration) *WaypointGenerateLoginTokenParams {
	return &WaypointGenerateLoginTokenParams{
		timeout: timeout,
	}
}

// NewWaypointGenerateLoginTokenParamsWithContext creates a new WaypointGenerateLoginTokenParams object
// with the ability to set a context for a request.
func NewWaypointGenerateLoginTokenParamsWithContext(ctx context.Context) *WaypointGenerateLoginTokenParams {
	return &WaypointGenerateLoginTokenParams{
		Context: ctx,
	}
}

// NewWaypointGenerateLoginTokenParamsWithHTTPClient creates a new WaypointGenerateLoginTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGenerateLoginTokenParamsWithHTTPClient(client *http.Client) *WaypointGenerateLoginTokenParams {
	return &WaypointGenerateLoginTokenParams{
		HTTPClient: client,
	}
}

/*
WaypointGenerateLoginTokenParams contains all the parameters to send to the API endpoint

	for the waypoint generate login token operation.

	Typically these are written to a http.Request.
*/
type WaypointGenerateLoginTokenParams struct {

	// Body.
	Body *models.HashicorpWaypointLoginTokenRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint generate login token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGenerateLoginTokenParams) WithDefaults() *WaypointGenerateLoginTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint generate login token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGenerateLoginTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) WithTimeout(timeout time.Duration) *WaypointGenerateLoginTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) WithContext(ctx context.Context) *WaypointGenerateLoginTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) WithHTTPClient(client *http.Client) *WaypointGenerateLoginTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) WithBody(body *models.HashicorpWaypointLoginTokenRequest) *WaypointGenerateLoginTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) SetBody(body *models.HashicorpWaypointLoginTokenRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) WithNamespaceID(namespaceID string) *WaypointGenerateLoginTokenParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint generate login token params
func (o *WaypointGenerateLoginTokenParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGenerateLoginTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
