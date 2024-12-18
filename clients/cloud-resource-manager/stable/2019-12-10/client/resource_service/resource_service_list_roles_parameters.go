// Code generated by go-swagger; DO NOT EDIT.

package resource_service

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

// NewResourceServiceListRolesParams creates a new ResourceServiceListRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceServiceListRolesParams() *ResourceServiceListRolesParams {
	return &ResourceServiceListRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceServiceListRolesParamsWithTimeout creates a new ResourceServiceListRolesParams object
// with the ability to set a timeout on a request.
func NewResourceServiceListRolesParamsWithTimeout(timeout time.Duration) *ResourceServiceListRolesParams {
	return &ResourceServiceListRolesParams{
		timeout: timeout,
	}
}

// NewResourceServiceListRolesParamsWithContext creates a new ResourceServiceListRolesParams object
// with the ability to set a context for a request.
func NewResourceServiceListRolesParamsWithContext(ctx context.Context) *ResourceServiceListRolesParams {
	return &ResourceServiceListRolesParams{
		Context: ctx,
	}
}

// NewResourceServiceListRolesParamsWithHTTPClient creates a new ResourceServiceListRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceServiceListRolesParamsWithHTTPClient(client *http.Client) *ResourceServiceListRolesParams {
	return &ResourceServiceListRolesParams{
		HTTPClient: client,
	}
}

/*
ResourceServiceListRolesParams contains all the parameters to send to the API endpoint

	for the resource service list roles operation.

	Typically these are written to a http.Request.
*/
type ResourceServiceListRolesParams struct {

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

	/* ResourceID.

	   Id is the resource id of the resource.
	*/
	ResourceID *string

	/* ResourceName.

	   Name is the resource name of the resource.
	*/
	ResourceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource service list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceListRolesParams) WithDefaults() *ResourceServiceListRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource service list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceListRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithTimeout(timeout time.Duration) *ResourceServiceListRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithContext(ctx context.Context) *ResourceServiceListRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithHTTPClient(client *http.Client) *ResourceServiceListRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ResourceServiceListRolesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithPaginationPageSize(paginationPageSize *int64) *ResourceServiceListRolesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ResourceServiceListRolesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithResourceID adds the resourceID to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithResourceID(resourceID *string) *ResourceServiceListRolesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceName adds the resourceName to the resource service list roles params
func (o *ResourceServiceListRolesParams) WithResourceName(resourceName *string) *ResourceServiceListRolesParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the resource service list roles params
func (o *ResourceServiceListRolesParams) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceServiceListRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ResourceID != nil {

		// query param resource.id
		var qrResourceID string

		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {

			if err := r.SetQueryParam("resource.id", qResourceID); err != nil {
				return err
			}
		}
	}

	if o.ResourceName != nil {

		// query param resource.name
		var qrResourceName string

		if o.ResourceName != nil {
			qrResourceName = *o.ResourceName
		}
		qResourceName := qrResourceName
		if qResourceName != "" {

			if err := r.SetQueryParam("resource.name", qResourceName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
