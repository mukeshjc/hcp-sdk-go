// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

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

// NewBoundaryServiceAvailableUpdateGetParams creates a new BoundaryServiceAvailableUpdateGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBoundaryServiceAvailableUpdateGetParams() *BoundaryServiceAvailableUpdateGetParams {
	return &BoundaryServiceAvailableUpdateGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBoundaryServiceAvailableUpdateGetParamsWithTimeout creates a new BoundaryServiceAvailableUpdateGetParams object
// with the ability to set a timeout on a request.
func NewBoundaryServiceAvailableUpdateGetParamsWithTimeout(timeout time.Duration) *BoundaryServiceAvailableUpdateGetParams {
	return &BoundaryServiceAvailableUpdateGetParams{
		timeout: timeout,
	}
}

// NewBoundaryServiceAvailableUpdateGetParamsWithContext creates a new BoundaryServiceAvailableUpdateGetParams object
// with the ability to set a context for a request.
func NewBoundaryServiceAvailableUpdateGetParamsWithContext(ctx context.Context) *BoundaryServiceAvailableUpdateGetParams {
	return &BoundaryServiceAvailableUpdateGetParams{
		Context: ctx,
	}
}

// NewBoundaryServiceAvailableUpdateGetParamsWithHTTPClient creates a new BoundaryServiceAvailableUpdateGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewBoundaryServiceAvailableUpdateGetParamsWithHTTPClient(client *http.Client) *BoundaryServiceAvailableUpdateGetParams {
	return &BoundaryServiceAvailableUpdateGetParams{
		HTTPClient: client,
	}
}

/*
BoundaryServiceAvailableUpdateGetParams contains all the parameters to send to the API endpoint

	for the boundary service available update get operation.

	Typically these are written to a http.Request.
*/
type BoundaryServiceAvailableUpdateGetParams struct {

	/* ClusterID.

	   cluster_id is the id of the cluster set by user on creation.
	*/
	ClusterID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the boundary service available update get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BoundaryServiceAvailableUpdateGetParams) WithDefaults() *BoundaryServiceAvailableUpdateGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the boundary service available update get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BoundaryServiceAvailableUpdateGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithTimeout(timeout time.Duration) *BoundaryServiceAvailableUpdateGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithContext(ctx context.Context) *BoundaryServiceAvailableUpdateGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithHTTPClient(client *http.Client) *BoundaryServiceAvailableUpdateGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithClusterID(clusterID string) *BoundaryServiceAvailableUpdateGetParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithLocationOrganizationID(locationOrganizationID string) *BoundaryServiceAvailableUpdateGetParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithLocationProjectID(locationProjectID string) *BoundaryServiceAvailableUpdateGetParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithLocationRegionProvider(locationRegionProvider *string) *BoundaryServiceAvailableUpdateGetParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) WithLocationRegionRegion(locationRegionRegion *string) *BoundaryServiceAvailableUpdateGetParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the boundary service available update get params
func (o *BoundaryServiceAvailableUpdateGetParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *BoundaryServiceAvailableUpdateGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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
