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

// NewGetGitHubRepositoriesParams creates a new GetGitHubRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGitHubRepositoriesParams() *GetGitHubRepositoriesParams {
	return &GetGitHubRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitHubRepositoriesParamsWithTimeout creates a new GetGitHubRepositoriesParams object
// with the ability to set a timeout on a request.
func NewGetGitHubRepositoriesParamsWithTimeout(timeout time.Duration) *GetGitHubRepositoriesParams {
	return &GetGitHubRepositoriesParams{
		timeout: timeout,
	}
}

// NewGetGitHubRepositoriesParamsWithContext creates a new GetGitHubRepositoriesParams object
// with the ability to set a context for a request.
func NewGetGitHubRepositoriesParamsWithContext(ctx context.Context) *GetGitHubRepositoriesParams {
	return &GetGitHubRepositoriesParams{
		Context: ctx,
	}
}

// NewGetGitHubRepositoriesParamsWithHTTPClient creates a new GetGitHubRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGitHubRepositoriesParamsWithHTTPClient(client *http.Client) *GetGitHubRepositoriesParams {
	return &GetGitHubRepositoriesParams{
		HTTPClient: client,
	}
}

/*
GetGitHubRepositoriesParams contains all the parameters to send to the API endpoint

	for the get git hub repositories operation.

	Typically these are written to a http.Request.
*/
type GetGitHubRepositoriesParams struct {

	// InstallationName.
	InstallationName *string

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

// WithDefaults hydrates default values in the get git hub repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubRepositoriesParams) WithDefaults() *GetGitHubRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get git hub repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubRepositoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithTimeout(timeout time.Duration) *GetGitHubRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithContext(ctx context.Context) *GetGitHubRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithHTTPClient(client *http.Client) *GetGitHubRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallationName adds the installationName to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithInstallationName(installationName *string) *GetGitHubRepositoriesParams {
	o.SetInstallationName(installationName)
	return o
}

// SetInstallationName adds the installationName to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetInstallationName(installationName *string) {
	o.InstallationName = installationName
}

// WithLocationOrganizationID adds the locationOrganizationID to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithLocationOrganizationID(locationOrganizationID string) *GetGitHubRepositoriesParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithLocationProjectID(locationProjectID string) *GetGitHubRepositoriesParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithLocationRegionProvider(locationRegionProvider *string) *GetGitHubRepositoriesParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) WithLocationRegionRegion(locationRegionRegion *string) *GetGitHubRepositoriesParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get git hub repositories params
func (o *GetGitHubRepositoriesParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitHubRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InstallationName != nil {

		// query param installation_name
		var qrInstallationName string

		if o.InstallationName != nil {
			qrInstallationName = *o.InstallationName
		}
		qInstallationName := qrInstallationName
		if qInstallationName != "" {

			if err := r.SetQueryParam("installation_name", qInstallationName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}