// Code generated by go-swagger; DO NOT EDIT.

package action

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

// NewActionRunActionParams creates a new ActionRunActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionRunActionParams() *ActionRunActionParams {
	return &ActionRunActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionRunActionParamsWithTimeout creates a new ActionRunActionParams object
// with the ability to set a timeout on a request.
func NewActionRunActionParamsWithTimeout(timeout time.Duration) *ActionRunActionParams {
	return &ActionRunActionParams{
		timeout: timeout,
	}
}

// NewActionRunActionParamsWithContext creates a new ActionRunActionParams object
// with the ability to set a context for a request.
func NewActionRunActionParamsWithContext(ctx context.Context) *ActionRunActionParams {
	return &ActionRunActionParams{
		Context: ctx,
	}
}

// NewActionRunActionParamsWithHTTPClient creates a new ActionRunActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionRunActionParamsWithHTTPClient(client *http.Client) *ActionRunActionParams {
	return &ActionRunActionParams{
		HTTPClient: client,
	}
}

/*
ActionRunActionParams contains all the parameters to send to the API endpoint

	for the action run action operation.

	Typically these are written to a http.Request.
*/
type ActionRunActionParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointActionsdriverRunActionRequest

	/* NamespaceID.

	   The namespace the action will run in
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action run action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionRunActionParams) WithDefaults() *ActionRunActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action run action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionRunActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action run action params
func (o *ActionRunActionParams) WithTimeout(timeout time.Duration) *ActionRunActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action run action params
func (o *ActionRunActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action run action params
func (o *ActionRunActionParams) WithContext(ctx context.Context) *ActionRunActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action run action params
func (o *ActionRunActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action run action params
func (o *ActionRunActionParams) WithHTTPClient(client *http.Client) *ActionRunActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action run action params
func (o *ActionRunActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the action run action params
func (o *ActionRunActionParams) WithBody(body *models.HashicorpCloudWaypointActionsdriverRunActionRequest) *ActionRunActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the action run action params
func (o *ActionRunActionParams) SetBody(body *models.HashicorpCloudWaypointActionsdriverRunActionRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the action run action params
func (o *ActionRunActionParams) WithNamespaceID(namespaceID string) *ActionRunActionParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the action run action params
func (o *ActionRunActionParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionRunActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
