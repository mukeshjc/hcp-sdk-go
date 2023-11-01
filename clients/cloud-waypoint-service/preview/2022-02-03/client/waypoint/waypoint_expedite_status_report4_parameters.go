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

// NewWaypointExpediteStatusReport4Params creates a new WaypointExpediteStatusReport4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointExpediteStatusReport4Params() *WaypointExpediteStatusReport4Params {
	return &WaypointExpediteStatusReport4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointExpediteStatusReport4ParamsWithTimeout creates a new WaypointExpediteStatusReport4Params object
// with the ability to set a timeout on a request.
func NewWaypointExpediteStatusReport4ParamsWithTimeout(timeout time.Duration) *WaypointExpediteStatusReport4Params {
	return &WaypointExpediteStatusReport4Params{
		timeout: timeout,
	}
}

// NewWaypointExpediteStatusReport4ParamsWithContext creates a new WaypointExpediteStatusReport4Params object
// with the ability to set a context for a request.
func NewWaypointExpediteStatusReport4ParamsWithContext(ctx context.Context) *WaypointExpediteStatusReport4Params {
	return &WaypointExpediteStatusReport4Params{
		Context: ctx,
	}
}

// NewWaypointExpediteStatusReport4ParamsWithHTTPClient creates a new WaypointExpediteStatusReport4Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointExpediteStatusReport4ParamsWithHTTPClient(client *http.Client) *WaypointExpediteStatusReport4Params {
	return &WaypointExpediteStatusReport4Params{
		HTTPClient: client,
	}
}

/*
WaypointExpediteStatusReport4Params contains all the parameters to send to the API endpoint

	for the waypoint expedite status report4 operation.

	Typically these are written to a http.Request.
*/
type WaypointExpediteStatusReport4Params struct {

	// Body.
	Body *models.HashicorpWaypointExpediteStatusReportRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// ReleaseSequenceApplicationApplication.
	ReleaseSequenceApplicationApplication string

	// ReleaseSequenceApplicationProject.
	ReleaseSequenceApplicationProject string

	// ReleaseSequenceNumber.
	//
	// Format: uint64
	ReleaseSequenceNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint expedite status report4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointExpediteStatusReport4Params) WithDefaults() *WaypointExpediteStatusReport4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint expedite status report4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointExpediteStatusReport4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithTimeout(timeout time.Duration) *WaypointExpediteStatusReport4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithContext(ctx context.Context) *WaypointExpediteStatusReport4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithHTTPClient(client *http.Client) *WaypointExpediteStatusReport4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithBody(body *models.HashicorpWaypointExpediteStatusReportRequest) *WaypointExpediteStatusReport4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetBody(body *models.HashicorpWaypointExpediteStatusReportRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithNamespaceID(namespaceID string) *WaypointExpediteStatusReport4Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithReleaseSequenceApplicationApplication adds the releaseSequenceApplicationApplication to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithReleaseSequenceApplicationApplication(releaseSequenceApplicationApplication string) *WaypointExpediteStatusReport4Params {
	o.SetReleaseSequenceApplicationApplication(releaseSequenceApplicationApplication)
	return o
}

// SetReleaseSequenceApplicationApplication adds the releaseSequenceApplicationApplication to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetReleaseSequenceApplicationApplication(releaseSequenceApplicationApplication string) {
	o.ReleaseSequenceApplicationApplication = releaseSequenceApplicationApplication
}

// WithReleaseSequenceApplicationProject adds the releaseSequenceApplicationProject to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithReleaseSequenceApplicationProject(releaseSequenceApplicationProject string) *WaypointExpediteStatusReport4Params {
	o.SetReleaseSequenceApplicationProject(releaseSequenceApplicationProject)
	return o
}

// SetReleaseSequenceApplicationProject adds the releaseSequenceApplicationProject to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetReleaseSequenceApplicationProject(releaseSequenceApplicationProject string) {
	o.ReleaseSequenceApplicationProject = releaseSequenceApplicationProject
}

// WithReleaseSequenceNumber adds the releaseSequenceNumber to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) WithReleaseSequenceNumber(releaseSequenceNumber string) *WaypointExpediteStatusReport4Params {
	o.SetReleaseSequenceNumber(releaseSequenceNumber)
	return o
}

// SetReleaseSequenceNumber adds the releaseSequenceNumber to the waypoint expedite status report4 params
func (o *WaypointExpediteStatusReport4Params) SetReleaseSequenceNumber(releaseSequenceNumber string) {
	o.ReleaseSequenceNumber = releaseSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointExpediteStatusReport4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param release.sequence.application.application
	if err := r.SetPathParam("release.sequence.application.application", o.ReleaseSequenceApplicationApplication); err != nil {
		return err
	}

	// path param release.sequence.application.project
	if err := r.SetPathParam("release.sequence.application.project", o.ReleaseSequenceApplicationProject); err != nil {
		return err
	}

	// path param release.sequence.number
	if err := r.SetPathParam("release.sequence.number", o.ReleaseSequenceNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
