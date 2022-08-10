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

// NewPackerServiceListBucketAncestryParams creates a new PackerServiceListBucketAncestryParams object
// with the default values initialized.
func NewPackerServiceListBucketAncestryParams() *PackerServiceListBucketAncestryParams {
	var ()
	return &PackerServiceListBucketAncestryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceListBucketAncestryParamsWithTimeout creates a new PackerServiceListBucketAncestryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceListBucketAncestryParamsWithTimeout(timeout time.Duration) *PackerServiceListBucketAncestryParams {
	var ()
	return &PackerServiceListBucketAncestryParams{

		timeout: timeout,
	}
}

// NewPackerServiceListBucketAncestryParamsWithContext creates a new PackerServiceListBucketAncestryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceListBucketAncestryParamsWithContext(ctx context.Context) *PackerServiceListBucketAncestryParams {
	var ()
	return &PackerServiceListBucketAncestryParams{

		Context: ctx,
	}
}

// NewPackerServiceListBucketAncestryParamsWithHTTPClient creates a new PackerServiceListBucketAncestryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceListBucketAncestryParamsWithHTTPClient(client *http.Client) *PackerServiceListBucketAncestryParams {
	var ()
	return &PackerServiceListBucketAncestryParams{
		HTTPClient: client,
	}
}

/*PackerServiceListBucketAncestryParams contains all the parameters to send to the API endpoint
for the packer service list bucket ancestry operation typically these are written to a http.Request
*/
type PackerServiceListBucketAncestryParams struct {

	/*BucketSlug
	  Human-readable name for the bucket you want to list ancestry for.

	*/
	BucketSlug string
	/*Channel
	  The channel of the bucket used to filter children.
	All listed children were built from iterations when assigned to this channel.
	By default, an unspecified channel will list all children based on any iteration of this bucket.

	*/
	Channel *string
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
	/*Type
	  The type of ancestry relations to list. Type 'parents' will list all the ancestral relations where this bucket is a child.
	Type 'children' will list all the ancestral relations where this bucket is a parent.
	By default, an unspecified type will list the ancestral relations where this bucket is either a parent or a child.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithTimeout(timeout time.Duration) *PackerServiceListBucketAncestryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithContext(ctx context.Context) *PackerServiceListBucketAncestryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithHTTPClient(client *http.Client) *PackerServiceListBucketAncestryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithBucketSlug(bucketSlug string) *PackerServiceListBucketAncestryParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithChannel adds the channel to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithChannel(channel *string) *PackerServiceListBucketAncestryParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetChannel(channel *string) {
	o.Channel = channel
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceListBucketAncestryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithLocationProjectID(locationProjectID string) *PackerServiceListBucketAncestryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceListBucketAncestryParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceListBucketAncestryParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithPaginationNextPageToken(paginationNextPageToken *string) *PackerServiceListBucketAncestryParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithPaginationPageSize(paginationPageSize *int64) *PackerServiceListBucketAncestryParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *PackerServiceListBucketAncestryParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithType adds the typeVar to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) WithType(typeVar *string) *PackerServiceListBucketAncestryParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the packer service list bucket ancestry params
func (o *PackerServiceListBucketAncestryParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceListBucketAncestryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
		return err
	}

	if o.Channel != nil {

		// query param channel
		var qrChannel string
		if o.Channel != nil {
			qrChannel = *o.Channel
		}
		qChannel := qrChannel
		if qChannel != "" {
			if err := r.SetQueryParam("channel", qChannel); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
