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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/models"
)

// NewPackerServiceCreateRegistryParams creates a new PackerServiceCreateRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceCreateRegistryParams() *PackerServiceCreateRegistryParams {
	return &PackerServiceCreateRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceCreateRegistryParamsWithTimeout creates a new PackerServiceCreateRegistryParams object
// with the ability to set a timeout on a request.
func NewPackerServiceCreateRegistryParamsWithTimeout(timeout time.Duration) *PackerServiceCreateRegistryParams {
	return &PackerServiceCreateRegistryParams{
		timeout: timeout,
	}
}

// NewPackerServiceCreateRegistryParamsWithContext creates a new PackerServiceCreateRegistryParams object
// with the ability to set a context for a request.
func NewPackerServiceCreateRegistryParamsWithContext(ctx context.Context) *PackerServiceCreateRegistryParams {
	return &PackerServiceCreateRegistryParams{
		Context: ctx,
	}
}

// NewPackerServiceCreateRegistryParamsWithHTTPClient creates a new PackerServiceCreateRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceCreateRegistryParamsWithHTTPClient(client *http.Client) *PackerServiceCreateRegistryParams {
	return &PackerServiceCreateRegistryParams{
		HTTPClient: client,
	}
}

/*
PackerServiceCreateRegistryParams contains all the parameters to send to the API endpoint

	for the packer service create registry operation.

	Typically these are written to a http.Request.
*/
type PackerServiceCreateRegistryParams struct {

	// Body.
	Body *models.HashicorpCloudPacker20220411CreateRegistryRequest

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service create registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceCreateRegistryParams) WithDefaults() *PackerServiceCreateRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service create registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceCreateRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithTimeout(timeout time.Duration) *PackerServiceCreateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithContext(ctx context.Context) *PackerServiceCreateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithHTTPClient(client *http.Client) *PackerServiceCreateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithBody(body *models.HashicorpCloudPacker20220411CreateRegistryRequest) *PackerServiceCreateRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetBody(body *models.HashicorpCloudPacker20220411CreateRegistryRequest) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceCreateRegistryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) WithLocationProjectID(locationProjectID string) *PackerServiceCreateRegistryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service create registry params
func (o *PackerServiceCreateRegistryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceCreateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}