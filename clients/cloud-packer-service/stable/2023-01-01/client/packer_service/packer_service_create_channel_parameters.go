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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
)

// NewPackerServiceCreateChannelParams creates a new PackerServiceCreateChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceCreateChannelParams() *PackerServiceCreateChannelParams {
	return &PackerServiceCreateChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceCreateChannelParamsWithTimeout creates a new PackerServiceCreateChannelParams object
// with the ability to set a timeout on a request.
func NewPackerServiceCreateChannelParamsWithTimeout(timeout time.Duration) *PackerServiceCreateChannelParams {
	return &PackerServiceCreateChannelParams{
		timeout: timeout,
	}
}

// NewPackerServiceCreateChannelParamsWithContext creates a new PackerServiceCreateChannelParams object
// with the ability to set a context for a request.
func NewPackerServiceCreateChannelParamsWithContext(ctx context.Context) *PackerServiceCreateChannelParams {
	return &PackerServiceCreateChannelParams{
		Context: ctx,
	}
}

// NewPackerServiceCreateChannelParamsWithHTTPClient creates a new PackerServiceCreateChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceCreateChannelParamsWithHTTPClient(client *http.Client) *PackerServiceCreateChannelParams {
	return &PackerServiceCreateChannelParams{
		HTTPClient: client,
	}
}

/*
PackerServiceCreateChannelParams contains all the parameters to send to the API endpoint

	for the packer service create channel operation.

	Typically these are written to a http.Request.
*/
type PackerServiceCreateChannelParams struct {

	// Body.
	Body *models.HashicorpCloudPacker20230101CreateChannelBody

	// BucketName.
	BucketName string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service create channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceCreateChannelParams) WithDefaults() *PackerServiceCreateChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service create channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceCreateChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithTimeout(timeout time.Duration) *PackerServiceCreateChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithContext(ctx context.Context) *PackerServiceCreateChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithHTTPClient(client *http.Client) *PackerServiceCreateChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithBody(body *models.HashicorpCloudPacker20230101CreateChannelBody) *PackerServiceCreateChannelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetBody(body *models.HashicorpCloudPacker20230101CreateChannelBody) {
	o.Body = body
}

// WithBucketName adds the bucketName to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithBucketName(bucketName string) *PackerServiceCreateChannelParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceCreateChannelParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithLocationProjectID(locationProjectID string) *PackerServiceCreateChannelParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceCreateChannelParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service create channel params
func (o *PackerServiceCreateChannelParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceCreateChannelParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service create channel params
func (o *PackerServiceCreateChannelParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceCreateChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_name
	if err := r.SetPathParam("bucket_name", o.BucketName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}