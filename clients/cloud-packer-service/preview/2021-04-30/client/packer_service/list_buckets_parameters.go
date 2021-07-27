// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

// NewListBucketsParams creates a new ListBucketsParams object
// with the default values initialized.
func NewListBucketsParams() *ListBucketsParams {
	var ()
	return &ListBucketsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBucketsParamsWithTimeout creates a new ListBucketsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBucketsParamsWithTimeout(timeout time.Duration) *ListBucketsParams {
	var ()
	return &ListBucketsParams{

		timeout: timeout,
	}
}

// NewListBucketsParamsWithContext creates a new ListBucketsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBucketsParamsWithContext(ctx context.Context) *ListBucketsParams {
	var ()
	return &ListBucketsParams{

		Context: ctx,
	}
}

// NewListBucketsParamsWithHTTPClient creates a new ListBucketsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBucketsParamsWithHTTPClient(client *http.Client) *ListBucketsParams {
	var ()
	return &ListBucketsParams{
		HTTPClient: client,
	}
}

/*ListBucketsParams contains all the parameters to send to the API endpoint
for the list buckets operation typically these are written to a http.Request
*/
type ListBucketsParams struct {

	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*LocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	LocationRegionProvider *string
	/*LocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	LocationRegionRegion *string
	/*PaginationNextPageToken
	  Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.

	*/
	PaginationNextPageToken *string
	/*PaginationPageSize
	  The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	*/
	PaginationPageSize *int64
	/*PaginationPreviousPageToken
	  Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.

	*/
	PaginationPreviousPageToken *string
	/*SortingOrderBy
	  Specifies the list of per field ordering that should be used for sorting.
	The order matters as rows are sorted in order by fields and when the field
	matches, the next field is used to tie break the ordering.
	The per field default ordering is ascending.

	The fields should be immutabile, unique, and orderable. If the field is
	not unique, more than one sort fields should be passed.

	Example: oder_by=name,age desc,created_at asc
	In that case, 'name' will get the default 'ascending' order.

	*/
	SortingOrderBy []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list buckets params
func (o *ListBucketsParams) WithTimeout(timeout time.Duration) *ListBucketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list buckets params
func (o *ListBucketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list buckets params
func (o *ListBucketsParams) WithContext(ctx context.Context) *ListBucketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list buckets params
func (o *ListBucketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list buckets params
func (o *ListBucketsParams) WithHTTPClient(client *http.Client) *ListBucketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list buckets params
func (o *ListBucketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the list buckets params
func (o *ListBucketsParams) WithLocationOrganizationID(locationOrganizationID string) *ListBucketsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list buckets params
func (o *ListBucketsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list buckets params
func (o *ListBucketsParams) WithLocationProjectID(locationProjectID string) *ListBucketsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list buckets params
func (o *ListBucketsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list buckets params
func (o *ListBucketsParams) WithLocationRegionProvider(locationRegionProvider *string) *ListBucketsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list buckets params
func (o *ListBucketsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list buckets params
func (o *ListBucketsParams) WithLocationRegionRegion(locationRegionRegion *string) *ListBucketsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list buckets params
func (o *ListBucketsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list buckets params
func (o *ListBucketsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListBucketsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list buckets params
func (o *ListBucketsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list buckets params
func (o *ListBucketsParams) WithPaginationPageSize(paginationPageSize *int64) *ListBucketsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list buckets params
func (o *ListBucketsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list buckets params
func (o *ListBucketsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListBucketsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list buckets params
func (o *ListBucketsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithSortingOrderBy adds the sortingOrderBy to the list buckets params
func (o *ListBucketsParams) WithSortingOrderBy(sortingOrderBy []string) *ListBucketsParams {
	o.SetSortingOrderBy(sortingOrderBy)
	return o
}

// SetSortingOrderBy adds the sortingOrderBy to the list buckets params
func (o *ListBucketsParams) SetSortingOrderBy(sortingOrderBy []string) {
	o.SortingOrderBy = sortingOrderBy
}

// WriteToRequest writes these params to a swagger request
func (o *ListBucketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	valuesSortingOrderBy := o.SortingOrderBy

	joinedSortingOrderBy := swag.JoinByFormat(valuesSortingOrderBy, "multi")
	// query array param sorting.order_by
	if err := r.SetQueryParam("sorting.order_by", joinedSortingOrderBy...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
