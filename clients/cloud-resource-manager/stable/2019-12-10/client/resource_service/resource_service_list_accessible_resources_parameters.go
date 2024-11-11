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

// NewResourceServiceListAccessibleResourcesParams creates a new ResourceServiceListAccessibleResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceServiceListAccessibleResourcesParams() *ResourceServiceListAccessibleResourcesParams {
	return &ResourceServiceListAccessibleResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceServiceListAccessibleResourcesParamsWithTimeout creates a new ResourceServiceListAccessibleResourcesParams object
// with the ability to set a timeout on a request.
func NewResourceServiceListAccessibleResourcesParamsWithTimeout(timeout time.Duration) *ResourceServiceListAccessibleResourcesParams {
	return &ResourceServiceListAccessibleResourcesParams{
		timeout: timeout,
	}
}

// NewResourceServiceListAccessibleResourcesParamsWithContext creates a new ResourceServiceListAccessibleResourcesParams object
// with the ability to set a context for a request.
func NewResourceServiceListAccessibleResourcesParamsWithContext(ctx context.Context) *ResourceServiceListAccessibleResourcesParams {
	return &ResourceServiceListAccessibleResourcesParams{
		Context: ctx,
	}
}

// NewResourceServiceListAccessibleResourcesParamsWithHTTPClient creates a new ResourceServiceListAccessibleResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceServiceListAccessibleResourcesParamsWithHTTPClient(client *http.Client) *ResourceServiceListAccessibleResourcesParams {
	return &ResourceServiceListAccessibleResourcesParams{
		HTTPClient: client,
	}
}

/*
ResourceServiceListAccessibleResourcesParams contains all the parameters to send to the API endpoint

	for the resource service list accessible resources operation.

	Typically these are written to a http.Request.
*/
type ResourceServiceListAccessibleResourcesParams struct {

	/* LeafNameFilter.

	     LeafNameFilter optionally allows for filtering results based on a
	substring match of the final/leaf name within the resource's resource
	name.
	*/
	LeafNameFilter *string

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

	/* Permission.

	   Permission is the permission to check.
	*/
	Permission *string

	/* ResourceType.

	   ResourceType is the resource type of the resources to list.
	*/
	ResourceType *string

	/* ScopeResourceID.

	     ScopeResourceId is resource ID of the parent object to list resources under.
	Either this or the ResourceName needs to be provided.
	*/
	ScopeResourceID *string

	/* ScopeResourceName.

	     ScopeResourceName is resource name of the parent object to list resources under.
	Either this or the ResourceId needs to be provided.
	*/
	ScopeResourceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource service list accessible resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceListAccessibleResourcesParams) WithDefaults() *ResourceServiceListAccessibleResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource service list accessible resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceListAccessibleResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithTimeout(timeout time.Duration) *ResourceServiceListAccessibleResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithContext(ctx context.Context) *ResourceServiceListAccessibleResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithHTTPClient(client *http.Client) *ResourceServiceListAccessibleResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeafNameFilter adds the leafNameFilter to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithLeafNameFilter(leafNameFilter *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetLeafNameFilter(leafNameFilter)
	return o
}

// SetLeafNameFilter adds the leafNameFilter to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetLeafNameFilter(leafNameFilter *string) {
	o.LeafNameFilter = leafNameFilter
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithPaginationPageSize(paginationPageSize *int64) *ResourceServiceListAccessibleResourcesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPermission adds the permission to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithPermission(permission *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WithResourceType adds the resourceType to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithResourceType(resourceType *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithScopeResourceID adds the scopeResourceID to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithScopeResourceID(scopeResourceID *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetScopeResourceID(scopeResourceID)
	return o
}

// SetScopeResourceID adds the scopeResourceId to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetScopeResourceID(scopeResourceID *string) {
	o.ScopeResourceID = scopeResourceID
}

// WithScopeResourceName adds the scopeResourceName to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) WithScopeResourceName(scopeResourceName *string) *ResourceServiceListAccessibleResourcesParams {
	o.SetScopeResourceName(scopeResourceName)
	return o
}

// SetScopeResourceName adds the scopeResourceName to the resource service list accessible resources params
func (o *ResourceServiceListAccessibleResourcesParams) SetScopeResourceName(scopeResourceName *string) {
	o.ScopeResourceName = scopeResourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceServiceListAccessibleResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LeafNameFilter != nil {

		// query param leaf_name_filter
		var qrLeafNameFilter string

		if o.LeafNameFilter != nil {
			qrLeafNameFilter = *o.LeafNameFilter
		}
		qLeafNameFilter := qrLeafNameFilter
		if qLeafNameFilter != "" {

			if err := r.SetQueryParam("leaf_name_filter", qLeafNameFilter); err != nil {
				return err
			}
		}
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

	if o.Permission != nil {

		// query param permission
		var qrPermission string

		if o.Permission != nil {
			qrPermission = *o.Permission
		}
		qPermission := qrPermission
		if qPermission != "" {

			if err := r.SetQueryParam("permission", qPermission); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// query param resource_type
		var qrResourceType string

		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {

			if err := r.SetQueryParam("resource_type", qResourceType); err != nil {
				return err
			}
		}
	}

	if o.ScopeResourceID != nil {

		// query param scope_resource_id
		var qrScopeResourceID string

		if o.ScopeResourceID != nil {
			qrScopeResourceID = *o.ScopeResourceID
		}
		qScopeResourceID := qrScopeResourceID
		if qScopeResourceID != "" {

			if err := r.SetQueryParam("scope_resource_id", qScopeResourceID); err != nil {
				return err
			}
		}
	}

	if o.ScopeResourceName != nil {

		// query param scope_resource_name
		var qrScopeResourceName string

		if o.ScopeResourceName != nil {
			qrScopeResourceName = *o.ScopeResourceName
		}
		qScopeResourceName := qrScopeResourceName
		if qScopeResourceName != "" {

			if err := r.SetQueryParam("scope_resource_name", qScopeResourceName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}