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

// NewWaypointExpediteStatusReport3Params creates a new WaypointExpediteStatusReport3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointExpediteStatusReport3Params() *WaypointExpediteStatusReport3Params {
	return &WaypointExpediteStatusReport3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointExpediteStatusReport3ParamsWithTimeout creates a new WaypointExpediteStatusReport3Params object
// with the ability to set a timeout on a request.
func NewWaypointExpediteStatusReport3ParamsWithTimeout(timeout time.Duration) *WaypointExpediteStatusReport3Params {
	return &WaypointExpediteStatusReport3Params{
		timeout: timeout,
	}
}

// NewWaypointExpediteStatusReport3ParamsWithContext creates a new WaypointExpediteStatusReport3Params object
// with the ability to set a context for a request.
func NewWaypointExpediteStatusReport3ParamsWithContext(ctx context.Context) *WaypointExpediteStatusReport3Params {
	return &WaypointExpediteStatusReport3Params{
		Context: ctx,
	}
}

// NewWaypointExpediteStatusReport3ParamsWithHTTPClient creates a new WaypointExpediteStatusReport3Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointExpediteStatusReport3ParamsWithHTTPClient(client *http.Client) *WaypointExpediteStatusReport3Params {
	return &WaypointExpediteStatusReport3Params{
		HTTPClient: client,
	}
}

/*
WaypointExpediteStatusReport3Params contains all the parameters to send to the API endpoint

	for the waypoint expedite status report3 operation.

	Typically these are written to a http.Request.
*/
type WaypointExpediteStatusReport3Params struct {

	// Body.
	Body *models.HashicorpWaypointExpediteStatusReportRequest

	// DeploymentSequenceApplicationApplication.
	DeploymentSequenceApplicationApplication string

	// DeploymentSequenceApplicationProject.
	DeploymentSequenceApplicationProject string

	// DeploymentSequenceNumber.
	//
	// Format: uint64
	DeploymentSequenceNumber string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint expedite status report3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointExpediteStatusReport3Params) WithDefaults() *WaypointExpediteStatusReport3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint expedite status report3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointExpediteStatusReport3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithTimeout(timeout time.Duration) *WaypointExpediteStatusReport3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithContext(ctx context.Context) *WaypointExpediteStatusReport3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithHTTPClient(client *http.Client) *WaypointExpediteStatusReport3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithBody(body *models.HashicorpWaypointExpediteStatusReportRequest) *WaypointExpediteStatusReport3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetBody(body *models.HashicorpWaypointExpediteStatusReportRequest) {
	o.Body = body
}

// WithDeploymentSequenceApplicationApplication adds the deploymentSequenceApplicationApplication to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithDeploymentSequenceApplicationApplication(deploymentSequenceApplicationApplication string) *WaypointExpediteStatusReport3Params {
	o.SetDeploymentSequenceApplicationApplication(deploymentSequenceApplicationApplication)
	return o
}

// SetDeploymentSequenceApplicationApplication adds the deploymentSequenceApplicationApplication to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetDeploymentSequenceApplicationApplication(deploymentSequenceApplicationApplication string) {
	o.DeploymentSequenceApplicationApplication = deploymentSequenceApplicationApplication
}

// WithDeploymentSequenceApplicationProject adds the deploymentSequenceApplicationProject to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithDeploymentSequenceApplicationProject(deploymentSequenceApplicationProject string) *WaypointExpediteStatusReport3Params {
	o.SetDeploymentSequenceApplicationProject(deploymentSequenceApplicationProject)
	return o
}

// SetDeploymentSequenceApplicationProject adds the deploymentSequenceApplicationProject to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetDeploymentSequenceApplicationProject(deploymentSequenceApplicationProject string) {
	o.DeploymentSequenceApplicationProject = deploymentSequenceApplicationProject
}

// WithDeploymentSequenceNumber adds the deploymentSequenceNumber to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithDeploymentSequenceNumber(deploymentSequenceNumber string) *WaypointExpediteStatusReport3Params {
	o.SetDeploymentSequenceNumber(deploymentSequenceNumber)
	return o
}

// SetDeploymentSequenceNumber adds the deploymentSequenceNumber to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetDeploymentSequenceNumber(deploymentSequenceNumber string) {
	o.DeploymentSequenceNumber = deploymentSequenceNumber
}

// WithNamespaceID adds the namespaceID to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) WithNamespaceID(namespaceID string) *WaypointExpediteStatusReport3Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint expedite status report3 params
func (o *WaypointExpediteStatusReport3Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointExpediteStatusReport3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment.sequence.application.application
	if err := r.SetPathParam("deployment.sequence.application.application", o.DeploymentSequenceApplicationApplication); err != nil {
		return err
	}

	// path param deployment.sequence.application.project
	if err := r.SetPathParam("deployment.sequence.application.project", o.DeploymentSequenceApplicationProject); err != nil {
		return err
	}

	// path param deployment.sequence.number
	if err := r.SetPathParam("deployment.sequence.number", o.DeploymentSequenceNumber); err != nil {
		return err
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