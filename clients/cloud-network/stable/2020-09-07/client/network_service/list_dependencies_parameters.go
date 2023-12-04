// Code generated by go-swagger; DO NOT EDIT.

package network_service

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

// NewListDependenciesParams creates a new ListDependenciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDependenciesParams() *ListDependenciesParams {
	return &ListDependenciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDependenciesParamsWithTimeout creates a new ListDependenciesParams object
// with the ability to set a timeout on a request.
func NewListDependenciesParamsWithTimeout(timeout time.Duration) *ListDependenciesParams {
	return &ListDependenciesParams{
		timeout: timeout,
	}
}

// NewListDependenciesParamsWithContext creates a new ListDependenciesParams object
// with the ability to set a context for a request.
func NewListDependenciesParamsWithContext(ctx context.Context) *ListDependenciesParams {
	return &ListDependenciesParams{
		Context: ctx,
	}
}

// NewListDependenciesParamsWithHTTPClient creates a new ListDependenciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDependenciesParamsWithHTTPClient(client *http.Client) *ListDependenciesParams {
	return &ListDependenciesParams{
		HTTPClient: client,
	}
}

/*
ListDependenciesParams contains all the parameters to send to the API endpoint

	for the list dependencies operation.

	Typically these are written to a http.Request.
*/
type ListDependenciesParams struct {

	/* ID.

	   id the id of the HVN we are listing dependencies for.
	*/
	ID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

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

// WithDefaults hydrates default values in the list dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDependenciesParams) WithDefaults() *ListDependenciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDependenciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list dependencies params
func (o *ListDependenciesParams) WithTimeout(timeout time.Duration) *ListDependenciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dependencies params
func (o *ListDependenciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dependencies params
func (o *ListDependenciesParams) WithContext(ctx context.Context) *ListDependenciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dependencies params
func (o *ListDependenciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dependencies params
func (o *ListDependenciesParams) WithHTTPClient(client *http.Client) *ListDependenciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dependencies params
func (o *ListDependenciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list dependencies params
func (o *ListDependenciesParams) WithID(id string) *ListDependenciesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list dependencies params
func (o *ListDependenciesParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the list dependencies params
func (o *ListDependenciesParams) WithLocationOrganizationID(locationOrganizationID string) *ListDependenciesParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list dependencies params
func (o *ListDependenciesParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list dependencies params
func (o *ListDependenciesParams) WithLocationProjectID(locationProjectID string) *ListDependenciesParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list dependencies params
func (o *ListDependenciesParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list dependencies params
func (o *ListDependenciesParams) WithLocationRegionProvider(locationRegionProvider *string) *ListDependenciesParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list dependencies params
func (o *ListDependenciesParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list dependencies params
func (o *ListDependenciesParams) WithLocationRegionRegion(locationRegionRegion *string) *ListDependenciesParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list dependencies params
func (o *ListDependenciesParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list dependencies params
func (o *ListDependenciesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListDependenciesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list dependencies params
func (o *ListDependenciesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list dependencies params
func (o *ListDependenciesParams) WithPaginationPageSize(paginationPageSize *int64) *ListDependenciesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list dependencies params
func (o *ListDependenciesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list dependencies params
func (o *ListDependenciesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListDependenciesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list dependencies params
func (o *ListDependenciesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListDependenciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
