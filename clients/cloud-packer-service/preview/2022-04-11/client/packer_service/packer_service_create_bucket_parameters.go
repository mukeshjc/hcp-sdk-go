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

// NewPackerServiceCreateBucketParams creates a new PackerServiceCreateBucketParams object
// with the default values initialized.
func NewPackerServiceCreateBucketParams() *PackerServiceCreateBucketParams {
	var ()
	return &PackerServiceCreateBucketParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceCreateBucketParamsWithTimeout creates a new PackerServiceCreateBucketParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceCreateBucketParamsWithTimeout(timeout time.Duration) *PackerServiceCreateBucketParams {
	var ()
	return &PackerServiceCreateBucketParams{

		timeout: timeout,
	}
}

// NewPackerServiceCreateBucketParamsWithContext creates a new PackerServiceCreateBucketParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceCreateBucketParamsWithContext(ctx context.Context) *PackerServiceCreateBucketParams {
	var ()
	return &PackerServiceCreateBucketParams{

		Context: ctx,
	}
}

// NewPackerServiceCreateBucketParamsWithHTTPClient creates a new PackerServiceCreateBucketParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceCreateBucketParamsWithHTTPClient(client *http.Client) *PackerServiceCreateBucketParams {
	var ()
	return &PackerServiceCreateBucketParams{
		HTTPClient: client,
	}
}

/*PackerServiceCreateBucketParams contains all the parameters to send to the API endpoint
for the packer service create bucket operation typically these are written to a http.Request
*/
type PackerServiceCreateBucketParams struct {

	/*Body*/
	Body *models.HashicorpCloudPacker20220411CreateBucketRequest
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithTimeout(timeout time.Duration) *PackerServiceCreateBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithContext(ctx context.Context) *PackerServiceCreateBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithHTTPClient(client *http.Client) *PackerServiceCreateBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithBody(body *models.HashicorpCloudPacker20220411CreateBucketRequest) *PackerServiceCreateBucketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetBody(body *models.HashicorpCloudPacker20220411CreateBucketRequest) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceCreateBucketParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) WithLocationProjectID(locationProjectID string) *PackerServiceCreateBucketParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service create bucket params
func (o *PackerServiceCreateBucketParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceCreateBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
