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
	"github.com/go-openapi/swag"
)

// NewWaypointListReleases2Params creates a new WaypointListReleases2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListReleases2Params() *WaypointListReleases2Params {
	return &WaypointListReleases2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListReleases2ParamsWithTimeout creates a new WaypointListReleases2Params object
// with the ability to set a timeout on a request.
func NewWaypointListReleases2ParamsWithTimeout(timeout time.Duration) *WaypointListReleases2Params {
	return &WaypointListReleases2Params{
		timeout: timeout,
	}
}

// NewWaypointListReleases2ParamsWithContext creates a new WaypointListReleases2Params object
// with the ability to set a context for a request.
func NewWaypointListReleases2ParamsWithContext(ctx context.Context) *WaypointListReleases2Params {
	return &WaypointListReleases2Params{
		Context: ctx,
	}
}

// NewWaypointListReleases2ParamsWithHTTPClient creates a new WaypointListReleases2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListReleases2ParamsWithHTTPClient(client *http.Client) *WaypointListReleases2Params {
	return &WaypointListReleases2Params{
		HTTPClient: client,
	}
}

/*
WaypointListReleases2Params contains all the parameters to send to the API endpoint

	for the waypoint list releases2 operation.

	Typically these are written to a http.Request.
*/
type WaypointListReleases2Params struct {

	// ApplicationApplication.
	ApplicationApplication string

	// ApplicationProject.
	ApplicationProject string

	/* LoadDetails.

	     Load additional details about the release. These will become available
	in the Preload section.

	     Default: "NONE"
	*/
	LoadDetails *string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// OrderDesc.
	OrderDesc *bool

	/* OrderLimit.

	   Limit the number of results.

	   Format: int64
	*/
	OrderLimit *int64

	/* OrderOrder.

	   Order for the results.

	   Default: "UNSET"
	*/
	OrderOrder *string

	/* PhysicalState.

	     The physical state to filter for. If this is zero or unset then no
	filtering on physical state will be done.

	     Default: "UNKNOWN"
	*/
	PhysicalState *string

	// WorkspaceWorkspace.
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list releases2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListReleases2Params) WithDefaults() *WaypointListReleases2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list releases2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListReleases2Params) SetDefaults() {
	var (
		loadDetailsDefault = string("NONE")

		orderOrderDefault = string("UNSET")

		physicalStateDefault = string("UNKNOWN")
	)

	val := WaypointListReleases2Params{
		LoadDetails:   &loadDetailsDefault,
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithTimeout(timeout time.Duration) *WaypointListReleases2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithContext(ctx context.Context) *WaypointListReleases2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithHTTPClient(client *http.Client) *WaypointListReleases2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithApplicationApplication(applicationApplication string) *WaypointListReleases2Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithApplicationProject(applicationProject string) *WaypointListReleases2Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithLoadDetails adds the loadDetails to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithLoadDetails(loadDetails *string) *WaypointListReleases2Params {
	o.SetLoadDetails(loadDetails)
	return o
}

// SetLoadDetails adds the loadDetails to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetLoadDetails(loadDetails *string) {
	o.LoadDetails = loadDetails
}

// WithNamespaceID adds the namespaceID to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithNamespaceID(namespaceID string) *WaypointListReleases2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithOrderDesc adds the orderDesc to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithOrderDesc(orderDesc *bool) *WaypointListReleases2Params {
	o.SetOrderDesc(orderDesc)
	return o
}

// SetOrderDesc adds the orderDesc to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetOrderDesc(orderDesc *bool) {
	o.OrderDesc = orderDesc
}

// WithOrderLimit adds the orderLimit to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithOrderLimit(orderLimit *int64) *WaypointListReleases2Params {
	o.SetOrderLimit(orderLimit)
	return o
}

// SetOrderLimit adds the orderLimit to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetOrderLimit(orderLimit *int64) {
	o.OrderLimit = orderLimit
}

// WithOrderOrder adds the orderOrder to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithOrderOrder(orderOrder *string) *WaypointListReleases2Params {
	o.SetOrderOrder(orderOrder)
	return o
}

// SetOrderOrder adds the orderOrder to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetOrderOrder(orderOrder *string) {
	o.OrderOrder = orderOrder
}

// WithPhysicalState adds the physicalState to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithPhysicalState(physicalState *string) *WaypointListReleases2Params {
	o.SetPhysicalState(physicalState)
	return o
}

// SetPhysicalState adds the physicalState to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetPhysicalState(physicalState *string) {
	o.PhysicalState = physicalState
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list releases2 params
func (o *WaypointListReleases2Params) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointListReleases2Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list releases2 params
func (o *WaypointListReleases2Params) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListReleases2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.application
	if err := r.SetPathParam("application.application", o.ApplicationApplication); err != nil {
		return err
	}

	// path param application.project
	if err := r.SetPathParam("application.project", o.ApplicationProject); err != nil {
		return err
	}

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

	if o.OrderDesc != nil {

		// query param order.desc
		var qrOrderDesc bool

		if o.OrderDesc != nil {
			qrOrderDesc = *o.OrderDesc
		}
		qOrderDesc := swag.FormatBool(qrOrderDesc)
		if qOrderDesc != "" {

			if err := r.SetQueryParam("order.desc", qOrderDesc); err != nil {
				return err
			}
		}
	}

	if o.OrderLimit != nil {

		// query param order.limit
		var qrOrderLimit int64

		if o.OrderLimit != nil {
			qrOrderLimit = *o.OrderLimit
		}
		qOrderLimit := swag.FormatInt64(qrOrderLimit)
		if qOrderLimit != "" {

			if err := r.SetQueryParam("order.limit", qOrderLimit); err != nil {
				return err
			}
		}
	}

	if o.OrderOrder != nil {

		// query param order.order
		var qrOrderOrder string

		if o.OrderOrder != nil {
			qrOrderOrder = *o.OrderOrder
		}
		qOrderOrder := qrOrderOrder
		if qOrderOrder != "" {

			if err := r.SetQueryParam("order.order", qOrderOrder); err != nil {
				return err
			}
		}
	}

	if o.PhysicalState != nil {

		// query param physical_state
		var qrPhysicalState string

		if o.PhysicalState != nil {
			qrPhysicalState = *o.PhysicalState
		}
		qPhysicalState := qrPhysicalState
		if qPhysicalState != "" {

			if err := r.SetQueryParam("physical_state", qPhysicalState); err != nil {
				return err
			}
		}
	}

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
