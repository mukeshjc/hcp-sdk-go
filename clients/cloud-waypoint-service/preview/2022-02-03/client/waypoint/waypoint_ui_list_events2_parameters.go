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

// NewWaypointUIListEvents2Params creates a new WaypointUIListEvents2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUIListEvents2Params() *WaypointUIListEvents2Params {
	return &WaypointUIListEvents2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIListEvents2ParamsWithTimeout creates a new WaypointUIListEvents2Params object
// with the ability to set a timeout on a request.
func NewWaypointUIListEvents2ParamsWithTimeout(timeout time.Duration) *WaypointUIListEvents2Params {
	return &WaypointUIListEvents2Params{
		timeout: timeout,
	}
}

// NewWaypointUIListEvents2ParamsWithContext creates a new WaypointUIListEvents2Params object
// with the ability to set a context for a request.
func NewWaypointUIListEvents2ParamsWithContext(ctx context.Context) *WaypointUIListEvents2Params {
	return &WaypointUIListEvents2Params{
		Context: ctx,
	}
}

// NewWaypointUIListEvents2ParamsWithHTTPClient creates a new WaypointUIListEvents2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUIListEvents2ParamsWithHTTPClient(client *http.Client) *WaypointUIListEvents2Params {
	return &WaypointUIListEvents2Params{
		HTTPClient: client,
	}
}

/*
WaypointUIListEvents2Params contains all the parameters to send to the API endpoint

	for the waypoint UI list events2 operation.

	Typically these are written to a http.Request.
*/
type WaypointUIListEvents2Params struct {

	// ApplicationApplication.
	ApplicationApplication string

	// ApplicationProject.
	ApplicationProject string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	// ProjectProject.
	ProjectProject *string

	/* SortingOrderBy.

	     Specifies the list of per field ordering that should be used for sorting.
	The order matters as rows are sorted in order by fields and when the field
	matches, the next field is used to tie break the ordering.
	The per field default ordering is ascending.

	The fields should be immutable, unique, and orderable. If the field is
	not unique, more than one sort fields should be passed.

	Example: order_by=name,age desc,created_at asc
	In that case, 'name' will get the default 'ascending' order.
	*/
	SortingOrderBy []string

	// WorkspaceWorkspace.
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint UI list events2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListEvents2Params) WithDefaults() *WaypointUIListEvents2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint UI list events2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListEvents2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithTimeout(timeout time.Duration) *WaypointUIListEvents2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithContext(ctx context.Context) *WaypointUIListEvents2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithHTTPClient(client *http.Client) *WaypointUIListEvents2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithApplicationApplication(applicationApplication string) *WaypointUIListEvents2Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithApplicationProject(applicationProject string) *WaypointUIListEvents2Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithNamespaceID adds the namespaceID to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithNamespaceID(namespaceID string) *WaypointUIListEvents2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointUIListEvents2Params {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithPaginationPageSize(paginationPageSize *int64) *WaypointUIListEvents2Params {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointUIListEvents2Params {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithProjectProject adds the projectProject to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithProjectProject(projectProject *string) *WaypointUIListEvents2Params {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithSortingOrderBy adds the sortingOrderBy to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithSortingOrderBy(sortingOrderBy []string) *WaypointUIListEvents2Params {
	o.SetSortingOrderBy(sortingOrderBy)
	return o
}

// SetSortingOrderBy adds the sortingOrderBy to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetSortingOrderBy(sortingOrderBy []string) {
	o.SortingOrderBy = sortingOrderBy
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointUIListEvents2Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list events2 params
func (o *WaypointUIListEvents2Params) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIListEvents2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.ProjectProject != nil {

		// query param project.project
		var qrProjectProject string

		if o.ProjectProject != nil {
			qrProjectProject = *o.ProjectProject
		}
		qProjectProject := qrProjectProject
		if qProjectProject != "" {

			if err := r.SetQueryParam("project.project", qProjectProject); err != nil {
				return err
			}
		}
	}

	if o.SortingOrderBy != nil {

		// binding items for sorting.order_by
		joinedSortingOrderBy := o.bindParamSortingOrderBy(reg)

		// query array param sorting.order_by
		if err := r.SetQueryParam("sorting.order_by", joinedSortingOrderBy...); err != nil {
			return err
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

// bindParamWaypointUIListEvents2 binds the parameter sorting.order_by
func (o *WaypointUIListEvents2Params) bindParamSortingOrderBy(formats strfmt.Registry) []string {
	sortingOrderByIR := o.SortingOrderBy

	var sortingOrderByIC []string
	for _, sortingOrderByIIR := range sortingOrderByIR { // explode []string

		sortingOrderByIIV := sortingOrderByIIR // string as string
		sortingOrderByIC = append(sortingOrderByIC, sortingOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	sortingOrderByIS := swag.JoinByFormat(sortingOrderByIC, "multi")

	return sortingOrderByIS
}