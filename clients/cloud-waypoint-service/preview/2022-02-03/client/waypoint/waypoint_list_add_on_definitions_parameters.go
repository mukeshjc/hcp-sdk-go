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

// NewWaypointListAddOnDefinitionsParams creates a new WaypointListAddOnDefinitionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListAddOnDefinitionsParams() *WaypointListAddOnDefinitionsParams {
	return &WaypointListAddOnDefinitionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListAddOnDefinitionsParamsWithTimeout creates a new WaypointListAddOnDefinitionsParams object
// with the ability to set a timeout on a request.
func NewWaypointListAddOnDefinitionsParamsWithTimeout(timeout time.Duration) *WaypointListAddOnDefinitionsParams {
	return &WaypointListAddOnDefinitionsParams{
		timeout: timeout,
	}
}

// NewWaypointListAddOnDefinitionsParamsWithContext creates a new WaypointListAddOnDefinitionsParams object
// with the ability to set a context for a request.
func NewWaypointListAddOnDefinitionsParamsWithContext(ctx context.Context) *WaypointListAddOnDefinitionsParams {
	return &WaypointListAddOnDefinitionsParams{
		Context: ctx,
	}
}

// NewWaypointListAddOnDefinitionsParamsWithHTTPClient creates a new WaypointListAddOnDefinitionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListAddOnDefinitionsParamsWithHTTPClient(client *http.Client) *WaypointListAddOnDefinitionsParams {
	return &WaypointListAddOnDefinitionsParams{
		HTTPClient: client,
	}
}

/*
WaypointListAddOnDefinitionsParams contains all the parameters to send to the API endpoint

	for the waypoint list add on definitions operation.

	Typically these are written to a http.Request.
*/
type WaypointListAddOnDefinitionsParams struct {

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

// WithDefaults hydrates default values in the waypoint list add on definitions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListAddOnDefinitionsParams) WithDefaults() *WaypointListAddOnDefinitionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list add on definitions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListAddOnDefinitionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithTimeout(timeout time.Duration) *WaypointListAddOnDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithContext(ctx context.Context) *WaypointListAddOnDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithHTTPClient(client *http.Client) *WaypointListAddOnDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithNamespaceID(namespaceID string) *WaypointListAddOnDefinitionsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointListAddOnDefinitionsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithPaginationPageSize(paginationPageSize *int64) *WaypointListAddOnDefinitionsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointListAddOnDefinitionsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list add on definitions params
func (o *WaypointListAddOnDefinitionsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListAddOnDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
