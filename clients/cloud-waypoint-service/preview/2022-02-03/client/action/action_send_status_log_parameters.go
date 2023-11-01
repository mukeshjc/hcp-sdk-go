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

// NewActionSendStatusLogParams creates a new ActionSendStatusLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionSendStatusLogParams() *ActionSendStatusLogParams {
	return &ActionSendStatusLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionSendStatusLogParamsWithTimeout creates a new ActionSendStatusLogParams object
// with the ability to set a timeout on a request.
func NewActionSendStatusLogParamsWithTimeout(timeout time.Duration) *ActionSendStatusLogParams {
	return &ActionSendStatusLogParams{
		timeout: timeout,
	}
}

// NewActionSendStatusLogParamsWithContext creates a new ActionSendStatusLogParams object
// with the ability to set a context for a request.
func NewActionSendStatusLogParamsWithContext(ctx context.Context) *ActionSendStatusLogParams {
	return &ActionSendStatusLogParams{
		Context: ctx,
	}
}

// NewActionSendStatusLogParamsWithHTTPClient creates a new ActionSendStatusLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionSendStatusLogParamsWithHTTPClient(client *http.Client) *ActionSendStatusLogParams {
	return &ActionSendStatusLogParams{
		HTTPClient: client,
	}
}

/*
ActionSendStatusLogParams contains all the parameters to send to the API endpoint

	for the action send status log operation.

	Typically these are written to a http.Request.
*/
type ActionSendStatusLogParams struct {

	/* ActionConfigID.

	   ULID of the action config id
	*/
	ActionConfigID string

	/* ActionRunSeq.

	   The run sequence to attach this run to

	   Format: uint64
	*/
	ActionRunSeq string

	// Body.
	Body *models.HashicorpCloudWaypointActionsdriverSendStatusLogRequest

	/* NamespaceID.

	   The namespace the action to be listed in
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action send status log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionSendStatusLogParams) WithDefaults() *ActionSendStatusLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action send status log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionSendStatusLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action send status log params
func (o *ActionSendStatusLogParams) WithTimeout(timeout time.Duration) *ActionSendStatusLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action send status log params
func (o *ActionSendStatusLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action send status log params
func (o *ActionSendStatusLogParams) WithContext(ctx context.Context) *ActionSendStatusLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action send status log params
func (o *ActionSendStatusLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action send status log params
func (o *ActionSendStatusLogParams) WithHTTPClient(client *http.Client) *ActionSendStatusLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action send status log params
func (o *ActionSendStatusLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionConfigID adds the actionConfigID to the action send status log params
func (o *ActionSendStatusLogParams) WithActionConfigID(actionConfigID string) *ActionSendStatusLogParams {
	o.SetActionConfigID(actionConfigID)
	return o
}

// SetActionConfigID adds the actionConfigId to the action send status log params
func (o *ActionSendStatusLogParams) SetActionConfigID(actionConfigID string) {
	o.ActionConfigID = actionConfigID
}

// WithActionRunSeq adds the actionRunSeq to the action send status log params
func (o *ActionSendStatusLogParams) WithActionRunSeq(actionRunSeq string) *ActionSendStatusLogParams {
	o.SetActionRunSeq(actionRunSeq)
	return o
}

// SetActionRunSeq adds the actionRunSeq to the action send status log params
func (o *ActionSendStatusLogParams) SetActionRunSeq(actionRunSeq string) {
	o.ActionRunSeq = actionRunSeq
}

// WithBody adds the body to the action send status log params
func (o *ActionSendStatusLogParams) WithBody(body *models.HashicorpCloudWaypointActionsdriverSendStatusLogRequest) *ActionSendStatusLogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the action send status log params
func (o *ActionSendStatusLogParams) SetBody(body *models.HashicorpCloudWaypointActionsdriverSendStatusLogRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the action send status log params
func (o *ActionSendStatusLogParams) WithNamespaceID(namespaceID string) *ActionSendStatusLogParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the action send status log params
func (o *ActionSendStatusLogParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionSendStatusLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_config.id
	if err := r.SetPathParam("action_config.id", o.ActionConfigID); err != nil {
		return err
	}

	// path param action_run_seq
	if err := r.SetPathParam("action_run_seq", o.ActionRunSeq); err != nil {
		return err
	}
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
