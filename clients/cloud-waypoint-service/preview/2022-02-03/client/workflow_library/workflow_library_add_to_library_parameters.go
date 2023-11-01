// Code generated by go-swagger; DO NOT EDIT.

package workflow_library

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

// NewWorkflowLibraryAddToLibraryParams creates a new WorkflowLibraryAddToLibraryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowLibraryAddToLibraryParams() *WorkflowLibraryAddToLibraryParams {
	return &WorkflowLibraryAddToLibraryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowLibraryAddToLibraryParamsWithTimeout creates a new WorkflowLibraryAddToLibraryParams object
// with the ability to set a timeout on a request.
func NewWorkflowLibraryAddToLibraryParamsWithTimeout(timeout time.Duration) *WorkflowLibraryAddToLibraryParams {
	return &WorkflowLibraryAddToLibraryParams{
		timeout: timeout,
	}
}

// NewWorkflowLibraryAddToLibraryParamsWithContext creates a new WorkflowLibraryAddToLibraryParams object
// with the ability to set a context for a request.
func NewWorkflowLibraryAddToLibraryParamsWithContext(ctx context.Context) *WorkflowLibraryAddToLibraryParams {
	return &WorkflowLibraryAddToLibraryParams{
		Context: ctx,
	}
}

// NewWorkflowLibraryAddToLibraryParamsWithHTTPClient creates a new WorkflowLibraryAddToLibraryParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowLibraryAddToLibraryParamsWithHTTPClient(client *http.Client) *WorkflowLibraryAddToLibraryParams {
	return &WorkflowLibraryAddToLibraryParams{
		HTTPClient: client,
	}
}

/*
WorkflowLibraryAddToLibraryParams contains all the parameters to send to the API endpoint

	for the workflow library add to library operation.

	Typically these are written to a http.Request.
*/
type WorkflowLibraryAddToLibraryParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointAddToLibraryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow library add to library params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowLibraryAddToLibraryParams) WithDefaults() *WorkflowLibraryAddToLibraryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow library add to library params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowLibraryAddToLibraryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) WithTimeout(timeout time.Duration) *WorkflowLibraryAddToLibraryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) WithContext(ctx context.Context) *WorkflowLibraryAddToLibraryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) WithHTTPClient(client *http.Client) *WorkflowLibraryAddToLibraryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) WithBody(body *models.HashicorpCloudWaypointAddToLibraryRequest) *WorkflowLibraryAddToLibraryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the workflow library add to library params
func (o *WorkflowLibraryAddToLibraryParams) SetBody(body *models.HashicorpCloudWaypointAddToLibraryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowLibraryAddToLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
