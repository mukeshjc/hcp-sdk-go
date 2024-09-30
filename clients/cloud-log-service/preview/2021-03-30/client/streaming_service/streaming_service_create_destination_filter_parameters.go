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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// NewStreamingServiceCreateDestinationFilterParams creates a new StreamingServiceCreateDestinationFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamingServiceCreateDestinationFilterParams() *StreamingServiceCreateDestinationFilterParams {
	return &StreamingServiceCreateDestinationFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamingServiceCreateDestinationFilterParamsWithTimeout creates a new StreamingServiceCreateDestinationFilterParams object
// with the ability to set a timeout on a request.
func NewStreamingServiceCreateDestinationFilterParamsWithTimeout(timeout time.Duration) *StreamingServiceCreateDestinationFilterParams {
	return &StreamingServiceCreateDestinationFilterParams{
		timeout: timeout,
	}
}

// NewStreamingServiceCreateDestinationFilterParamsWithContext creates a new StreamingServiceCreateDestinationFilterParams object
// with the ability to set a context for a request.
func NewStreamingServiceCreateDestinationFilterParamsWithContext(ctx context.Context) *StreamingServiceCreateDestinationFilterParams {
	return &StreamingServiceCreateDestinationFilterParams{
		Context: ctx,
	}
}

// NewStreamingServiceCreateDestinationFilterParamsWithHTTPClient creates a new StreamingServiceCreateDestinationFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamingServiceCreateDestinationFilterParamsWithHTTPClient(client *http.Client) *StreamingServiceCreateDestinationFilterParams {
	return &StreamingServiceCreateDestinationFilterParams{
		HTTPClient: client,
	}
}

/*
StreamingServiceCreateDestinationFilterParams contains all the parameters to send to the API endpoint

	for the streaming service create destination filter operation.

	Typically these are written to a http.Request.
*/
type StreamingServiceCreateDestinationFilterParams struct {

	// Body.
	Body *models.LogService20210330CreateDestinationFilterRequest

	/* DestinationID.

	   destination_id is the destination associated with the filter
	*/
	DestinationID string

	/* OrganizationID.

	   organization_id is the organization holding the destination and filter
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the streaming service create destination filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceCreateDestinationFilterParams) WithDefaults() *StreamingServiceCreateDestinationFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the streaming service create destination filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceCreateDestinationFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithTimeout(timeout time.Duration) *StreamingServiceCreateDestinationFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithContext(ctx context.Context) *StreamingServiceCreateDestinationFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithHTTPClient(client *http.Client) *StreamingServiceCreateDestinationFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithBody(body *models.LogService20210330CreateDestinationFilterRequest) *StreamingServiceCreateDestinationFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetBody(body *models.LogService20210330CreateDestinationFilterRequest) {
	o.Body = body
}

// WithDestinationID adds the destinationID to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithDestinationID(destinationID string) *StreamingServiceCreateDestinationFilterParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithOrganizationID adds the organizationID to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) WithOrganizationID(organizationID string) *StreamingServiceCreateDestinationFilterParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the streaming service create destination filter params
func (o *StreamingServiceCreateDestinationFilterParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *StreamingServiceCreateDestinationFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
