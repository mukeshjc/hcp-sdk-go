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
)

// NewWaypointUIGetDeployment2Params creates a new WaypointUIGetDeployment2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUIGetDeployment2Params() *WaypointUIGetDeployment2Params {
	return &WaypointUIGetDeployment2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIGetDeployment2ParamsWithTimeout creates a new WaypointUIGetDeployment2Params object
// with the ability to set a timeout on a request.
func NewWaypointUIGetDeployment2ParamsWithTimeout(timeout time.Duration) *WaypointUIGetDeployment2Params {
	return &WaypointUIGetDeployment2Params{
		timeout: timeout,
	}
}

// NewWaypointUIGetDeployment2ParamsWithContext creates a new WaypointUIGetDeployment2Params object
// with the ability to set a context for a request.
func NewWaypointUIGetDeployment2ParamsWithContext(ctx context.Context) *WaypointUIGetDeployment2Params {
	return &WaypointUIGetDeployment2Params{
		Context: ctx,
	}
}

// NewWaypointUIGetDeployment2ParamsWithHTTPClient creates a new WaypointUIGetDeployment2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUIGetDeployment2ParamsWithHTTPClient(client *http.Client) *WaypointUIGetDeployment2Params {
	return &WaypointUIGetDeployment2Params{
		HTTPClient: client,
	}
}

/*
WaypointUIGetDeployment2Params contains all the parameters to send to the API endpoint

	for the waypoint UI get deployment2 operation.

	Typically these are written to a http.Request.
*/
type WaypointUIGetDeployment2Params struct {

	/* LoadDetails.

	     Indicate if the fetched deployments should include additional information
	about each deployment.

	     Default: "NONE"
	*/
	LoadDetails *string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// RefID.
	RefID *string

	// RefSequenceApplicationApplication.
	RefSequenceApplicationApplication string

	// RefSequenceApplicationProject.
	RefSequenceApplicationProject string

	// RefSequenceNumber.
	//
	// Format: uint64
	RefSequenceNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint UI get deployment2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIGetDeployment2Params) WithDefaults() *WaypointUIGetDeployment2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint UI get deployment2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIGetDeployment2Params) SetDefaults() {
	var (
		loadDetailsDefault = string("NONE")
	)

	val := WaypointUIGetDeployment2Params{
		LoadDetails: &loadDetailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithTimeout(timeout time.Duration) *WaypointUIGetDeployment2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithContext(ctx context.Context) *WaypointUIGetDeployment2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithHTTPClient(client *http.Client) *WaypointUIGetDeployment2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoadDetails adds the loadDetails to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithLoadDetails(loadDetails *string) *WaypointUIGetDeployment2Params {
	o.SetLoadDetails(loadDetails)
	return o
}

// SetLoadDetails adds the loadDetails to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetLoadDetails(loadDetails *string) {
	o.LoadDetails = loadDetails
}

// WithNamespaceID adds the namespaceID to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithNamespaceID(namespaceID string) *WaypointUIGetDeployment2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithRefID(refID *string) *WaypointUIGetDeployment2Params {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetRefID(refID *string) {
	o.RefID = refID
}

// WithRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithRefSequenceApplicationApplication(refSequenceApplicationApplication string) *WaypointUIGetDeployment2Params {
	o.SetRefSequenceApplicationApplication(refSequenceApplicationApplication)
	return o
}

// SetRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetRefSequenceApplicationApplication(refSequenceApplicationApplication string) {
	o.RefSequenceApplicationApplication = refSequenceApplicationApplication
}

// WithRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithRefSequenceApplicationProject(refSequenceApplicationProject string) *WaypointUIGetDeployment2Params {
	o.SetRefSequenceApplicationProject(refSequenceApplicationProject)
	return o
}

// SetRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetRefSequenceApplicationProject(refSequenceApplicationProject string) {
	o.RefSequenceApplicationProject = refSequenceApplicationProject
}

// WithRefSequenceNumber adds the refSequenceNumber to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) WithRefSequenceNumber(refSequenceNumber string) *WaypointUIGetDeployment2Params {
	o.SetRefSequenceNumber(refSequenceNumber)
	return o
}

// SetRefSequenceNumber adds the refSequenceNumber to the waypoint UI get deployment2 params
func (o *WaypointUIGetDeployment2Params) SetRefSequenceNumber(refSequenceNumber string) {
	o.RefSequenceNumber = refSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIGetDeployment2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LoadDetails != nil {

		// query param load_details
		var qrLoadDetails string

		if o.LoadDetails != nil {
			qrLoadDetails = *o.LoadDetails
		}
		qLoadDetails := qrLoadDetails
		if qLoadDetails != "" {

			if err := r.SetQueryParam("load_details", qLoadDetails); err != nil {
				return err
			}
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.RefID != nil {

		// query param ref.id
		var qrRefID string

		if o.RefID != nil {
			qrRefID = *o.RefID
		}
		qRefID := qrRefID
		if qRefID != "" {

			if err := r.SetQueryParam("ref.id", qRefID); err != nil {
				return err
			}
		}
	}

	// path param ref.sequence.application.application
	if err := r.SetPathParam("ref.sequence.application.application", o.RefSequenceApplicationApplication); err != nil {
		return err
	}

	// path param ref.sequence.application.project
	if err := r.SetPathParam("ref.sequence.application.project", o.RefSequenceApplicationProject); err != nil {
		return err
	}

	// path param ref.sequence.number
	if err := r.SetPathParam("ref.sequence.number", o.RefSequenceNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
