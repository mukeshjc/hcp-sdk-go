// Code generated by go-swagger; DO NOT EDIT.

package streaming_service

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

// NewStreamingServiceGetDestinationParams creates a new StreamingServiceGetDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamingServiceGetDestinationParams() *StreamingServiceGetDestinationParams {
	return &StreamingServiceGetDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamingServiceGetDestinationParamsWithTimeout creates a new StreamingServiceGetDestinationParams object
// with the ability to set a timeout on a request.
func NewStreamingServiceGetDestinationParamsWithTimeout(timeout time.Duration) *StreamingServiceGetDestinationParams {
	return &StreamingServiceGetDestinationParams{
		timeout: timeout,
	}
}

// NewStreamingServiceGetDestinationParamsWithContext creates a new StreamingServiceGetDestinationParams object
// with the ability to set a context for a request.
func NewStreamingServiceGetDestinationParamsWithContext(ctx context.Context) *StreamingServiceGetDestinationParams {
	return &StreamingServiceGetDestinationParams{
		Context: ctx,
	}
}

// NewStreamingServiceGetDestinationParamsWithHTTPClient creates a new StreamingServiceGetDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamingServiceGetDestinationParamsWithHTTPClient(client *http.Client) *StreamingServiceGetDestinationParams {
	return &StreamingServiceGetDestinationParams{
		HTTPClient: client,
	}
}

/*
StreamingServiceGetDestinationParams contains all the parameters to send to the API endpoint

	for the streaming service get destination operation.

	Typically these are written to a http.Request.
*/
type StreamingServiceGetDestinationParams struct {

	/* DestinationID.

	   destination_id uniquely identifies the destination.
	*/
	DestinationID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the streaming service get destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceGetDestinationParams) WithDefaults() *StreamingServiceGetDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the streaming service get destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceGetDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) WithTimeout(timeout time.Duration) *StreamingServiceGetDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) WithContext(ctx context.Context) *StreamingServiceGetDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) WithHTTPClient(client *http.Client) *StreamingServiceGetDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) WithDestinationID(destinationID string) *StreamingServiceGetDestinationParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithOrganizationID adds the organizationID to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) WithOrganizationID(organizationID string) *StreamingServiceGetDestinationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the streaming service get destination params
func (o *StreamingServiceGetDestinationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *StreamingServiceGetDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param destination_id
	if err := r.SetPathParam("destination_id", o.DestinationID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}