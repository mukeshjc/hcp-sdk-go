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

// NewGetBuildParams creates a new GetBuildParams object
// with the default values initialized.
func NewGetBuildParams() *GetBuildParams {
	var ()
	return &GetBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildParamsWithTimeout creates a new GetBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildParamsWithTimeout(timeout time.Duration) *GetBuildParams {
	var ()
	return &GetBuildParams{

		timeout: timeout,
	}
}

// NewGetBuildParamsWithContext creates a new GetBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildParamsWithContext(ctx context.Context) *GetBuildParams {
	var ()
	return &GetBuildParams{

		Context: ctx,
	}
}

// NewGetBuildParamsWithHTTPClient creates a new GetBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildParamsWithHTTPClient(client *http.Client) *GetBuildParams {
	var ()
	return &GetBuildParams{
		HTTPClient: client,
	}
}

/*GetBuildParams contains all the parameters to send to the API endpoint
for the get build operation typically these are written to a http.Request
*/
type GetBuildParams struct {

	/*BuildID
	  build ULID

	*/
	BuildID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build params
func (o *GetBuildParams) WithTimeout(timeout time.Duration) *GetBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build params
func (o *GetBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build params
func (o *GetBuildParams) WithContext(ctx context.Context) *GetBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build params
func (o *GetBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build params
func (o *GetBuildParams) WithHTTPClient(client *http.Client) *GetBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build params
func (o *GetBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the get build params
func (o *GetBuildParams) WithBuildID(buildID string) *GetBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the get build params
func (o *GetBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithLocationOrganizationID adds the locationOrganizationID to the get build params
func (o *GetBuildParams) WithLocationOrganizationID(locationOrganizationID string) *GetBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get build params
func (o *GetBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get build params
func (o *GetBuildParams) WithLocationProjectID(locationProjectID string) *GetBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get build params
func (o *GetBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get build params
func (o *GetBuildParams) WithLocationRegionProvider(locationRegionProvider *string) *GetBuildParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get build params
func (o *GetBuildParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get build params
func (o *GetBuildParams) WithLocationRegionRegion(locationRegionRegion *string) *GetBuildParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get build params
func (o *GetBuildParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
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