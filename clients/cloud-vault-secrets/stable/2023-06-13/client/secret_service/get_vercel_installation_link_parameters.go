// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewGetVercelInstallationLinkParams creates a new GetVercelInstallationLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVercelInstallationLinkParams() *GetVercelInstallationLinkParams {
	return &GetVercelInstallationLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVercelInstallationLinkParamsWithTimeout creates a new GetVercelInstallationLinkParams object
// with the ability to set a timeout on a request.
func NewGetVercelInstallationLinkParamsWithTimeout(timeout time.Duration) *GetVercelInstallationLinkParams {
	return &GetVercelInstallationLinkParams{
		timeout: timeout,
	}
}

// NewGetVercelInstallationLinkParamsWithContext creates a new GetVercelInstallationLinkParams object
// with the ability to set a context for a request.
func NewGetVercelInstallationLinkParamsWithContext(ctx context.Context) *GetVercelInstallationLinkParams {
	return &GetVercelInstallationLinkParams{
		Context: ctx,
	}
}

// NewGetVercelInstallationLinkParamsWithHTTPClient creates a new GetVercelInstallationLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVercelInstallationLinkParamsWithHTTPClient(client *http.Client) *GetVercelInstallationLinkParams {
	return &GetVercelInstallationLinkParams{
		HTTPClient: client,
	}
}

/*
GetVercelInstallationLinkParams contains all the parameters to send to the API endpoint

	for the get vercel installation link operation.

	Typically these are written to a http.Request.
*/
type GetVercelInstallationLinkParams struct {

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

// WithDefaults hydrates default values in the get vercel installation link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVercelInstallationLinkParams) WithDefaults() *GetVercelInstallationLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vercel installation link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVercelInstallationLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithTimeout(timeout time.Duration) *GetVercelInstallationLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithContext(ctx context.Context) *GetVercelInstallationLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithHTTPClient(client *http.Client) *GetVercelInstallationLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithLocationOrganizationID(locationOrganizationID string) *GetVercelInstallationLinkParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithLocationProjectID(locationProjectID string) *GetVercelInstallationLinkParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithLocationRegionProvider(locationRegionProvider *string) *GetVercelInstallationLinkParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) WithLocationRegionRegion(locationRegionRegion *string) *GetVercelInstallationLinkParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get vercel installation link params
func (o *GetVercelInstallationLinkParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetVercelInstallationLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
