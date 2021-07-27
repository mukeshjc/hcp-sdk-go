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

// NewListIterationsParams creates a new ListIterationsParams object
// with the default values initialized.
func NewListIterationsParams() *ListIterationsParams {
	var ()
	return &ListIterationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListIterationsParamsWithTimeout creates a new ListIterationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIterationsParamsWithTimeout(timeout time.Duration) *ListIterationsParams {
	var ()
	return &ListIterationsParams{

		timeout: timeout,
	}
}

// NewListIterationsParamsWithContext creates a new ListIterationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIterationsParamsWithContext(ctx context.Context) *ListIterationsParams {
	var ()
	return &ListIterationsParams{

		Context: ctx,
	}
}

// NewListIterationsParamsWithHTTPClient creates a new ListIterationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIterationsParamsWithHTTPClient(client *http.Client) *ListIterationsParams {
	var ()
	return &ListIterationsParams{
		HTTPClient: client,
	}
}

/*ListIterationsParams contains all the parameters to send to the API endpoint
for the list iterations operation typically these are written to a http.Request
*/
type ListIterationsParams struct {

	/*BucketSlug*/
	BucketSlug string
	/*IncludeIncomplete*/
	IncludeIncomplete *bool
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

// WithTimeout adds the timeout to the list iterations params
func (o *ListIterationsParams) WithTimeout(timeout time.Duration) *ListIterationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list iterations params
func (o *ListIterationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list iterations params
func (o *ListIterationsParams) WithContext(ctx context.Context) *ListIterationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list iterations params
func (o *ListIterationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list iterations params
func (o *ListIterationsParams) WithHTTPClient(client *http.Client) *ListIterationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list iterations params
func (o *ListIterationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the list iterations params
func (o *ListIterationsParams) WithBucketSlug(bucketSlug string) *ListIterationsParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the list iterations params
func (o *ListIterationsParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithIncludeIncomplete adds the includeIncomplete to the list iterations params
func (o *ListIterationsParams) WithIncludeIncomplete(includeIncomplete *bool) *ListIterationsParams {
	o.SetIncludeIncomplete(includeIncomplete)
	return o
}

// SetIncludeIncomplete adds the includeIncomplete to the list iterations params
func (o *ListIterationsParams) SetIncludeIncomplete(includeIncomplete *bool) {
	o.IncludeIncomplete = includeIncomplete
}

// WithLocationOrganizationID adds the locationOrganizationID to the list iterations params
func (o *ListIterationsParams) WithLocationOrganizationID(locationOrganizationID string) *ListIterationsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list iterations params
func (o *ListIterationsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list iterations params
func (o *ListIterationsParams) WithLocationProjectID(locationProjectID string) *ListIterationsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list iterations params
func (o *ListIterationsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list iterations params
func (o *ListIterationsParams) WithLocationRegionProvider(locationRegionProvider *string) *ListIterationsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list iterations params
func (o *ListIterationsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list iterations params
func (o *ListIterationsParams) WithLocationRegionRegion(locationRegionRegion *string) *ListIterationsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list iterations params
func (o *ListIterationsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list iterations params
func (o *ListIterationsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListIterationsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list iterations params
func (o *ListIterationsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list iterations params
func (o *ListIterationsParams) WithPaginationPageSize(paginationPageSize *int64) *ListIterationsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list iterations params
func (o *ListIterationsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list iterations params
func (o *ListIterationsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListIterationsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list iterations params
func (o *ListIterationsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithSortingOrderBy adds the sortingOrderBy to the list iterations params
func (o *ListIterationsParams) WithSortingOrderBy(sortingOrderBy []string) *ListIterationsParams {
	o.SetSortingOrderBy(sortingOrderBy)
	return o
}

// SetSortingOrderBy adds the sortingOrderBy to the list iterations params
func (o *ListIterationsParams) SetSortingOrderBy(sortingOrderBy []string) {
	o.SortingOrderBy = sortingOrderBy
}

// WriteToRequest writes these params to a swagger request
func (o *ListIterationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
		return err
	}

	if o.IncludeIncomplete != nil {

		// query param include_incomplete
		var qrIncludeIncomplete bool
		if o.IncludeIncomplete != nil {
			qrIncludeIncomplete = *o.IncludeIncomplete
		}
		qIncludeIncomplete := swag.FormatBool(qrIncludeIncomplete)
		if qIncludeIncomplete != "" {
			if err := r.SetQueryParam("include_incomplete", qIncludeIncomplete); err != nil {
				return err
			}
		}

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
