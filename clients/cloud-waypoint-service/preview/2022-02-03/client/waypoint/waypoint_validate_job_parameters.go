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

// NewWaypointValidateJobParams creates a new WaypointValidateJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointValidateJobParams() *WaypointValidateJobParams {
	return &WaypointValidateJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointValidateJobParamsWithTimeout creates a new WaypointValidateJobParams object
// with the ability to set a timeout on a request.
func NewWaypointValidateJobParamsWithTimeout(timeout time.Duration) *WaypointValidateJobParams {
	return &WaypointValidateJobParams{
		timeout: timeout,
	}
}

// NewWaypointValidateJobParamsWithContext creates a new WaypointValidateJobParams object
// with the ability to set a context for a request.
func NewWaypointValidateJobParamsWithContext(ctx context.Context) *WaypointValidateJobParams {
	return &WaypointValidateJobParams{
		Context: ctx,
	}
}

// NewWaypointValidateJobParamsWithHTTPClient creates a new WaypointValidateJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointValidateJobParamsWithHTTPClient(client *http.Client) *WaypointValidateJobParams {
	return &WaypointValidateJobParams{
		HTTPClient: client,
	}
}

/*
WaypointValidateJobParams contains all the parameters to send to the API endpoint

	for the waypoint validate job operation.

	Typically these are written to a http.Request.
*/
type WaypointValidateJobParams struct {

	// Body.
	Body *models.HashicorpWaypointValidateJobRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint validate job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointValidateJobParams) WithDefaults() *WaypointValidateJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint validate job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointValidateJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint validate job params
func (o *WaypointValidateJobParams) WithTimeout(timeout time.Duration) *WaypointValidateJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint validate job params
func (o *WaypointValidateJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint validate job params
func (o *WaypointValidateJobParams) WithContext(ctx context.Context) *WaypointValidateJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint validate job params
func (o *WaypointValidateJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint validate job params
func (o *WaypointValidateJobParams) WithHTTPClient(client *http.Client) *WaypointValidateJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint validate job params
func (o *WaypointValidateJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint validate job params
func (o *WaypointValidateJobParams) WithBody(body *models.HashicorpWaypointValidateJobRequest) *WaypointValidateJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint validate job params
func (o *WaypointValidateJobParams) SetBody(body *models.HashicorpWaypointValidateJobRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint validate job params
func (o *WaypointValidateJobParams) WithNamespaceID(namespaceID string) *WaypointValidateJobParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint validate job params
func (o *WaypointValidateJobParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointValidateJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
