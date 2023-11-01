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

// NewWaypointListProjectsParams creates a new WaypointListProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListProjectsParams() *WaypointListProjectsParams {
	return &WaypointListProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListProjectsParamsWithTimeout creates a new WaypointListProjectsParams object
// with the ability to set a timeout on a request.
func NewWaypointListProjectsParamsWithTimeout(timeout time.Duration) *WaypointListProjectsParams {
	return &WaypointListProjectsParams{
		timeout: timeout,
	}
}

// NewWaypointListProjectsParamsWithContext creates a new WaypointListProjectsParams object
// with the ability to set a context for a request.
func NewWaypointListProjectsParamsWithContext(ctx context.Context) *WaypointListProjectsParams {
	return &WaypointListProjectsParams{
		Context: ctx,
	}
}

// NewWaypointListProjectsParamsWithHTTPClient creates a new WaypointListProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListProjectsParamsWithHTTPClient(client *http.Client) *WaypointListProjectsParams {
	return &WaypointListProjectsParams{
		HTTPClient: client,
	}
}

/*
WaypointListProjectsParams contains all the parameters to send to the API endpoint

	for the waypoint list projects operation.

	Typically these are written to a http.Request.
*/
type WaypointListProjectsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListProjectsParams) WithDefaults() *WaypointListProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list projects params
func (o *WaypointListProjectsParams) WithTimeout(timeout time.Duration) *WaypointListProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list projects params
func (o *WaypointListProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list projects params
func (o *WaypointListProjectsParams) WithContext(ctx context.Context) *WaypointListProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list projects params
func (o *WaypointListProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list projects params
func (o *WaypointListProjectsParams) WithHTTPClient(client *http.Client) *WaypointListProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list projects params
func (o *WaypointListProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint list projects params
func (o *WaypointListProjectsParams) WithNamespaceID(namespaceID string) *WaypointListProjectsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list projects params
func (o *WaypointListProjectsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint list projects params
func (o *WaypointListProjectsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointListProjectsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint list projects params
func (o *WaypointListProjectsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint list projects params
func (o *WaypointListProjectsParams) WithPaginationPageSize(paginationPageSize *int64) *WaypointListProjectsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint list projects params
func (o *WaypointListProjectsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list projects params
func (o *WaypointListProjectsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointListProjectsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list projects params
func (o *WaypointListProjectsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
