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
)

// NewPackerServiceUpdateChannelParams creates a new PackerServiceUpdateChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceUpdateChannelParams() *PackerServiceUpdateChannelParams {
	return &PackerServiceUpdateChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceUpdateChannelParamsWithTimeout creates a new PackerServiceUpdateChannelParams object
// with the ability to set a timeout on a request.
func NewPackerServiceUpdateChannelParamsWithTimeout(timeout time.Duration) *PackerServiceUpdateChannelParams {
	return &PackerServiceUpdateChannelParams{
		timeout: timeout,
	}
}

// NewPackerServiceUpdateChannelParamsWithContext creates a new PackerServiceUpdateChannelParams object
// with the ability to set a context for a request.
func NewPackerServiceUpdateChannelParamsWithContext(ctx context.Context) *PackerServiceUpdateChannelParams {
	return &PackerServiceUpdateChannelParams{
		Context: ctx,
	}
}

// NewPackerServiceUpdateChannelParamsWithHTTPClient creates a new PackerServiceUpdateChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceUpdateChannelParamsWithHTTPClient(client *http.Client) *PackerServiceUpdateChannelParams {
	return &PackerServiceUpdateChannelParams{
		HTTPClient: client,
	}
}

/*
PackerServiceUpdateChannelParams contains all the parameters to send to the API endpoint

	for the packer service update channel operation.

	Typically these are written to a http.Request.
*/
type PackerServiceUpdateChannelParams struct {

	// Body.
	Body PackerServiceUpdateChannelBody

	/* BucketSlug.

	   Human-readable name for the bucket that the channel is associated with.
	*/
	BucketSlug string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* Slug.

	   Human-readable name for the channel.
	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service update channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateChannelParams) WithDefaults() *PackerServiceUpdateChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service update channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithTimeout(timeout time.Duration) *PackerServiceUpdateChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithContext(ctx context.Context) *PackerServiceUpdateChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithHTTPClient(client *http.Client) *PackerServiceUpdateChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithBody(body PackerServiceUpdateChannelBody) *PackerServiceUpdateChannelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetBody(body PackerServiceUpdateChannelBody) {
	o.Body = body
}

// WithBucketSlug adds the bucketSlug to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithBucketSlug(bucketSlug string) *PackerServiceUpdateChannelParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceUpdateChannelParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithLocationProjectID(locationProjectID string) *PackerServiceUpdateChannelParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithSlug adds the slug to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) WithSlug(slug string) *PackerServiceUpdateChannelParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the packer service update channel params
func (o *PackerServiceUpdateChannelParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceUpdateChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
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

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
